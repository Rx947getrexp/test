// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TActivity is the golang structure for table t_activity.
type TActivity struct {
	Id        int64       `description:"自增id"`
	UserId    int64       `description:"用户id"`
	Status    int         `description:"状态:1-success；2-fail"`
	GiftSec   int         `description:"赠送时间（失败为0）"`
	CreatedAt *gtime.Time `description:"创建时间"`
	UpdatedAt *gtime.Time `description:"更新时间"`
	Comment   string      `description:"备注信息"`
}
