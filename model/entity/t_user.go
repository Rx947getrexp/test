// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TUser is the golang structure for table t_user.
type TUser struct {
	Id               int64       `orm:"id"                 description:"自增id"`
	Uname            string      `orm:"uname"              description:"用户名"`
	Passwd           string      `orm:"passwd"             description:"用户密码"`
	Email            string      `orm:"email"              description:"邮件"`
	Phone            string      `orm:"phone"              description:"电话"`
	Level            int         `orm:"level"              description:"等级：0-vip0；1-vip1；2-vip2"`
	ExpiredTime      int64       `orm:"expired_time"       description:"vip到期时间"`
	V2RayUuid        string      `orm:"v2ray_uuid"         description:"节点UUID"`
	V2RayTag         int         `orm:"v2ray_tag"          description:"v2ray存在UUID标签:1-有；2-无"`
	Channel          string      `orm:"channel"            description:""`
	ChannelId        int         `orm:"channel_id"         description:"渠道id"`
	Status           int         `orm:"status"             description:"冻结状态：0-正常；1-冻结"`
	CreatedAt        *gtime.Time `orm:"created_at"         description:"创建时间"`
	UpdatedAt        *gtime.Time `orm:"updated_at"         description:"更新时间"`
	Comment          string      `orm:"comment"            description:"备注信息"`
	ClientId         string      `orm:"client_id"          description:""`
	LastLoginIp      string      `orm:"last_login_ip"      description:"最近一次登录的ip"`
	LastLoginCountry string      `orm:"last_login_country" description:"最近一次登录的国家"`
	PreferredCountry string      `orm:"preferred_country"  description:"用户选择的国家（国家名称）"`
}
