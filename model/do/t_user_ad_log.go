// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TUserAdLog is the golang structure of table t_user_ad_log for DAO operations like Where/Data.
type TUserAdLog struct {
	g.Meta     `orm:"table:t_user_ad_log, do:true"`
	Id         interface{} // 自增id
	UserId     interface{} // 用户id
	AdLocation interface{} // 广告位的位置
	AdName     interface{} // 广告名称
	DeviceType interface{} // 设备类型
	AppVersion interface{} // APP版本
	ClientId   interface{} // 设备号
	Type       interface{} //
	Content    interface{} //
	Result     interface{} //
	ReportTime interface{} // 提交时间
	CreatedAt  *gtime.Time // 记录创建时间
	AppName    interface{} // app_name
}
