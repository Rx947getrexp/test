package service

import (
	"fmt"
	"go-speed/global"
	"go-speed/model"
	"go-speed/model/request"
	"xorm.io/xorm"
)

func AdAdminList(param *request.AdListAdminRequest, user *model.AdminUser) *xorm.Session {
	session := global.Db.Table("t_ad as ad")
	session.Where("ad.status = 1")
	if param.Name != "" {
		session.Where("ad.name like ?", "%"+param.Name+"%")
	}
	if param.Tag != "" {
		session.Where("ad.tag like ?", "%"+param.Tag+"%")
	}
	if param.AdType > 0 {
		session.Where("ad.ad_type = ?", param.AdType)
	}
	return session
}

func NoticeAdminList(param *request.NoticeListAdminRequest, user *model.AdminUser) *xorm.Session {
	session := global.Db.Table("t_notice as n")
	session.Where("n.status = 1")
	if param.Title != "" {
		session.Where("n.title like ?", "%"+param.Title+"%")
	}
	if param.Tag != "" {
		session.Where("n.tag like ?", "%"+param.Tag+"%")
	}
	return session
}

func GoodsAdminList(param *request.GoodsListAdminRequest, user *model.AdminUser) *xorm.Session {
	session := global.Db.Table("t_goods as g")
	session.Where("g.status = 1")
	if param.Title != "" {
		session.Where("g.title like ?", "%"+param.Title+"%")
	}
	return session
}

func NodeAdminList(param *request.NodeListAdminRequest, user *model.AdminUser) *xorm.Session {
	session := global.Db.Table("t_node as n")
	session.Where("n.status = 1")
	if param.Id > 0 {
		session.Where("id = ?", param.Id)
	}
	if param.Name != "" {
		session.Where("n.name like ?", "%"+param.Name+"%")
	}
	if param.Title != "" {
		session.Where("n.title like ?", "%"+param.Title+"%")
	}
	if param.Country != "" {
		session.Where("n.country like ?", "%"+param.Country+"%")
	}
	if param.Ip != "" {
		session.Where("n.ip like ?", "%"+param.Ip+"%")
	}
	return session
}

func NodeUuidAdminList(param *request.NodeUuidListAdminRequest, user *model.AdminUser) *xorm.Session {
	session := global.Db.Table("t_node_uuid as nu")
	session.Join("LEFT", "t_user as u", "u.id = nu.user_id")
	session.Join("LEFT", "t_node as n", "n.id = nu.node_id")
	session.Where("nu.status = 1")
	if param.Id > 0 {
		session.Where("nu.id = ?", param.Id)
	}
	if param.NodeId > 0 {
		session.Where("nu.node_id = ?", param.NodeId)
	}
	if param.UName != "" {
		session.Where("u.uname = ?", param.UName)
	}
	if param.NodeName != "" {
		session.Where("n.name like ?", "%"+param.NodeName+"%")
	}
	if param.Uuid != "" {
		session.Where("nu.uuid = ?", param.Uuid)
	}
	return session
}

func OrderAdminList(param *request.OrderListAdminRequest, user *model.AdminUser) *xorm.Session {
	session := global.Db.Table("t_order as o")
	session.Join("LEFT", "t_user as u", "u.id = o.user_id")
	if param.Id > 0 {
		session.Where("o.id = ?", param.Id)
	}
	if param.UserId > 0 {
		session.Where("o.user_id = ?", param.UserId)
	}
	if param.Uname != "" {
		session.Where("u.uname = ?", param.Uname)
	}
	if param.StartTime != "" {
		session.Where("w.created_at >= ?", param.StartTime)
	}
	if param.EndTime != "" {
		session.Where("w.created_at < ?", param.EndTime)
	}
	return session
}

func SiteAdminList(param *request.SiteListAdminRequest, user *model.AdminUser) *xorm.Session {
	session := global.Db.Table("t_site as s")
	session.Where("s.status = 1")
	if param.Site != "" {
		session.Where("s.site = ?", param.Site)
	}
	if param.Ip != "" {
		session.Where("s.ip = ?", param.Ip)
	}
	return session
}

func AppDnsAdminList(param *request.AppDnsListAdminRequest, user *model.AdminUser) *xorm.Session {
	session := global.Db.Table("t_app_dns as d")
	session.Where("d.status = 1")
	if param.Dns != "" {
		session.Where("d.dns = ?", param.Dns)
	}
	if param.Ip != "" {
		session.Where("s.ip = ?", param.Ip)
	}
	if param.SiteType > 0 {
		session.Where("site_type = ?", param.SiteType)
	}
	if param.Level > 0 {
		session.Where("level = ?", param.Level)
	}
	return session
}

func NodeDnsAdminList(param *request.NodeDnsListAdminRequest, user *model.AdminUser) *xorm.Session {
	session := global.Db.Table("t_node_dns as d")
	session.Where("d.status = 1")
	if param.Dns != "" {
		session.Where("d.dns = ?", param.Dns)
	}
	if param.Ip != "" {
		session.Where("s.ip = ?", param.Ip)
	}
	if param.NodeId > 0 {
		session.Where("node_id = ?", param.NodeId)
	}
	if param.Level > 0 {
		session.Where("level = ?", param.Level)
	}
	return session
}

func GiftAdminList(param *request.GiftListAdminRequest, user *model.AdminUser) *xorm.Session {
	session := global.Db.Table("t_gift as g")
	session.Join("LEFT", "t_user as u", "u.id = g.user_id")
	if param.UserId > 0 {
		session.Where("g.user_id = ?", param.UserId)
	}
	if param.GType > 0 {
		session.Where("g.g_type = ?", param.GType)
	}
	if param.Uname != "" {
		session.Where("u.uname = ?", param.Uname)
	}
	return session
}

func ActivityAdminList(param *request.ActivityListAdminRequest, user *model.AdminUser) *xorm.Session {
	session := global.Db.Table("t_activity as a")
	session.Join("LEFT", "t_user as u", "u.id = a.user_id")
	if param.UserId > 0 {
		session.Where("a.user_id = ?", param.UserId)
	}
	if param.Uname != "" {
		session.Where("u.uname = ?", param.Uname)
	}
	return session
}

func SpeedLogsAdminList(param *request.SpeedLogsAdminRequest, user *model.AdminUser) *xorm.Session {
	session := global.Db.Table("t_work_log as w")
	session.Join("LEFT", "t_user as u", "u.id = w.user_id")
	session.Join("LEFT", "t_dev as dev", "dev.id = w.dev_id")
	session.Join("LEFT", "t_node as node", "node.id = w.node_id")
	if param.UserId > 0 {
		session.Where("w.user_id = ?", param.UserId)
	}
	if param.DevId > 0 {
		session.Where("w.dev_id = ?", param.DevId)
	}
	if param.Uname != "" {
		session.Where("u.uname = ?", param.Uname)
	}
	if param.StartTime != "" {
		session.Where("w.created_at >= ?", param.StartTime)
	}
	if param.EndTime != "" {
		session.Where("w.created_at < ?", param.EndTime)
	}
	return session
}

func DevLogsAdminList(param *request.DevLogsAdminRequest, user *model.AdminUser) *xorm.Session {
	session := global.Db.Table("t_upload_log as up")
	session.Join("LEFT", "t_user as u", "u.id = up.user_id")
	session.Join("LEFT", "t_dev as dev", "dev.id = up.dev_id")
	if param.UserId > 0 {
		session.Where("up.user_id = ?", param.UserId)
	}
	if param.DevId > 0 {
		session.Where("up.dev_id = ?", param.DevId)
	}
	if param.Uname != "" {
		session.Where("u.uname = ?", param.Uname)
	}
	if param.StartTime != "" {
		session.Where("up.created_at >= ?", param.StartTime)
	}
	if param.EndTime != "" {
		session.Where("up.created_at < ?", param.EndTime)
	}
	return session
}

func ChannelAdminList(param *request.ChannelListAdminRequest, user *model.AdminUser) *xorm.Session {
	session := global.Db.Table("t_channel as c")
	session.Where("c.status = 1")
	if param.Name != "" {
		session.Where("c.name = ?", param.Name)
	}
	return session
}

func MemberAdminList(param *request.MemberListAdminRequest, user *model.AdminUser) *xorm.Session {
	session := global.Db.Table("t_user_team as t")
	session.Join("LEFT", "t_user as u1", "u1.id = t.user_id")
	session.Join("LEFT", "t_user as u2", "u2.id = t.direct_id")
	if param.UserId > 0 {
		session.Where("t.user_id = ?", param.UserId)
	}
	if param.Uname != "" {
		session.Where("u1.uname = ?", param.Uname)
	}
	if param.DirectId > 0 {
		session.Where("t.direct_id = ?", param.DirectId)
	}
	if param.DirectName != "" {
		session.Where("u2.uname = ?", param.DirectName)
	}
	if param.TeamId > 0 {
		session.Where("t.direct_tree like ?", "%"+fmt.Sprint(param.TeamId)+"%")
	}
	return session
}

func MemberDevAdminList(param *request.MemberDevListAdminRequest, user *model.AdminUser) *xorm.Session {
	session := global.Db.Table("t_user_dev as ud")
	session.Join("LEFT", "t_user as u", "u.id = ud.user_id")
	session.Join("LEFT", "t_dev as d", "d.id = ud.dev_id")
	session.Where("ud.status = 1")
	if param.UserId > 0 {
		session.Where("ud.user_id = ?", param.UserId)
	}
	if param.Uname != "" {
		session.Where("u.uname = ?", param.Uname)
	}
	if param.DevId > 0 {
		session.Where("ud.dev_id = ?", param.DevId)
	}
	return session
}

func AppVersionAdminList(param *request.AppVersionListAdminRequest, user *model.AdminUser) *xorm.Session {
	session := global.Db.Table("t_app_version as v")
	session.Where("v.status = 1")
	if param.AppType > 0 {
		session.Where("v.app_type = ?", param.AppType)
	}
	return session
}

func IosAccountAdminList(param *request.IosAccountListAdminRequest, user *model.AdminUser) *xorm.Session {
	session := global.Db.Table("t_ios_account as a")
	session.Where("a.status = 1 ")
	if param.Name != "" {
		session.Where("c.name = ?", param.Name)
	}
	if param.Account != "" {
		session.Where("c.account = ?", param.Account)
	}
	if param.AccountType > 0 {
		session.Where("account_type = ?", param.AccountType)
	}
	return session
}
