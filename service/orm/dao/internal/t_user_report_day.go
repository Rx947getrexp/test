// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TUserReportDayDao is the data access object for table t_user_report_day.
type TUserReportDayDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns TUserReportDayColumns // columns contains all the column names of Table for convenient usage.
}

// TUserReportDayColumns defines and stores column names for table t_user_report_day.
type TUserReportDayColumns struct {
	Id        string // 自增id
	Date      string // 数据日期, 20230101
	ChannelId string // 渠道id
	Total     string // 用户总量
	New       string // 新增用户
	Retained  string // 留存
	CreatedAt string // 记录创建时间
}

// tUserReportDayColumns holds the columns for table t_user_report_day.
var tUserReportDayColumns = TUserReportDayColumns{
	Id:        "id",
	Date:      "date",
	ChannelId: "channel_id",
	Total:     "total",
	New:       "new",
	Retained:  "retained",
	CreatedAt: "created_at",
}

// NewTUserReportDayDao creates and returns a new DAO object for table data access.
func NewTUserReportDayDao() *TUserReportDayDao {
	return &TUserReportDayDao{
		group:   "oss",
		table:   "t_user_report_day",
		columns: tUserReportDayColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TUserReportDayDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TUserReportDayDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TUserReportDayDao) Columns() TUserReportDayColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TUserReportDayDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TUserReportDayDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TUserReportDayDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
