// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TCountry is the golang structure for table t_country.
type TCountry struct {
	Id        uint64      `orm:"id"         description:"自增id"`
	Name      string      `orm:"name"       description:"国家名称"`
	NameCn    string      `orm:"name_cn"    description:"国家名称中文"`
	CreatedAt *gtime.Time `orm:"created_at" description:"记录创建时间"`
	UpdatedAt *gtime.Time `orm:"updated_at" description:"更新时间"`
}
