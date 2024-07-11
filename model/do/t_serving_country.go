// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TServingCountry is the golang structure of table t_serving_country for DAO operations like Where/Data.
type TServingCountry struct {
	g.Meta      `orm:"table:t_serving_country, do:true"`
	Id          interface{} // 自增id
	Name        interface{} // 国家名称，不可以修改，作为ID用
	Display     interface{} // 用于在用户侧展示的国家名称
	LogoLink    interface{} // 国家图片地址
	PingUrl     interface{} // 前端使用
	IsRecommend interface{} // 推荐节点1-是；2-否
	Weight      interface{} // 权重
	Status      interface{} // 状态:1-正常；2-已软删
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	Level       interface{} // 等级：0-所有用户都可以选择；1-青铜、铂金会员可选择；2-铂金会员可选择
	IsFree      interface{} // 是否为免费站点，0: 不免费,1: 免费
}
