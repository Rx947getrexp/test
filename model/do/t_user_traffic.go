// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TUserTraffic is the golang structure of table t_user_traffic for DAO operations like Where/Data.
type TUserTraffic struct {
	g.Meta    `orm:"table:t_user_traffic, do:true"`
	Id        interface{} // 自增id
	Email     interface{} // 邮件
	Ip        interface{} // ip地址
	Date      interface{} // 数据日期, 20230101
	Uplink    interface{} // 上行流量
	Downlink  interface{} // 下行流量
	CreatedAt *gtime.Time // 记录创建时间
	UpdatedAt *gtime.Time // 记录更新时间
}
