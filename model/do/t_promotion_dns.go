// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TPromotionDns is the golang structure of table t_promotion_dns for DAO operations like Where/Data.
type TPromotionDns struct {
	g.Meta         `orm:"table:t_promotion_dns, do:true"`
	Id             interface{} // 自增id
	Dns            interface{} // 域名
	Ip             interface{} // ip地址
	MacChannel     interface{} // 苹果电脑渠道
	WinChannel     interface{} // windows电脑渠道
	AndroidChannel interface{} // 安卓渠道
	Promoter       interface{} // 推广人员
	Status         interface{} // 状态:1-正常；2-已软删
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
	Author         interface{} // 作者
	Comment        interface{} // 备注信息
}
