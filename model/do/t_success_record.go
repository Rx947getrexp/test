// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TSuccessRecord is the golang structure of table t_success_record for DAO operations like Where/Data.
type TSuccessRecord struct {
	g.Meta     `orm:"table:t_success_record, do:true"`
	Id         interface{} // 自增id
	UserId     interface{} // 用户id
	OrderId    interface{} // 订单id
	StartTime  interface{} // 本次计费开始时间戳
	EndTime    interface{} // 本次计费结束时间戳
	SurplusSec interface{} // 剩余时长(s)
	TotalSec   interface{} // 订单总时长(s）
	GoodsDay   interface{} // 套餐天数
	SendDay    interface{} // 赠送天数
	PayType    interface{} // 订单状态:1-银行卡；2-支付宝；3-微信支付
	Status     interface{} // 1-using使用中；2-wait等待; 3-end已结束
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
	Comment    interface{} // 备注信息
}
