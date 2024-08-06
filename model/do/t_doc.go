// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TDoc is the golang structure of table t_doc for DAO operations like Where/Data.
type TDoc struct {
	g.Meta    `orm:"table:t_doc, do:true"`
	Id        interface{} // 自增id
	Type      interface{} //
	Name      interface{} //
	Desc      interface{} //
	Content   interface{} // content
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
