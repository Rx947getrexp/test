// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TNode is the golang structure for table t_node.
type TNode struct {
	Id          int64       `orm:"id"           description:"自增id"`
	Name        string      `orm:"name"         description:"节点名称"`
	Title       string      `orm:"title"        description:"节点标题"`
	TitleEn     string      `orm:"title_en"     description:"节点标题（英文）"`
	TitleRus    string      `orm:"title_rus"    description:"节点标题（俄文)"`
	Country     string      `orm:"country"      description:"国家"`
	CountryEn   string      `orm:"country_en"   description:"国家（英文）"`
	CountryRus  string      `orm:"country_rus"  description:"国家（俄文)"`
	Ip          string      `orm:"ip"           description:"内网IP"`
	Server      string      `orm:"server"       description:"公网域名"`
	NodeType    int         `orm:"node_type"    description:"节点类别:1-常规；2-高带宽...(根据情况而定)"`
	Port        int         `orm:"port"         description:"公网端口"`
	Cpu         int         `orm:"cpu"          description:"cpu核数量（单位个）"`
	Flow        int64       `orm:"flow"         description:"流量带宽"`
	Disk        int64       `orm:"disk"         description:"磁盘容量（单位B）"`
	Memory      int64       `orm:"memory"       description:"内存大小（单位B）"`
	MinPort     int         `orm:"min_port"     description:"最小端口"`
	MaxPort     int         `orm:"max_port"     description:"最大端口"`
	Path        string      `orm:"path"         description:"ws路径"`
	IsRecommend int         `orm:"is_recommend" description:"推荐节点1-是；2-否"`
	ChannelId   int         `orm:"channel_id"   description:"市场渠道（默认0）-优选节点有效"`
	Status      int         `orm:"status"       description:"状态:1-正常；2-已软删"`
	CreatedAt   *gtime.Time `orm:"created_at"   description:"创建时间"`
	UpdatedAt   *gtime.Time `orm:"updated_at"   description:"更新时间"`
	Author      string      `orm:"author"       description:"作者"`
	Comment     string      `orm:"comment"      description:"备注信息"`
	Weight      uint        `orm:"weight"       description:"权重"`
}
