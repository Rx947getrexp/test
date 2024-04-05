// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TNodeUuid is the golang structure for table t_node_uuid.
type TNodeUuid struct {
	Id        int64       `description:"自增id"`
	UserId    int64       `description:"用户id"`
	NodeId    int64       `description:"节点id"`
	Email     string      `description:"节点邮箱，用于区分流量"`
	V2RayUuid string      `description:"节点UUID"`
	Server    string      `description:"公网域名"`
	Port      int         `description:"公网端口"`
	UsedFlow  int64       `description:"已使用流量（单位B）"`
	Status    int         `description:"状态:1-正常；2-已软删"`
	CreatedAt *gtime.Time `description:"创建时间"`
	UpdatedAt *gtime.Time `description:"更新时间"`
	Comment   string      `description:"备注信息"`
}
