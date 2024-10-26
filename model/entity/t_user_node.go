// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TUserNode is the golang structure for table t_user_node.
type TUserNode struct {
	Id        int64       `description:"自增id"`
	UserId    uint64      `description:"用户uid"`
	Email     string      `description:"用户邮箱"`
	Ip        string      `description:"节点IP"`
	V2RayUuid string      `description:"uuid"`
	Status    int         `description:"状态：0-未写入节点配置；1-已经写入到节点配置"`
	CreatedAt *gtime.Time `description:"创建时间"`
	UpdatedAt *gtime.Time `description:"更新时间"`
}
