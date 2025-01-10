// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TDailyAdStatisticsDao is the data access object for the table t_daily_ad_statistics.
type TDailyAdStatisticsDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of the current DAO.
	columns TDailyAdStatisticsColumns // columns contains all the column names of Table for convenient usage.
}

// TDailyAdStatisticsColumns defines and stores column names for the table t_daily_ad_statistics.
type TDailyAdStatisticsColumns struct {
	Id        string // 主键ID
	AdId      string // 广告ID
	AdName    string // 广告名称
	Date      string // 统计数据日期，整数类型，格式为 YYYYMMDD，例如20250102表示2025年1月2日
	Exposure  string // 广告的曝光量，默认值为0，表示当天的曝光次数
	Clicks    string // 广告的点击量，默认值为0，表示当天的点击次数
	Rewards   string // 广告完播后获赠时长的用户数，默认值为0，表示当天广告完播后获赠时长的用户数
	CreatedAt string // 记录创建时间，默认值为当前时间
	UpdatedAt string // 记录更新时间，默认值为当前时间，并在每次更新时自动更新
}

// tDailyAdStatisticsColumns holds the columns for the table t_daily_ad_statistics.
var tDailyAdStatisticsColumns = TDailyAdStatisticsColumns{
	Id:        "id",
	AdId:      "ad_id",
	AdName:    "ad_name",
	Date:      "date",
	Exposure:  "exposure",
	Clicks:    "clicks",
	Rewards:   "rewards",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewTDailyAdStatisticsDao creates and returns a new DAO object for table data access.
func NewTDailyAdStatisticsDao() *TDailyAdStatisticsDao {
	return &TDailyAdStatisticsDao{
		group:   "speed-report",
		table:   "t_daily_ad_statistics",
		columns: tDailyAdStatisticsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TDailyAdStatisticsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TDailyAdStatisticsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TDailyAdStatisticsDao) Columns() TDailyAdStatisticsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TDailyAdStatisticsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TDailyAdStatisticsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TDailyAdStatisticsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
