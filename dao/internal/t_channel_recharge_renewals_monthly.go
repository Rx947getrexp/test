// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TChannelRechargeRenewalsMonthlyDao is the data access object for the table t_channel_recharge_renewals_monthly.
type TChannelRechargeRenewalsMonthlyDao struct {
	table   string                                 // table is the underlying table name of the DAO.
	group   string                                 // group is the database configuration group name of the current DAO.
	columns TChannelRechargeRenewalsMonthlyColumns // columns contains all the column names of Table for convenient usage.
}

// TChannelRechargeRenewalsMonthlyColumns defines and stores column names for the table t_channel_recharge_renewals_monthly.
type TChannelRechargeRenewalsMonthlyColumns struct {
	Id             string // 主键ID
	Month          string // 统计数据月份，整数类型，格式为 YYYYMM，例如202501
	Channel        string // 渠道id
	RechargeUsers  string // 付费用户数量
	RechargeAmount string // 付费用户充值总金额
	Retained       string // 充值用户次月留存数量
	RenewalsUsers  string // 次月续费人数
	RenewalsAmount string // 次月续费充值总金额
	CreatedAt      string // 记录创建时间，默认值为当前时间
	UpdatedAt      string // 记录更新时间，默认值为当前时间，并在每次更新时自动更新
}

// tChannelRechargeRenewalsMonthlyColumns holds the columns for the table t_channel_recharge_renewals_monthly.
var tChannelRechargeRenewalsMonthlyColumns = TChannelRechargeRenewalsMonthlyColumns{
	Id:             "id",
	Month:          "month",
	Channel:        "channel",
	RechargeUsers:  "recharge_users",
	RechargeAmount: "recharge_amount",
	Retained:       "retained",
	RenewalsUsers:  "renewals_users",
	RenewalsAmount: "renewals_amount",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewTChannelRechargeRenewalsMonthlyDao creates and returns a new DAO object for table data access.
func NewTChannelRechargeRenewalsMonthlyDao() *TChannelRechargeRenewalsMonthlyDao {
	return &TChannelRechargeRenewalsMonthlyDao{
		group:   "speed-report",
		table:   "t_channel_recharge_renewals_monthly",
		columns: tChannelRechargeRenewalsMonthlyColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TChannelRechargeRenewalsMonthlyDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TChannelRechargeRenewalsMonthlyDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TChannelRechargeRenewalsMonthlyDao) Columns() TChannelRechargeRenewalsMonthlyColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TChannelRechargeRenewalsMonthlyDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TChannelRechargeRenewalsMonthlyDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TChannelRechargeRenewalsMonthlyDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
