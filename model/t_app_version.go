package model

import (
	"time"
)

type TAppVersion struct {
	Id        int64     `xorm:"pk autoincr comment('自增id') BIGINT"`
	AppType   int       `xorm:"not null comment('1-ios;2-安卓；3-h5zip') INT"`
	Version   string    `xorm:"comment('版本号') VARCHAR(64)"`
	Link      string    `xorm:"comment('超链地址') VARCHAR(64)"`
	Status    int       `xorm:"not null comment('状态:1-正常；2-已软删') INT"`
	CreatedAt time.Time `xorm:"comment('创建时间') TIMESTAMP"`
	UpdatedAt time.Time `xorm:"comment('更新时间') TIMESTAMP"`
	Author    string    `xorm:"comment('作者') VARCHAR(255)"`
	Comment   string    `xorm:"comment('备注信息') VARCHAR(255)"`
}
