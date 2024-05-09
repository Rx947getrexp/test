// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TPaymentChannels is the golang structure for table t_payment_channels.
type TPaymentChannels struct {
	Id            int64       `orm:"id"              description:"自增id"`
	Name          string      `orm:"name"            description:"支付通道名称"`
	IsActive      int         `orm:"is_active"       description:"支付通道是否可用，1表示可用,2表示不可用"`
	FreeTrialDays int         `orm:"free_trial_days" description:"赠送的免费时长（以天为单位）"`
	CreatedAt     *gtime.Time `orm:"created_at"      description:"创建时间"`
	UpdatedAt     *gtime.Time `orm:"updated_at"      description:"更新时间"`
}
