// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TUserOnlineDayDao is the data access object for table t_user_online_day.
type TUserOnlineDayDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns TUserOnlineDayColumns // columns contains all the column names of Table for convenient usage.
}

// TUserOnlineDayColumns defines and stores column names for table t_user_online_day.
type TUserOnlineDayColumns struct {
	Id               string // 自增id
	Date             string // 数据日期, 20230101
	Email            string // 邮件
	ChannelId        string // 渠道id
	OnlineDuration   string // 在线时间长度
	Uplink           string // 上行流量
	Downlink         string // 下行流量
	CreatedAt        string // 记录创建时间
	LastLoginCountry string // 最后登陆的国家
}

// tUserOnlineDayColumns holds the columns for table t_user_online_day.
var tUserOnlineDayColumns = TUserOnlineDayColumns{
	Id:               "id",
	Date:             "date",
	Email:            "email",
	ChannelId:        "channel_id",
	OnlineDuration:   "online_duration",
	Uplink:           "uplink",
	Downlink:         "downlink",
	CreatedAt:        "created_at",
	LastLoginCountry: "last_login_country",
}

// NewTUserOnlineDayDao creates and returns a new DAO object for table data access.
func NewTUserOnlineDayDao() *TUserOnlineDayDao {
	return &TUserOnlineDayDao{
		group:   "speed-report",
		table:   "t_user_online_day",
		columns: tUserOnlineDayColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TUserOnlineDayDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TUserOnlineDayDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TUserOnlineDayDao) Columns() TUserOnlineDayColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TUserOnlineDayDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TUserOnlineDayDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TUserOnlineDayDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
