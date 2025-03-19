// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TChannelRetaindDailyDao is the data access object for the table t_channel_retaind_daily.
type TChannelRetaindDailyDao struct {
	table   string                      // table is the underlying table name of the DAO.
	group   string                      // group is the database configuration group name of the current DAO.
	columns TChannelRetaindDailyColumns // columns contains all the column names of Table for convenient usage.
}

// TChannelRetaindDailyColumns defines and stores column names for the table t_channel_retaind_daily.
type TChannelRetaindDailyColumns struct {
	Id            string // 主键ID
	Date          string // 统计数据日期，整数类型，格式为 YYYYMMDD，例如20250102表示2025年1月2日
	Channel       string // 渠道id
	NewUsers      string // 新增用户数量
	Day2Retained  string // 新增用户次日留存数量
	Day7Retained  string // 新增用户7日留存数量
	Day15Retained string // 新增用户15日留存数量
	Day30Retained string // 新增用户30日留存数量
	CreatedAt     string // 记录创建时间，默认值为当前时间
	UpdatedAt     string // 记录更新时间，默认值为当前时间，并在每次更新时自动更新
}

// tChannelRetaindDailyColumns holds the columns for the table t_channel_retaind_daily.
var tChannelRetaindDailyColumns = TChannelRetaindDailyColumns{
	Id:            "id",
	Date:          "date",
	Channel:       "channel",
	NewUsers:      "new_users",
	Day2Retained:  "day_2_retained",
	Day7Retained:  "day_7_retained",
	Day15Retained: "day_15_retained",
	Day30Retained: "day_30_retained",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewTChannelRetaindDailyDao creates and returns a new DAO object for table data access.
func NewTChannelRetaindDailyDao() *TChannelRetaindDailyDao {
	return &TChannelRetaindDailyDao{
		group:   "speed-report",
		table:   "t_channel_retaind_daily",
		columns: tChannelRetaindDailyColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TChannelRetaindDailyDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TChannelRetaindDailyDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TChannelRetaindDailyDao) Columns() TChannelRetaindDailyColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TChannelRetaindDailyDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TChannelRetaindDailyDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TChannelRetaindDailyDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
