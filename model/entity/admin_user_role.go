// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminUserRole is the golang structure for table admin_user_role.
type AdminUserRole struct {
	Id        int64       `description:"自增序列"`
	Uid       int         `description:"用户id"`
	RoleId    int         `description:"角色id"`
	CreatedAt *gtime.Time `description:"创建时间"`
	UpdatedAt *gtime.Time `description:"更新时间"`
	IsDel     int         `description:"软删：0-未删；1-已删"`
	Author    string      `description:"更新人"`
}
