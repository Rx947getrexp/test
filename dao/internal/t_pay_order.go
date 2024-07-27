// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TPayOrderDao is the data access object for table t_pay_order.
type TPayOrderDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns TPayOrderColumns // columns contains all the column names of Table for convenient usage.
}

// TPayOrderColumns defines and stores column names for table t_pay_order.
type TPayOrderColumns struct {
	Id                 string // 自增id
	UserId             string // 用户uid
	Email              string // 用户邮箱
	OrderNo            string // 订单号
	PaymentChannelId   string // 支付通道ID
	GoodsId            string // 套餐ID
	OrderAmount        string // 交易金额
	Currency           string // 交易币种
	PayTypeCode        string // 支付类型编码
	Status             string // 状态:1-正常；2-已软删
	ReturnStatus       string // 支付平台返回的结果
	StatusMes          string // 状态:1-正常；2-已软删
	OrderData          string // 创建订单时支付平台返回的信息
	ResultStatus       string // 查询结果，实际订单状态
	OrderRealityAmount string // 实际交易金额
	PaymentProof       string // 支付凭证地址
	PaymentChannelErr  string // 通道错误
	CreatedAt          string // 创建时间
	UpdatedAt          string // 更新时间
	Version            string // 数据版本号
	Commission         string // 手续费
}

// tPayOrderColumns holds the columns for table t_pay_order.
var tPayOrderColumns = TPayOrderColumns{
	Id:                 "id",
	UserId:             "user_id",
	Email:              "email",
	OrderNo:            "order_no",
	PaymentChannelId:   "payment_channel_id",
	GoodsId:            "goods_id",
	OrderAmount:        "order_amount",
	Currency:           "currency",
	PayTypeCode:        "pay_type_code",
	Status:             "status",
	ReturnStatus:       "return_status",
	StatusMes:          "status_mes",
	OrderData:          "order_data",
	ResultStatus:       "result_status",
	OrderRealityAmount: "order_reality_amount",
	PaymentProof:       "payment_proof",
	PaymentChannelErr:  "payment_channel_err",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
	Version:            "version",
	Commission:         "commission",
}

// NewTPayOrderDao creates and returns a new DAO object for table data access.
func NewTPayOrderDao() *TPayOrderDao {
	return &TPayOrderDao{
		group:   "speed",
		table:   "t_pay_order",
		columns: tPayOrderColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TPayOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TPayOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TPayOrderDao) Columns() TPayOrderColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TPayOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TPayOrderDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TPayOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
