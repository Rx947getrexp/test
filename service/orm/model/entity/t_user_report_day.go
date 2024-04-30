// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TUserReportDay is the golang structure for table t_user_report_day.
type TUserReportDay struct {
	Id        uint64      `description:"自增id"`
	Date      uint        `description:"数据日期, 20230101"`
	ChannelId uint        `description:"渠道id"`
	Total     uint        `description:"用户总量"`
	New       uint        `description:"新增用户"`
	Retained  uint        `description:"留存"`
	CreatedAt *gtime.Time `description:"记录创建时间"`
}
