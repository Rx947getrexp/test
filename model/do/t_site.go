// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TSite is the golang structure of table t_site for DAO operations like Where/Data.
type TSite struct {
	g.Meta    `orm:"table:t_site, do:true"`
	Id        interface{} // 自增id
	Site      interface{} // 域名
	Ip        interface{} // ip
	Status    interface{} // 1-正常；2-软删
	Author    interface{} // 作者
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	Comment   interface{} // 备注信息
}
