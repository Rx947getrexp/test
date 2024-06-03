package model

import (
	"time"
)

type AdminUser struct {
	Id        int       `xorm:"not null pk autoincr comment('用户id') INT"`
	Uname     string    `xorm:"not null comment('用户名') unique VARCHAR(64)"`
	Passwd    string    `xorm:"not null comment('用户密码') VARCHAR(64)"`
	Nickname  string    `xorm:"comment('昵称') VARCHAR(64)"`
	CreatedAt time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	UpdatedAt time.Time `xorm:"comment('更新时间') TIMESTAMP"`
	Pwd2      string    `xorm:"comment('二级密码') VARCHAR(64)"`
	Authkey   string    `xorm:"comment('谷歌验证码私钥') VARCHAR(255)"`
	Status    int       `xorm:"default 0 comment('冻结状态：0-正常；1-冻结') INT"`
	IsDel     int       `xorm:"default 0 comment('0-正常；1-软删') INT"`
	IsReset   int       `xorm:"comment('0-否；1-代表需要重置两步验证码') INT"`
	IsFirst   int       `xorm:"comment('0-否；1-代表首次登录需要修改密码') INT"`
	Channel   string    `xorm:"comment('用户可查看的范围') VARCHAR(32)"`
}
