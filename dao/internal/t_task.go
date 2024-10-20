// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TTaskDao is the data access object for table t_task.
type TTaskDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns TTaskColumns // columns contains all the column names of Table for convenient usage.
}

// TTaskColumns defines and stores column names for table t_task.
type TTaskColumns struct {
	Id        string // 自增id
	Ip        string // 节点IP
	Date      string // 任务日期, 20230101
	UserCnt   string // 用户数量
	Status    string // 状态：0-初始状态；1-完成
	Type      string // 任务类型
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// tTaskColumns holds the columns for table t_task.
var tTaskColumns = TTaskColumns{
	Id:        "id",
	Ip:        "ip",
	Date:      "date",
	UserCnt:   "user_cnt",
	Status:    "status",
	Type:      "type",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewTTaskDao creates and returns a new DAO object for table data access.
func NewTTaskDao() *TTaskDao {
	return &TTaskDao{
		group:   "speed_collector",
		table:   "t_task",
		columns: tTaskColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TTaskDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TTaskDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TTaskDao) Columns() TTaskColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TTaskDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TTaskDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TTaskDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
