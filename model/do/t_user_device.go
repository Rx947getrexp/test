// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TUserDevice is the golang structure of table t_user_device for DAO operations like Where/Data.
type TUserDevice struct {
	g.Meta    `orm:"table:t_user_device, do:true"`
	Id        interface{} // 自增id
	UserId    interface{} // 用户uid
	ClientId  interface{} //
	Os        interface{} // 客户端设备系统os
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	Kicked    interface{} // 剔除状态, 0:正常，1:被剔除
}
