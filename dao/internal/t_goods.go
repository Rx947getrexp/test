// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TGoodsDao is the data access object for table t_goods.
type TGoodsDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns TGoodsColumns // columns contains all the column names of Table for convenient usage.
}

// TGoodsColumns defines and stores column names for table t_goods.
type TGoodsColumns struct {
	Id               string // 自增id
	MType            string // 会员类型：1-vip1；2-vip2
	Title            string // 套餐标题
	TitleEn          string // 套餐标题（英文）
	TitleRus         string // 套餐标题（俄文）
	Price            string // 单价(U)
	PriceUnit        string // 价格单位
	Period           string // 有效期（天）
	DevLimit         string // 设备限制数
	FlowLimit        string // 流量限制数；单位：字节；0-不限制
	IsDiscount       string // 是否优惠:1-是；2-否
	Low              string // 最低赠送(天)
	High             string // 最高赠送(天)
	Status           string // 状态:1-正常；2-已软删
	CreatedAt        string // 创建时间
	UpdatedAt        string // 更新时间
	Author           string // 作者
	Comment          string // 备注信息
	UsdPayPrice      string // usd_pay价格(U)
	UsdPriceUnit     string // USD支付的价格单位
	WebmoneyPayPrice string // webmoney价格(wmz)
	PriceRub         string // 卢布价格(RUB)
	PriceWmz         string // WMZ价格(WMZ)
	PriceUsd         string // USD价格(USD)
	PriceUah         string // UAH价格(UAH)
}

// tGoodsColumns holds the columns for table t_goods.
var tGoodsColumns = TGoodsColumns{
	Id:               "id",
	MType:            "m_type",
	Title:            "title",
	TitleEn:          "title_en",
	TitleRus:         "title_rus",
	Price:            "price",
	PriceUnit:        "price_unit",
	Period:           "period",
	DevLimit:         "dev_limit",
	FlowLimit:        "flow_limit",
	IsDiscount:       "is_discount",
	Low:              "low",
	High:             "high",
	Status:           "status",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	Author:           "author",
	Comment:          "comment",
	UsdPayPrice:      "usd_pay_price",
	UsdPriceUnit:     "usd_price_unit",
	WebmoneyPayPrice: "webmoney_pay_price",
	PriceRub:         "price_rub",
	PriceWmz:         "price_wmz",
	PriceUsd:         "price_usd",
	PriceUah:         "price_uah",
}

// NewTGoodsDao creates and returns a new DAO object for table data access.
func NewTGoodsDao() *TGoodsDao {
	return &TGoodsDao{
		group:   "speed",
		table:   "t_goods",
		columns: tGoodsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TGoodsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TGoodsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TGoodsDao) Columns() TGoodsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TGoodsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TGoodsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TGoodsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
