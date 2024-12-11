// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TUserReportDay is the golang structure of table t_user_report_day for DAO operations like Where/Data.
type TUserReportDay struct {
	g.Meta        `orm:"table:t_user_report_day, do:true"`
	Id            interface{} // 自增id
	Date          interface{} // 数据日期, 20230101
	ChannelId     interface{} // 渠道id
	Total         interface{} // 用户总量
	New           interface{} // 新增用户
	Retained      interface{} // 留存
	MonthRetained interface{} // 月留存
	CreatedAt     *gtime.Time // 记录创建时间
}
