// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TActivityDao is the data access object for table t_activity.
type TActivityDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns TActivityColumns // columns contains all the column names of Table for convenient usage.
}

// TActivityColumns defines and stores column names for table t_activity.
type TActivityColumns struct {
	Id        string // 自增id
	UserId    string // 用户id
	Status    string // 状态:1-success；2-fail
	GiftSec   string // 赠送时间（失败为0）
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	Comment   string // 备注信息
}

// tActivityColumns holds the columns for table t_activity.
var tActivityColumns = TActivityColumns{
	Id:        "id",
	UserId:    "user_id",
	Status:    "status",
	GiftSec:   "gift_sec",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	Comment:   "comment",
}

// NewTActivityDao creates and returns a new DAO object for table data access.
func NewTActivityDao() *TActivityDao {
	return &TActivityDao{
		group:   "speed",
		table:   "t_activity",
		columns: tActivityColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TActivityDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TActivityDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TActivityDao) Columns() TActivityColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TActivityDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TActivityDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TActivityDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
