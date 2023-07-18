package model

import (
	"time"
)

type TAppDns struct {
	Id        int64     `xorm:"pk autoincr comment('自增id') BIGINT"`
	SiteType  int       `xorm:"not null comment('站点类型:1-appapi域名；2-offsite域名；3-onlineservice在线客服；4-管理后台...') INT"`
	Dns       string    `xorm:"comment('域名') VARCHAR(64)"`
	Ip        string    `xorm:"comment('ip地址') VARCHAR(64)"`
	Level     int       `xorm:"not null comment('线路级别:1,2,3...用于白名单机制') INT"`
	Status    int       `xorm:"not null comment('状态:1-正常；2-已软删') INT"`
	CreatedAt time.Time `xorm:"comment('创建时间') TIMESTAMP"`
	UpdatedAt time.Time `xorm:"comment('更新时间') TIMESTAMP"`
	Author    string    `xorm:"comment('作者') VARCHAR(255)"`
	Comment   string    `xorm:"comment('备注信息') VARCHAR(255)"`
}
