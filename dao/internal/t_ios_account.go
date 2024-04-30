// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TIosAccountDao is the data access object for table t_ios_account.
type TIosAccountDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns TIosAccountColumns // columns contains all the column names of Table for convenient usage.
}

// TIosAccountColumns defines and stores column names for table t_ios_account.
type TIosAccountColumns struct {
	Id          string // 编号id
	Account     string // ios账号
	Pass        string // 密码
	Name        string // 别名
	AccountType string // 1-国区；2-海外
	Status      string // 1-正常；2-下架
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
	Author      string // 作者
	Comment     string // 备注
}

// tIosAccountColumns holds the columns for table t_ios_account.
var tIosAccountColumns = TIosAccountColumns{
	Id:          "id",
	Account:     "account",
	Pass:        "pass",
	Name:        "name",
	AccountType: "account_type",
	Status:      "status",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	Author:      "author",
	Comment:     "comment",
}

// NewTIosAccountDao creates and returns a new DAO object for table data access.
func NewTIosAccountDao() *TIosAccountDao {
	return &TIosAccountDao{
		group:   "speed",
		table:   "t_ios_account",
		columns: tIosAccountColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TIosAccountDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TIosAccountDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TIosAccountDao) Columns() TIosAccountColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TIosAccountDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TIosAccountDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TIosAccountDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
