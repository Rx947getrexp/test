// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TUserTrafficLogDao is the data access object for table t_user_traffic_log.
type TUserTrafficLogDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns TUserTrafficLogColumns // columns contains all the column names of Table for convenient usage.
}

// TUserTrafficLogColumns defines and stores column names for table t_user_traffic_log.
type TUserTrafficLogColumns struct {
	Id        string // 自增id
	Email     string // 邮件
	Ip        string // ip地址
	DateTime  string // 数据采集时间
	Uplink    string // 上行流量
	Downlink  string // 下行流量
	CreatedAt string // 记录创建时间
}

// tUserTrafficLogColumns holds the columns for table t_user_traffic_log.
var tUserTrafficLogColumns = TUserTrafficLogColumns{
	Id:        "id",
	Email:     "email",
	Ip:        "ip",
	DateTime:  "date_time",
	Uplink:    "uplink",
	Downlink:  "downlink",
	CreatedAt: "created_at",
}

// NewTUserTrafficLogDao creates and returns a new DAO object for table data access.
func NewTUserTrafficLogDao() *TUserTrafficLogDao {
	return &TUserTrafficLogDao{
		group:   "speed",
		table:   "t_user_traffic_log",
		columns: tUserTrafficLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TUserTrafficLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TUserTrafficLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TUserTrafficLogDao) Columns() TUserTrafficLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TUserTrafficLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TUserTrafficLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TUserTrafficLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
