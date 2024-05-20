// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TUserVipAttrRecordDao is the data access object for table t_user_vip_attr_record.
type TUserVipAttrRecordDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns TUserVipAttrRecordColumns // columns contains all the column names of Table for convenient usage.
}

// TUserVipAttrRecordColumns defines and stores column names for table t_user_vip_attr_record.
type TUserVipAttrRecordColumns struct {
	Id              string // 自增id
	Email           string // 用户邮箱
	Source          string // 来源
	OrderNo         string // 订单号
	ExpiredTimeFrom string // 会员到期时间-原值
	ExpiredTimeTo   string // 会员到期时间-新值
	Desc            string // 记录描述
	CreatedAt       string // 创建时间
	IsRevert        string // 是否被回滚
}

// tUserVipAttrRecordColumns holds the columns for table t_user_vip_attr_record.
var tUserVipAttrRecordColumns = TUserVipAttrRecordColumns{
	Id:              "id",
	Email:           "email",
	Source:          "source",
	OrderNo:         "order_no",
	ExpiredTimeFrom: "expired_time_from",
	ExpiredTimeTo:   "expired_time_to",
	Desc:            "desc",
	CreatedAt:       "created_at",
	IsRevert:        "is_revert",
}

// NewTUserVipAttrRecordDao creates and returns a new DAO object for table data access.
func NewTUserVipAttrRecordDao() *TUserVipAttrRecordDao {
	return &TUserVipAttrRecordDao{
		group:   "speed",
		table:   "t_user_vip_attr_record",
		columns: tUserVipAttrRecordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TUserVipAttrRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TUserVipAttrRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TUserVipAttrRecordDao) Columns() TUserVipAttrRecordColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TUserVipAttrRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TUserVipAttrRecordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TUserVipAttrRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
