// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TAppDns is the golang structure of table t_app_dns for DAO operations like Where/Data.
type TAppDns struct {
	g.Meta    `orm:"table:t_app_dns, do:true"`
	Id        interface{} // 自增id
	SiteType  interface{} // 站点类型:1-appapi域名；2-offsite域名；3-onlineservice在线客服；4-管理后台...
	Dns       interface{} // 域名
	Ip        interface{} // ip地址
	Level     interface{} // 线路级别:1,2,3...用于白名单机制
	Status    interface{} // 状态:1-正常；2-已软删
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	Author    interface{} // 作者
	Comment   interface{} // 备注信息
}
