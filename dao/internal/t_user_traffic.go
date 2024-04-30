// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TUserTrafficDao is the data access object for table t_user_traffic.
type TUserTrafficDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns TUserTrafficColumns // columns contains all the column names of Table for convenient usage.
}

// TUserTrafficColumns defines and stores column names for table t_user_traffic.
type TUserTrafficColumns struct {
	Id        string // 自增id
	Email     string // 邮件
	Ip        string // ip地址
	Date      string // 数据日期, 20230101
	Uplink    string // 上行流量
	Downlink  string // 下行流量
	CreatedAt string // 记录创建时间
	UpdatedAt string // 记录更新时间
}

// tUserTrafficColumns holds the columns for table t_user_traffic.
var tUserTrafficColumns = TUserTrafficColumns{
	Id:        "id",
	Email:     "email",
	Ip:        "ip",
	Date:      "date",
	Uplink:    "uplink",
	Downlink:  "downlink",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewTUserTrafficDao creates and returns a new DAO object for table data access.
func NewTUserTrafficDao() *TUserTrafficDao {
	return &TUserTrafficDao{
		group:   "speed",
		table:   "t_user_traffic",
		columns: tUserTrafficColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TUserTrafficDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TUserTrafficDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TUserTrafficDao) Columns() TUserTrafficColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TUserTrafficDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TUserTrafficDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TUserTrafficDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
