// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminUserRoleDao is the data access object for table admin_user_role.
type AdminUserRoleDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns AdminUserRoleColumns // columns contains all the column names of Table for convenient usage.
}

// AdminUserRoleColumns defines and stores column names for table admin_user_role.
type AdminUserRoleColumns struct {
	Id        string // 自增序列
	Uid       string // 用户id
	RoleId    string // 角色id
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	IsDel     string // 软删：0-未删；1-已删
	Author    string // 更新人
}

// adminUserRoleColumns holds the columns for table admin_user_role.
var adminUserRoleColumns = AdminUserRoleColumns{
	Id:        "id",
	Uid:       "uid",
	RoleId:    "role_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	IsDel:     "is_del",
	Author:    "author",
}

// NewAdminUserRoleDao creates and returns a new DAO object for table data access.
func NewAdminUserRoleDao() *AdminUserRoleDao {
	return &AdminUserRoleDao{
		group:   "speed",
		table:   "admin_user_role",
		columns: adminUserRoleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AdminUserRoleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AdminUserRoleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AdminUserRoleDao) Columns() AdminUserRoleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AdminUserRoleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AdminUserRoleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AdminUserRoleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
