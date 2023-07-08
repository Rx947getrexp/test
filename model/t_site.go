package model

import (
	"time"
)

type TSite struct {
	Id        int64     `xorm:"pk autoincr comment('自增id') BIGINT"`
	Site      string    `xorm:"not null comment('域名') VARCHAR(255)"`
	Ip        string    `xorm:"comment('ip') VARCHAR(255)"`
	Status    int       `xorm:"comment('1-正常；2-软删') INT"`
	Author    string    `xorm:"comment('作者') VARCHAR(255)"`
	CreatedAt time.Time `xorm:"comment('创建时间') TIMESTAMP"`
	UpdatedAt time.Time `xorm:"comment('更新时间') TIMESTAMP"`
	Comment   string    `xorm:"comment('备注信息') VARCHAR(255)"`
}
