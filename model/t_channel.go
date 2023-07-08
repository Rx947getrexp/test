package model

import (
	"time"
)

type TChannel struct {
	Id        int64     `xorm:"pk autoincr comment('编号id') BIGINT"`
	Name      string    `xorm:"not null comment('渠道名称') VARCHAR(255)"`
	Code      string    `xorm:"not null comment('渠道编号') VARCHAR(255)"`
	Link      string    `xorm:"comment('渠道链接') VARCHAR(255)"`
	Status    int       `xorm:"comment('状态:1-正常；2-已软删') INT"`
	CreatedAt time.Time `xorm:"comment('创建时间') TIMESTAMP"`
	UpdatedAt time.Time `xorm:"comment('更新时间') TIMESTAMP"`
	Author    string    `xorm:"comment('作者') VARCHAR(255)"`
	Comment   string    `xorm:"comment('备注') VARCHAR(255)"`
}
