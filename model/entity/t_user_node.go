// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TUserNode is the golang structure for table t_user_node.
type TUserNode struct {
	Id        uint64      `orm:"id"         description:"自增id"`
	Email     string      `orm:"email"      description:"用户邮箱"`
	Ip        string      `orm:"ip"         description:"节点IP"`
	V2RayUuid string      `orm:"v2ray_uuid" description:"uuid"`
	CreatedAt *gtime.Time `orm:"created_at" description:"创建时间"`
	UserId    uint64      `orm:"user_id"    description:"用户uid"`
}
