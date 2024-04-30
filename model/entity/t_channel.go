// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TChannel is the golang structure for table t_channel.
type TChannel struct {
	Id        int64       `description:"编号id"`
	Name      string      `description:"渠道名称"`
	Code      string      `description:"渠道编号"`
	Link      string      `description:"渠道链接"`
	Status    int         `description:"状态:1-正常；2-已软删"`
	CreatedAt *gtime.Time `description:"创建时间"`
	UpdatedAt *gtime.Time `description:"更新时间"`
	Author    string      `description:"作者"`
	Comment   string      `description:"备注"`
}
