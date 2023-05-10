package model

import (
	"time"
)

type AdminRoleRes struct {
	RoleId    int       `xorm:"not null pk comment('角色id') INT"`
	ResIds    string    `xorm:"not null comment('资源id列表') TEXT"`
	ResTree   string    `xorm:"comment('资源菜单json树') TEXT"`
	CreatedAt time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	UpdatedAt time.Time `xorm:"comment('更新时间') TIMESTAMP"`
	Author    string    `xorm:"comment('更新人') VARCHAR(255)"`
}
