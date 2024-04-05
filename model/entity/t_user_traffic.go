// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TUserTraffic is the golang structure for table t_user_traffic.
type TUserTraffic struct {
	Id        uint64      `description:"自增id"`
	Email     string      `description:"邮件"`
	Ip        string      `description:"ip地址"`
	Date      uint        `description:"数据日期, 20230101"`
	Uplink    uint64      `description:"上行流量"`
	Downlink  uint64      `description:"下行流量"`
	CreatedAt *gtime.Time `description:"记录创建时间"`
	UpdatedAt *gtime.Time `description:"记录更新时间"`
}
