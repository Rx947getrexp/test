package model

import (
	"time"
)

type TPromotionChannels struct {
	Id              int64     `xorm:"pk autoincr comment('自增id') BIGINT"`
	PromoterName    string    `xorm:"not null comment('推广人姓名') VARCHAR(100)"`
	PromotionDomain string    `xorm:"not null comment('推广域名') VARCHAR(255)"`
	Channel         string    `xorm:"not null comment('推广域名对应渠道') VARCHAR(50)"`
	CreatedAt       time.Time `xorm:"comment('创建时间') TIMESTAMP"`
}
