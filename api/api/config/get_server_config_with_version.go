package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-speed/api/api/common"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/response"
	"go-speed/service"
	"go-speed/util"
	"go-speed/util/geo"
	"net/http"
	"sync"
)

var (
	UserGetConfigCounter map[string]uint64
	mu                   sync.Mutex // 使用Mutex来保证线程安全
)

func init() {
	UserGetConfigCounter = make(map[string]uint64)
}

// GetCounter 安全地获取当前计数器的值
func GetCounter(country string) uint64 {
	mu.Lock()
	defer mu.Unlock()
	if count, ok := UserGetConfigCounter[country]; ok {
		return count
	} else {
		return 0
	}
}

func IncrementCounter(country string) uint64 {
	mu.Lock()
	defer mu.Unlock()
	if count, ok := UserGetConfigCounter[country]; ok {
		if count > 100000000000 {
			count = 0
		}
		UserGetConfigCounter[country] = count + 1
	} else {
		UserGetConfigCounter[country] = 1
	}
	return UserGetConfigCounter[country]
}

type GetServerConfigWithoutRulesReq struct {
	UserId      uint64 `form:"user_id" binding:"required" json:"user_id"`
	CountryName string `form:"country_name" json:"country_name"`
}

type GetServerConfigWithoutRulesRes struct {
}

func GetServerConfigWithoutRules(ctx *gin.Context) {
	var (
		err          error
		req          = new(GetServerConfigWithoutRulesReq)
		userEntity   *entity.TUser
		winCountry   entity.TServingCountry
		nodeEntities []entity.TNode
	)
	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
		response.RespFail(ctx, i18n.RetMsgParamParseErr, nil)
		return
	}
	global.MyLogger(ctx).Info().Msgf(">>> req: %+v", *req)

	userEntity, err = common.CheckUserByUserId(ctx, req.UserId)
	if err != nil {
		return
	}

	winCountry, nodeEntities, err = chooseCountryForUser(ctx, req.UserId, req.CountryName)
	if err != nil {
		return
	}
	nodesLen := len(nodeEntities)
	global.MyLogger(ctx).Info().Msgf(">>> winCountry: %s, len(nodeEntities): %d", winCountry.Name, nodesLen)

	dnss := make([]string, 0)
	v2rayServers := make([]Server, 0)

	index, counter, nodeID := 0, GetCounter(winCountry.Name), int64(0)
	if nodesLen > 1 {
		index = int(counter % uint64(nodesLen))
		nodeID, _ = service.GetMinLoadNode(ctx, nodeEntities)
	}

	global.MyLogger(ctx).Info().Msgf("[choose-node-for-user-1] userId: %d, index: %d, counter: %d, nodesLen: %d, country: %s, nodeID: %d", userEntity.Id, index, counter, nodesLen, winCountry.Name, nodeID)
	for i, item := range nodeEntities {
		if nodeID == 0 {
			if i != index {
				continue
			}
		} else {
			if item.Id != nodeID {
				continue
			}
		}

		global.MyLogger(ctx).Info().Msgf("[choose-node-for-user-2] (%s) (%s) (%s) (%s) (nodeID: %d)", userEntity.Email, item.Ip, item.CountryEn, winCountry.Name, nodeID)

		err = service.AddUserConfigToNode(ctx, userEntity, &item)
		if err != nil {
			global.MyLogger(ctx).Err(err).Msgf("AddUserConfigToNode failed, node: %s", item.Ip)
			continue
		}

		nodeId := item.Id
		nodePorts := []int{item.Port}
		for x := item.MinPort; x <= item.MaxPort; x++ {
			nodePorts = append(nodePorts, x)
		}
		global.MyLogger(ctx).Info().Msgf(">>>>> nodeId:%d, nodePorts: %+v", nodeId, nodePorts)
		var dnsList []entity.TNodeDns
		err = dao.TNodeDns.Ctx(ctx).Where(do.TNodeDns{
			NodeId: nodeId,
			Level:  userEntity.Level + 1,
		}).Scan(&dnsList)
		if err != nil {
			global.MyLogger(ctx).Err(err).Msgf(">>>>> FindNodeDnsByNodeId failed: %+v", err.Error())
			continue
		}
		for _, dns := range dnsList {
			for _, nodePort := range nodePorts {
				v2rayServers = append(v2rayServers, Server{
					Email:    util.GetUserV2rayConfigEmail(userEntity.Email),
					Password: util.GetUserV2rayConfigUUID(userEntity.V2RayUuid),
					Port:     nodePort,
					Address:  dns.Dns,
				})
			}
			if !geo.IsInArrayIgnoreCase(dns.Dns, dnss) {
				dnss = append(dnss, dns.Dns)
			}
		}
	}
	global.MyLogger(ctx).Info().Msgf(">>>>> dns: %+v", dnss)
	global.MyLogger(ctx).Info().Msgf(">>>>> v2rayServers: %+v", v2rayServers)

	if len(v2rayServers) == 0 {
		global.MyLogger(ctx).Err(errors.New("获取V2ray配置失败")).Msgf("v2rayServers is empty")
		response.RespFail(ctx, i18n.RetMsgGetV2rayConfigFailed, nil)
		return
	}
	v, err := json.Marshal(GenV2rayConfig(ctx, v2rayServers, winCountry.Name, true))
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("GenV2rayConfig failed")
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}
	global.MyLogger(ctx).Info().Msgf(">>>>> V2rayConfig: %s", string(v))
	ctx.String(http.StatusOK, fmt.Sprintf(string(v)))
}
