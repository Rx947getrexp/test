// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TOrderDao is the data access object for table t_order.
type TOrderDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns TOrderColumns // columns contains all the column names of Table for convenient usage.
}

// TOrderColumns defines and stores column names for table t_order.
type TOrderColumns struct {
	Id         string // 自增id
	UserId     string // 用户id
	GoodsId    string // 商品id
	Title      string // 商品标题
	Price      string // 单价(U)
	PriceCny   string // 折合RMB单价(CNY)
	Status     string // 订单状态:1-init；2-success；3-cancel
	FinishedAt string // 完成时间
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
	Comment    string // 备注信息
}

// tOrderColumns holds the columns for table t_order.
var tOrderColumns = TOrderColumns{
	Id:         "id",
	UserId:     "user_id",
	GoodsId:    "goods_id",
	Title:      "title",
	Price:      "price",
	PriceCny:   "price_cny",
	Status:     "status",
	FinishedAt: "finished_at",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	Comment:    "comment",
}

// NewTOrderDao creates and returns a new DAO object for table data access.
func NewTOrderDao() *TOrderDao {
	return &TOrderDao{
		group:   "speed",
		table:   "t_order",
		columns: tOrderColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TOrderDao) Columns() TOrderColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TOrderDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
