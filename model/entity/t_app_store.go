// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TAppStore is the golang structure for table t_app_store.
type TAppStore struct {
	Id        int64       `orm:"id"         description:"自增id"`
	TitleCn   string      `orm:"title_cn"   description:"商店名称(中文)"`
	TitleEn   string      `orm:"title_en"   description:"商店名称(英文)"`
	TitleRu   string      `orm:"title_ru"   description:"商店名称(俄语)"`
	Type      string      `orm:"type"       description:"商店类型，ios(苹果)，android(安卓)..."`
	Url       string      `orm:"url"        description:"商店地址"`
	Cover     string      `orm:"cover"      description:"商店图标"`
	Status    int         `orm:"status"     description:"状态:1-正常；2-已软删"`
	CreatedAt *gtime.Time `orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `orm:"updated_at" description:"更新时间"`
	Author    string      `orm:"author"     description:"作者"`
	Comment   string      `orm:"comment"    description:"备注信息"`
}
