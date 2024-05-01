// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TPayOrder is the golang structure of table t_pay_order for DAO operations like Where/Data.
type TPayOrder struct {
	g.Meta             `orm:"table:t_pay_order, do:true"`
	Id                 interface{} // 自增id
	UserId             interface{} // 用户uid
	Email              interface{} // 用户邮箱
	OrderNo            interface{} // 订单号
	OrderAmount        interface{} // 交易金额
	Currency           interface{} // 交易币种
	PayTypeCode        interface{} // 支付类型编码
	Status             interface{} // 状态:1-正常；2-已软删
	ReturnStatus       interface{} // 支付平台返回的结果
	StatusMes          interface{} // 状态:1-正常；2-已软删
	OrderData          interface{} // 创建订单时支付平台返回的信息
	ResultStatus       interface{} // 查询结果，实际订单状态
	OrderRealityAmount interface{} // 实际交易金额
	CreatedAt          *gtime.Time // 创建时间
	UpdatedAt          *gtime.Time // 更新时间
	Version            interface{} // 数据版本号
}
