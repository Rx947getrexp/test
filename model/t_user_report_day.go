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
type TUserChannelDay struct {
	Id                 int64     `xorm:"pk autoincr comment('自增id') BIGINT"`
	Date               string    `xorm:"INT"`
	Channel            string    `xorm:"VARCHAR(32)"`
	Total              int       `xorm:"INT"`
	New                int       `xorm:"INT"`
	Retained           int       `xorm:"INT"`
	TotalRecharge      int       `xorm:"INT"`
	TotalRechargeMoney float64   `xorm:"decimal(10,2)"`
	NewRechargeMoney   float64   `xorm:"decimal(10,2)"`
	CreatedAt          time.Time `xorm:"comment('创建时间') TIMESTAMP"`
}
type TUserOnlineDay struct {
	Id               int64     `xorm:"pk autoincr comment('自增id') BIGINT"`
	Date             int       `xorm:"INT"`
	Email            string    `xorm:"INT"`
	Channel          string    `xorm:"VARCHAR(32)"`
	OnlineDuration   int       `xorm:"INT"`
	Uplink           int64     `xorm:"INT"`
	Downlink         int64     `xorm:"INT"`
	CreatedAt        time.Time `xorm:"comment('创建时间') TIMESTAMP"`
	LastLoginCountry string    `xorm:"comment('最后登陆国家') VARCHAR(64)"`
}
type TUserNodeDay struct {
	Id        int64     `xorm:"pk autoincr comment('自增id') BIGINT"`
	Date      int       `xorm:"INT"`
	Ip        string    `xorm:"VARCHAR(64)"`
	Total     int       `xorm:"INT"`
	New       int       `xorm:"INT"`
	Retained  int       `xorm:"INT"`
	CreatedAt time.Time `xorm:"comment('创建时间') TIMESTAMP"`
}
type TUserNodeOnlineDay struct {
	Id             int64     `xorm:"pk autoincr comment('自增id') BIGINT"`
	Date           int       `xorm:"INT"`
	Email          string    `xorm:"INT"`
	Channel        string    `xorm:"VARCHAR(32)"`
	OnlineDuration int       `xorm:"INT"`
	Uplink         int64     `xorm:"INT"`
	Downlink       int64     `xorm:"INT"`
	Node           string    `xorm:"VARCHAR(64)"`
	RegisterDate   time.Time `xorm:"comment('创建时间') TIMESTAMP"`
	CreatedAt      time.Time `xorm:"comment('创建时间') TIMESTAMP"`
}
type TUserRechargeReportDay struct {
	Id        int64     `xorm:"pk autoincr comment('自增id') BIGINT"`
	Date      int       `xorm:"INT"`
	GoodsId   int       `xorm:"INT"`
	Total     int       `xorm:"INT"`
	New       int       `xorm:"INT"`
	CreatedAt time.Time `xorm:"comment('创建时间') TIMESTAMP"`
}
type TUserRechargeTimesReportDay struct {
	Id        int64     `xorm:"pk autoincr comment('自增id') BIGINT"`
	Date      int       `xorm:"INT"`
	GoodsId   int       `xorm:"INT"`
	Total     int       `xorm:"INT"`
	New       int       `xorm:"INT"`
	CreatedAt time.Time `xorm:"comment('创建时间') TIMESTAMP"`
}
type TUserChannelRechargeDay struct {
	Id        int64     `xorm:"pk autoincr comment('自增id') BIGINT"`
	Date      int       `xorm:"INT"`
	Channel   string    `xorm:"VARCHAR(32)"`
	GoodsName string    `xorm:"VARCHAR(32)"`
	UsdTotal  int       `xorm:"INT"`
	UsdNew    int       `xorm:"INT"`
	RubTotal  int       `xorm:"INT"`
	RubNew    int       `xorm:"INT"`
	CreatedAt time.Time `xorm:"comment('创建时间') TIMESTAMP"`
}
