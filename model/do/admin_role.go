// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminRole is the golang structure of table admin_role for DAO operations like Where/Data.
type AdminRole struct {
	g.Meta    `orm:"table:admin_role, do:true"`
	Id        interface{} // 角色id
	Name      interface{} // 角色名称
	IsDel     interface{} // 0-正常；1-软删
	IsUsed    interface{} // 1-已启用；2-未启用
	Remark    interface{} // 备注
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
