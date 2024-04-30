// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TSite is the golang structure for table t_site.
type TSite struct {
	Id        int64       `description:"自增id"`
	Site      string      `description:"域名"`
	Ip        string      `description:"ip"`
	Status    int         `description:"1-正常；2-软删"`
	Author    string      `description:"作者"`
	CreatedAt *gtime.Time `description:"创建时间"`
	UpdatedAt *gtime.Time `description:"更新时间"`
	Comment   string      `description:"备注信息"`
}
