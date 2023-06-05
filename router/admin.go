package router

import (
	"github.com/gin-gonic/gin"
	"go-speed/api"
	"go-speed/api/admin"
)

func AdminRoute(group *gin.RouterGroup) {
	group.POST("login", admin.LoginAdmin)
	group.POST("upload", admin.Upload)

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
		orderGroup.GET("order_summary", admin.OrderSummary)

		//广告管理
		adGroup := group.Group("ad")
		adGroup.GET("ad_list", admin.AdList)
		adGroup.POST("add_ad", admin.AddAd)
		adGroup.POST("edit_ad", admin.EditAd)
		adGroup.GET("ad_summary", admin.AdSummary)

		//监控管理
		monitorGroup := group.Group("monitor")
		monitorGroup.GET("node_list", admin.NodeList)
		monitorGroup.POST("add_node", admin.AddNode)
		monitorGroup.POST("edit_node", admin.EditNode)

		//站点管理
		siteGroup := group.Group("site")
		siteGroup.GET("link_detail", admin.LinkDetail)
		siteGroup.POST("edit_link", admin.EditLink)
		siteGroup.GET("site_list", admin.SiteList)
		siteGroup.POST("add_site", admin.AddSite)
		siteGroup.POST("edit_site", admin.EditSite)
		siteGroup.POST("del_site", admin.DelSite)

		//奖励管理
		giveGroup := group.Group("give")
		giveGroup.GET("give_team", admin.GiveTeam)
		giveGroup.GET("give_activity", admin.GiveActivity)
		giveGroup.GET("give_combo", admin.GiveCombo)
		giveGroup.GET("give_summary", admin.GiveSummary)

		//平台报表
		plantGroup := group.Group("plant")
		plantGroup.GET("plant_day_summary", admin.PlantDaySummary)
		plantGroup.GET("plant_month_summary", admin.PlantMonthSummary)

	}
}
