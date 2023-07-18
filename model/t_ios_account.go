package model

import (
	"time"
)

type TIosAccount struct {
	Id          int64     `xorm:"pk autoincr comment('编号id') BIGINT"`
	Account     string    `xorm:"not null comment('ios账号') unique VARCHAR(64)"`
	Pass        string    `xorm:"not null comment('密码') VARCHAR(64)"`
	Name        string    `xorm:"comment('别名') VARCHAR(64)"`
	AccountType int       `xorm:"comment('1-国区；2-海外') INT"`
	Status      int       `xorm:"comment('1-正常；2-下架') INT"`
	CreatedAt   time.Time `xorm:"comment('创建时间') TIMESTAMP"`
	UpdatedAt   time.Time `xorm:"comment('更新时间') TIMESTAMP"`
	Author      string    `xorm:"comment('作者') VARCHAR(64)"`
	Comment     string    `xorm:"comment('备注') VARCHAR(255)"`
}
