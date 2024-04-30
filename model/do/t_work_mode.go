// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TWorkMode is the golang structure of table t_work_mode for DAO operations like Where/Data.
type TWorkMode struct {
	g.Meta    `orm:"table:t_work_mode, do:true"`
	Id        interface{} // 自增id
	UserId    interface{} // 用户id
	DevId     interface{} // 设备id
	ModeType  interface{} // 模式类别:1-智能；2-手选
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	Comment   interface{} // 备注信息
}
