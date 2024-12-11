// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TUserRechargeReportDay is the golang structure of table t_user_recharge_report_day for DAO operations like Where/Data.
type TUserRechargeReportDay struct {
	g.Meta    `orm:"table:t_user_recharge_report_day, do:true"`
	Id        interface{} // 自增id
	Date      interface{} // 数据日期, 20230101
	GoodsId   interface{} // 商品套餐id
	Total     interface{} // 用户充值总量
	New       interface{} // 新增用户充值数量
	CreatedAt *gtime.Time // 记录创建时间
}
