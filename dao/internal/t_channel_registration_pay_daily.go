// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TChannelRegistrationPayDailyDao is the data access object for the table t_channel_registration_pay_daily.
type TChannelRegistrationPayDailyDao struct {
	table   string                              // table is the underlying table name of the DAO.
	group   string                              // group is the database configuration group name of the current DAO.
	columns TChannelRegistrationPayDailyColumns // columns contains all the column names of Table for convenient usage.
}

// TChannelRegistrationPayDailyColumns defines and stores column names for the table t_channel_registration_pay_daily.
type TChannelRegistrationPayDailyColumns struct {
	Id                  string // 主键ID
	Date                string // 统计数据日期，整数类型，格式为 YYYYMMDD，例如20250102表示2025年1月2日
	Channel             string // 渠道id
	NewUsers            string // 新增用户数量
	DailyActiveUsers    string // 日活用户数量
	MonthlyActiveUsers  string // 月活用户数量
	TotalRechargeUsers  string // 充值用户数量
	TotalRechargeAmount string // 付费金额数量
	CreatedAt           string // 记录创建时间，默认值为当前时间
	UpdatedAt           string // 记录更新时间，默认值为当前时间，并在每次更新时自动更新
}

// tChannelRegistrationPayDailyColumns holds the columns for the table t_channel_registration_pay_daily.
var tChannelRegistrationPayDailyColumns = TChannelRegistrationPayDailyColumns{
	Id:                  "id",
	Date:                "date",
	Channel:             "channel",
	NewUsers:            "new_users",
	DailyActiveUsers:    "daily_active_users",
	MonthlyActiveUsers:  "monthly_active_users",
	TotalRechargeUsers:  "total_recharge_users",
	TotalRechargeAmount: "total_recharge_amount",
	CreatedAt:           "created_at",
	UpdatedAt:           "updated_at",
}

// NewTChannelRegistrationPayDailyDao creates and returns a new DAO object for table data access.
func NewTChannelRegistrationPayDailyDao() *TChannelRegistrationPayDailyDao {
	return &TChannelRegistrationPayDailyDao{
		group:   "speed-report",
		table:   "t_channel_registration_pay_daily",
		columns: tChannelRegistrationPayDailyColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TChannelRegistrationPayDailyDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TChannelRegistrationPayDailyDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TChannelRegistrationPayDailyDao) Columns() TChannelRegistrationPayDailyColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TChannelRegistrationPayDailyDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TChannelRegistrationPayDailyDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TChannelRegistrationPayDailyDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
