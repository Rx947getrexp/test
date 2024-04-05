// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminRes is the golang structure of table admin_res for DAO operations like Where/Data.
type AdminRes struct {
	g.Meta    `orm:"table:admin_res, do:true"`
	Id        interface{} // 资源id
	Name      interface{} // 资源名称
	ResType   interface{} // 类型：1-菜单；2-接口；3-按钮
	Pid       interface{} // 上级id（没有默认为0）
	Url       interface{} // url地址
	Sort      interface{} // 排序
	Icon      interface{} // 图标
	IsDel     interface{} // 软删状态：0-未删（默认）；1-已删
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	Author    interface{} // 修改人
}
