// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminUser is the golang structure of table admin_user for DAO operations like Where/Data.
type AdminUser struct {
	g.Meta    `orm:"table:admin_user, do:true"`
	Id        interface{} // 用户id
	Uname     interface{} // 用户名
	Passwd    interface{} // 用户密码
	Nickname  interface{} // 昵称
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	Pwd2      interface{} // 二级密码
	Authkey   interface{} // 谷歌验证码私钥
	Status    interface{} // 冻结状态：0-正常；1-冻结
	IsDel     interface{} // 0-正常；1-软删
	IsReset   interface{} // 0-否；1-代表需要重置两步验证码
	IsFirst   interface{} // 0-否；1-代表首次登录需要修改密码
}
