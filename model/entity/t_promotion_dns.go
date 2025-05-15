// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TPromotionDns is the golang structure for table t_promotion_dns.
type TPromotionDns struct {
	Id             int64       `orm:"id"              description:"自增id"`
	Dns            string      `orm:"dns"             description:"域名"`
	Ip             string      `orm:"ip"              description:"ip地址"`
	MacChannel     string      `orm:"mac_channel"     description:"苹果电脑渠道"`
	WinChannel     string      `orm:"win_channel"     description:"windows电脑渠道"`
	AndroidChannel string      `orm:"android_channel" description:"安卓渠道"`
	Promoter       string      `orm:"promoter"        description:"推广人员"`
	Status         int         `orm:"status"          description:"状态:1-正常；2-已软删"`
	CreatedAt      *gtime.Time `orm:"created_at"      description:"创建时间"`
	UpdatedAt      *gtime.Time `orm:"updated_at"      description:"更新时间"`
	Author         string      `orm:"author"          description:"作者"`
	Comment        string      `orm:"comment"         description:"备注信息"`
}
