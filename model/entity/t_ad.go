// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TAd is the golang structure for table t_ad.
type TAd struct {
	Id            uint64      `description:"自增id"`
	Advertiser    string      `description:"广告主，客户名称"`
	Name          string      `description:"广告名称"`
	Type          string      `description:"广告类型. enum: text,image,video"`
	Url           string      `description:"广告内容地址"`
	Logo          string      `description:"logo"`
	SlotLocations string      `description:"广告位的位置，包括权重"`
	Devices       string      `description:"广告位的位置，包括权重"`
	TargetUrls    string      `description:"跳转地址，包括：pc,ios,android"`
	Labels        string      `description:"标签"`
	ExposureTime  int         `description:"单次曝光时间，单位秒"`
	UserLevels    string      `description:"用户等级"`
	StartTime     *gtime.Time `description:"广告生效时间"`
	EndTime       *gtime.Time `description:"广告失效时间"`
	Status        int         `description:"状态:1-上架；2-下架"`
	GiftDuration  int         `description:"赠送时间"`
	CreatedAt     *gtime.Time `description:"创建时间"`
	UpdatedAt     *gtime.Time `description:"更新时间"`
}
