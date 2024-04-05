// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TUploadLog is the golang structure for table t_upload_log.
type TUploadLog struct {
	Id        int64       `description:"自增id"`
	UserId    int64       `description:"用户id"`
	DevId     int64       `description:"设备id"`
	Content   string      `description:"日志内容"`
	CreatedAt *gtime.Time `description:"创建时间"`
	UpdatedAt *gtime.Time `description:"更新时间"`
	Comment   string      `description:"备注信息"`
}
