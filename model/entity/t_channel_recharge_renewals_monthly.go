// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TChannelRechargeRenewalsMonthly is the golang structure for table t_channel_recharge_renewals_monthly.
type TChannelRechargeRenewalsMonthly struct {
	Id             int64       `orm:"id"              description:"主键ID"`
	Month          int         `orm:"month"           description:"统计数据月份，整数类型，格式为 YYYYMM，例如202501"`
	Channel        string      `orm:"channel"         description:"渠道id"`
	RechargeUsers  int         `orm:"recharge_users"  description:"付费用户数量"`
	RechargeAmount string      `orm:"recharge_amount" description:"付费用户充值总金额"`
	Retained       int         `orm:"retained"        description:"充值用户次月留存数量"`
	RenewalsUsers  int         `orm:"renewals_users"  description:"次月续费人数"`
	RenewalsAmount string      `orm:"renewals_amount" description:"次月续费充值总金额"`
	CreatedAt      *gtime.Time `orm:"created_at"      description:"记录创建时间，默认值为当前时间"`
	UpdatedAt      *gtime.Time `orm:"updated_at"      description:"记录更新时间，默认值为当前时间，并在每次更新时自动更新"`
}
