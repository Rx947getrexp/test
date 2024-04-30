// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TUserDev is the golang structure of table t_user_dev for DAO operations like Where/Data.
type TUserDev struct {
	g.Meta    `orm:"table:t_user_dev, do:true"`
	Id        interface{} // 自增id
	UserId    interface{} // 用户id
	DevId     interface{} // 设备id
	Status    interface{} // 状态:1-正常；2-已踢
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	Comment   interface{} // 备注信息
}
