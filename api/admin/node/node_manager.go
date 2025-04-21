package node

import (
	"context"
	"fmt"
	"go-speed/constant"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/response"
	"go-speed/service"
	"go-speed/util"
	"io"
	"math"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
)

type NodeManagerReq struct {
	Server string `form:"server" binding:"required" json:"server" dc:"公网域名"`
	Status int    `form:"status" binding:"required" json:"status" dc:"状态。1-正常；2-异常"`
}

const (
	secretKey      = "@hsspeed2025#"
	timeoutSeconds = 300 // 5 分钟
)

var nodeDebounceCache = make(map[string]time.Time)
var debounceMutex sync.Mutex

const debounceDuration = 5 * time.Minute

const (
	CountryStatusActive   = 1 // 国家上架
	CountryStatusInactive = 2 // 国家下架

	NodeStatusActive   = 1 // 节点上架
	NodeStatusInactive = 2 // 节点下架
)

// 根据上报的节点状态，自动上下架节点机器和对应国家
func ReportNodeStatus(c *gin.Context) {

	ip := c.ClientIP()
	if ip != "185.22.154.21" { //探测机器的IP
		global.Logger.Warn().Msgf("非法请求IP:%v", ip)
		response.RespFail(c, i18n.RetMsgOperateFailed, nil)
		return
	}

	timestamp := c.GetHeader("X-Timestamp")
	nonce := c.GetHeader("X-Nonce")
	signature := c.GetHeader("X-Signature")

	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		global.Logger.Warn().Msgf("读取请求体失败:%s", err.Error())
		response.RespFail(c, i18n.RetMsgOperateFailed, nil)
		return
	}
	// 检查时间戳有效性
	tsInt, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil || math.Abs(float64(time.Now().Unix()-tsInt)) > timeoutSeconds {
		global.Logger.Warn().Msgf("X-Timestamp 已过期:%s", timestamp)
		response.RespFail(c, i18n.RetMsgOperateFailed, nil)
		return
	}
	// 校验签名
	if !util.CheckSignature(secretKey, timestamp, nonce, string(bodyBytes), signature) {
		global.Logger.Warn().Msgf("签名验证失败:%s", err.Error())
		response.RespFail(c, i18n.RetMsgOperateFailed, nil)
		return
	}

	c.Request.Body = io.NopCloser(strings.NewReader(string(bodyBytes)))

	req := new(NodeManagerReq)
	if err := c.ShouldBind(req); err != nil {
		global.Logger.Err(err).Msgf("参数校验失败，err:%v", err.Error())
		response.RespFail(c, i18n.RetMsgParamInvalid, nil)
		return
	}

	// 通过dns获取节点IP
	nodeIp, err := service.GetNodeIpByServer(c, req.Server)
	if err != nil {
		global.Logger.Err(err).Msgf("通过dns获取节点IP失败，err:%v", err.Error())
		response.RespFail(c, i18n.RetMsgOperateFailed, nil)
		return
	}

	nodeIp = strings.TrimSpace(nodeIp) //还真出现了存到库里的IP前后有空格，导致拼接ping url出现问题，导致ping报错。
	global.Logger.Info().Msgf("节点查询结果成功，nodeIp： %s。", nodeIp)

	// === 防抖逻辑 ===
	debounceMutex.Lock()
	lastHandled, exists := nodeDebounceCache[nodeIp]
	if exists && time.Since(lastHandled) < debounceDuration {
		debounceMutex.Unlock()
		global.Logger.Warn().Msgf("节点 %s 操作过于频繁，上次处理时间: %v", nodeIp, lastHandled)
		response.RespFail(c, i18n.RetMsgOperateFailed, nil)
		return
	}
	nodeDebounceCache[nodeIp] = time.Now()
	debounceMutex.Unlock()

	// 拨测
	ctx, _ := gin.CreateTestContext(nil)
	defer ctx.Done()
	_, probeErr := service.GetSysStatsByIp(ctx, nodeIp) //通过节点ip ping节点状态

	switch req.Status {
	case 1: //上报节点正常
		if probeErr != nil {
			global.Logger.Err(err).Msgf("拨测失败，节点未恢复，忽略上报，err:%v", err.Error())
			response.RespFail(c, i18n.RetMsgOperateFailed, nil)
			return
		}
		err = updateNodeStatus(c, nodeIp, NodeStatusActive)
		if err != nil {
			global.Logger.Err(err).Msgf("更新节点状态失败，err:%v", err.Error())
			response.RespFail(c, i18n.RetMsgOperateFailed, nil)
			return
		}
		global.Logger.Info().Msgf("节点 %s 已恢复，已上线", nodeIp)
		response.RespOk(c, i18n.RetMsgSuccess, &response.Response{
			Code: 1,
		})

	case 2: //上报节点异常
		// 如果节点是正常的，直接返回正常
		if probeErr == nil {
			response.RespOk(c, i18n.RetMsgSuccess, nil)
			return
		}
		err = updateNodeStatus(c, nodeIp, NodeStatusInactive) //下架机器
		if err != nil {
			global.Logger.Err(err).Msgf("更新节点状态失败，err:%v", err.Error())
			response.RespFail(c, i18n.RetMsgOperateFailed, nil)
			return
		}
		global.Logger.Info().Msgf("节点 %s 异常已确认，已下架", nodeIp)
		response.RespOk(c, i18n.RetMsgSuccess, &response.Response{
			Code: 1,
		})

	default:
		response.RespFail(c, i18n.RetMsgOperateFailed, nil)
	}
}

// 节点机器上下架
func updateNodeStatus(c *gin.Context, ip string, status int) error {
	if ip == "" || status == 0 {
		global.Logger.Warn().Msgf("参数异常：ip[%s],status[%d]", ip, status)
		return fmt.Errorf("参数异常：ip[%s],status[%d]", ip, status)
	}

	global.Logger.Info().Msgf("更新节点 [%s] 状态为 [%d]", ip, status)
	ctx := c
	// 在事务中执行更新逻辑
	err := dao.TNode.Ctx(c).Transaction(c, func(c context.Context, tx gdb.TX) error {
		// 使用事务更新节点状态
		_, err := dao.TNode.Ctx(c).Where(do.TNode{Ip: ip}).Data(do.TNode{
			Status:    status,
			UpdatedAt: gtime.Now(),
		}).Update()
		if err != nil {
			global.MyLogger(ctx).Err(err).Msgf("事务中更新节点状态失败，err:%v", err)
			return err
		}
		//更新国家状态
		if err := updateCountryStatusByNode(ctx, ip); err != nil {
			global.MyLogger(ctx).Err(err).Msgf("事务中更新国家状态失败，err:%v", err)
			return err
		}
		return nil
	})
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("updateNodeStatus 更新节点【%s】失败。", ip)
		return err
	}

	return nil
}

// 国家自动上下架
func updateCountryStatusByNode(c *gin.Context, ip string) error {
	var (
		countryEntities *entity.TNode
		currentEntity   *entity.TServingCountry
	)
	// 查询节点的国家字段
	err := dao.TNode.Ctx(c).
		Where(do.TNode{Ip: ip}).
		Order(dao.TNode.Columns().UpdatedAt, constant.OrderTypeDesc).
		Limit(1).
		Scan(&countryEntities)
	if err != nil {
		global.Logger.Err(err).Msgf("查询节点 [%s] 信息失败：%v", ip, err)
		return err
	}

	country := countryEntities.CountryEn

	if country == "" {
		global.Logger.Warn().Msgf("节点 [%s] 获取国家失败", ip)
		return fmt.Errorf("节点 [%s] 获取国家失败", ip)
	}

	// 查出该国家下是否有“至少一个”状态正常的节点
	count, err := dao.TNode.Ctx(c).
		Where(do.TNode{Status: NodeStatusActive, CountryEn: country}).
		Count()
	if err != nil {
		global.Logger.Err(err).Msgf("统计国家下节点状态失败：%v", err)
		return err
	}

	// 更新国家状态
	newStatus := CountryStatusInactive // 默认下架
	if count > 0 {
		newStatus = CountryStatusActive // 有至少一个正常节点
	}

	// 先获取状态，看是不是要更新
	err = dao.TServingCountry.Ctx(c).
		Where(do.TServingCountry{Name: country}).
		Limit(1).
		Scan(&currentEntity)
	if err != nil {
		global.Logger.Err(err).Msgf("查询国家当前状态失败：%v", err)
		return err
	}
	if currentEntity.Status == newStatus {
		return nil // 无需更新
	}

	// 实施更新（上下架）
	_, err = dao.TServingCountry.Ctx(c).Where(do.TServingCountry{Name: country}).Data(do.TServingCountry{
		Status:    newStatus,
		UpdatedAt: gtime.Now(),
	}).Update()

	if err != nil {
		global.Logger.Err(err).Msgf("更新国家状态失败：%v", err)
		return err
	}

	global.Logger.Info().Msgf("国家 [%s] 状态已更新为 %d", country, newStatus)
	return nil
}
