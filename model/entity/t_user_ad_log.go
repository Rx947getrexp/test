// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TUserAdLog is the golang structure for table t_user_ad_log.
type TUserAdLog struct {
	Id         uint64      `description:"自增id"`
	UserId     uint64      `description:"用户id"`
	AdLocation string      `description:"广告位的位置"`
	AdName     string      `description:"广告名称"`
	DeviceType string      `description:"设备类型"`
	AppVersion string      `description:"APP版本"`
	ClientId   string      `description:"设备号"`
	Type       string      `description:""`
	Content    string      `description:""`
	Result     string      `description:""`
	ReportTime string      `description:"提交时间"`
	CreatedAt  *gtime.Time `description:"记录创建时间"`
}
