// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TDoc is the golang structure for table t_doc.
type TDoc struct {
	Id        uint64      `orm:"id"         description:"自增id"`
	Type      string      `orm:"type"       description:""`
	Name      string      `orm:"name"       description:""`
	Desc      string      `orm:"desc"       description:""`
	Content   string      `orm:"content"    description:"content"`
	CreatedAt *gtime.Time `orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `orm:"updated_at" description:"更新时间"`
}
