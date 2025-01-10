// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TDailyAdStatistics is the golang structure of table t_daily_ad_statistics for DAO operations like Where/Data.
type TDailyAdStatistics struct {
	g.Meta    `orm:"table:t_daily_ad_statistics, do:true"`
	Id        interface{} // 主键ID
	AdId      interface{} // 广告ID
	AdName    interface{} // 广告名称
	Date      interface{} // 统计数据日期，整数类型，格式为 YYYYMMDD，例如20250102表示2025年1月2日
	Exposure  interface{} // 广告的曝光量，默认值为0，表示当天的曝光次数
	Clicks    interface{} // 广告的点击量，默认值为0，表示当天的点击次数
	Rewards   interface{} // 广告完播后获赠时长的用户数，默认值为0，表示当天广告完播后获赠时长的用户数
	CreatedAt *gtime.Time // 记录创建时间，默认值为当前时间
	UpdatedAt *gtime.Time // 记录更新时间，默认值为当前时间，并在每次更新时自动更新
}
