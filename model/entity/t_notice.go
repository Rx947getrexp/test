// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TNotice is the golang structure for table t_notice.
type TNotice struct {
	Id         int64       `description:"自增id"`
	Title      string      `description:"标题"`
	TitleEn    string      `description:"标题（英文）"`
	TitleRus   string      `description:"标题（俄文）"`
	Tag        string      `description:"标签"`
	TagEn      string      `description:"标签（英文）"`
	TagRus     string      `description:"标签（俄文）"`
	Content    string      `description:"正文内容"`
	ContentEn  string      `description:"正文内容（英文）"`
	ContentRus string      `description:"正文内容（俄文）"`
	Author     string      `description:"作者"`
	CreatedAt  *gtime.Time `description:"创建时间"`
	UpdatedAt  *gtime.Time `description:"更新时间"`
	Status     int         `description:"状态:1-发布；2-软删"`
	Comment    string      `description:"备注信息"`
}
