// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TAdGift is the golang structure of table t_ad_gift for DAO operations like Where/Data.
type TAdGift struct {
	g.Meta       `orm:"table:t_ad_gift, do:true"`
	Id           interface{} // 自增id
	UserId       interface{} // 用户id
	AdId         interface{} // 广告ID
	AdName       interface{} // 广告名称
	ExposureTime interface{} // 单次曝光时间，单位秒
	GiftDuration interface{} // 赠送时间
	CreatedAt    *gtime.Time // 创建时间
}
