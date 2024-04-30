// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TUserTeam is the golang structure for table t_user_team.
type TUserTeam struct {
	Id         int64       `description:"自增id"`
	UserId     int64       `description:"用户id"`
	DirectId   int64       `description:"上级id"`
	DirectTree string      `description:"上级列表"`
	CreatedAt  *gtime.Time `description:"创建时间"`
	UpdatedAt  *gtime.Time `description:"更新时间"`
	Comment    string      `description:"备注信息"`
}
