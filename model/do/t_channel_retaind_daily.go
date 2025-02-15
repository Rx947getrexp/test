// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TChannelRetaindDaily is the golang structure of table t_channel_retaind_daily for DAO operations like Where/Data.
type TChannelRetaindDaily struct {
	g.Meta        `orm:"table:t_channel_retaind_daily, do:true"`
	Id            interface{} // 主键ID
	Date          interface{} // 统计数据日期，整数类型，格式为 YYYYMMDD，例如20250102表示2025年1月2日
	Channel       interface{} // 渠道id
	NewUsers      interface{} // 新增用户数量
	Day2Retained  interface{} // 新增用户次日留存数量
	Day7Retained  interface{} // 新增用户7日留存数量
	Day15Retained interface{} // 新增用户15日留存数量
	Day30Retained interface{} // 新增用户30日留存数量
	CreatedAt     *gtime.Time // 记录创建时间，默认值为当前时间
	UpdatedAt     *gtime.Time // 记录更新时间，默认值为当前时间，并在每次更新时自动更新
}
