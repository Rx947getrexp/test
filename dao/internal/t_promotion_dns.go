// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TPromotionDnsDao is the data access object for the table t_promotion_dns.
type TPromotionDnsDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of the current DAO.
	columns TPromotionDnsColumns // columns contains all the column names of Table for convenient usage.
}

// TPromotionDnsColumns defines and stores column names for the table t_promotion_dns.
type TPromotionDnsColumns struct {
	Id             string // 自增id
	Dns            string // 域名
	Ip             string // ip地址
	MacChannel     string // 苹果电脑渠道
	WinChannel     string // windows电脑渠道
	AndroidChannel string // 安卓渠道
	Promoter       string // 推广人员
	Status         string // 状态:1-正常；2-已软删
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
	Author         string // 作者
	Comment        string // 备注信息
}

// tPromotionDnsColumns holds the columns for the table t_promotion_dns.
var tPromotionDnsColumns = TPromotionDnsColumns{
	Id:             "id",
	Dns:            "dns",
	Ip:             "ip",
	MacChannel:     "mac_channel",
	WinChannel:     "win_channel",
	AndroidChannel: "android_channel",
	Promoter:       "promoter",
	Status:         "status",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	Author:         "author",
	Comment:        "comment",
}

// NewTPromotionDnsDao creates and returns a new DAO object for table data access.
func NewTPromotionDnsDao() *TPromotionDnsDao {
	return &TPromotionDnsDao{
		group:   "speed",
		table:   "t_promotion_dns",
		columns: tPromotionDnsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TPromotionDnsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TPromotionDnsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TPromotionDnsDao) Columns() TPromotionDnsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TPromotionDnsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TPromotionDnsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TPromotionDnsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
