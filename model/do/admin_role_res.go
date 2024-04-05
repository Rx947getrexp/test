// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminRoleRes is the golang structure of table admin_role_res for DAO operations like Where/Data.
type AdminRoleRes struct {
	g.Meta    `orm:"table:admin_role_res, do:true"`
	RoleId    interface{} // 角色id
	ResIds    interface{} // 资源id列表
	ResTree   interface{} // 资源菜单json树
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	Author    interface{} // 更新人
}
