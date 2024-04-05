// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TUserConfig is the golang structure for table t_user_config.
type TUserConfig struct {
	Id        int64       `description:"id"`
	UserId    int64       `description:"id"`
	NodeId    int64       `description:"ID"`
	Status    int         `description:":1-2-"`
	CreatedAt *gtime.Time `description:""`
	UpdatedAt *gtime.Time `description:""`
}
