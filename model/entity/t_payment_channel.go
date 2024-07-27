// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TPaymentChannel is the golang structure for table t_payment_channel.
type TPaymentChannel struct {
	Id                  int64       `orm:"id"                    description:"自增id"`
	ChannelId           string      `orm:"channel_id"            description:"支付通道ID"`
	ChannelName         string      `orm:"channel_name"          description:"支付通道名称"`
	IsActive            int         `orm:"is_active"             description:"支付通道是否可用，1表示可用,2表示不可用"`
	FreeTrialDays       int         `orm:"free_trial_days"       description:"赠送的免费时长（以天为单位）"`
	TimeoutDuration     int         `orm:"timeout_duration"      description:"订单未支付超时关闭时间（单位分钟）"`
	PaymentQrCode       string      `orm:"payment_qr_code"       description:"支付收款码. eg: U支付收款码"`
	PaymentQrUrl        string      `orm:"payment_qr_url"        description:"支付收款链接"`
	BankCardInfo        string      `orm:"bank_card_info"        description:"银行卡信息"`
	CustomerServiceInfo string      `orm:"customer_service_info" description:"客服信息"`
	MerNo               string      `orm:"mer_no"                description:"mer_no"`
	PayTypeCode         string      `orm:"pay_type_code"         description:"pay_type_code"`
	Weight              int         `orm:"weight"                description:"权重，排序使用"`
	CreatedAt           *gtime.Time `orm:"created_at"            description:"创建时间"`
	UpdatedAt           *gtime.Time `orm:"updated_at"            description:"更新时间"`
	UsdNetwork          string      `orm:"usd_network"           description:"USD支付网络"`
	CurrencyType        string      `orm:"currency_type"         description:"支付渠道币种"`
	FreekassaCode       string      `orm:"freekassa_code"        description:"freekassa支付通道"`
	CommissionRate      float64     `orm:"commission_rate"       description:"手续费比例"`
	Commission          float64     `orm:"commission"            description:"手续费"`
	MinPayAmount        float64     `orm:"min_pay_amount"        description:"最低支付金额"`
}
