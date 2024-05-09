// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TPayOrder is the golang structure for table t_pay_order.
type TPayOrder struct {
	Id                 int64       `orm:"id"                   description:"自增id"`
	UserId             uint64      `orm:"user_id"              description:"用户uid"`
	Email              string      `orm:"email"                description:"用户邮箱"`
	OrderNo            string      `orm:"order_no"             description:"订单号"`
	OrderAmount        string      `orm:"order_amount"         description:"交易金额"`
	Currency           string      `orm:"currency"             description:"交易币种"`
	PayTypeCode        string      `orm:"pay_type_code"        description:"支付类型编码"`
	Status             string      `orm:"status"               description:"状态:1-正常；2-已软删"`
	ReturnStatus       string      `orm:"return_status"        description:"支付平台返回的结果"`
	StatusMes          string      `orm:"status_mes"           description:"状态:1-正常；2-已软删"`
	OrderData          string      `orm:"order_data"           description:"创建订单时支付平台返回的信息"`
	ResultStatus       string      `orm:"result_status"        description:"查询结果，实际订单状态"`
	OrderRealityAmount string      `orm:"order_reality_amount" description:"实际交易金额"`
	CreatedAt          *gtime.Time `orm:"created_at"           description:"创建时间"`
	UpdatedAt          *gtime.Time `orm:"updated_at"           description:"更新时间"`
	Version            int         `orm:"version"              description:"数据版本号"`
	PaymentChannelName string      `orm:"payment_channel_name" description:"支付通道名称"`
}
