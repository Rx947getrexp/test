// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TUserTeam is the golang structure of table t_user_team for DAO operations like Where/Data.
type TUserTeam struct {
	g.Meta     `orm:"table:t_user_team, do:true"`
	Id         interface{} // 自增id
	UserId     interface{} // 用户id
	DirectId   interface{} // 上级id
	DirectTree interface{} // 上级列表
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
	Comment    interface{} // 备注信息
}
