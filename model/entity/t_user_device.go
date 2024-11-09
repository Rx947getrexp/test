// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TUserDevice is the golang structure for table t_user_device.
type TUserDevice struct {
	Id        uint64      `description:"自增id"`
	UserId    uint64      `description:"用户uid"`
	ClientId  string      `description:""`
	Os        string      `description:"客户端设备系统os"`
	CreatedAt *gtime.Time `description:"创建时间"`
	UpdatedAt *gtime.Time `description:"更新时间"`
}
