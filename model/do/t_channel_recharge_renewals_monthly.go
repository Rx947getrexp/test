// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TChannelRechargeRenewalsMonthly is the golang structure of table t_channel_recharge_renewals_monthly for DAO operations like Where/Data.
type TChannelRechargeRenewalsMonthly struct {
	g.Meta         `orm:"table:t_channel_recharge_renewals_monthly, do:true"`
	Id             interface{} // 主键ID
	Month          interface{} // 统计数据月份，整数类型，格式为 YYYYMM，例如202501
	Channel        interface{} // 渠道id
	RechargeUsers  interface{} // 付费用户数量
	RechargeAmount interface{} // 付费用户充值总金额
	Retained       interface{} // 充值用户次月留存数量
	RenewalsUsers  interface{} // 次月续费人数
	RenewalsAmount interface{} // 次月续费充值总金额
	CreatedAt      *gtime.Time // 记录创建时间，默认值为当前时间
	UpdatedAt      *gtime.Time // 记录更新时间，默认值为当前时间，并在每次更新时自动更新
}
