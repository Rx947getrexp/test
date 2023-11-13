package model

import (
	"time"
)

type TUserConfig struct {
	Id        int64     `xorm:"pk autoincr comment('自增id') BIGINT"`
	UserId    int64     `xorm:"not null comment('用户id') BIGINT"`
	NodeId    int64     `xorm:"not null comment('节点id') BIGINT"`
	Status    int       `xorm:"not null comment('状态:1-正常；2-删除') INT"`
	CreatedAt time.Time `xorm:"comment('创建时间') TIMESTAMP"`
	UpdatedAt time.Time `xorm:"comment('更新时间') TIMESTAMP"`
}
