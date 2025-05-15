// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TUserOpLogDao is the data access object for table t_user_op_log.
type TUserOpLogDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns TUserOpLogColumns // columns contains all the column names of Table for convenient usage.
}

// TUserOpLogColumns defines and stores column names for table t_user_op_log.
type TUserOpLogColumns struct {
	Id           string // 自增id
	Email        string // 用户账号
	DeviceId     string // 设备ID
	DeviceType   string // 设备类型
	PageName     string // page_name
	Result       string // result
	Content      string // content
	Version      string //
	CreateTime   string // 提交时间
	CreatedAt    string // 记录创建时间
	InterfaceUrl string // 接口地址
	ServerCode   string // 后端状态码
	HttpCode     string // HTTP状态码
	TraceId      string // TraceId
	UserId       string // 用户uid
	AppName      string // app_name
}

// tUserOpLogColumns holds the columns for table t_user_op_log.
var tUserOpLogColumns = TUserOpLogColumns{
	Id:           "id",
	Email:        "email",
	DeviceId:     "device_id",
	DeviceType:   "device_type",
	PageName:     "page_name",
	Result:       "result",
	Content:      "content",
	Version:      "version",
	CreateTime:   "create_time",
	CreatedAt:    "created_at",
	InterfaceUrl: "interface_url",
	ServerCode:   "server_code",
	HttpCode:     "http_code",
	TraceId:      "trace_id",
	UserId:       "user_id",
	AppName:      "app_name",
}

// NewTUserOpLogDao creates and returns a new DAO object for table data access.
func NewTUserOpLogDao() *TUserOpLogDao {
	return &TUserOpLogDao{
		group:   "speed-report",
		table:   "t_user_op_log",
		columns: tUserOpLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TUserOpLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TUserOpLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TUserOpLogDao) Columns() TUserOpLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TUserOpLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TUserOpLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TUserOpLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
