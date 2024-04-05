// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TUserOnlineDay is the golang structure for table t_user_online_day.
type TUserOnlineDay struct {
	Id             uint64      `description:"自增id"`
	Date           uint        `description:"数据日期, 20230101"`
	Email          string      `description:"邮件"`
	ChannelId      uint        `description:"渠道id"`
	OnlineDuration uint        `description:"在线时间长度"`
	Uplink         uint64      `description:"上行流量"`
	Downlink       uint64      `description:"下行流量"`
	CreatedAt      *gtime.Time `description:"记录创建时间"`
}
