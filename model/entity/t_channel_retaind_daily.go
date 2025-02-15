// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TChannelRetaindDaily is the golang structure for table t_channel_retaind_daily.
type TChannelRetaindDaily struct {
	Id            int64       `orm:"id"              description:"主键ID"`
	Date          int         `orm:"date"            description:"统计数据日期，整数类型，格式为 YYYYMMDD，例如20250102表示2025年1月2日"`
	Channel       string      `orm:"channel"         description:"渠道id"`
	NewUsers      int         `orm:"new_users"       description:"新增用户数量"`
	Day2Retained  int         `orm:"day_2_retained"  description:"新增用户次日留存数量"`
	Day7Retained  int         `orm:"day_7_retained"  description:"新增用户7日留存数量"`
	Day15Retained int         `orm:"day_15_retained" description:"新增用户15日留存数量"`
	Day30Retained int         `orm:"day_30_retained" description:"新增用户30日留存数量"`
	CreatedAt     *gtime.Time `orm:"created_at"      description:"记录创建时间，默认值为当前时间"`
	UpdatedAt     *gtime.Time `orm:"updated_at"      description:"记录更新时间，默认值为当前时间，并在每次更新时自动更新"`
}
