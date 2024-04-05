// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TNodeDns is the golang structure for table t_node_dns.
type TNodeDns struct {
	Id        int64       `description:"自增id"`
	NodeId    int64       `description:"节点id"`
	Dns       string      `description:"域名"`
	Ip        string      `description:"ip地址"`
	Level     int         `description:"线路级别:1,2,3...用于白名单机制"`
	Status    int         `description:"状态:1-正常；2-已软删"`
	CreatedAt *gtime.Time `description:"创建时间"`
	UpdatedAt *gtime.Time `description:"更新时间"`
	Author    string      `description:"作者"`
	Comment   string      `description:"备注信息"`
	IsMachine int         `description:"是否为真实机器"`
}
