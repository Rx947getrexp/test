// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TUserAdLogDao is the data access object for table t_user_ad_log.
type TUserAdLogDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns TUserAdLogColumns // columns contains all the column names of Table for convenient usage.
}

// TUserAdLogColumns defines and stores column names for table t_user_ad_log.
type TUserAdLogColumns struct {
	Id         string // 自增id
	UserId     string // 用户id
	AdLocation string // 广告位的位置
	AdName     string // 广告名称
	DeviceType string // 设备类型
	AppVersion string // APP版本
	ClientId   string // 设备号
	Type       string //
	Content    string //
	Result     string //
	ReportTime string // 提交时间
	CreatedAt  string // 记录创建时间
	AppName    string // app_name
}

// tUserAdLogColumns holds the columns for table t_user_ad_log.
var tUserAdLogColumns = TUserAdLogColumns{
	Id:         "id",
	UserId:     "user_id",
	AdLocation: "ad_location",
	AdName:     "ad_name",
	DeviceType: "device_type",
	AppVersion: "app_version",
	ClientId:   "client_id",
	Type:       "type",
	Content:    "content",
	Result:     "result",
	ReportTime: "report_time",
	CreatedAt:  "created_at",
	AppName:    "app_name",
}

// NewTUserAdLogDao creates and returns a new DAO object for table data access.
func NewTUserAdLogDao() *TUserAdLogDao {
	return &TUserAdLogDao{
		group:   "speed-report",
		table:   "t_user_ad_log",
		columns: tUserAdLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TUserAdLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TUserAdLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TUserAdLogDao) Columns() TUserAdLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TUserAdLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TUserAdLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TUserAdLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
