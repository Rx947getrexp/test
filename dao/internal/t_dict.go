// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TDictDao is the data access object for table t_dict.
type TDictDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns TDictColumns // columns contains all the column names of Table for convenient usage.
}

// TDictColumns defines and stores column names for table t_dict.
type TDictColumns struct {
	KeyId     string // 键
	Value     string // 值
	Note      string // 描述
	IsDel     string // 0-正常；1-软删
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// tDictColumns holds the columns for table t_dict.
var tDictColumns = TDictColumns{
	KeyId:     "key_id",
	Value:     "value",
	Note:      "note",
	IsDel:     "is_del",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewTDictDao creates and returns a new DAO object for table data access.
func NewTDictDao() *TDictDao {
	return &TDictDao{
		group:   "speed",
		table:   "t_dict",
		columns: tDictColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TDictDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TDictDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TDictDao) Columns() TDictColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TDictDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TDictDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TDictDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
