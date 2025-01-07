// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TAdGift is the golang structure for table t_ad_gift.
type TAdGift struct {
	Id           uint64      `description:"自增id"`
	UserId       uint64      `description:"用户id"`
	AdId         uint64      `description:"广告ID"`
	AdName       string      `description:"广告名称"`
	ExposureTime int         `description:"单次曝光时间，单位秒"`
	GiftDuration int         `description:"赠送时间"`
	CreatedAt    *gtime.Time `description:"创建时间"`
}
