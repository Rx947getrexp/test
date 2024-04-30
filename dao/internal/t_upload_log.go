// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TUploadLogDao is the data access object for table t_upload_log.
type TUploadLogDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns TUploadLogColumns // columns contains all the column names of Table for convenient usage.
}

// TUploadLogColumns defines and stores column names for table t_upload_log.
type TUploadLogColumns struct {
	Id        string // 自增id
	UserId    string // 用户id
	DevId     string // 设备id
	Content   string // 日志内容
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	Comment   string // 备注信息
}

// tUploadLogColumns holds the columns for table t_upload_log.
var tUploadLogColumns = TUploadLogColumns{
	Id:        "id",
	UserId:    "user_id",
	DevId:     "dev_id",
	Content:   "content",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	Comment:   "comment",
}

// NewTUploadLogDao creates and returns a new DAO object for table data access.
func NewTUploadLogDao() *TUploadLogDao {
	return &TUploadLogDao{
		group:   "speed",
		table:   "t_upload_log",
		columns: tUploadLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TUploadLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TUploadLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TUploadLogDao) Columns() TUploadLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TUploadLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TUploadLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TUploadLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
