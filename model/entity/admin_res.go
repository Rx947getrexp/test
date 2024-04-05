// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminRes is the golang structure for table admin_res.
type AdminRes struct {
	Id        int         `description:"资源id"`
	Name      string      `description:"资源名称"`
	ResType   int         `description:"类型：1-菜单；2-接口；3-按钮"`
	Pid       int         `description:"上级id（没有默认为0）"`
	Url       string      `description:"url地址"`
	Sort      int         `description:"排序"`
	Icon      string      `description:"图标"`
	IsDel     int         `description:"软删状态：0-未删（默认）；1-已删"`
	CreatedAt *gtime.Time `description:"创建时间"`
	UpdatedAt *gtime.Time `description:"更新时间"`
	Author    string      `description:"修改人"`
}
