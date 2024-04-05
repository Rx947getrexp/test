// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TDevDao is the data access object for table t_dev.
type TDevDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns TDevColumns // columns contains all the column names of Table for convenient usage.
}

// TDevColumns defines and stores column names for table t_dev.
type TDevColumns struct {
	Id        string // 自增id
	ClientId  string //
	Os        string // 客户端设备系统os
	IsSend    string // 1-已赠送时间；2-未赠送
	Network   string // 网络模式（1-自动；2-手动）
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	Comment   string // 备注信息
}

// tDevColumns holds the columns for table t_dev.
var tDevColumns = TDevColumns{
	Id:        "id",
	ClientId:  "client_id",
	Os:        "os",
	IsSend:    "is_send",
	Network:   "network",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	Comment:   "comment",
}

// NewTDevDao creates and returns a new DAO object for table data access.
func NewTDevDao() *TDevDao {
	return &TDevDao{
		group:   "speed",
		table:   "t_dev",
		columns: tDevColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TDevDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TDevDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TDevDao) Columns() TDevColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TDevDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TDevDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TDevDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
