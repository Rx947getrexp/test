package model

import (
	"time"
)

type UserLogs struct {
	Id        int64     `xorm:"pk autoincr comment('自增id') BIGINT"`
	UserId    int64     `xorm:"not null comment('用户id') unique(log_user_date) BIGINT"`
	Datestr   string    `xorm:"not null comment('日期') unique(log_user_date) VARCHAR(32)"`
	Ip        string    `xorm:"not null comment('IP地址') VARCHAR(32)"`
	UserAgent string    `xorm:"comment('请求头user-agent') VARCHAR(1000)"`
	CreatedAt time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	UpdatedAt time.Time `xorm:"comment('更新时间') TIMESTAMP"`
	Comment   string    `xorm:"comment('备注信息') VARCHAR(255)"`
}
