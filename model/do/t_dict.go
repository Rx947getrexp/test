// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TDict is the golang structure of table t_dict for DAO operations like Where/Data.
type TDict struct {
	g.Meta    `orm:"table:t_dict, do:true"`
	KeyId     interface{} // 键
	Value     interface{} // 值
	Note      interface{} // 描述
	IsDel     interface{} // 0-正常；1-软删
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
