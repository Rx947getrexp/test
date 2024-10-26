// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TTask is the golang structure of table t_task for DAO operations like Where/Data.
type TTask struct {
	g.Meta    `orm:"table:t_task, do:true"`
	Id        interface{} // 自增id
	Ip        interface{} // 节点IP
	Date      interface{} // 任务日期, 20230101
	UserCnt   interface{} // 用户数量
	Status    interface{} // 状态：0-初始状态；1-完成
	Type      interface{} // 任务类型
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
