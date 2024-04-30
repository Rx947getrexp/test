// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TSuccessRecord is the golang structure for table t_success_record.
type TSuccessRecord struct {
	Id         int64       `description:"自增id"`
	UserId     int64       `description:"用户id"`
	OrderId    int64       `description:"订单id"`
	StartTime  int64       `description:"本次计费开始时间戳"`
	EndTime    int64       `description:"本次计费结束时间戳"`
	SurplusSec int64       `description:"剩余时长(s)"`
	TotalSec   int64       `description:"订单总时长(s）"`
	GoodsDay   int         `description:"套餐天数"`
	SendDay    int         `description:"赠送天数"`
	PayType    int         `description:"订单状态:1-银行卡；2-支付宝；3-微信支付"`
	Status     int         `description:"1-using使用中；2-wait等待; 3-end已结束"`
	CreatedAt  *gtime.Time `description:"创建时间"`
	UpdatedAt  *gtime.Time `description:"更新时间"`
	Comment    string      `description:"备注信息"`
}
