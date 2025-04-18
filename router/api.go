package router

import (
	apiPath "go-speed/api"
	"go-speed/api/admin/official_docs"
	"go-speed/api/api"
	"go-speed/api/api/ad"
	"go-speed/api/api/config"
	"go-speed/api/api/country"
	"go-speed/api/api/device"
	"go-speed/api/api/goods"
	"go-speed/api/api/node"
	"go-speed/api/api/order"
	"go-speed/api/api/report"
	"go-speed/api/api/user"

	"github.com/gin-gonic/gin"
)

func ApiRoute(group *gin.RouterGroup) {
	//group.Use(api.PrintParam())
	//{}
	internalGroup := group.Group("internal")
	internalGroup.Use(api.InternalAuth())
	{
		internalGroup.POST("describe_user_info", user.DescribeUserInfo)
		internalGroup.POST("describe_node_list", node.DescribeNodeList)
		internalGroup.POST("delete_cancelled_user", user.DeleteCancelledUser)
	}
	group.POST("ad_list_test", ad.ADList)
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
	group.GET("app_info", api.AppInfo)                           // call
	group.GET("pc_app_info", api.PCAppInfo)                      // call
	group.GET("app_filter", api.AppFilter)                       //策略审核
	group.GET("list_node_for_report", node.ListNodeForReport)    //获取节点ip列表，上报ping结果
	group.POST("report_node_ping_result", apiPath.APIDeprecated) // 前端上报日志，已经迁移到 collector服务
	group.POST("report_user_op_log", apiPath.APIDeprecated)      // 前端上报日志, 已经迁移到 collector服务
	group.POST("report_user_ad_log", report.ReportUserADLog)
	group.GET("get_rules", config.GetRules) // 获取ip和域名列表
	group.POST("pay_notify", order.PayNotify)
	group.POST("goods_list", goods.GoodsList)
	group.POST("payment_channel_list", order.PaymentChannelList)
	group.GET("promoter_channel_mapping", api.GetPromotionDnsMapping) //官网接口，获取后台配置的推广人员与渠道映射关系
	group.GET("promoter_shop_mapping", api.GetPromotionShopMapping)   //官网接口，下载页面的各个商店的推广链接
	//签名验证
	switchStateGroup := group.Group("switch")
	switchStateGroup.Use(api.Verify)
	{
		switchStateGroup.POST("machine_states_witching", api.ServerStateSwitching) //踢机器的接口
	}
	group.GET("pay_notify", order.PayNotify)
	group.POST("russpay_callback", order.RussPayCallback)
	group.POST("get_official_docs", official_docs.OfficialDocsList)
	group.POST("get_official_docs_by_id", official_docs.OfficialDocById)
	group.Use(api.JWTAuth())
	{
		group.POST("change_passwd", api.ChangePasswd)
		group.GET("user_info", api.UserInfo) // call
		group.GET("team_list", api.TeamList)
		group.GET("team_info", api.TeamInfo)                      //分享vpn送时长
		group.POST("receive_free", api.ReceiveFree)               //免费领取会员
		group.GET("receive_free_summary", api.ReceiveFreeSummary) //免费领取会员
		group.POST("upload_log", api.UploadLog)
		//group.POST("create_order", api.CreateOrder)
		group.GET("order_list", api.OrderList)
		group.GET("dev_list", api.DevList)                     // call
		group.POST("ban_dev", api.BanDev)                      // call
		group.POST("device_list", device.DeviceList)           // call
		group.POST("kick_device", device.KickDevice)           // call
		group.POST("change_network", api.ChangeNetwork)        //暂没用到
		group.POST("connect", api.Connect)                     // call
		group.GET("get_conf", api.GetConfig)                   // call
		group.GET("cancel_account", api.CancelAccount)         //call
		group.POST("save_user_config", api.SaveUserConfig)     // call
		group.POST("get_user_config", api.GetUserConfig)       // call
		group.POST("delete_user_config", api.DeleteUserConfig) // call
		group.GET("get_traffic_list", api.TrafficList)

		// 新版本接口
		group.GET("get_serving_country_list", country.ServingCountryList) // 查询在线的国家列表
		group.POST("set_default_country", user.SetPreferredCountry)       // 用户设置默认国家
		group.GET("get_server_config", config.GetServerConfig)            // 查询v2ray代理配置
		group.POST("connect_server", config.ConnectServer)                // 连接代理

		group.GET("get_server_config_without_rules", config.GetServerConfigWithoutRules) // 获取配置不带ip和域名池

		// 支付相关
		group.POST("create_order", order.CreateOrder)
		group.POST("upload_payment_proof", order.UploadPaymentProof)
		group.POST("confirm_order", order.ConfirmOrder)
		group.POST("cancel_order", order.CancelOrder)
		group.POST("order_list", order.GetOrderList)
		group.POST("query_order", order.QueryOrder)

		// 广告
		group.POST("ad_list", ad.ADList)
		group.POST("ad_completion_notify", ad.ADCompletionNotify)
	}
}
