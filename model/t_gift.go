package model

import (
	"time"
)

type TGift struct {
	Id        int64     `xorm:"pk autoincr comment('自增id') BIGINT"`
	UserId    int64     `xorm:"not null comment('用户id') BIGINT"`
	OpId      string    `xorm:"comment('业务id') VARCHAR(128)"`
	OpUid     int64     `xorm:"comment('业务uid') BIGINT"`
	Title     string    `xorm:"not null comment('赠送标题') VARCHAR(128)"`
	GiftSec   int       `xorm:"not null comment('赠送时间（单位s）') INT"`
	GType     int       `xorm:"not null comment('赠送类别（1-注册；2-推荐；3-日常活动；4-充值）') INT"`
	CreatedAt time.Time `xorm:"comment('创建时间') TIMESTAMP"`
	UpdatedAt time.Time `xorm:"comment('更新时间') TIMESTAMP"`
	Comment   string    `xorm:"comment('备注信息') VARCHAR(255)"`
}
