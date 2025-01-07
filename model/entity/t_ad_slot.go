// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TAdSlot is the golang structure for table t_ad_slot.
type TAdSlot struct {
	Id        uint64      `description:"自增id"`
	Location  string      `description:"广告位的位置，相当于ID"`
	Name      string      `description:"广告位名称"`
	Desc      string      `description:"广告位描述"`
	Status    int         `description:"状态:1-上架；2-下架"`
	CreatedAt *gtime.Time `description:"创建时间"`
	UpdatedAt *gtime.Time `description:"更新时间"`
}
