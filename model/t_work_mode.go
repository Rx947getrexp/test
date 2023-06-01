package model

import (
	"time"
)

type TWorkMode struct {
	Id        int64     `xorm:"pk autoincr comment('自增id') BIGINT"`
	UserId    int64     `xorm:"not null comment('用户id') BIGINT"`
	DevId     int64     `xorm:"not null comment('设备id') unique BIGINT"`
	ModeType  int       `xorm:"not null comment('模式类别:1-智能；2-手选') INT"`
	CreatedAt time.Time `xorm:"comment('创建时间') TIMESTAMP"`
	UpdatedAt time.Time `xorm:"comment('更新时间') TIMESTAMP"`
	Comment   string    `xorm:"comment('备注信息') VARCHAR(255)"`
}
