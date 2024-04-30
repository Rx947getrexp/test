// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TWorkLog is the golang structure for table t_work_log.
type TWorkLog struct {
	Id        int64       `description:"自增id"`
	UserId    int64       `description:"用户id"`
	DevId     int64       `description:"设备id"`
	ModeType  int         `description:"模式类别:1-智能；2-手选"`
	NodeId    int64       `description:"工作节点"`
	Flow      int64       `description:"使用流量（字节）"`
	CreatedAt *gtime.Time `description:"创建时间"`
	UpdatedAt *gtime.Time `description:"更新时间"`
	Comment   string      `description:"备注信息"`
}
