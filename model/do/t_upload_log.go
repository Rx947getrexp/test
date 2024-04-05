// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TUploadLog is the golang structure of table t_upload_log for DAO operations like Where/Data.
type TUploadLog struct {
	g.Meta    `orm:"table:t_upload_log, do:true"`
	Id        interface{} // 自增id
	UserId    interface{} // 用户id
	DevId     interface{} // 设备id
	Content   interface{} // 日志内容
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	Comment   interface{} // 备注信息
}
