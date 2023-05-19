package model

import (
	"time"
)

type TUserTeam struct {
	Id         int64     `xorm:"pk autoincr comment('自增id') BIGINT"`
	UserId     int64     `xorm:"not null comment('用户id') BIGINT"`
	DirectId   int64     `xorm:"not null comment('上级id') BIGINT"`
	DirectTree string    `xorm:"not null comment('上级列表') TEXT"`
	CreatedAt  time.Time `xorm:"comment('创建时间') TIMESTAMP"`
	UpdatedAt  time.Time `xorm:"comment('更新时间') TIMESTAMP"`
	Comment    string    `xorm:"comment('备注信息') VARCHAR(255)"`
}
