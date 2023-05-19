package model

import (
	"time"
)

type TActivity struct {
	Id        int64     `xorm:"pk autoincr comment('自增id') BIGINT"`
	UserId    int64     `xorm:"not null comment('用户id') BIGINT"`
	Status    int       `xorm:"not null comment('状态:1-success；2-fail') INT"`
	GiftSec   int       `xorm:"not null comment('赠送时间（失败为0）') INT"`
	CreatedAt time.Time `xorm:"comment('创建时间') TIMESTAMP"`
	UpdatedAt time.Time `xorm:"comment('更新时间') TIMESTAMP"`
	Comment   string    `xorm:"comment('备注信息') VARCHAR(255)"`
}
