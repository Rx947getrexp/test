package model

import (
	"time"
)

type TGoods struct {
	Id               int64     `xorm:"pk autoincr comment('自增id') BIGINT"`
	MType            int       `xorm:"not null comment('会员类型：1-vip1；2-vip2') INT"`
	Title            string    `xorm:"not null comment('套餐标题') VARCHAR(128)"`
	TitleEn          string    `xorm:"comment('套餐标题（英文）') VARCHAR(128)"`
	TitleRus         string    `xorm:"comment('套餐标题（俄文)') VARCHAR(128)"`
	Price            string    `xorm:"not null comment('单价(U)') DECIMAL(10,6)"`
	UsdPayPrice      string    `xorm:"not null comment('Usd支付单价(U)') DECIMAL(10,6)"`
	WebmoneyPayPrice string    `xorm:"not null comment('webmoney支付单价(kwm)') DECIMAL(10,6)"`
	PriceUnit        string    `xorm:"not null comment('价格的单位')"`
	UsdPriceUnit     string    `xorm:"not null comment('Usd价格的单位') DECIMAL(10,6)"`
	Period           int       `xorm:"not null comment('有效期（天）') INT"`
	DevLimit         int       `xorm:"not null comment('设备限制数') INT"`
	FlowLimit        int64     `xorm:"not null comment('流量限制数；单位：字节；0-不限制') BIGINT"`
	IsDiscount       int       `xorm:"comment('是否优惠:1-是；2-否') INT"`
	Low              int       `xorm:"comment('最低赠送(天)') INT"`
	High             int       `xorm:"comment('最高赠送(天)') INT"`
	Status           int       `xorm:"not null comment('状态:1-正常；2-已软删') INT"`
	CreatedAt        time.Time `xorm:"comment('创建时间') TIMESTAMP"`
	UpdatedAt        time.Time `xorm:"comment('更新时间') TIMESTAMP"`
	Author           string    `xorm:"comment('作者') VARCHAR(255)"`
	Comment          string    `xorm:"comment('备注信息') VARCHAR(255)"`
}
