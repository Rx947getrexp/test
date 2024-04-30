// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TActivity is the golang structure of table t_activity for DAO operations like Where/Data.
type TActivity struct {
	g.Meta    `orm:"table:t_activity, do:true"`
	Id        interface{} // 自增id
	UserId    interface{} // 用户id
	Status    interface{} // 状态:1-success；2-fail
	GiftSec   interface{} // 赠送时间（失败为0）
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	Comment   interface{} // 备注信息
}
