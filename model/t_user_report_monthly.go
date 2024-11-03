package model

import (
	"time"
)

type TUserReportMonthly struct {
	Id            uint64    `xorm:"pk autoincr comment('自增id') BIGINT"`
	StatMonth     uint32    `xorm:"not null comment('统计月份') INT(11)"`
	Os            string    `xorm:"not null comment('设备类型') VARCHAR(128)"`
	UserCount     uint32    `xorm:"not null comment('用户总数') INT(11)"`
	NewUsers      uint32    `xorm:"not null comment('新增用户量') INT(11)"`
	RetainedUsers uint32    `xorm:"not null comment('次月留存') INT(11)"`
	CreatedAt     time.Time `xorm:"comment('创建时间') TIMESTAMP"`
}
