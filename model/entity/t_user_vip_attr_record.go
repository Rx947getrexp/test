// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TUserVipAttrRecord is the golang structure for table t_user_vip_attr_record.
type TUserVipAttrRecord struct {
	Id              int64       `orm:"id"                description:"自增id"`
	Email           string      `orm:"email"             description:"用户邮箱"`
	Source          string      `orm:"source"            description:"来源"`
	OrderNo         string      `orm:"order_no"          description:"订单号"`
	ExpiredTimeFrom int         `orm:"expired_time_from" description:"会员到期时间-原值"`
	ExpiredTimeTo   int         `orm:"expired_time_to"   description:"会员到期时间-新值"`
	Desc            string      `orm:"desc"              description:"记录描述"`
	CreatedAt       *gtime.Time `orm:"created_at"        description:"创建时间"`
	IsRevert        int         `orm:"is_revert"         description:"是否被回滚"`
}
