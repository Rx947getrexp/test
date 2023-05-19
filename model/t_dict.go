package model

import (
	"time"
)

type TDict struct {
	KeyId     string    `xorm:"not null pk comment('键') VARCHAR(255)"`
	Value     string    `xorm:"not null comment('值') TEXT"`
	Note      string    `xorm:"comment('描述') VARCHAR(255)"`
	IsDel     int       `xorm:"comment('0-正常；1-软删') INT"`
	CreatedAt time.Time `xorm:"comment('创建时间') TIMESTAMP"`
	UpdatedAt time.Time `xorm:"comment('更新时间') TIMESTAMP"`
}
