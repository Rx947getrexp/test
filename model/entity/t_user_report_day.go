// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TUserReportDay is the golang structure for table t_user_report_day.
type TUserReportDay struct {
	Id            uint64      `orm:"id"             description:"自增id"`
	Date          uint        `orm:"date"           description:"数据日期, 20230101"`
	ChannelId     uint        `orm:"channel_id"     description:"渠道id"`
	Total         uint        `orm:"total"          description:"用户总量"`
	New           uint        `orm:"new"            description:"新增用户"`
	Retained      uint        `orm:"retained"       description:"留存"`
	MonthRetained uint        `orm:"month_retained" description:"月留存"`
	CreatedAt     *gtime.Time `orm:"created_at"     description:"记录创建时间"`
}
