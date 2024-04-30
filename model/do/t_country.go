// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TCountry is the golang structure of table t_country for DAO operations like Where/Data.
type TCountry struct {
	g.Meta    `orm:"table:t_country, do:true"`
	Id        interface{} // 自增id
	Name      interface{} // 国家名称
	NameCn    interface{} // 国家名称中文
	CreatedAt *gtime.Time // 记录创建时间
	UpdatedAt *gtime.Time // 更新时间
}
