// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TAppDnsDao is the data access object for table t_app_dns.
type TAppDnsDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns TAppDnsColumns // columns contains all the column names of Table for convenient usage.
}

// TAppDnsColumns defines and stores column names for table t_app_dns.
type TAppDnsColumns struct {
	Id        string // 自增id
	SiteType  string // 站点类型:1-appapi域名；2-offsite域名；3-onlineservice在线客服；4-管理后台...
	Dns       string // 域名
	Ip        string // ip地址
	Level     string // 线路级别:1,2,3...用于白名单机制
	Status    string // 状态:1-正常；2-已软删
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	Author    string // 作者
	Comment   string // 备注信息
}

// tAppDnsColumns holds the columns for table t_app_dns.
var tAppDnsColumns = TAppDnsColumns{
	Id:        "id",
	SiteType:  "site_type",
	Dns:       "dns",
	Ip:        "ip",
	Level:     "level",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	Author:    "author",
	Comment:   "comment",
}

// NewTAppDnsDao creates and returns a new DAO object for table data access.
func NewTAppDnsDao() *TAppDnsDao {
	return &TAppDnsDao{
		group:   "speed",
		table:   "t_app_dns",
		columns: tAppDnsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TAppDnsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TAppDnsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TAppDnsDao) Columns() TAppDnsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TAppDnsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TAppDnsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TAppDnsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
