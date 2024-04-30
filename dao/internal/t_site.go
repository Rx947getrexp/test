// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TSiteDao is the data access object for table t_site.
type TSiteDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns TSiteColumns // columns contains all the column names of Table for convenient usage.
}

// TSiteColumns defines and stores column names for table t_site.
type TSiteColumns struct {
	Id        string // 自增id
	Site      string // 域名
	Ip        string // ip
	Status    string // 1-正常；2-软删
	Author    string // 作者
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	Comment   string // 备注信息
}

// tSiteColumns holds the columns for table t_site.
var tSiteColumns = TSiteColumns{
	Id:        "id",
	Site:      "site",
	Ip:        "ip",
	Status:    "status",
	Author:    "author",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	Comment:   "comment",
}

// NewTSiteDao creates and returns a new DAO object for table data access.
func NewTSiteDao() *TSiteDao {
	return &TSiteDao{
		group:   "speed",
		table:   "t_site",
		columns: tSiteColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TSiteDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TSiteDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TSiteDao) Columns() TSiteColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TSiteDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TSiteDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TSiteDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
