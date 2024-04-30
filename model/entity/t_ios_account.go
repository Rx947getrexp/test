// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TIosAccount is the golang structure for table t_ios_account.
type TIosAccount struct {
	Id          int64       `description:"编号id"`
	Account     string      `description:"ios账号"`
	Pass        string      `description:"密码"`
	Name        string      `description:"别名"`
	AccountType int         `description:"1-国区；2-海外"`
	Status      int         `description:"1-正常；2-下架"`
	CreatedAt   *gtime.Time `description:"创建时间"`
	UpdatedAt   *gtime.Time `description:"更新时间"`
	Author      string      `description:"作者"`
	Comment     string      `description:"备注"`
}
