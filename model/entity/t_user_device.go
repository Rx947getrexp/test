// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TUserDevice is the golang structure for table t_user_device.
type TUserDevice struct {
	Id        uint64      `orm:"id"         description:"自增id"`
	UserId    uint64      `orm:"user_id"    description:"用户uid"`
	ClientId  string      `orm:"client_id"  description:""`
	Os        string      `orm:"os"         description:"客户端设备系统os"`
	CreatedAt *gtime.Time `orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `orm:"updated_at" description:"更新时间"`
	Kicked    int         `orm:"kicked"     description:"剔除状态, 0:正常，1:被剔除"`
}
