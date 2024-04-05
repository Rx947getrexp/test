// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TUserDao is the data access object for table t_user.
type TUserDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns TUserColumns // columns contains all the column names of Table for convenient usage.
}

// TUserColumns defines and stores column names for table t_user.
type TUserColumns struct {
	Id          string // 自增id
	Uname       string // 用户名
	Passwd      string // 用户密码
	Email       string // 邮件
	Phone       string // 电话
	Level       string // 等级：0-vip0；1-vip1；2-vip2
	ExpiredTime string // vip到期时间
	V2RayUuid   string // 节点UUID
	V2RayTag    string // v2ray存在UUID标签:1-有；2-无
	Channel     string //
	ChannelId   string // 渠道id
	Status      string // 冻结状态：0-正常；1-冻结
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
	Comment     string // 备注信息
	ClientId    string //
}

// tUserColumns holds the columns for table t_user.
var tUserColumns = TUserColumns{
	Id:          "id",
	Uname:       "uname",
	Passwd:      "passwd",
	Email:       "email",
	Phone:       "phone",
	Level:       "level",
	ExpiredTime: "expired_time",
	V2RayUuid:   "v2ray_uuid",
	V2RayTag:    "v2ray_tag",
	Channel:     "channel",
	ChannelId:   "channel_id",
	Status:      "status",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	Comment:     "comment",
	ClientId:    "client_id",
}

// NewTUserDao creates and returns a new DAO object for table data access.
func NewTUserDao() *TUserDao {
	return &TUserDao{
		group:   "speed",
		table:   "t_user",
		columns: tUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TUserDao) Columns() TUserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
