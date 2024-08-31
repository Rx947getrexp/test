// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TV2RayUserTrafficDao is the data access object for table t_v2ray_user_traffic.
type TV2RayUserTrafficDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns TV2RayUserTrafficColumns // columns contains all the column names of Table for convenient usage.
}

// TV2RayUserTrafficColumns defines and stores column names for table t_v2ray_user_traffic.
type TV2RayUserTrafficColumns struct {
	Id        string // 自增id
	Email     string // 邮件
	Date      string // 数据日期, 20230101
	Ip        string // ip地址
	Uplink    string // 上行流量
	Downlink  string // 下行流量
	CreatedAt string // 记录创建时间
	UpdatedAt string // 记录更新时间
}

// tV2RayUserTrafficColumns holds the columns for table t_v2ray_user_traffic.
var tV2RayUserTrafficColumns = TV2RayUserTrafficColumns{
	Id:        "id",
	Email:     "email",
	Date:      "date",
	Ip:        "ip",
	Uplink:    "uplink",
	Downlink:  "downlink",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewTV2RayUserTrafficDao creates and returns a new DAO object for table data access.
func NewTV2RayUserTrafficDao() *TV2RayUserTrafficDao {
	return &TV2RayUserTrafficDao{
		group:   "speed_collector",
		table:   "t_v2ray_user_traffic",
		columns: tV2RayUserTrafficColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TV2RayUserTrafficDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TV2RayUserTrafficDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TV2RayUserTrafficDao) Columns() TV2RayUserTrafficColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TV2RayUserTrafficDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TV2RayUserTrafficDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TV2RayUserTrafficDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
