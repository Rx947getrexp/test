// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TUserReportMonthlyDao is the data access object for table t_user_report_monthly.
type TUserReportMonthlyDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns TUserReportMonthlyColumns // columns contains all the column names of Table for convenient usage.
}

// TUserReportMonthlyColumns defines and stores column names for table t_user_report_monthly.
type TUserReportMonthlyColumns struct {
	Id            string //
	StatMonth     string // 统计月份
	Os            string // 设备类型
	UserCount     string // 用户总数
	NewUsers      string // 新增用户量
	RetainedUsers string // 次月留存
	CreatedAt     string // 记录创建时间
}

// tUserReportMonthlyColumns holds the columns for table t_user_report_monthly.
var tUserReportMonthlyColumns = TUserReportMonthlyColumns{
	Id:            "id",
	StatMonth:     "stat_month",
	Os:            "os",
	UserCount:     "user_count",
	NewUsers:      "new_users",
	RetainedUsers: "retained_users",
	CreatedAt:     "created_at",
}

// NewTUserReportMonthlyDao creates and returns a new DAO object for table data access.
func NewTUserReportMonthlyDao() *TUserReportMonthlyDao {
	return &TUserReportMonthlyDao{
		group:   "speed-report",
		table:   "t_user_report_monthly",
		columns: tUserReportMonthlyColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TUserReportMonthlyDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TUserReportMonthlyDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TUserReportMonthlyDao) Columns() TUserReportMonthlyColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TUserReportMonthlyDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TUserReportMonthlyDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TUserReportMonthlyDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
