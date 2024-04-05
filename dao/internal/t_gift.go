// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TGiftDao is the data access object for table t_gift.
type TGiftDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns TGiftColumns // columns contains all the column names of Table for convenient usage.
}

// TGiftColumns defines and stores column names for table t_gift.
type TGiftColumns struct {
	Id        string // 自增id
	UserId    string // 用户id
	OpId      string // 业务id
	OpUid     string // 业务uid
	Title     string // 赠送标题
	GiftSec   string // 赠送时间（单位s）
	GType     string // 赠送类别（1-注册；2-推荐；3-日常活动；4-充值）
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	Comment   string // 备注信息
}

// tGiftColumns holds the columns for table t_gift.
var tGiftColumns = TGiftColumns{
	Id:        "id",
	UserId:    "user_id",
	OpId:      "op_id",
	OpUid:     "op_uid",
	Title:     "title",
	GiftSec:   "gift_sec",
	GType:     "g_type",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	Comment:   "comment",
}

// NewTGiftDao creates and returns a new DAO object for table data access.
func NewTGiftDao() *TGiftDao {
	return &TGiftDao{
		group:   "speed",
		table:   "t_gift",
		columns: tGiftColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TGiftDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TGiftDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TGiftDao) Columns() TGiftColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TGiftDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TGiftDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TGiftDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
