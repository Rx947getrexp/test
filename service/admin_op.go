package service

import (
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
	session.Where("ad.status = 1")
	if param.Title != "" {
		session.Where("g.title like ?", "%"+param.Title+"%")
	}
	return session
}

func NodeAdminList(param *request.NodeListAdminRequest, user *model.AdminUser) *xorm.Session {
	session := global.Db.Table("t_node as n")
	session.Where("n.status = 1")
	if param.Name != "" {
		session.Where("n.name like ?", "%"+param.Name+"%")
	}
	if param.Title != "" {
		session.Where("n.title like ?", "%"+param.Title+"%")
	}
	if param.Country != "" {
		session.Where("n.country like ?", "%"+param.Country+"%")
	}
	return session
}
