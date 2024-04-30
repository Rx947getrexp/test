// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TGift is the golang structure for table t_gift.
type TGift struct {
	Id        int64       `description:"自增id"`
	UserId    int64       `description:"用户id"`
	OpId      string      `description:"业务id"`
	OpUid     int64       `description:"业务uid"`
	Title     string      `description:"赠送标题"`
	GiftSec   int         `description:"赠送时间（单位s）"`
	GType     int         `description:"赠送类别（1-注册；2-推荐；3-日常活动；4-充值）"`
	CreatedAt *gtime.Time `description:"创建时间"`
	UpdatedAt *gtime.Time `description:"更新时间"`
	Comment   string      `description:"备注信息"`
}
