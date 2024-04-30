// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TUserPingDao is the data access object for table t_user_ping.
type TUserPingDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns TUserPingColumns // columns contains all the column names of Table for convenient usage.
}

// TUserPingColumns defines and stores column names for table t_user_ping.
type TUserPingColumns struct {
	Id        string // 自增id
	Email     string // 用户邮箱
	Host      string // 节点host, ip or dns
	Code      string // ping的结果
	Cost      string // ping耗时
	Time      string // 上报时间
	CreatedAt string // 记录创建时间
}

// tUserPingColumns holds the columns for table t_user_ping.
var tUserPingColumns = TUserPingColumns{
	Id:        "id",
	Email:     "email",
	Host:      "host",
	Code:      "code",
	Cost:      "cost",
	Time:      "time",
	CreatedAt: "created_at",
}

// NewTUserPingDao creates and returns a new DAO object for table data access.
func NewTUserPingDao() *TUserPingDao {
	return &TUserPingDao{
		group:   "speed-report",
		table:   "t_user_ping",
		columns: tUserPingColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TUserPingDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TUserPingDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TUserPingDao) Columns() TUserPingColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TUserPingDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TUserPingDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TUserPingDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
