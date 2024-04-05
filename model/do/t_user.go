// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TUser is the golang structure of table t_user for DAO operations like Where/Data.
type TUser struct {
	g.Meta      `orm:"table:t_user, do:true"`
	Id          interface{} // 自增id
	Uname       interface{} // 用户名
	Passwd      interface{} // 用户密码
	Email       interface{} // 邮件
	Phone       interface{} // 电话
	Level       interface{} // 等级：0-vip0；1-vip1；2-vip2
	ExpiredTime interface{} // vip到期时间
	V2RayUuid   interface{} // 节点UUID
	V2RayTag    interface{} // v2ray存在UUID标签:1-有；2-无
	Channel     interface{} //
	ChannelId   interface{} // 渠道id
	Status      interface{} // 冻结状态：0-正常；1-冻结
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	Comment     interface{} // 备注信息
	ClientId    interface{} //
}
