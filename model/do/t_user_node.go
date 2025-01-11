// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TUserNode is the golang structure of table t_user_node for DAO operations like Where/Data.
type TUserNode struct {
	g.Meta    `orm:"table:t_user_node, do:true"`
	Id        interface{} // 自增id
	Email     interface{} // 用户邮箱
	Ip        interface{} // 节点IP
	V2RayUuid interface{} // uuid
	CreatedAt *gtime.Time // 创建时间
	UserId    interface{} // 用户uid
}
