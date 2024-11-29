// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TUserRechargeReportDay is the golang structure for table t_user_recharge_report_day.
type TUserRechargeReportDay struct {
	Id        uint64      `orm:"id"         description:"自增id"`
	Date      uint        `orm:"date"       description:"数据日期, 20230101"`
	GoodsId   uint        `orm:"goods_id"   description:"商品套餐id"`
	Total     uint        `orm:"total"      description:"用户充值总量"`
	New       uint        `orm:"new"        description:"新增用户充值数量"`
	CreatedAt *gtime.Time `orm:"created_at" description:"记录创建时间"`
}
