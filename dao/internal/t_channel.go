// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TChannelDao is the data access object for table t_channel.
type TChannelDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns TChannelColumns // columns contains all the column names of Table for convenient usage.
}

// TChannelColumns defines and stores column names for table t_channel.
type TChannelColumns struct {
	Id        string // 编号id
	Name      string // 渠道名称
	Code      string // 渠道编号
	Link      string // 渠道链接
	Status    string // 状态:1-正常；2-已软删
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	Author    string // 作者
	Comment   string // 备注
}

// tChannelColumns holds the columns for table t_channel.
var tChannelColumns = TChannelColumns{
	Id:        "id",
	Name:      "name",
	Code:      "code",
	Link:      "link",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	Author:    "author",
	Comment:   "comment",
}

// NewTChannelDao creates and returns a new DAO object for table data access.
func NewTChannelDao() *TChannelDao {
	return &TChannelDao{
		group:   "speed",
		table:   "t_channel",
		columns: tChannelColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TChannelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TChannelDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TChannelDao) Columns() TChannelColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TChannelDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TChannelDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TChannelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
