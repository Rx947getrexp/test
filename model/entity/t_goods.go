// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TGoods is the golang structure for table t_goods.
type TGoods struct {
	Id          int64       `orm:"id"            description:"自增id"`
	MType       int         `orm:"m_type"        description:"会员类型：1-vip1；2-vip2"`
	Title       string      `orm:"title"         description:"套餐标题"`
	TitleEn     string      `orm:"title_en"      description:"套餐标题（英文）"`
	TitleRus    string      `orm:"title_rus"     description:"套餐标题（俄文）"`
	Price       float64     `orm:"price"         description:"单价(U)"`
	Period      int         `orm:"period"        description:"有效期（天）"`
	DevLimit    int         `orm:"dev_limit"     description:"设备限制数"`
	FlowLimit   int64       `orm:"flow_limit"    description:"流量限制数；单位：字节；0-不限制"`
	IsDiscount  int         `orm:"is_discount"   description:"是否优惠:1-是；2-否"`
	Low         int         `orm:"low"           description:"最低赠送(天)"`
	High        int         `orm:"high"          description:"最高赠送(天)"`
	Status      int         `orm:"status"        description:"状态:1-正常；2-已软删"`
	CreatedAt   *gtime.Time `orm:"created_at"    description:"创建时间"`
	UpdatedAt   *gtime.Time `orm:"updated_at"    description:"更新时间"`
	Author      string      `orm:"author"        description:"作者"`
	Comment     string      `orm:"comment"       description:"备注信息"`
	UsdPayPrice float64     `orm:"usd_pay_price" description:"usd_pay价格(U)"`
}
