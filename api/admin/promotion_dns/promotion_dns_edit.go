package promotion_dns

import (
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/response"
	"go-speed/service"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/os/gtime"
)

type PromotionDnsEditRequest struct {
	Id             int64  `form:"id" binding:"required" json:"id" dc:"机器id"`
	Dns            string `form:"dns" json:"dns" dc:"机器域名"`
	Ip             string `form:"ip" json:"ip" dc:"ip地址"`
	Status         int64  `form:"status" json:"status" dc:"状态"`
	MacChannel     string `form:"mac_channel" json:"mac_channel" dc:"苹果电脑渠道"`
	WinChannel     string `form:"win_channel" json:"win_channel" dc:"windows电脑渠道"`
	AndroidChannel string `form:"android_channel" json:"android_channel" dc:"安卓渠道"`
	Promoter       string `form:"promoter" json:"promoter" dc:"推广人员"`
	Comment        string `form:"comment" json:"comment" dc:"备注信息"`
}

func PromotionDnsEdit(c *gin.Context) {
	// 定义局部变量
	var (
		err error
		req = new(PromotionDnsEditRequest)
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

	// 更新域名信息
	updateData := make(map[string]interface{}) // 存储要更新的字段
	time := gtime.Now()
	updateData["updated_at"] = time        // 更新时间必须更新
	updateData["author"] = adminUser.Uname // 更新此次操作人

	if req.Dns != "" {
		updateData["dns"] = req.Dns
	}

	if req.Ip != "" {
		updateData["ip"] = req.Ip
	}

	if req.Status != 0 {
		updateData["status"] = req.Status
	}

	if req.MacChannel != "" {
		updateData["mac_channel"] = req.MacChannel
	}

	if req.WinChannel != "" {
		updateData["win_channel"] = req.WinChannel
	}

	if req.AndroidChannel != "" {
		updateData["android_channel"] = req.AndroidChannel
	}

	if req.Promoter != "" {
		updateData["promoter"] = req.Promoter
	}

	if req.Comment != "" {
		updateData["comment"] = req.Comment
	}

	// 更新数据
	_, err = dao.TPromotionDns.Ctx(c).Where("id", req.Id).Data(updateData).Update()
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("更新数据失败，error: %v", err)
		response.RespFail(c, "更新数据失败", nil)
		return
	}

	// 返回成功
	response.RespOk(c, i18n.RetMsgSuccess, nil)
}
