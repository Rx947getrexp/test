// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TDailyAdStatistics is the golang structure for table t_daily_ad_statistics.
type TDailyAdStatistics struct {
	Id        int64       `orm:"id"         description:"主键ID"`
	AdId      int         `orm:"ad_id"      description:"广告ID"`
	AdName    string      `orm:"ad_name"    description:"广告名称"`
	Date      int         `orm:"date"       description:"统计数据日期，整数类型，格式为 YYYYMMDD，例如20250102表示2025年1月2日"`
	Exposure  int         `orm:"exposure"   description:"广告的曝光量，默认值为0，表示当天的曝光次数"`
	Clicks    int         `orm:"clicks"     description:"广告的点击量，默认值为0，表示当天的点击次数"`
	Rewards   int         `orm:"rewards"    description:"广告完播后获赠时长的用户数，默认值为0，表示当天广告完播后获赠时长的用户数"`
	CreatedAt *gtime.Time `orm:"created_at" description:"记录创建时间，默认值为当前时间"`
	UpdatedAt *gtime.Time `orm:"updated_at" description:"记录更新时间，默认值为当前时间，并在每次更新时自动更新"`
}
