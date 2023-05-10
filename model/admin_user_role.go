package model

import (
	"time"
)

type AdminUserRole struct {
	Id        int64     `xorm:"pk autoincr comment('自增序列') BIGINT"`
	Uid       int       `xorm:"not null comment('用户id') unique INT"`
	RoleId    int       `xorm:"not null comment('角色id') INT"`
	CreatedAt time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	UpdatedAt time.Time `xorm:"comment('更新时间') TIMESTAMP"`
	IsDel     int       `xorm:"not null comment('软删：0-未删；1-已删') INT"`
	Author    string    `xorm:"comment('更新人') VARCHAR(255)"`
}
