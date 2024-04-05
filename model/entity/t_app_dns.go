// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TAppDns is the golang structure for table t_app_dns.
type TAppDns struct {
	Id        int64       `description:"自增id"`
	SiteType  int         `description:"站点类型:1-appapi域名；2-offsite域名；3-onlineservice在线客服；4-管理后台..."`
	Dns       string      `description:"域名"`
	Ip        string      `description:"ip地址"`
	Level     int         `description:"线路级别:1,2,3...用于白名单机制"`
	Status    int         `description:"状态:1-正常；2-已软删"`
	CreatedAt *gtime.Time `description:"创建时间"`
	UpdatedAt *gtime.Time `description:"更新时间"`
	Author    string      `description:"作者"`
	Comment   string      `description:"备注信息"`
}
