// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TUserOnlineDay is the golang structure of table t_user_online_day for DAO operations like Where/Data.
type TUserOnlineDay struct {
	g.Meta         `orm:"table:t_user_online_day, do:true"`
	Id             interface{} // 自增id
	Date           interface{} // 数据日期, 20230101
	Email          interface{} // 邮件
	ChannelId      interface{} // 渠道id
	OnlineDuration interface{} // 在线时间长度
	Uplink         interface{} // 上行流量
	Downlink       interface{} // 下行流量
	CreatedAt      *gtime.Time // 记录创建时间
}
