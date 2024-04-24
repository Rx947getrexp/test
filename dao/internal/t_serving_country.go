// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TServingCountryDao is the data access object for table t_serving_country.
type TServingCountryDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns TServingCountryColumns // columns contains all the column names of Table for convenient usage.
}

// TServingCountryColumns defines and stores column names for table t_serving_country.
type TServingCountryColumns struct {
	Id          string // 自增id
	Name        string // 国家名称，不可以修改，作为ID用
	Display     string // 用于在用户侧展示的国家名称
	LogoLink    string // 国家图片地址
	PingUrl     string // 前端使用
	IsRecommend string // 推荐节点1-是；2-否
	Weight      string // 权重
	Status      string // 状态:1-正常；2-已软删
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
}

// tServingCountryColumns holds the columns for table t_serving_country.
var tServingCountryColumns = TServingCountryColumns{
	Id:          "id",
	Name:        "name",
	Display:     "display",
	LogoLink:    "logo_link",
	PingUrl:     "ping_url",
	IsRecommend: "is_recommend",
	Weight:      "weight",
	Status:      "status",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewTServingCountryDao creates and returns a new DAO object for table data access.
func NewTServingCountryDao() *TServingCountryDao {
	return &TServingCountryDao{
		group:   "speed",
		table:   "t_serving_country",
		columns: tServingCountryColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TServingCountryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TServingCountryDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TServingCountryDao) Columns() TServingCountryColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TServingCountryDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TServingCountryDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TServingCountryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
