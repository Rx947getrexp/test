// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TAppVersion is the golang structure of table t_app_version for DAO operations like Where/Data.
type TAppVersion struct {
	g.Meta    `orm:"table:t_app_version, do:true"`
	Id        interface{} // 自增id
	AppType   interface{} // 1-ios;2-安卓；3-h5zip
	Version   interface{} // 版本号
	Link      interface{} // 超链地址
	Status    interface{} // 状态:1-正常；2-已软删
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	Author    interface{} // 作者
	Comment   interface{} // 备注信息
}
