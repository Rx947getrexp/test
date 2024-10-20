// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TTask is the golang structure for table t_task.
type TTask struct {
	Id        int64       `description:"自增id"`
	Ip        string      `description:"节点IP"`
	Date      uint        `description:"任务日期, 20230101"`
	UserCnt   uint        `description:"用户数量"`
	Status    uint        `description:"状态：0-初始状态；1-完成"`
	Type      string      `description:"任务类型"`
	CreatedAt *gtime.Time `description:"创建时间"`
	UpdatedAt *gtime.Time `description:"更新时间"`
}
