// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TChannelRegistrationPayDaily is the golang structure of table t_channel_registration_pay_daily for DAO operations like Where/Data.
type TChannelRegistrationPayDaily struct {
	g.Meta              `orm:"table:t_channel_registration_pay_daily, do:true"`
	Id                  interface{} // 主键ID
	Date                interface{} // 统计数据日期，整数类型，格式为 YYYYMMDD，例如20250102表示2025年1月2日
	Channel             interface{} // 渠道id
	NewUsers            interface{} // 新增用户数量
	DailyActiveUsers    interface{} // 日活用户数量
	MonthlyActiveUsers  interface{} // 月活用户数量
	TotalRechargeUsers  interface{} // 充值用户数量
	TotalRechargeAmount interface{} // 付费金额数量
	CreatedAt           *gtime.Time // 记录创建时间，默认值为当前时间
	UpdatedAt           *gtime.Time // 记录更新时间，默认值为当前时间，并在每次更新时自动更新
}
