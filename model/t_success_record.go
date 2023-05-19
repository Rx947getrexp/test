package model

import (
	"time"
)

type TSuccessRecord struct {
	Id        int64     `xorm:"pk autoincr comment('自增id') BIGINT"`
	UserId    int64     `xorm:"not null comment('用户id') BIGINT"`
	OrderId   int64     `xorm:"not null comment('订单id') BIGINT"`
	StartTime int64     `xorm:"not null comment('套餐开始时间（时间戳）') BIGINT"`
	EndTime   int64     `xorm:"not null comment('套餐结束时间（时间戳）') BIGINT"`
	Title     string    `xorm:"not null comment('商品标题') VARCHAR(128)"`
	Price     string    `xorm:"not null comment('单价(U)') DECIMAL(10,6)"`
	PriceCny  string    `xorm:"not null comment('折合RMB单价(CNY)') DECIMAL(10,2)"`
	MType     int       `xorm:"not null comment('会员类型：1-vip1；2-vip2') INT"`
	PayType   int       `xorm:"not null comment('订单状态:1-银行卡；2-支付宝；3-微信支付') INT"`
	CreatedAt time.Time `xorm:"comment('创建时间') TIMESTAMP"`
	UpdatedAt time.Time `xorm:"comment('更新时间') TIMESTAMP"`
	Comment   string    `xorm:"comment('备注信息') VARCHAR(255)"`
}
