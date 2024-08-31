// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TV2RayUserTraffic is the golang structure for table t_v2ray_user_traffic.
type TV2RayUserTraffic struct {
	Id        uint64      `orm:"id"         description:"自增id"`
	Email     string      `orm:"email"      description:"邮件"`
	Date      uint        `orm:"date"       description:"数据日期, 20230101"`
	Ip        string      `orm:"ip"         description:"ip地址"`
	Uplink    uint64      `orm:"uplink"     description:"上行流量"`
	Downlink  uint64      `orm:"downlink"   description:"下行流量"`
	CreatedAt *gtime.Time `orm:"created_at" description:"记录创建时间"`
	UpdatedAt *gtime.Time `orm:"updated_at" description:"记录更新时间"`
}
