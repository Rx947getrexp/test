package model

import (
	"time"
)

type TDev struct {
	Id        int64     `xorm:"pk autoincr comment('自增id') BIGINT"`
	Os        string    `xorm:"not null comment('客户端设备系统os') VARCHAR(64)"`
	CreatedAt time.Time `xorm:"comment('创建时间') TIMESTAMP"`
	UpdatedAt time.Time `xorm:"comment('更新时间') TIMESTAMP"`
	Comment   string    `xorm:"comment('备注信息') VARCHAR(255)"`
}
