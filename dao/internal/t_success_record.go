// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TSuccessRecordDao is the data access object for table t_success_record.
type TSuccessRecordDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns TSuccessRecordColumns // columns contains all the column names of Table for convenient usage.
}

// TSuccessRecordColumns defines and stores column names for table t_success_record.
type TSuccessRecordColumns struct {
	Id         string // 自增id
	UserId     string // 用户id
	OrderId    string // 订单id
	StartTime  string // 本次计费开始时间戳
	EndTime    string // 本次计费结束时间戳
	SurplusSec string // 剩余时长(s)
	TotalSec   string // 订单总时长(s）
	GoodsDay   string // 套餐天数
	SendDay    string // 赠送天数
	PayType    string // 订单状态:1-银行卡；2-支付宝；3-微信支付
	Status     string // 1-using使用中；2-wait等待; 3-end已结束
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
	Comment    string // 备注信息
}

// tSuccessRecordColumns holds the columns for table t_success_record.
var tSuccessRecordColumns = TSuccessRecordColumns{
	Id:         "id",
	UserId:     "user_id",
	OrderId:    "order_id",
	StartTime:  "start_time",
	EndTime:    "end_time",
	SurplusSec: "surplus_sec",
	TotalSec:   "total_sec",
	GoodsDay:   "goods_day",
	SendDay:    "send_day",
	PayType:    "pay_type",
	Status:     "status",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	Comment:    "comment",
}

// NewTSuccessRecordDao creates and returns a new DAO object for table data access.
func NewTSuccessRecordDao() *TSuccessRecordDao {
	return &TSuccessRecordDao{
		group:   "speed",
		table:   "t_success_record",
		columns: tSuccessRecordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TSuccessRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TSuccessRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TSuccessRecordDao) Columns() TSuccessRecordColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TSuccessRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TSuccessRecordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TSuccessRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
