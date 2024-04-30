// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminUser is the golang structure for table admin_user.
type AdminUser struct {
	Id        int         `description:"用户id"`
	Uname     string      `description:"用户名"`
	Passwd    string      `description:"用户密码"`
	Nickname  string      `description:"昵称"`
	CreatedAt *gtime.Time `description:"创建时间"`
	UpdatedAt *gtime.Time `description:"更新时间"`
	Pwd2      string      `description:"二级密码"`
	Authkey   string      `description:"谷歌验证码私钥"`
	Status    int         `description:"冻结状态：0-正常；1-冻结"`
	IsDel     int         `description:"0-正常；1-软删"`
	IsReset   int         `description:"0-否；1-代表需要重置两步验证码"`
	IsFirst   int         `description:"0-否；1-代表首次登录需要修改密码"`
}
