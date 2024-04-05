// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TDev is the golang structure of table t_dev for DAO operations like Where/Data.
type TDev struct {
	g.Meta    `orm:"table:t_dev, do:true"`
	Id        interface{} // 自增id
	ClientId  interface{} //
	Os        interface{} // 客户端设备系统os
	IsSend    interface{} // 1-已赠送时间；2-未赠送
	Network   interface{} // 网络模式（1-自动；2-手动）
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	Comment   interface{} // 备注信息
}
