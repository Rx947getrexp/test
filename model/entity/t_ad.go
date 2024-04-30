// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TAd is the golang structure for table t_ad.
type TAd struct {
	Id        int64       `description:"自增id"`
	Status    int         `description:"状态:1-上架；2-下架"`
	Sort      int         `description:"排序"`
	Name      string      `description:"广告名称"`
	Logo      string      `description:"广告logo"`
	Link      string      `description:"广告链接"`
	AdType    int         `description:"广告分类：1-社交；2-游戏；3-漫画；4-视频..."`
	Tag       string      `description:"标签标题"`
	Content   string      `description:"正文介绍"`
	Author    string      `description:"作者"`
	CreatedAt *gtime.Time `description:"创建时间"`
	UpdatedAt *gtime.Time `description:"更新时间"`
	Comment   string      `description:"备注信息"`
}
