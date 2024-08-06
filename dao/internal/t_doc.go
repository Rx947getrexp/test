// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TDocDao is the data access object for table t_doc.
type TDocDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns TDocColumns // columns contains all the column names of Table for convenient usage.
}

// TDocColumns defines and stores column names for table t_doc.
type TDocColumns struct {
	Id        string // 自增id
	Type      string //
	Name      string //
	Desc      string //
	Content   string // content
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// tDocColumns holds the columns for table t_doc.
var tDocColumns = TDocColumns{
	Id:        "id",
	Type:      "type",
	Name:      "name",
	Desc:      "desc",
	Content:   "content",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewTDocDao creates and returns a new DAO object for table data access.
func NewTDocDao() *TDocDao {
	return &TDocDao{
		group:   "speed",
		table:   "t_doc",
		columns: tDocColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TDocDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TDocDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TDocDao) Columns() TDocColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TDocDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TDocDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TDocDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
