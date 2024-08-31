// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TV2RayUserTrafficLog is the golang structure of table t_v2ray_user_traffic_log for DAO operations like Where/Data.
type TV2RayUserTrafficLog struct {
	g.Meta    `orm:"table:t_v2ray_user_traffic_log, do:true"`
	Id        interface{} // 自增id
	Email     interface{} // 邮件
	Ip        interface{} // ip地址
	DateTime  interface{} // 数据采集时间
	Uplink    interface{} // 上行流量
	Downlink  interface{} // 下行流量
	CreatedAt *gtime.Time // 记录创建时间
}
