// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TUserPing is the golang structure of table t_user_ping for DAO operations like Where/Data.
type TUserPing struct {
	g.Meta    `orm:"table:t_user_ping, do:true"`
	Id        interface{} // 自增id
	Email     interface{} // 用户邮箱
	Host      interface{} // 节点host, ip or dns
	Code      interface{} // ping的结果
	Cost      interface{} // ping耗时
	Time      interface{} // 上报时间
	CreatedAt *gtime.Time // 记录创建时间
}
