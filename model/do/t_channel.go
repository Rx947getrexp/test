// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TChannel is the golang structure of table t_channel for DAO operations like Where/Data.
type TChannel struct {
	g.Meta    `orm:"table:t_channel, do:true"`
	Id        interface{} // 编号id
	Name      interface{} // 渠道名称
	Code      interface{} // 渠道编号
	Link      interface{} // 渠道链接
	Status    interface{} // 状态:1-正常；2-已软删
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	Author    interface{} // 作者
	Comment   interface{} // 备注
}
