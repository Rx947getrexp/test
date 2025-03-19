// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TDailyPaymentTotalByChannel is the golang structure of table t_daily_payment_total_by_channel for DAO operations like Where/Data.
type TDailyPaymentTotalByChannel struct {
	g.Meta    `orm:"table:t_daily_payment_total_by_channel, do:true"`
	Id        interface{} // 主键ID
	Date      interface{} // 统计数据日期，整数类型，格式为 YYYYMMDD，例如20250102表示2025年1月2日
	Channel   interface{} // 支付渠道名称
	Amount    interface{} // 支付金额统计
	CreatedAt *gtime.Time // 记录创建时间，默认值为当前时间
	UpdatedAt *gtime.Time // 记录更新时间，默认值为当前时间，并在每次更新时自动更新
}
