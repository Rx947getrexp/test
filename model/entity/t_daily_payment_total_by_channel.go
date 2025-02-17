// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TDailyPaymentTotalByChannel is the golang structure for table t_daily_payment_total_by_channel.
type TDailyPaymentTotalByChannel struct {
	Id        int64       `orm:"id"         description:"主键ID"`
	Date      int         `orm:"date"       description:"统计数据日期，整数类型，格式为 YYYYMMDD，例如20250102表示2025年1月2日"`
	Channel   string      `orm:"channel"    description:"支付渠道名称"`
	Amount    float64     `orm:"amount"     description:"支付金额统计"`
	CreatedAt *gtime.Time `orm:"created_at" description:"记录创建时间，默认值为当前时间"`
	UpdatedAt *gtime.Time `orm:"updated_at" description:"记录更新时间，默认值为当前时间，并在每次更新时自动更新"`
}
