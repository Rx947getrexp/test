package model

import (
	"time"
)

type AdminRole struct {
	Id        int       `xorm:"not null pk autoincr comment('角色id') INT"`
	Name      string    `xorm:"not null comment('角色名称') VARCHAR(32)"`
	IsDel     int       `xorm:"not null default 0 comment('0-正常；1-软删') INT"`
	IsUsed    int       `xorm:"not null comment('1-已启用；2-未启用') INT"`
	Remark    string    `xorm:"comment('备注') VARCHAR(255)"`
	CreatedAt time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	UpdatedAt time.Time `xorm:"comment('更新时间') TIMESTAMP"`
}
