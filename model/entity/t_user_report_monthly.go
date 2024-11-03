// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TUserReportMonthly is the golang structure for table t_user_report_monthly.
type TUserReportMonthly struct {
	Id            uint64      `orm:"id"             description:""`
	StatMonth     uint        `orm:"stat_month"     description:"统计月份"`
	Os            string      `orm:"os"             description:"设备类型"`
	UserCount     uint        `orm:"user_count"     description:"用户总数"`
	NewUsers      uint        `orm:"new_users"      description:"新增用户量"`
	RetainedUsers uint        `orm:"retained_users" description:"次月留存"`
	CreatedAt     *gtime.Time `orm:"created_at"     description:"记录创建时间"`
}
