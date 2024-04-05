// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TNoticeDao is the data access object for table t_notice.
type TNoticeDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns TNoticeColumns // columns contains all the column names of Table for convenient usage.
}

// TNoticeColumns defines and stores column names for table t_notice.
type TNoticeColumns struct {
	Id         string // 自增id
	Title      string // 标题
	TitleEn    string // 标题（英文）
	TitleRus   string // 标题（俄文）
	Tag        string // 标签
	TagEn      string // 标签（英文）
	TagRus     string // 标签（俄文）
	Content    string // 正文内容
	ContentEn  string // 正文内容（英文）
	ContentRus string // 正文内容（俄文）
	Author     string // 作者
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
	Status     string // 状态:1-发布；2-软删
	Comment    string // 备注信息
}

// tNoticeColumns holds the columns for table t_notice.
var tNoticeColumns = TNoticeColumns{
	Id:         "id",
	Title:      "title",
	TitleEn:    "title_en",
	TitleRus:   "title_rus",
	Tag:        "tag",
	TagEn:      "tag_en",
	TagRus:     "tag_rus",
	Content:    "content",
	ContentEn:  "content_en",
	ContentRus: "content_rus",
	Author:     "author",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	Status:     "status",
	Comment:    "comment",
}

// NewTNoticeDao creates and returns a new DAO object for table data access.
func NewTNoticeDao() *TNoticeDao {
	return &TNoticeDao{
		group:   "speed",
		table:   "t_notice",
		columns: tNoticeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TNoticeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TNoticeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TNoticeDao) Columns() TNoticeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TNoticeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TNoticeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TNoticeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
