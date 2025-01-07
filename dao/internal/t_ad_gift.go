// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TAdGiftDao is the data access object for table t_ad_gift.
type TAdGiftDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns TAdGiftColumns // columns contains all the column names of Table for convenient usage.
}

// TAdGiftColumns defines and stores column names for table t_ad_gift.
type TAdGiftColumns struct {
	Id           string // 自增id
	UserId       string // 用户id
	AdId         string // 广告ID
	AdName       string // 广告名称
	ExposureTime string // 单次曝光时间，单位秒
	GiftDuration string // 赠送时间
	CreatedAt    string // 创建时间
}

// tAdGiftColumns holds the columns for table t_ad_gift.
var tAdGiftColumns = TAdGiftColumns{
	Id:           "id",
	UserId:       "user_id",
	AdId:         "ad_id",
	AdName:       "ad_name",
	ExposureTime: "exposure_time",
	GiftDuration: "gift_duration",
	CreatedAt:    "created_at",
}

// NewTAdGiftDao creates and returns a new DAO object for table data access.
func NewTAdGiftDao() *TAdGiftDao {
	return &TAdGiftDao{
		group:   "speed",
		table:   "t_ad_gift",
		columns: tAdGiftColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TAdGiftDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TAdGiftDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TAdGiftDao) Columns() TAdGiftColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TAdGiftDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TAdGiftDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TAdGiftDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
