// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TUserOnlineDay is the golang structure for table t_user_online_day.
type TUserOnlineDay struct {
	Id               uint64      `orm:"id"                 description:"自增id"`
	Date             uint        `orm:"date"               description:"数据日期, 20230101"`
	Email            string      `orm:"email"              description:"邮件"`
	ChannelId        uint        `orm:"channel_id"         description:"渠道id"`
	OnlineDuration   uint        `orm:"online_duration"    description:"在线时间长度"`
	Uplink           uint64      `orm:"uplink"             description:"上行流量"`
	Downlink         uint64      `orm:"downlink"           description:"下行流量"`
	CreatedAt        *gtime.Time `orm:"created_at"         description:"记录创建时间"`
	LastLoginCountry string      `orm:"last_login_country" description:"最后登陆的国家"`
}
