// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TGoods is the golang structure of table t_goods for DAO operations like Where/Data.
type TGoods struct {
	g.Meta           `orm:"table:t_goods, do:true"`
	Id               interface{} // 自增id
	MType            interface{} // 会员类型：1-vip1；2-vip2
	Title            interface{} // 套餐标题
	TitleEn          interface{} // 套餐标题（英文）
	TitleRus         interface{} // 套餐标题（俄文）
	Price            interface{} // 单价(U)
	PriceUnit        interface{} // 价格单位
	Period           interface{} // 有效期（天）
	DevLimit         interface{} // 设备限制数
	FlowLimit        interface{} // 流量限制数；单位：字节；0-不限制
	IsDiscount       interface{} // 是否优惠:1-是；2-否
	Low              interface{} // 最低赠送(天)
	High             interface{} // 最高赠送(天)
	Status           interface{} // 状态:1-正常；2-已软删
	CreatedAt        *gtime.Time // 创建时间
	UpdatedAt        *gtime.Time // 更新时间
	Author           interface{} // 作者
	Comment          interface{} // 备注信息
	UsdPayPrice      interface{} // usd_pay价格(U)
	UsdPriceUnit     interface{} // USD支付的价格单位
	WebmoneyPayPrice interface{} // webmoney价格(wmz)
	PriceRub         interface{} // 卢布价格(RUB)
	PriceWmz         interface{} // WMZ价格(WMZ)
	PriceUsd         interface{} // USD价格(USD)
	PriceBtc         interface{} // BTC价格(BTC)
	PriceUah         interface{} // UAH价格(UAH)
}
