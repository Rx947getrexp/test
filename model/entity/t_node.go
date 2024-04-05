// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TNode is the golang structure for table t_node.
type TNode struct {
	Id          int64       `description:"自增id"`
	Name        string      `description:"节点名称"`
	Title       string      `description:"节点标题"`
	TitleEn     string      `description:"节点标题（英文）"`
	TitleRus    string      `description:"节点标题（俄文)"`
	Country     string      `description:"国家"`
	CountryEn   string      `description:"国家（英文）"`
	CountryRus  string      `description:"国家（俄文)"`
	Ip          string      `description:"内网IP"`
	Server      string      `description:"公网域名"`
	NodeType    int         `description:"节点类别:1-常规；2-高带宽...(根据情况而定)"`
	Port        int         `description:"公网端口"`
	Cpu         int         `description:"cpu核数量（单位个）"`
	Flow        int64       `description:"流量带宽"`
	Disk        int64       `description:"磁盘容量（单位B）"`
	Memory      int64       `description:"内存大小（单位B）"`
	MinPort     int         `description:"最小端口"`
	MaxPort     int         `description:"最大端口"`
	Path        string      `description:"ws路径"`
	IsRecommend int         `description:"推荐节点1-是；2-否"`
	ChannelId   int         `description:"市场渠道（默认0）-优选节点有效"`
	Status      int         `description:"状态:1-正常；2-已软删"`
	CreatedAt   *gtime.Time `description:"创建时间"`
	UpdatedAt   *gtime.Time `description:"更新时间"`
	Author      string      `description:"作者"`
	Comment     string      `description:"备注信息"`
}
