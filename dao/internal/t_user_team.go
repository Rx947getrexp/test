// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TUserTeamDao is the data access object for table t_user_team.
type TUserTeamDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns TUserTeamColumns // columns contains all the column names of Table for convenient usage.
}

// TUserTeamColumns defines and stores column names for table t_user_team.
type TUserTeamColumns struct {
	Id         string // 自增id
	UserId     string // 用户id
	DirectId   string // 上级id
	DirectTree string // 上级列表
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
	Comment    string // 备注信息
}

// tUserTeamColumns holds the columns for table t_user_team.
var tUserTeamColumns = TUserTeamColumns{
	Id:         "id",
	UserId:     "user_id",
	DirectId:   "direct_id",
	DirectTree: "direct_tree",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	Comment:    "comment",
}

// NewTUserTeamDao creates and returns a new DAO object for table data access.
func NewTUserTeamDao() *TUserTeamDao {
	return &TUserTeamDao{
		group:   "speed",
		table:   "t_user_team",
		columns: tUserTeamColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TUserTeamDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TUserTeamDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TUserTeamDao) Columns() TUserTeamColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TUserTeamDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TUserTeamDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TUserTeamDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
