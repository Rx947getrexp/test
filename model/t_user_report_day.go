package model

import "time"

type TUserReportDay struct {
	Id        int64     `xorm:"pk autoincr comment('自增id') BIGINT"`
	Date      int       `xorm:"INT"`
	ChannelId int       `xorm:"INT"`
	Total     int       `xorm:"INT"`
	New       int       `xorm:"INT"`
	Retained  int       `xorm:"INT"`
	CreatedAt time.Time `xorm:"comment('创建时间') TIMESTAMP"`
}

type TUserOnlineDay struct {
	Id             int64     `xorm:"pk autoincr comment('自增id') BIGINT"`
	Date           int       `xorm:"INT"`
	Email          string    `xorm:"INT"`
	ChannelId      int       `xorm:"INT"`
	OnlineDuration int       `xorm:"INT"`
	Uplink         int64     `xorm:"INT"`
	Downlink       int64     `xorm:"INT"`
	CreatedAt      time.Time `xorm:"comment('创建时间') TIMESTAMP"`
}
