package model

import (
	"time"
)

type TDev struct {
	Id        int64     `xorm:"pk autoincr comment('自增id') BIGINT"`
	ClientId  string    `xorm:"comment('客户端自身设备ID') VARCHAR(64)"`
	Os        string    `xorm:"not null comment('客户端设备系统os') VARCHAR(64)"`
	IsSend    int       `xorm:"comment('1-已赠送时间；2-未赠送') INT"`
	Network   int       `xorm:"default 1 comment('网络模式（1-自动；2-手动）') INT"`
	CreatedAt time.Time `xorm:"comment('创建时间') TIMESTAMP"`
	UpdatedAt time.Time `xorm:"comment('更新时间') TIMESTAMP"`
	Comment   string    `xorm:"comment('备注信息') VARCHAR(255)"`
}
