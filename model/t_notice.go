package model

import (
	"time"
)

type TNotice struct {
	Id        int64     `xorm:"pk autoincr comment('自增id') BIGINT"`
	Title     string    `xorm:"not null comment('标题') VARCHAR(128)"`
	Content   string    `xorm:"not null comment('正文内容') TEXT"`
	Author    string    `xorm:"not null comment('作者') VARCHAR(128)"`
	CreatedAt time.Time `xorm:"comment('创建时间') TIMESTAMP"`
	UpdatedAt time.Time `xorm:"comment('更新时间') TIMESTAMP"`
	Comment   string    `xorm:"comment('备注信息') VARCHAR(255)"`
}
