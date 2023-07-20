package model

import (
	"time"
)

type TSuccessRecord struct {
	Id         int64     `xorm:"pk autoincr comment('自增id') BIGINT"`
	UserId     int64     `xorm:"not null comment('用户id') BIGINT"`
	OrderId    int64     `xorm:"not null comment('订单id') BIGINT"`
	StartTime  int64     `xorm:"not null comment('本次计费开始时间戳') BIGINT"`
	EndTime    int64     `xorm:"not null comment('本次计费结束时间戳') BIGINT"`
	SurplusSec int64     `xorm:"not null comment('剩余时长(s)') BIGINT"`
	TotalSec   int64     `xorm:"comment('订单总时长(s）') BIGINT"`
	GoodsDay   int       `xorm:"comment('套餐天数') INT"`
	SendDay    int       `xorm:"comment('赠送天数') INT"`
	PayType    int       `xorm:"not null comment('订单状态:1-银行卡；2-支付宝；3-微信支付') INT"`
	Status     int       `xorm:"comment('1-using使用中；2-wait等待; 3-end已结束') INT"`
	CreatedAt  time.Time `xorm:"comment('创建时间') TIMESTAMP"`
	UpdatedAt  time.Time `xorm:"comment('更新时间') TIMESTAMP"`
	Comment    string    `xorm:"comment('备注信息') VARCHAR(255)"`
}
