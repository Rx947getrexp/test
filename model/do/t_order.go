// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TOrder is the golang structure of table t_order for DAO operations like Where/Data.
type TOrder struct {
	g.Meta     `orm:"table:t_order, do:true"`
	Id         interface{} // 自增id
	UserId     interface{} // 用户id
	GoodsId    interface{} // 商品id
	Title      interface{} // 商品标题
	Price      interface{} // 单价(U)
	PriceCny   interface{} // 折合RMB单价(CNY)
	Status     interface{} // 订单状态:1-init；2-success；3-cancel
	FinishedAt *gtime.Time // 完成时间
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
	Comment    interface{} // 备注信息
}
