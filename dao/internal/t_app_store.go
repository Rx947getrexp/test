// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TAppStoreDao is the data access object for the table t_app_store.
type TAppStoreDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of the current DAO.
	columns TAppStoreColumns // columns contains all the column names of Table for convenient usage.
}

// TAppStoreColumns defines and stores column names for the table t_app_store.
type TAppStoreColumns struct {
	Id        string // 自增id
	TitleCn   string // 商店名称(中文)
	TitleEn   string // 商店名称(英文)
	TitleRu   string // 商店名称(俄语)
	Type      string // 商店类型，ios(苹果)，android(安卓)...
	Url       string // 商店地址
	Cover     string // 商店图标
	Status    string // 状态:1-正常；2-已软删
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	Author    string // 作者
	Comment   string // 备注信息
}

// tAppStoreColumns holds the columns for the table t_app_store.
var tAppStoreColumns = TAppStoreColumns{
	Id:        "id",
	TitleCn:   "title_cn",
	TitleEn:   "title_en",
	TitleRu:   "title_ru",
	Type:      "type",
	Url:       "url",
	Cover:     "cover",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	Author:    "author",
	Comment:   "comment",
}

// NewTAppStoreDao creates and returns a new DAO object for table data access.
func NewTAppStoreDao() *TAppStoreDao {
	return &TAppStoreDao{
		group:   "speed",
		table:   "t_app_store",
		columns: tAppStoreColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TAppStoreDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TAppStoreDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TAppStoreDao) Columns() TAppStoreColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TAppStoreDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TAppStoreDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TAppStoreDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
