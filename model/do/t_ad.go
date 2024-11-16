// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TAd is the golang structure of table t_ad for DAO operations like Where/Data.
type TAd struct {
	g.Meta        `orm:"table:t_ad, do:true"`
	Id            interface{} // 自增id
	Advertiser    interface{} // 广告主，客户名称
	Name          interface{} // 广告名称
	Type          interface{} // 广告类型. enum: text,image,video
	Url           interface{} // 广告内容地址
	Logo          interface{} // logo
	SlotLocations interface{} // 广告位的位置，包括权重
	Devices       interface{} // 广告位的位置，包括权重
	TargetUrls    interface{} // 跳转地址，包括：pc,ios,android
	Labels        interface{} // 标签
	ExposureTime  interface{} // 单次曝光时间，单位秒
	UserLevels    interface{} // 用户等级
	StartTime     *gtime.Time // 广告生效时间
	EndTime       *gtime.Time // 广告失效时间
	Status        interface{} // 状态:1-上架；2-下架
	GiftDuration  interface{} // 赠送时间
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
}
