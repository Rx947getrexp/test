package promotion_dns

import (
	"fmt"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/response"
	"go-speed/service"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/os/gtime"
)

type PromotionDnsAddRequest struct {
	Dns            string `form:"dns" binding:"required" json:"dns" dc:"机器域名"`
	Ip             string `form:"ip" binding:"required" json:"ip" dc:"ip地址"`
	MacChannel     string `form:"mac_channel" binding:"required" json:"mac_channel" dc:"苹果电脑渠道"`
	WinChannel     string `form:"win_channel" binding:"required" json:"win_channel" dc:"windows电脑渠道"`
	AndroidChannel string `form:"android_channel" binding:"required" json:"android_channel" dc:"安卓渠道"`
	Promoter       string `form:"promoter" binding:"required" json:"promoter" dc:"推广人员"`
	Comment        string `form:"comment" json:"comment" dc:"备注信息"`
}

func PromotionDnsAdd(c *gin.Context) {
	// 定义局部变量
	var (
		err    error
		req    = new(PromotionDnsAddRequest)
		entity *do.TPromotionDns
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
		global.Logger.Err(err).Msg("用户不合法！")
		response.RespFail(c, "用户不合法！", nil)
		return
	}

	err = dao.TPromotionDns.Ctx(c).Where(do.TPromotionDns{Dns: req.Dns}).Scan(&entity)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("查询数据失败，error: %v", err)
		response.RespFail(c, "查询数据失败", nil)
		return
	}
	if entity != nil {
		msg := fmt.Sprintf("推广域名:%s，已存在。", req.Dns)
		global.MyLogger(c).Warn().Msgf(msg)
		response.RespFail(c, msg, nil)
		return
	}

	// 获取当前时间
	now := gtime.Now()
	// 插入新的域名机器记录，并获取插入数据的ID
	lastInsertId, err := dao.TPromotionDns.Ctx(c).Data(do.TPromotionDns{
		Dns:            req.Dns,
		Ip:             req.Ip,
		MacChannel:     req.MacChannel,
		WinChannel:     req.WinChannel,
		AndroidChannel: req.AndroidChannel,
		Promoter:       req.Promoter,
		Status:         1,
		Author:         adminUser.Uname,
		Comment:        req.Comment,
		CreatedAt:      now,
		UpdatedAt:      now,
	}).InsertAndGetId()

	if err != nil || lastInsertId == 0 {
		global.MyLogger(c).Err(err).Msgf("添加数据失败，error: %v", err)
		response.RespFail(c, "添加数据失败", nil)
		return
	}

	// 返回成功响应
	response.RespOk(c, i18n.RetMsgSuccess, nil)
}
