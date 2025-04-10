// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TAppStore is the golang structure of table t_app_store for DAO operations like Where/Data.
type TAppStore struct {
	g.Meta    `orm:"table:t_app_store, do:true"`
	Id        interface{} // 自增id
	TitleCn   interface{} // 商店名称(中文)
	TitleEn   interface{} // 商店名称(英文)
	TitleRu   interface{} // 商店名称(俄语)
	Type      interface{} // 商店类型，ios(苹果)，android(安卓)...
	Url       interface{} // 商店地址
	Cover     interface{} // 商店图标
	Status    interface{} // 状态:1-正常；2-已软删
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	Author    interface{} // 作者
	Comment   interface{} // 备注信息
}
