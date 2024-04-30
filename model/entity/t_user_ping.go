// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TUserPing is the golang structure for table t_user_ping.
type TUserPing struct {
	Id        uint64      `description:"自增id"`
	Email     string      `description:"用户邮箱"`
	Host      string      `description:"节点host, ip or dns"`
	Code      string      `description:"ping的结果"`
	Cost      string      `description:"ping耗时"`
	Time      string      `description:"上报时间"`
	CreatedAt *gtime.Time `description:"记录创建时间"`
}
