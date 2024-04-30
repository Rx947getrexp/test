// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TGift is the golang structure of table t_gift for DAO operations like Where/Data.
type TGift struct {
	g.Meta    `orm:"table:t_gift, do:true"`
	Id        interface{} // 自增id
	UserId    interface{} // 用户id
	OpId      interface{} // 业务id
	OpUid     interface{} // 业务uid
	Title     interface{} // 赠送标题
	GiftSec   interface{} // 赠送时间（单位s）
	GType     interface{} // 赠送类别（1-注册；2-推荐；3-日常活动；4-充值）
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	Comment   interface{} // 备注信息
}
