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
	g.Meta    `orm:"table:t_ad, do:true"`
	Id        interface{} // 自增id
	Status    interface{} // 状态:1-上架；2-下架
	Sort      interface{} // 排序
	Name      interface{} // 广告名称
	Logo      interface{} // 广告logo
	Link      interface{} // 广告链接
	AdType    interface{} // 广告分类：1-社交；2-游戏；3-漫画；4-视频...
	Tag       interface{} // 标签标题
	Content   interface{} // 正文介绍
	Author    interface{} // 作者
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	Comment   interface{} // 备注信息
}
