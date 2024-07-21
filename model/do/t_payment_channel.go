// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TPaymentChannel is the golang structure of table t_payment_channel for DAO operations like Where/Data.
type TPaymentChannel struct {
	g.Meta              `orm:"table:t_payment_channel, do:true"`
	Id                  interface{} // 自增id
	ChannelId           interface{} // 支付通道ID
	ChannelName         interface{} // 支付通道名称
	IsActive            interface{} // 支付通道是否可用，1表示可用,2表示不可用
	FreeTrialDays       interface{} // 赠送的免费时长（以天为单位）
	TimeoutDuration     interface{} // 订单未支付超时关闭时间（单位分钟）
	PaymentQrCode       interface{} // 支付收款码. eg: U支付收款码
	PaymentQrUrl        interface{} // 支付收款链接
	BankCardInfo        interface{} // 银行卡信息
	CustomerServiceInfo interface{} // 客服信息
	MerNo               interface{} // mer_no
	PayTypeCode         interface{} // pay_type_code
	Weight              interface{} // 权重，排序使用
	CreatedAt           *gtime.Time // 创建时间
	UpdatedAt           *gtime.Time // 更新时间
	UsdNetwork          interface{} // USD支付网络
	CurrencyType        interface{} // 支付渠道币种
	FreekassaCode       interface{} // freekassa支付通道
	CommissionRate      interface{} // 手续费比例
	Commission          interface{} // 手续费
}
