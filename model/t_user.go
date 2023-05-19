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
	CreatedAt   time.Time `xorm:"comment('创建时间') TIMESTAMP"`
	UpdatedAt   time.Time `xorm:"comment('更新时间') TIMESTAMP"`
	Comment     string    `xorm:"comment('备注信息') VARCHAR(255)"`
}
