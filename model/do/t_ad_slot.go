// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TAdSlot is the golang structure of table t_ad_slot for DAO operations like Where/Data.
type TAdSlot struct {
	g.Meta    `orm:"table:t_ad_slot, do:true"`
	Id        interface{} // 自增id
	Location  interface{} // 广告位的位置，相当于ID
	Name      interface{} // 广告位名称
	Desc      interface{} // 广告位描述
	Status    interface{} // 状态:1-上架；2-下架
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
