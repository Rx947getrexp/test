// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TUserVipAttrRecordsDao is the data access object for table t_user_vip_attr_records.
type TUserVipAttrRecordsDao struct {
	table   string                     // table is the underlying table name of the DAO.
	group   string                     // group is the database configuration group name of current DAO.
	columns TUserVipAttrRecordsColumns // columns contains all the column names of Table for convenient usage.
}

// TUserVipAttrRecordsColumns defines and stores column names for table t_user_vip_attr_records.
type TUserVipAttrRecordsColumns struct {
	Id          string // 自增id
	Email       string // 用户邮箱
	Source      string // 来源
	OrderNo     string // 订单号
	ExpiredTime string // 会员到期时间
	Desc        string // 记录描述
	CreatedAt   string // 创建时间
}

// tUserVipAttrRecordsColumns holds the columns for table t_user_vip_attr_records.
var tUserVipAttrRecordsColumns = TUserVipAttrRecordsColumns{
	Id:          "id",
	Email:       "email",
	Source:      "source",
	OrderNo:     "order_no",
	ExpiredTime: "expired_time",
	Desc:        "desc",
	CreatedAt:   "created_at",
}

// NewTUserVipAttrRecordsDao creates and returns a new DAO object for table data access.
func NewTUserVipAttrRecordsDao() *TUserVipAttrRecordsDao {
	return &TUserVipAttrRecordsDao{
		group:   "speed",
		table:   "t_user_vip_attr_records",
		columns: tUserVipAttrRecordsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TUserVipAttrRecordsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TUserVipAttrRecordsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TUserVipAttrRecordsDao) Columns() TUserVipAttrRecordsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TUserVipAttrRecordsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TUserVipAttrRecordsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TUserVipAttrRecordsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
