// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserLogsDao is the data access object for table user_logs.
type UserLogsDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns UserLogsColumns // columns contains all the column names of Table for convenient usage.
}

// UserLogsColumns defines and stores column names for table user_logs.
type UserLogsColumns struct {
	Id        string // 自增id
	UserId    string // 用户id
	Datestr   string // 日期
	Ip        string // IP地址
	UserAgent string // 请求头user-agent
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	Comment   string // 备注信息
}

// userLogsColumns holds the columns for table user_logs.
var userLogsColumns = UserLogsColumns{
	Id:        "id",
	UserId:    "user_id",
	Datestr:   "datestr",
	Ip:        "ip",
	UserAgent: "user_agent",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	Comment:   "comment",
}

// NewUserLogsDao creates and returns a new DAO object for table data access.
func NewUserLogsDao() *UserLogsDao {
	return &UserLogsDao{
		group:   "speed",
		table:   "user_logs",
		columns: userLogsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserLogsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserLogsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserLogsDao) Columns() UserLogsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserLogsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserLogsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserLogsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
