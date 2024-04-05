// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TUser is the golang structure for table t_user.
type TUser struct {
	Id          int64       `description:"自增id"`
	Uname       string      `description:"用户名"`
	Passwd      string      `description:"用户密码"`
	Email       string      `description:"邮件"`
	Phone       string      `description:"电话"`
	Level       int         `description:"等级：0-vip0；1-vip1；2-vip2"`
	ExpiredTime int64       `description:"vip到期时间"`
	V2RayUuid   string      `description:"节点UUID"`
	V2RayTag    int         `description:"v2ray存在UUID标签:1-有；2-无"`
	Channel     string      `description:""`
	ChannelId   int         `description:"渠道id"`
	Status      int         `description:"冻结状态：0-正常；1-冻结"`
	CreatedAt   *gtime.Time `description:"创建时间"`
	UpdatedAt   *gtime.Time `description:"更新时间"`
	Comment     string      `description:"备注信息"`
	ClientId    string      `description:""`
}
