// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TNodeUuid is the golang structure of table t_node_uuid for DAO operations like Where/Data.
type TNodeUuid struct {
	g.Meta    `orm:"table:t_node_uuid, do:true"`
	Id        interface{} // 自增id
	UserId    interface{} // 用户id
	NodeId    interface{} // 节点id
	Email     interface{} // 节点邮箱，用于区分流量
	V2RayUuid interface{} // 节点UUID
	Server    interface{} // 公网域名
	Port      interface{} // 公网端口
	UsedFlow  interface{} // 已使用流量（单位B）
	Status    interface{} // 状态:1-正常；2-已软删
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	Comment   interface{} // 备注信息
}
