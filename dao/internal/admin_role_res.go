// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminRoleResDao is the data access object for table admin_role_res.
type AdminRoleResDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns AdminRoleResColumns // columns contains all the column names of Table for convenient usage.
}

// AdminRoleResColumns defines and stores column names for table admin_role_res.
type AdminRoleResColumns struct {
	RoleId    string // 角色id
	ResIds    string // 资源id列表
	ResTree   string // 资源菜单json树
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	Author    string // 更新人
}

// adminRoleResColumns holds the columns for table admin_role_res.
var adminRoleResColumns = AdminRoleResColumns{
	RoleId:    "role_id",
	ResIds:    "res_ids",
	ResTree:   "res_tree",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	Author:    "author",
}

// NewAdminRoleResDao creates and returns a new DAO object for table data access.
func NewAdminRoleResDao() *AdminRoleResDao {
	return &AdminRoleResDao{
		group:   "speed",
		table:   "admin_role_res",
		columns: adminRoleResColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AdminRoleResDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AdminRoleResDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AdminRoleResDao) Columns() AdminRoleResColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AdminRoleResDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AdminRoleResDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AdminRoleResDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
