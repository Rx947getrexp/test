// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TUserConfigDao is the data access object for table t_user_config.
type TUserConfigDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns TUserConfigColumns // columns contains all the column names of Table for convenient usage.
}

// TUserConfigColumns defines and stores column names for table t_user_config.
type TUserConfigColumns struct {
	Id        string // id
	UserId    string // id
	NodeId    string // ID
	Status    string // :1-2-
	CreatedAt string //
	UpdatedAt string //
}

// tUserConfigColumns holds the columns for table t_user_config.
var tUserConfigColumns = TUserConfigColumns{
	Id:        "id",
	UserId:    "user_id",
	NodeId:    "node_id",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewTUserConfigDao creates and returns a new DAO object for table data access.
func NewTUserConfigDao() *TUserConfigDao {
	return &TUserConfigDao{
		group:   "speed",
		table:   "t_user_config",
		columns: tUserConfigColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TUserConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TUserConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TUserConfigDao) Columns() TUserConfigColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TUserConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TUserConfigDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TUserConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
