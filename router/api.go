package router

import (
	"go-speed/api/api"

	"github.com/gin-gonic/gin"
)

func ApiRoute(group *gin.RouterGroup) {
	//group.Use(api.PrintParam())
	//{}
	group.POST("generate_dev_id", api.GenerateDevId) // call
	group.POST("send_email", api.SendEmail)
	group.POST("reg", api.Reg)     // call
	group.POST("login", api.Login) // call
	group.POST("forget_passwd", api.ForgetPasswd)
	group.GET("notice_list", api.NoticeList)
	group.GET("notice_detail", api.NoticeDetail)
	group.GET("node_list", api.NodeList) // call
	group.GET("dns_list", api.DnsList)   // call
	group.GET("combo_list", api.ComboList)
	group.GET("ad_list", api.AdList)
	group.GET("expire_user", api.ExpireUserList)
	group.GET("app_info", api.AppInfo)      // call
	group.GET("pc_app_info", api.PCAppInfo) // call
	group.GET("app_filter", api.AppFilter)  //策略审核

	group.Use(api.JWTAuth())
	{
		group.POST("change_passwd", api.ChangePasswd)
		group.GET("user_info", api.UserInfo) // call
		group.GET("team_list", api.TeamList)
		group.GET("team_info", api.TeamInfo)
		group.POST("receive_free", api.ReceiveFree)
		group.GET("receive_free_summary", api.ReceiveFreeSummary)
		group.POST("upload_log", api.UploadLog)
		group.POST("create_order", api.CreateOrder)
		group.GET("order_list", api.OrderList)
		group.GET("dev_list", api.DevList)                     // call
		group.POST("ban_dev", api.BanDev)                      // call
		group.POST("change_network", api.ChangeNetwork)        //暂没用到
		group.POST("connect", api.Connect)                     // call
		group.GET("get_conf", api.GetConfig)                   // call
		group.GET("cancel_account", api.CancelAccount)         //call
		group.POST("save_user_config", api.SaveUserConfig)     // call
		group.POST("get_user_config", api.GetUserConfig)       // call
		group.POST("delete_user_config", api.DeleteUserConfig) // call
		group.GET("get_traffic_list", api.TrafficList)
	}

}
