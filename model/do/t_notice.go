// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TNotice is the golang structure of table t_notice for DAO operations like Where/Data.
type TNotice struct {
	g.Meta     `orm:"table:t_notice, do:true"`
	Id         interface{} // 自增id
	Title      interface{} // 标题
	TitleEn    interface{} // 标题（英文）
	TitleRus   interface{} // 标题（俄文）
	Tag        interface{} // 标签
	TagEn      interface{} // 标签（英文）
	TagRus     interface{} // 标签（俄文）
	Content    interface{} // 正文内容
	ContentEn  interface{} // 正文内容（英文）
	ContentRus interface{} // 正文内容（俄文）
	Author     interface{} // 作者
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
	Status     interface{} // 状态:1-发布；2-软删
	Comment    interface{} // 备注信息
}
