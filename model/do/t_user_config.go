// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TUserConfig is the golang structure of table t_user_config for DAO operations like Where/Data.
type TUserConfig struct {
	g.Meta    `orm:"table:t_user_config, do:true"`
	Id        interface{} // id
	UserId    interface{} // id
	NodeId    interface{} // ID
	Status    interface{} // :1-2-
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
