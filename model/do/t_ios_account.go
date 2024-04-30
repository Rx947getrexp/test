// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TIosAccount is the golang structure of table t_ios_account for DAO operations like Where/Data.
type TIosAccount struct {
	g.Meta      `orm:"table:t_ios_account, do:true"`
	Id          interface{} // 编号id
	Account     interface{} // ios账号
	Pass        interface{} // 密码
	Name        interface{} // 别名
	AccountType interface{} // 1-国区；2-海外
	Status      interface{} // 1-正常；2-下架
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	Author      interface{} // 作者
	Comment     interface{} // 备注
}
