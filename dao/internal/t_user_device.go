// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TUserDeviceDao is the data access object for table t_user_device.
type TUserDeviceDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns TUserDeviceColumns // columns contains all the column names of Table for convenient usage.
}

// TUserDeviceColumns defines and stores column names for table t_user_device.
type TUserDeviceColumns struct {
	Id        string // 自增id
	UserId    string // 用户uid
	ClientId  string //
	Os        string // 客户端设备系统os
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	Kicked    string // 剔除状态, 0:正常，1:被剔除
}

// tUserDeviceColumns holds the columns for table t_user_device.
var tUserDeviceColumns = TUserDeviceColumns{
	Id:        "id",
	UserId:    "user_id",
	ClientId:  "client_id",
	Os:        "os",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	Kicked:    "kicked",
}

// NewTUserDeviceDao creates and returns a new DAO object for table data access.
func NewTUserDeviceDao() *TUserDeviceDao {
	return &TUserDeviceDao{
		group:   "speed",
		table:   "t_user_device",
		columns: tUserDeviceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TUserDeviceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TUserDeviceDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TUserDeviceDao) Columns() TUserDeviceColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TUserDeviceDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TUserDeviceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TUserDeviceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
