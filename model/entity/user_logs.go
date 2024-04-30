// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserLogs is the golang structure for table user_logs.
type UserLogs struct {
	Id        int64       `description:"自增id"`
	UserId    int64       `description:"用户id"`
	Datestr   string      `description:"日期"`
	Ip        string      `description:"IP地址"`
	UserAgent string      `description:"请求头user-agent"`
	CreatedAt *gtime.Time `description:"创建时间"`
	UpdatedAt *gtime.Time `description:"更新时间"`
	Comment   string      `description:"备注信息"`
}
