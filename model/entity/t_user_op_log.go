// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TUserOpLog is the golang structure for table t_user_op_log.
type TUserOpLog struct {
	Id           uint64      `description:"自增id"`
	Email        string      `description:"用户账号"`
	DeviceId     string      `description:"设备ID"`
	DeviceType   string      `description:"设备类型"`
	PageName     string      `description:"page_name"`
	Result       string      `description:"result"`
	Content      string      `description:"content"`
	Version      string      `description:""`
	CreateTime   string      `description:"提交时间"`
	CreatedAt    *gtime.Time `description:"记录创建时间"`
	InterfaceUrl string      `description:"接口地址"`
	ServerCode   string      `description:"后端状态码"`
	HttpCode     string      `description:"HTTP状态码"`
	TraceId      string      `description:"TraceId"`
	UserId       uint64      `description:"用户uid"`
}
