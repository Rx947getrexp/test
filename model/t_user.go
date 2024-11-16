package model

import (
	"time"
)

type TUser struct {
	Id          int64     `xorm:"pk autoincr comment('自增id') BIGINT"`
	Uname       string    `xorm:"not null comment('用户名') unique VARCHAR(64)"`
	Passwd      string    `xorm:"not null comment('用户密码') VARCHAR(64)"`
	Email       string    `xorm:"comment('邮件') VARCHAR(64)"`
	Phone       string    `xorm:"comment('电话') VARCHAR(64)"`
	Level       int       `xorm:"comment('等级：0-vip0；1-vip1；2-vip2') INT"`
	ExpiredTime int64     `xorm:"comment('vip到期时间') BIGINT"`
	V2rayUuid   string    `xorm:"comment('节点UUID') VARCHAR(128)"`
	V2rayTag    int       `xorm:"comment('v2ray存在UUID标签:1-有；2-无') INT"`
	ChannelId   int       `xorm:"comment('渠道id') INT"`
	Channel     string    `xorm:"comment('渠道id') VARCHAR(128)"`
	Status      int       `xorm:"comment('冻结状态：0-正常；1-冻结') INT"`
	CreatedAt   time.Time `xorm:"comment('创建时间') TIMESTAMP"`
	UpdatedAt   time.Time `xorm:"comment('更新时间') TIMESTAMP"`
	Comment     string    `xorm:"comment('备注信息') VARCHAR(255)"`
	ClientId    string    `xorm:"comment('clientID') VARCHAR(128)"`
	Kicked      int       `xorm:"comment('kicked') tinyint(1)"`
}
