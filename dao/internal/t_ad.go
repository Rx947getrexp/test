// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TAdDao is the data access object for table t_ad.
type TAdDao struct {
	table   string     // table is the underlying table name of the DAO.
	group   string     // group is the database configuration group name of current DAO.
	columns TAdColumns // columns contains all the column names of Table for convenient usage.
}

// TAdColumns defines and stores column names for table t_ad.
type TAdColumns struct {
	Id        string // 自增id
	Status    string // 状态:1-上架；2-下架
	Sort      string // 排序
	Name      string // 广告名称
	Logo      string // 广告logo
	Link      string // 广告链接
	AdType    string // 广告分类：1-社交；2-游戏；3-漫画；4-视频...
	Tag       string // 标签标题
	Content   string // 正文介绍
	Author    string // 作者
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	Comment   string // 备注信息
}

// tAdColumns holds the columns for table t_ad.
var tAdColumns = TAdColumns{
	Id:        "id",
	Status:    "status",
	Sort:      "sort",
	Name:      "name",
	Logo:      "logo",
	Link:      "link",
	AdType:    "ad_type",
	Tag:       "tag",
	Content:   "content",
	Author:    "author",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	Comment:   "comment",
}

// NewTAdDao creates and returns a new DAO object for table data access.
func NewTAdDao() *TAdDao {
	return &TAdDao{
		group:   "speed",
		table:   "t_ad",
		columns: tAdColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TAdDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TAdDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TAdDao) Columns() TAdColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TAdDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TAdDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TAdDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
