package model

import (
	"time"
)

type TOrder struct {
	Id         int64     `xorm:"pk autoincr comment('自增id') BIGINT"`
	UserId     int64     `xorm:"not null comment('用户id') BIGINT"`
	GoodsId    int64     `xorm:"not null comment('商品id') BIGINT"`
	Title      string    `xorm:"not null comment('商品标题') VARCHAR(128)"`
	Price      string    `xorm:"not null comment('单价(U)') DECIMAL(10,6)"`
	PriceCny   string    `xorm:"not null comment('折合RMB单价(CNY)') DECIMAL(10,2)"`
	Status     int       `xorm:"not null comment('订单状态:1-init；2-success；3-cancel') INT"`
	FinishedAt time.Time `xorm:"comment('完成时间') TIMESTAMP"`
	CreatedAt  time.Time `xorm:"comment('创建时间') TIMESTAMP"`
	UpdatedAt  time.Time `xorm:"comment('更新时间') TIMESTAMP"`
	Comment    string    `xorm:"comment('备注信息') VARCHAR(255)"`
}
