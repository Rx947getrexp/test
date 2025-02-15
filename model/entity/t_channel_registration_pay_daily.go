// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TChannelRegistrationPayDaily is the golang structure for table t_channel_registration_pay_daily.
type TChannelRegistrationPayDaily struct {
	Id                  int64       `orm:"id"                    description:"主键ID"`
	Date                int         `orm:"date"                  description:"统计数据日期，整数类型，格式为 YYYYMMDD，例如20250102表示2025年1月2日"`
	Channel             string      `orm:"channel"               description:"渠道id"`
	NewUsers            int         `orm:"new_users"             description:"新增用户数量"`
	DailyActiveUsers    int         `orm:"daily_active_users"    description:"日活用户数量"`
	MonthlyActiveUsers  int         `orm:"monthly_active_users"  description:"月活用户数量"`
	TotalRechargeUsers  int         `orm:"total_recharge_users"  description:"充值用户数量"`
	TotalRechargeAmount string      `orm:"total_recharge_amount" description:"付费金额数量"`
	CreatedAt           *gtime.Time `orm:"created_at"            description:"记录创建时间，默认值为当前时间"`
	UpdatedAt           *gtime.Time `orm:"updated_at"            description:"记录更新时间，默认值为当前时间，并在每次更新时自动更新"`
}
