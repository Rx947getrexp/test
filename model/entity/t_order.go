// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TOrder is the golang structure for table t_order.
type TOrder struct {
	Id         int64       `description:"自增id"`
	UserId     int64       `description:"用户id"`
	GoodsId    int64       `description:"商品id"`
	Title      string      `description:"商品标题"`
	Price      float64     `description:"单价(U)"`
	PriceCny   float64     `description:"折合RMB单价(CNY)"`
	Status     int         `description:"订单状态:1-init；2-success；3-cancel"`
	FinishedAt *gtime.Time `description:"完成时间"`
	CreatedAt  *gtime.Time `description:"创建时间"`
	UpdatedAt  *gtime.Time `description:"更新时间"`
	Comment    string      `description:"备注信息"`
}
