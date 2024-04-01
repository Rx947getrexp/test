package router

import (
	"github.com/gin-gonic/gin"
	"go-speed/api"
	"go-speed/api/admin"
)

func AdminRoute(group *gin.RouterGroup) {
	group.POST("login", admin.LoginAdmin)
	group.POST("upload", admin.Upload)
	group.POST("edit_member_expired_time", admin.EditMemberExpiredTime)
	group.GET("get_report_user_day_list", admin.GetReportUserDayList)
	group.GET("get_online_user_day_list", admin.GetOnlineUserDayList)

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
		memberGroup.POST("edit_member_expired_time", admin.EditMemberExpiredTime)

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

	}
}
