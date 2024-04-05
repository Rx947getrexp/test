// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminRole is the golang structure for table admin_role.
type AdminRole struct {
	Id        int         `description:"角色id"`
	Name      string      `description:"角色名称"`
	IsDel     int         `description:"0-正常；1-软删"`
	IsUsed    int         `description:"1-已启用；2-未启用"`
	Remark    string      `description:"备注"`
	CreatedAt *gtime.Time `description:"创建时间"`
	UpdatedAt *gtime.Time `description:"更新时间"`
}
