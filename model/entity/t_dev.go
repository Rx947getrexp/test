// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TDev is the golang structure for table t_dev.
type TDev struct {
	Id        int64       `description:"自增id"`
	ClientId  string      `description:""`
	Os        string      `description:"客户端设备系统os"`
	IsSend    int         `description:"1-已赠送时间；2-未赠送"`
	Network   int         `description:"网络模式（1-自动；2-手动）"`
	CreatedAt *gtime.Time `description:"创建时间"`
	UpdatedAt *gtime.Time `description:"更新时间"`
	Comment   string      `description:"备注信息"`
}
