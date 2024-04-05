// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminRoleRes is the golang structure for table admin_role_res.
type AdminRoleRes struct {
	RoleId    int         `description:"角色id"`
	ResIds    string      `description:"资源id列表"`
	ResTree   string      `description:"资源菜单json树"`
	CreatedAt *gtime.Time `description:"创建时间"`
	UpdatedAt *gtime.Time `description:"更新时间"`
	Author    string      `description:"更新人"`
}
