// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TPaymentChannelDao is the data access object for table t_payment_channel.
type TPaymentChannelDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns TPaymentChannelColumns // columns contains all the column names of Table for convenient usage.
}

// TPaymentChannelColumns defines and stores column names for table t_payment_channel.
type TPaymentChannelColumns struct {
	Id                  string // 自增id
	ChannelId           string // 支付通道ID
	ChannelName         string // 支付通道名称
	IsActive            string // 支付通道是否可用，1表示可用,2表示不可用
	FreeTrialDays       string // 赠送的免费时长（以天为单位）
	TimeoutDuration     string // 订单未支付超时关闭时间（单位分钟）
	PaymentQrCode       string // 支付收款码. eg: U支付收款码
	BankCardInfo        string // 银行卡信息
	CustomerServiceInfo string // 客服信息
	MerNo               string // mer_no
	PayTypeCode         string // pay_type_code
	Weight              string // 权重，排序使用
	CreatedAt           string // 创建时间
	UpdatedAt           string // 更新时间
}

// tPaymentChannelColumns holds the columns for table t_payment_channel.
var tPaymentChannelColumns = TPaymentChannelColumns{
	Id:                  "id",
	ChannelId:           "channel_id",
	ChannelName:         "channel_name",
	IsActive:            "is_active",
	FreeTrialDays:       "free_trial_days",
	TimeoutDuration:     "timeout_duration",
	PaymentQrCode:       "payment_qr_code",
	BankCardInfo:        "bank_card_info",
	CustomerServiceInfo: "customer_service_info",
	MerNo:               "mer_no",
	PayTypeCode:         "pay_type_code",
	Weight:              "weight",
	CreatedAt:           "created_at",
	UpdatedAt:           "updated_at",
}

// NewTPaymentChannelDao creates and returns a new DAO object for table data access.
func NewTPaymentChannelDao() *TPaymentChannelDao {
	return &TPaymentChannelDao{
		group:   "speed",
		table:   "t_payment_channel",
		columns: tPaymentChannelColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TPaymentChannelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TPaymentChannelDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TPaymentChannelDao) Columns() TPaymentChannelColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TPaymentChannelDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TPaymentChannelDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TPaymentChannelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
