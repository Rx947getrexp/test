// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TNodeDao is the data access object for table t_node.
type TNodeDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns TNodeColumns // columns contains all the column names of Table for convenient usage.
}

// TNodeColumns defines and stores column names for table t_node.
type TNodeColumns struct {
	Id          string // 自增id
	Name        string // 节点名称
	Title       string // 节点标题
	TitleEn     string // 节点标题（英文）
	TitleRus    string // 节点标题（俄文)
	Country     string // 国家
	CountryEn   string // 国家（英文）
	CountryRus  string // 国家（俄文)
	Ip          string // 内网IP
	Server      string // 公网域名
	NodeType    string // 节点类别:1-常规；2-高带宽...(根据情况而定)
	Port        string // 公网端口
	Cpu         string // cpu核数量（单位个）
	Flow        string // 流量带宽
	Disk        string // 磁盘容量（单位B）
	Memory      string // 内存大小（单位B）
	MinPort     string // 最小端口
	MaxPort     string // 最大端口
	Path        string // ws路径
	IsRecommend string // 推荐节点1-是；2-否
	ChannelId   string // 市场渠道（默认0）-优选节点有效
	Status      string // 状态:1-正常；2-已软删
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
	Author      string // 作者
	Comment     string // 备注信息
}

// tNodeColumns holds the columns for table t_node.
var tNodeColumns = TNodeColumns{
	Id:          "id",
	Name:        "name",
	Title:       "title",
	TitleEn:     "title_en",
	TitleRus:    "title_rus",
	Country:     "country",
	CountryEn:   "country_en",
	CountryRus:  "country_rus",
	Ip:          "ip",
	Server:      "server",
	NodeType:    "node_type",
	Port:        "port",
	Cpu:         "cpu",
	Flow:        "flow",
	Disk:        "disk",
	Memory:      "memory",
	MinPort:     "min_port",
	MaxPort:     "max_port",
	Path:        "path",
	IsRecommend: "is_recommend",
	ChannelId:   "channel_id",
	Status:      "status",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	Author:      "author",
	Comment:     "comment",
}

// NewTNodeDao creates and returns a new DAO object for table data access.
func NewTNodeDao() *TNodeDao {
	return &TNodeDao{
		group:   "speed",
		table:   "t_node",
		columns: tNodeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TNodeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TNodeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TNodeDao) Columns() TNodeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TNodeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TNodeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TNodeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
