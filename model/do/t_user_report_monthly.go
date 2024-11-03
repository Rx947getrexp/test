// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TUserReportMonthly is the golang structure of table t_user_report_monthly for DAO operations like Where/Data.
type TUserReportMonthly struct {
	g.Meta        `orm:"table:t_user_report_monthly, do:true"`
	Id            interface{} //
	StatMonth     interface{} // 统计月份
	Os            interface{} // 设备类型
	UserCount     interface{} // 用户总数
	NewUsers      interface{} // 新增用户量
	RetainedUsers interface{} // 次月留存
	CreatedAt     *gtime.Time // 记录创建时间
}
