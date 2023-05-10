package model

import (
	"time"
)

type AdminRes struct {
	Id        int       `xorm:"not null pk autoincr comment('资源id') INT"`
	Name      string    `xorm:"not null comment('资源名称') VARCHAR(64)"`
	ResType   int       `xorm:"not null comment('类型：1-菜单；2-接口；3-按钮') INT"`
	Pid       int       `xorm:"not null comment('上级id（没有默认为0）') INT"`
	Url       string    `xorm:"comment('url地址') unique VARCHAR(255)"`
	Sort      int       `xorm:"not null comment('排序') INT"`
	Icon      string    `xorm:"comment('图标') VARCHAR(32)"`
	IsDel     int       `xorm:"not null comment('软删状态：0-未删（默认）；1-已删') INT"`
	CreatedAt time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	UpdatedAt time.Time `xorm:"comment('更新时间') TIMESTAMP"`
	Author    string    `xorm:"comment('修改人') VARCHAR(64)"`
}
