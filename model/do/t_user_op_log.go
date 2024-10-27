// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TUserOpLog is the golang structure of table t_user_op_log for DAO operations like Where/Data.
type TUserOpLog struct {
	g.Meta       `orm:"table:t_user_op_log, do:true"`
	Id           interface{} // 自增id
	Email        interface{} // 用户账号
	DeviceId     interface{} // 设备ID
	DeviceType   interface{} // 设备类型
	PageName     interface{} // page_name
	Result       interface{} // result
	Content      interface{} // content
	Version      interface{} //
	CreateTime   interface{} // 提交时间
	CreatedAt    *gtime.Time // 记录创建时间
	InterfaceUrl interface{} // 接口地址
	ServerCode   interface{} // 后端状态码
	HttpCode     interface{} // HTTP状态码
	TraceId      interface{} // TraceId
	UserId       interface{} // 用户uid
}
