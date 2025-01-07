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
	Id            string // 自增id
	Advertiser    string // 广告主，客户名称
	Name          string // 广告名称
	Type          string // 广告类型. enum: text,image,video
	Url           string // 广告内容地址
	Logo          string // logo
	SlotLocations string // 广告位的位置，包括权重
	Devices       string // 广告位的位置，包括权重
	TargetUrls    string // 跳转地址，包括：pc,ios,android
	Labels        string // 标签
	ExposureTime  string // 单次曝光时间，单位秒
	UserLevels    string // 用户等级
	StartTime     string // 广告生效时间
	EndTime       string // 广告失效时间
	Status        string // 状态:1-上架；2-下架
	GiftDuration  string // 赠送时间
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
}

// tAdColumns holds the columns for table t_ad.
var tAdColumns = TAdColumns{
	Id:            "id",
	Advertiser:    "advertiser",
	Name:          "name",
	Type:          "type",
	Url:           "url",
	Logo:          "logo",
	SlotLocations: "slot_locations",
	Devices:       "devices",
	TargetUrls:    "target_urls",
	Labels:        "labels",
	ExposureTime:  "exposure_time",
	UserLevels:    "user_levels",
	StartTime:     "start_time",
	EndTime:       "end_time",
	Status:        "status",
	GiftDuration:  "gift_duration",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
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
