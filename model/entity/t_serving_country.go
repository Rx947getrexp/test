// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TServingCountry is the golang structure for table t_serving_country.
type TServingCountry struct {
	Id          int64       `orm:"id"           description:"自增id"`
	Name        string      `orm:"name"         description:"国家名称，不可以修改，作为ID用"`
	Display     string      `orm:"display"      description:"用于在用户侧展示的国家名称"`
	LogoLink    string      `orm:"logo_link"    description:"国家图片地址"`
	PingUrl     string      `orm:"ping_url"     description:"前端使用"`
	IsRecommend int         `orm:"is_recommend" description:"推荐节点1-是；2-否"`
	Weight      int         `orm:"weight"       description:"权重"`
	Status      int         `orm:"status"       description:"状态:1-正常；2-已软删"`
	CreatedAt   *gtime.Time `orm:"created_at"   description:"创建时间"`
	UpdatedAt   *gtime.Time `orm:"updated_at"   description:"更新时间"`
	Level       int         `orm:"level"        description:"等级：0-所有用户都可以选择；1-青铜、铂金会员可选择；2-铂金会员可选择"`
}
