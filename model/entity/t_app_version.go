// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TAppVersion is the golang structure for table t_app_version.
type TAppVersion struct {
	Id        int64       `description:"自增id"`
	AppType   int         `description:"1-ios;2-安卓；3-h5zip"`
	Version   string      `description:"版本号"`
	Link      string      `description:"超链地址"`
	Status    int         `description:"状态:1-正常；2-已软删"`
	CreatedAt *gtime.Time `description:"创建时间"`
	UpdatedAt *gtime.Time `description:"更新时间"`
	Author    string      `description:"作者"`
	Comment   string      `description:"备注信息"`
}
