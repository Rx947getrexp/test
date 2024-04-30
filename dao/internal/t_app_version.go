// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TAppVersionDao is the data access object for table t_app_version.
type TAppVersionDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns TAppVersionColumns // columns contains all the column names of Table for convenient usage.
}

// TAppVersionColumns defines and stores column names for table t_app_version.
type TAppVersionColumns struct {
	Id        string // 自增id
	AppType   string // 1-ios;2-安卓；3-h5zip
	Version   string // 版本号
	Link      string // 超链地址
	Status    string // 状态:1-正常；2-已软删
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	Author    string // 作者
	Comment   string // 备注信息
}

// tAppVersionColumns holds the columns for table t_app_version.
var tAppVersionColumns = TAppVersionColumns{
	Id:        "id",
	AppType:   "app_type",
	Version:   "version",
	Link:      "link",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	Author:    "author",
	Comment:   "comment",
}

// NewTAppVersionDao creates and returns a new DAO object for table data access.
func NewTAppVersionDao() *TAppVersionDao {
	return &TAppVersionDao{
		group:   "speed",
		table:   "t_app_version",
		columns: tAppVersionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TAppVersionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TAppVersionDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TAppVersionDao) Columns() TAppVersionColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TAppVersionDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TAppVersionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TAppVersionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
