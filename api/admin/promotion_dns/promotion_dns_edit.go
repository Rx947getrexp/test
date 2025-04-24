package promotion_dns

import (
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/response"
	"go-speed/service"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/os/gtime"
)

type PromotionDnsEditRequest struct {
	Id             int64   `form:"id" binding:"required" json:"id" dc:"机器id"`
	Dns            *string `form:"dns" json:"dns" dc:"机器域名"`
	Ip             *string `form:"ip" json:"ip" dc:"ip地址"`
	Status         *int64  `form:"status" json:"status" dc:"状态"`
	MacChannel     *string `form:"mac_channel" json:"mac_channel" dc:"苹果电脑渠道"`
	WinChannel     *string `form:"win_channel" json:"win_channel" dc:"windows电脑渠道"`
	AndroidChannel *string `form:"android_channel" json:"android_channel" dc:"安卓渠道"`
	Promoter       *string `form:"promoter" json:"promoter" dc:"推广人员"`
	Comment        *string `form:"comment" json:"comment" dc:"备注信息"`
}

func PromotionDnsEdit(c *gin.Context) {
	// 定义局部变量
	var (
		err    error
		req    = new(PromotionDnsEditRequest)
		entity *entity.TPromotionDns
	)

	if err = c.ShouldBind(req); err != nil {
		global.Logger.Err(err).Msg(err.Error())
		response.RespFail(c, "绑定参数失败", nil)
		return
	}

	// 从上下文中获取用户信息
	claims := c.MustGet("claims").(*service.CustomClaims)
	// 根据用户信息获取管理员用户对象
	adminUser, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msgf("Illegal user, err: %v", err)
		response.RespFail(c, "用户不合法！", nil)
		return
	}

	// 对 req.Status 的合法值判断（1 或 2）
	if req.Status != nil && *req.Status != 1 && *req.Status != 2 {
		response.RespFail(c, "状态参数不合法，仅支持 1(上架) 或 2(下架)", nil)
		return
	}

	err = dao.TPromotionDns.Ctx(c).Where(do.TPromotionDns{Id: req.Id}).Scan(&entity)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("查询数据失败，error: %v", err)
		response.RespFail(c, "查询数据失败", nil)
		return
	}
	if entity == nil {
		global.MyLogger(c).Error().Msgf("数据不存在，id: %d", req.Id)
		response.RespFail(c, "数据不存在", nil)
		return
	}

	// 更新域名信息
	time := gtime.Now()
	// 存储要更新的字段
	updateDo := do.TPromotionDns{
		UpdatedAt: time,
		Author:    adminUser.Uname,
	}

	if req.Dns != nil && *req.Dns != "" {
		updateDo.Dns = *req.Dns
	}

	if req.Ip != nil && *req.Ip != "" {
		updateDo.Ip = *req.Ip
	}

	if req.Status != nil && *req.Status != 0 {
		updateDo.Status = *req.Status
	}

	if req.MacChannel != nil && *req.MacChannel != "" {
		updateDo.MacChannel = *req.MacChannel
	}

	if req.WinChannel != nil && *req.WinChannel != "" {
		updateDo.WinChannel = *req.WinChannel
	}

	if req.AndroidChannel != nil && *req.AndroidChannel != "" {
		updateDo.AndroidChannel = *req.AndroidChannel
	}

	if req.Promoter != nil {
		updateDo.Promoter = *req.Promoter
	}

	if req.Comment != nil {
		updateDo.Comment = *req.Comment
	}

	// 更新数据
	_, err = dao.TPromotionDns.Ctx(c).Where(do.TPromotionDns{Id: req.Id}).Data(updateDo).Update()
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("更新数据失败，error: %v", err)
		response.RespFail(c, "更新数据失败", nil)
		return
	}

	// 返回成功
	response.RespOk(c, i18n.RetMsgSuccess, nil)
}
