// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TPaymentChannels is the golang structure of table t_payment_channels for DAO operations like Where/Data.
type TPaymentChannels struct {
	g.Meta        `orm:"table:t_payment_channels, do:true"`
	Id            interface{} // 自增id
	Name          interface{} // 支付通道名称
	IsActive      interface{} // 支付通道是否可用，1表示可用,2表示不可用
	FreeTrialDays interface{} // 赠送的免费时长（以天为单位）
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
}
