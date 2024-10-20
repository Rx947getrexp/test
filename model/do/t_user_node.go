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
	UserId    interface{} // 用户uid
	Email     interface{} // 用户邮箱
	Ip        interface{} // 节点IP
	V2RayUuid interface{} // uuid
	Status    interface{} // 状态：0-未写入节点配置；1-已经写入到节点配置
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
