// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TNode is the golang structure of table t_node for DAO operations like Where/Data.
type TNode struct {
	g.Meta      `orm:"table:t_node, do:true"`
	Id          interface{} // 自增id
	Name        interface{} // 节点名称
	Title       interface{} // 节点标题
	TitleEn     interface{} // 节点标题（英文）
	TitleRus    interface{} // 节点标题（俄文)
	Country     interface{} // 国家
	CountryEn   interface{} // 国家（英文）
	CountryRus  interface{} // 国家（俄文)
	Ip          interface{} // 内网IP
	Server      interface{} // 公网域名
	NodeType    interface{} // 节点类别:1-常规；2-高带宽...(根据情况而定)
	Port        interface{} // 公网端口
	Cpu         interface{} // cpu核数量（单位个）
	Flow        interface{} // 流量带宽
	Disk        interface{} // 磁盘容量（单位B）
	Memory      interface{} // 内存大小（单位B）
	MinPort     interface{} // 最小端口
	MaxPort     interface{} // 最大端口
	Path        interface{} // ws路径
	IsRecommend interface{} // 推荐节点1-是；2-否
	ChannelId   interface{} // 市场渠道（默认0）-优选节点有效
	Status      interface{} // 状态:1-正常；2-已软删
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	Author      interface{} // 作者
	Comment     interface{} // 备注信息
}
