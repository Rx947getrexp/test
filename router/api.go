package router

import (
	"github.com/gin-gonic/gin"
	"go-speed/api/api"
)

func ApiRoute(group *gin.RouterGroup) {
	group.POST("generate_dev_id", api.GenerateDevId)
	group.POST("send_email", api.SendEmail)
	group.POST("reg", api.Reg)
	group.POST("login", api.Login)
	group.POST("forget_passwd", api.ForgetPasswd)

	group.GET("notice_list", api.NoticeList)
	group.GET("notice_detail", api.NoticeDetail)
	group.GET("node_list", api.NodeList)
	group.GET("dns_list", api.DnsList)
	group.GET("combo_list", api.ComboList)
	group.GET("ad_list", api.AdList)
	group.GET("app_info", api.AppInfo)
	group.GET("app_filter", api.AppFilter)

	group.Use(api.JWTAuth())
	{
		group.POST("change_passwd", api.ChangePasswd)
		group.GET("user_info", api.UserInfo)

		group.GET("team_list", api.TeamList)
		group.GET("team_info", api.TeamInfo)

		group.POST("receive_free", api.ReceiveFree)
		group.GET("receive_free_summary", api.ReceiveFreeSummary)

		group.POST("upload_log", api.UploadLog)

		group.POST("create_order", api.CreateOrder)
		group.GET("order_list", api.OrderList)

		group.GET("dev_list", api.DevList)
		group.POST("ban_dev", api.BanDev)

		group.POST("change_network", api.ChangeNetwork)
		//group.POST("switch_button_status", api.SwitchButtonStatus)

	}
}
