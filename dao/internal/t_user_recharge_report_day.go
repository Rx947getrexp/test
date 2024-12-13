// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TUserRechargeReportDayDao is the data access object for table t_user_recharge_report_day.
type TUserRechargeReportDayDao struct {
	table   string                        // table is the underlying table name of the DAO.
	group   string                        // group is the database configuration group name of current DAO.
	columns TUserRechargeReportDayColumns // columns contains all the column names of Table for convenient usage.
}

// TUserRechargeReportDayColumns defines and stores column names for table t_user_recharge_report_day.
type TUserRechargeReportDayColumns struct {
	Id        string // 自增id
	Date      string // 数据日期, 20230101
	GoodsId   string // 商品套餐id
	Total     string // 用户充值总量
	New       string // 新增用户充值数量
	CreatedAt string // 记录创建时间
}

// tUserRechargeReportDayColumns holds the columns for table t_user_recharge_report_day.
var tUserRechargeReportDayColumns = TUserRechargeReportDayColumns{
	Id:        "id",
	Date:      "date",
	GoodsId:   "goods_id",
	Total:     "total",
	New:       "new",
	CreatedAt: "created_at",
}

// NewTUserRechargeReportDayDao creates and returns a new DAO object for table data access.
func NewTUserRechargeReportDayDao() *TUserRechargeReportDayDao {
	return &TUserRechargeReportDayDao{
		group:   "speed-report",
		table:   "t_user_recharge_report_day",
		columns: tUserRechargeReportDayColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TUserRechargeReportDayDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TUserRechargeReportDayDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TUserRechargeReportDayDao) Columns() TUserRechargeReportDayColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TUserRechargeReportDayDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TUserRechargeReportDayDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TUserRechargeReportDayDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
