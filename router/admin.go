package router

import (
	"go-speed/api"
	"go-speed/api/admin"
	"go-speed/api/admin/country"
	"go-speed/api/admin/node"
	"go-speed/api/admin/official_docs"
	"go-speed/api/admin/order"
	"go-speed/api/admin/payment_channel"
	"go-speed/api/admin/promotion"
	"go-speed/api/admin/report"
	"go-speed/api/admin/vip"

	"github.com/gin-gonic/gin"
)

func AdminRoute(group *gin.RouterGroup) {
	group.POST("login", admin.LoginAdmin)
	group.POST("upload", admin.Upload)
	group.POST("edit_member_expired_time", vip.EditMemberExpiredTime)
	group.GET("get_report_user_day_list", admin.GetReportUserDayList)
	group.GET("get_channel_user_day_list", admin.GetChannelUserDayList)
	group.GET("get_promotion_channel_user_day_list", admin.GetPromotionChannelUserDayList)
	group.GET("get_online_user_day_list", admin.GetOnlineUserDayList)
	group.GET("get_node_day_list", admin.GetNodeDayList)
	group.GET("get_node_online_user_day_list", admin.GetNodeOnlineUserDayList)
	group.GET("get_user_recharge_list", admin.GetUserRechargeList)
	group.GET("get_user_recharge_times_list", admin.GetUserRechargeTimesList)
	group.GET("get_recharge_click_list", admin.GetRechargeClickByDeviceList)
	group.GET("get_report_device_day_list", admin.GetReportDeviceDayList)
	group.GET("get_channel_user_recharge_list", admin.GetChannelUserRechargeList)
	group.GET("get_channel_user_recharge_day_list", admin.GetChannelUserRecharge)
	group.GET("get_channel_user_recharge_month_list", admin.GetChannelUserRechargeByMonth)
	group.GET("get_user_op_log_list", report.GetUserOpLogList)

	group.GET("reset_cache_test", admin.ResetCache) //刷新缓存

	nodeReportGroup := group.Group("node_report")
	nodeReportGroup.Use(admin.NodeReportAuth())
	{
		nodeReportGroup.GET("test", api.Test)
		nodeReportGroup.POST("heartbeat", admin.Heartbeat) //5s一个心跳包
		nodeReportGroup.POST("data", admin.Data)           //30s上报一次数据
	}

	group.Use(admin.JWTAuth())
	{
		//公共
		group.GET("full_menu_tree", admin.GetFullMenuTree)
		group.GET("role_menu_tree", admin.GetRoleTree)
		group.GET("user_info", admin.UserInfo)
		group.GET("reset_cache", admin.ResetCache) //刷新缓存

		//首页

		//系统管理
		sysGroup := group.Group("sys")
		sysGroup.GET("full_tree", admin.GetFullTree)
		sysGroup.POST("add_resource", admin.AddResource)
		sysGroup.POST("del_resource", admin.DelResource)
		sysGroup.POST("edit_resource", admin.EditResource)
		sysGroup.POST("edit_passwd", admin.EditPasswd)
		sysGroup.POST("generate_auth2Key", admin.GenerateAuth2Key) //生成两步验证器私钥
		sysGroup.POST("set_auth2Key", admin.SetAuth2Key)           //设置两步验证器

		//管理员模块
		adminGroup := group.Group("admin")
		adminGroup.GET("user_list", admin.GetAdminUserList)
		adminGroup.POST("add_user", admin.AddAdminUser)
		adminGroup.POST("edit_user", admin.EditAdminUser)
		adminGroup.GET("role_list", admin.GetAdminRoleList)
		adminGroup.POST("add_role", admin.AddAdminRole)
		adminGroup.POST("edit_role", admin.EditAdminRole)

		//会员管理
		memberGroup := group.Group("member")
		memberGroup.GET("member_list", admin.MemberList)
		memberGroup.GET("member_dev_list", admin.MemberDevList)
		memberGroup.POST("edit_member", admin.EditMember)
		memberGroup.POST("edit_member_dev", admin.EditMemberDev)
		memberGroup.POST("edit_member_expired_time", vip.EditMemberExpiredTime)

		//套餐管理
		comboGroup := group.Group("combo")
		comboGroup.GET("combo_list", admin.ComboList)
		comboGroup.POST("add_combo", admin.AddCombo)
		comboGroup.POST("edit_combo", admin.EditCombo)

		//通知消息
		noticeGroup := group.Group("notice")
		noticeGroup.GET("notice_list", admin.NoticeList)
		noticeGroup.POST("add_notice", admin.AddNotice)
		noticeGroup.POST("edit_notice", admin.EditNotice)

		//订单管理
		orderGroup := group.Group("order")
		orderGroup.GET("order_list", admin.OrderList)
		orderGroup.POST("pay_order_list", order.PayOrderList)
		orderGroup.POST("sync_order_status", order.SyncOrderStatus)
		orderGroup.POST("confirm_order_status", order.ConfirmOrderStatus)

		//广告管理
		adGroup := group.Group("ad")
		adGroup.GET("ad_list", admin.AdList)
		adGroup.POST("add_ad", admin.AddAd)
		adGroup.POST("edit_ad", admin.EditAd)

		//监控管理
		monitorGroup := group.Group("monitor")
		monitorGroup.GET("node_list", admin.NodeList)
		monitorGroup.POST("add_node", admin.AddNode)
		monitorGroup.POST("edit_node", admin.EditNode)
		monitorGroup.GET("node_uuid_list", admin.NodeUuidList)

		//站点管理
		siteGroup := group.Group("site")
		siteGroup.GET("app_info", admin.AppInfo)
		siteGroup.POST("edit_app_info", admin.EditAppInfo)
		siteGroup.GET("site_list", admin.SiteList)
		siteGroup.POST("add_site", admin.AddSite)
		siteGroup.POST("edit_site", admin.EditSite)

		//奖励管理
		giveGroup := group.Group("give")
		giveGroup.GET("gift_list", admin.GiftList)
		giveGroup.GET("activity_list", admin.ActivityList) // 免费领会员

		//日志信息
		logGroup := group.Group("log")
		logGroup.GET("speed_logs", admin.SpeedLogs) //加速日志
		logGroup.GET("dev_logs", admin.DevLogs)     // 设备日志

		//渠道管理
		channelGroup := group.Group("channel")
		channelGroup.GET("channel_list", admin.ChannelList)
		channelGroup.POST("add_channel", admin.AddChannel)
		channelGroup.POST("edit_channel", admin.EditChannel)

		//APP版本管理
		versionGroup := group.Group("app_version")
		versionGroup.GET("version_list", admin.AppVersionList)
		versionGroup.POST("add_version", admin.AddAppVersion)
		versionGroup.POST("edit_version", admin.EditAppVersion)

		//域名管理
		dnsGroup := group.Group("dns")
		dnsGroup.GET("app_dns_list", admin.AppDnsList)
		dnsGroup.POST("add_app_dns", admin.AddAppDns)
		dnsGroup.POST("edit_app_dns", admin.EditAppDns)
		dnsGroup.GET("node_dns_list", admin.NodeDnsList)
		dnsGroup.POST("add_node_dns", admin.AddNodeDns)
		dnsGroup.POST("edit_node_dns", admin.EditNodeDns)

		//IOS海外账号管理
		iosAccountGroup := group.Group("ios_account")
		iosAccountGroup.GET("ios_account_list", admin.IosAccountList)
		iosAccountGroup.POST("add_ios_account", admin.AddIosAccount)
		iosAccountGroup.POST("edit_ios_account", admin.EditIosAccount)

		//平台报表
		plantGroup := group.Group("plant")
		giveGroup.GET("give_summary", admin.GiveSummary)
		adGroup.GET("ad_summary", admin.AdSummary)
		plantGroup.GET("plant_day_summary", admin.PlantDaySummary)
		plantGroup.GET("plant_month_summary", admin.PlantMonthSummary)
		orderGroup.GET("order_summary", admin.OrderSummary)

		/************************************** 新版本接口 ******************************************************/
		//国家标识图片管理
		// 国家管理
		countryGroup := group.Group("country")
		countryGroup.GET("list", country.CountryList) // 国家名称列表

		servingCountryGroup := group.Group("serving_country")
		servingCountryGroup.GET("list", country.ServingCountryList)    // 查询在线的国家列表
		servingCountryGroup.POST("add", country.ServingCountryCreate)  // 添加国家
		servingCountryGroup.POST("edit", country.ServingCountryModify) // 修改国家信息

		machineGroup := group.Group("machine")
		machineGroup.GET("list", node.NodeList)    // 查询机器列表
		machineGroup.POST("add", node.NodeCreate)  // 添加机器
		machineGroup.POST("edit", node.NodeModify) // 修改机器信息
		//支付通道管理

		paymentChannelGroup := group.Group("payment_channel")
		paymentChannelGroup.POST("list", payment_channel.PaymentChannelList)
		paymentChannelGroup.POST("edit", payment_channel.PaymentChannelModify)
		paymentChannelGroup.POST("upload", payment_channel.UploadPaymentQRCode)

		// 官方文档
		officialDocsGroup := group.Group("official_docs")
		officialDocsGroup.POST("add", official_docs.OfficialDocsCreate)
		officialDocsGroup.POST("delete", official_docs.OfficialDocsDelete)
		officialDocsGroup.POST("edit", official_docs.OfficialDocsModify)
		officialDocsGroup.POST("list", official_docs.OfficialDocsList)
		officialDocsGroup.POST("upload", official_docs.UploadOfficialDocsImage)
		//推广渠道管理
		promotionGroup := group.Group("promotion")
		promotionGroup.GET("promotion_list", promotion.PromotionList)
		promotionGroup.POST("add_promotion", promotion.AddPromotion)
		promotionGroup.POST("edit_promotion", promotion.EditPromotion)
		promotionGroup.POST("delete_promotion", promotion.DeletePromotion)
	}
}
