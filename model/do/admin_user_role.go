// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminUserRole is the golang structure of table admin_user_role for DAO operations like Where/Data.
type AdminUserRole struct {
	g.Meta    `orm:"table:admin_user_role, do:true"`
	Id        interface{} // 自增序列
	Uid       interface{} // 用户id
	RoleId    interface{} // 角色id
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	IsDel     interface{} // 软删：0-未删；1-已删
	Author    interface{} // 更新人
}
