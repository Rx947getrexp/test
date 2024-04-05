// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TGoods is the golang structure for table t_goods.
type TGoods struct {
	Id         int64       `description:"自增id"`
	MType      int         `description:"会员类型：1-vip1；2-vip2"`
	Title      string      `description:"套餐标题"`
	TitleEn    string      `description:"套餐标题（英文）"`
	TitleRus   string      `description:"套餐标题（俄文）"`
	Price      float64     `description:"单价(U)"`
	Period     int         `description:"有效期（天）"`
	DevLimit   int         `description:"设备限制数"`
	FlowLimit  int64       `description:"流量限制数；单位：字节；0-不限制"`
	IsDiscount int         `description:"是否优惠:1-是；2-否"`
	Low        int         `description:"最低赠送(天)"`
	High       int         `description:"最高赠送(天)"`
	Status     int         `description:"状态:1-正常；2-已软删"`
	CreatedAt  *gtime.Time `description:"创建时间"`
	UpdatedAt  *gtime.Time `description:"更新时间"`
	Author     string      `description:"作者"`
	Comment    string      `description:"备注信息"`
}
