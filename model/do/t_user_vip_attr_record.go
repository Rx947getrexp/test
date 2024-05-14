// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TUserVipAttrRecord is the golang structure of table t_user_vip_attr_record for DAO operations like Where/Data.
type TUserVipAttrRecord struct {
	g.Meta          `orm:"table:t_user_vip_attr_record, do:true"`
	Id              interface{} // 自增id
	Email           interface{} // 用户邮箱
	Source          interface{} // 来源
	OrderNo         interface{} // 订单号
	ExpiredTimeFrom interface{} // 会员到期时间-原值
	ExpiredTimeTo   interface{} // 会员到期时间-新值
	Desc            interface{} // 记录描述
	CreatedAt       *gtime.Time // 创建时间
}
