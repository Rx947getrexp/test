package model

import (
	"time"
)

type TUserDev struct {
	Id        int64     `xorm:"pk autoincr comment('自增id') BIGINT"`
	UserId    int64     `xorm:"not null comment('用户id') BIGINT"`
	DevId     int64     `xorm:"not null comment('设备id') BIGINT"`
	Status    int       `xorm:"not null comment('状态:1-正常；2-已踢') INT"`
	CreatedAt time.Time `xorm:"comment('创建时间') TIMESTAMP"`
	UpdatedAt time.Time `xorm:"comment('更新时间') TIMESTAMP"`
	Comment   string    `xorm:"comment('备注信息') VARCHAR(255)"`
}
