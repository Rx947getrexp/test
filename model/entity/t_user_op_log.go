// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TUserOpLog is the golang structure for table t_user_op_log.
type TUserOpLog struct {
	Id           uint64      `orm:"id"            description:"自增id"`
	Email        string      `orm:"email"         description:"用户账号"`
	DeviceId     string      `orm:"device_id"     description:"设备ID"`
	DeviceType   string      `orm:"device_type"   description:"设备类型"`
	PageName     string      `orm:"page_name"     description:"page_name"`
	Result       string      `orm:"result"        description:"result"`
	Content      string      `orm:"content"       description:"content"`
	Version      string      `orm:"version"       description:""`
	CreateTime   string      `orm:"create_time"   description:"提交时间"`
	CreatedAt    *gtime.Time `orm:"created_at"    description:"记录创建时间"`
	InterfaceUrl string      `orm:"interface_url" description:"接口地址"`
	ServerCode   string      `orm:"server_code"   description:"后端状态码"`
	HttpCode     string      `orm:"http_code"     description:"HTTP状态码"`
	TraceId      string      `orm:"trace_id"      description:"TraceId"`
	UserId       uint64      `orm:"user_id"       description:"用户uid"`
	AppName      string      `orm:"app_name"      description:"app_name"`
}
