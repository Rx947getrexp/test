// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TWorkModeDao is the data access object for table t_work_mode.
type TWorkModeDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns TWorkModeColumns // columns contains all the column names of Table for convenient usage.
}

// TWorkModeColumns defines and stores column names for table t_work_mode.
type TWorkModeColumns struct {
	Id        string // 自增id
	UserId    string // 用户id
	DevId     string // 设备id
	ModeType  string // 模式类别:1-智能；2-手选
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	Comment   string // 备注信息
}

// tWorkModeColumns holds the columns for table t_work_mode.
var tWorkModeColumns = TWorkModeColumns{
	Id:        "id",
	UserId:    "user_id",
	DevId:     "dev_id",
	ModeType:  "mode_type",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	Comment:   "comment",
}

// NewTWorkModeDao creates and returns a new DAO object for table data access.
func NewTWorkModeDao() *TWorkModeDao {
	return &TWorkModeDao{
		group:   "speed",
		table:   "t_work_mode",
		columns: tWorkModeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TWorkModeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TWorkModeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TWorkModeDao) Columns() TWorkModeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TWorkModeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TWorkModeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TWorkModeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
