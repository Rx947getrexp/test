// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TDailyPaymentTotalByChannelDao is the data access object for the table t_daily_payment_total_by_channel.
type TDailyPaymentTotalByChannelDao struct {
	table   string                             // table is the underlying table name of the DAO.
	group   string                             // group is the database configuration group name of the current DAO.
	columns TDailyPaymentTotalByChannelColumns // columns contains all the column names of Table for convenient usage.
}

// TDailyPaymentTotalByChannelColumns defines and stores column names for the table t_daily_payment_total_by_channel.
type TDailyPaymentTotalByChannelColumns struct {
	Id        string // 主键ID
	Date      string // 统计数据日期，整数类型，格式为 YYYYMMDD，例如20250102表示2025年1月2日
	Channel   string // 支付渠道名称
	Amount    string // 支付金额统计
	CreatedAt string // 记录创建时间，默认值为当前时间
	UpdatedAt string // 记录更新时间，默认值为当前时间，并在每次更新时自动更新
}

// tDailyPaymentTotalByChannelColumns holds the columns for the table t_daily_payment_total_by_channel.
var tDailyPaymentTotalByChannelColumns = TDailyPaymentTotalByChannelColumns{
	Id:        "id",
	Date:      "date",
	Channel:   "channel",
	Amount:    "amount",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewTDailyPaymentTotalByChannelDao creates and returns a new DAO object for table data access.
func NewTDailyPaymentTotalByChannelDao() *TDailyPaymentTotalByChannelDao {
	return &TDailyPaymentTotalByChannelDao{
		group:   "speed-report",
		table:   "t_daily_payment_total_by_channel",
		columns: tDailyPaymentTotalByChannelColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TDailyPaymentTotalByChannelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TDailyPaymentTotalByChannelDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TDailyPaymentTotalByChannelDao) Columns() TDailyPaymentTotalByChannelColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TDailyPaymentTotalByChannelDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TDailyPaymentTotalByChannelDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TDailyPaymentTotalByChannelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
