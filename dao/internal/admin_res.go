// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminResDao is the data access object for table admin_res.
type AdminResDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns AdminResColumns // columns contains all the column names of Table for convenient usage.
}

// AdminResColumns defines and stores column names for table admin_res.
type AdminResColumns struct {
	Id        string // 资源id
	Name      string // 资源名称
	ResType   string // 类型：1-菜单；2-接口；3-按钮
	Pid       string // 上级id（没有默认为0）
	Url       string // url地址
	Sort      string // 排序
	Icon      string // 图标
	IsDel     string // 软删状态：0-未删（默认）；1-已删
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	Author    string // 修改人
}

// adminResColumns holds the columns for table admin_res.
var adminResColumns = AdminResColumns{
	Id:        "id",
	Name:      "name",
	ResType:   "res_type",
	Pid:       "pid",
	Url:       "url",
	Sort:      "sort",
	Icon:      "icon",
	IsDel:     "is_del",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	Author:    "author",
}

// NewAdminResDao creates and returns a new DAO object for table data access.
func NewAdminResDao() *AdminResDao {
	return &AdminResDao{
		group:   "speed",
		table:   "admin_res",
		columns: adminResColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AdminResDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AdminResDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AdminResDao) Columns() AdminResColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AdminResDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AdminResDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AdminResDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
