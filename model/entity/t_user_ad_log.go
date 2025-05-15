// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TUserAdLog is the golang structure for table t_user_ad_log.
type TUserAdLog struct {
	Id         uint64      `orm:"id"          description:"自增id"`
	UserId     uint64      `orm:"user_id"     description:"用户id"`
	AdLocation string      `orm:"ad_location" description:"广告位的位置"`
	AdName     string      `orm:"ad_name"     description:"广告名称"`
	DeviceType string      `orm:"device_type" description:"设备类型"`
	AppVersion string      `orm:"app_version" description:"APP版本"`
	ClientId   string      `orm:"client_id"   description:"设备号"`
	Type       string      `orm:"type"        description:""`
	Content    string      `orm:"content"     description:""`
	Result     string      `orm:"result"      description:""`
	ReportTime string      `orm:"report_time" description:"提交时间"`
	CreatedAt  *gtime.Time `orm:"created_at"  description:"记录创建时间"`
	AppName    string      `orm:"app_name"    description:"app_name"`
}
