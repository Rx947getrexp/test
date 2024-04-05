// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserLogs is the golang structure of table user_logs for DAO operations like Where/Data.
type UserLogs struct {
	g.Meta    `orm:"table:user_logs, do:true"`
	Id        interface{} // 自增id
	UserId    interface{} // 用户id
	Datestr   interface{} // 日期
	Ip        interface{} // IP地址
	UserAgent interface{} // 请求头user-agent
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	Comment   interface{} // 备注信息
}
