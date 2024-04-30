// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TNodeUuidDao is the data access object for table t_node_uuid.
type TNodeUuidDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns TNodeUuidColumns // columns contains all the column names of Table for convenient usage.
}

// TNodeUuidColumns defines and stores column names for table t_node_uuid.
type TNodeUuidColumns struct {
	Id        string // 自增id
	UserId    string // 用户id
	NodeId    string // 节点id
	Email     string // 节点邮箱，用于区分流量
	V2RayUuid string // 节点UUID
	Server    string // 公网域名
	Port      string // 公网端口
	UsedFlow  string // 已使用流量（单位B）
	Status    string // 状态:1-正常；2-已软删
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	Comment   string // 备注信息
}

// tNodeUuidColumns holds the columns for table t_node_uuid.
var tNodeUuidColumns = TNodeUuidColumns{
	Id:        "id",
	UserId:    "user_id",
	NodeId:    "node_id",
	Email:     "email",
	V2RayUuid: "v2ray_uuid",
	Server:    "server",
	Port:      "port",
	UsedFlow:  "used_flow",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	Comment:   "comment",
}

// NewTNodeUuidDao creates and returns a new DAO object for table data access.
func NewTNodeUuidDao() *TNodeUuidDao {
	return &TNodeUuidDao{
		group:   "speed",
		table:   "t_node_uuid",
		columns: tNodeUuidColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TNodeUuidDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TNodeUuidDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TNodeUuidDao) Columns() TNodeUuidColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TNodeUuidDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TNodeUuidDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TNodeUuidDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
