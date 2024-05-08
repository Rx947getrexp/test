// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TPaymentChannelsDao is the data access object for table t_payment_channels.
type TPaymentChannelsDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns TPaymentChannelsColumns // columns contains all the column names of Table for convenient usage.
}

// TPaymentChannelsColumns defines and stores column names for table t_payment_channels.
type TPaymentChannelsColumns struct {
	Id            string // 自增id
	Name          string // 支付通道名称
	IsActive      string // 支付通道是否可用，1表示可用,2表示不可用
	FreeTrialDays string // 赠送的免费时长（以天为单位）
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
}

// tPaymentChannelsColumns holds the columns for table t_payment_channels.
var tPaymentChannelsColumns = TPaymentChannelsColumns{
	Id:            "id",
	Name:          "name",
	IsActive:      "is_active",
	FreeTrialDays: "free_trial_days",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewTPaymentChannelsDao creates and returns a new DAO object for table data access.
func NewTPaymentChannelsDao() *TPaymentChannelsDao {
	return &TPaymentChannelsDao{
		group:   "speed",
		table:   "t_payment_channels",
		columns: tPaymentChannelsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TPaymentChannelsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TPaymentChannelsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TPaymentChannelsDao) Columns() TPaymentChannelsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TPaymentChannelsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TPaymentChannelsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TPaymentChannelsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
