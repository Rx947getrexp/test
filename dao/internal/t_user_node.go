// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TUserNodeDao is the data access object for table t_user_node.
type TUserNodeDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns TUserNodeColumns // columns contains all the column names of Table for convenient usage.
}

// TUserNodeColumns defines and stores column names for table t_user_node.
type TUserNodeColumns struct {
	Id        string // 自增id
	UserId    string // 用户uid
	Email     string // 用户邮箱
	Ip        string // 节点IP
	V2RayUuid string // uuid
	Status    string // 状态：0-未写入节点配置；1-已经写入到节点配置
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// tUserNodeColumns holds the columns for table t_user_node.
var tUserNodeColumns = TUserNodeColumns{
	Id:        "id",
	UserId:    "user_id",
	Email:     "email",
	Ip:        "ip",
	V2RayUuid: "v2ray_uuid",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewTUserNodeDao creates and returns a new DAO object for table data access.
func NewTUserNodeDao() *TUserNodeDao {
	return &TUserNodeDao{
		group:   "speed_status",
		table:   "t_user_node",
		columns: tUserNodeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TUserNodeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TUserNodeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TUserNodeDao) Columns() TUserNodeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TUserNodeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TUserNodeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TUserNodeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
