// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TV2RayUserTrafficLogDao is the data access object for table t_v2ray_user_traffic_log.
type TV2RayUserTrafficLogDao struct {
	table   string                      // table is the underlying table name of the DAO.
	group   string                      // group is the database configuration group name of current DAO.
	columns TV2RayUserTrafficLogColumns // columns contains all the column names of Table for convenient usage.
}

// TV2RayUserTrafficLogColumns defines and stores column names for table t_v2ray_user_traffic_log.
type TV2RayUserTrafficLogColumns struct {
	Id        string // 自增id
	Email     string // 邮件
	Ip        string // ip地址
	DateTime  string // 数据采集时间
	Uplink    string // 上行流量
	Downlink  string // 下行流量
	CreatedAt string // 记录创建时间
}

// tV2RayUserTrafficLogColumns holds the columns for table t_v2ray_user_traffic_log.
var tV2RayUserTrafficLogColumns = TV2RayUserTrafficLogColumns{
	Id:        "id",
	Email:     "email",
	Ip:        "ip",
	DateTime:  "date_time",
	Uplink:    "uplink",
	Downlink:  "downlink",
	CreatedAt: "created_at",
}

// NewTV2RayUserTrafficLogDao creates and returns a new DAO object for table data access.
func NewTV2RayUserTrafficLogDao() *TV2RayUserTrafficLogDao {
	return &TV2RayUserTrafficLogDao{
		group:   "speed_collector",
		table:   "t_v2ray_user_traffic_log",
		columns: tV2RayUserTrafficLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TV2RayUserTrafficLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TV2RayUserTrafficLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TV2RayUserTrafficLogDao) Columns() TV2RayUserTrafficLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TV2RayUserTrafficLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TV2RayUserTrafficLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TV2RayUserTrafficLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
