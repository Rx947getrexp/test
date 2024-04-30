// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TDict is the golang structure for table t_dict.
type TDict struct {
	KeyId     string      `description:"键"`
	Value     string      `description:"值"`
	Note      string      `description:"描述"`
	IsDel     int         `description:"0-正常；1-软删"`
	CreatedAt *gtime.Time `description:"创建时间"`
	UpdatedAt *gtime.Time `description:"更新时间"`
}
