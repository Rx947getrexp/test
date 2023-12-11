package model

import "time"

type TUserTraffic struct {
	Id        uint64    `xorm:"pk autoincr comment('自增id') BIGINT unsigned"`
	Email     string    `xorm:"comment('邮箱') VARCHAR(128)"`
	Ip        string    `xorm:"comment('ip地址') VARCHAR(64)"`
	Date      int       `xorm:"comment('数据日期') VARCHAR(16)"`
	Uplink    uint64    `xorm:"comment('上行流量') BIGINT unsigned"`
	Downlink  uint64    `xorm:"comment('下行流量') BIGINT unsigned"`
	CreatedAt time.Time `xorm:"comment('创建时间') TIMESTAMP"`
	UpdatedAt time.Time `xorm:"comment('更新时间') TIMESTAMP"`
}

type TUserTrafficLog struct {
	Id        uint64    `xorm:"pk autoincr comment('自增id') BIGINT unsigned"`
	Email     string    `xorm:"comment('邮箱') VARCHAR(128)"`
	Ip        string    `xorm:"comment('ip地址') VARCHAR(64)"`
	DateTime  string    `xorm:"comment('采集时间') VARCHAR(64)"`
	Uplink    uint64    `xorm:"comment('上行流量') BIGINT unsigned"`
	Downlink  uint64    `xorm:"comment('下行流量') BIGINT unsigned"`
	CreatedAt time.Time `xorm:"comment('创建时间') TIMESTAMP"`
}
