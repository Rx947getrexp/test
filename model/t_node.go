package model

import (
	"time"
)

type TNode struct {
	Id        int64     `xorm:"pk autoincr comment('自增id') BIGINT"`
	Name      string    `xorm:"not null comment('节点名称') VARCHAR(64)"`
	Title     string    `xorm:"not null comment('节点标题') VARCHAR(128)"`
	TitleEn   string    `xorm:"not null comment('节点标题（英文）') VARCHAR(128)"`
	Country   string    `xorm:"not null comment('国家') VARCHAR(64)"`
	CountryEn string    `xorm:"not null comment('国家（英文）') VARCHAR(64)"`
	Ip        string    `xorm:"not null comment('内网IP') VARCHAR(64)"`
	Server    string    `xorm:"not null comment('公网域名') VARCHAR(64)"`
	Port      int       `xorm:"not null comment('公网端口') INT"`
	Cpu       int       `xorm:"not null comment('cpu核数量（单位个）') INT"`
	Flow      int64     `xorm:"not null comment('流量带宽') BIGINT"`
	Disk      int64     `xorm:"not null comment('磁盘容量（单位B）') BIGINT"`
	Memory    int64     `xorm:"not null comment('内存大小（单位B）') BIGINT"`
	Status    int       `xorm:"not null comment('状态:1-正常；2-已软删') INT"`
	CreatedAt time.Time `xorm:"comment('创建时间') TIMESTAMP"`
	UpdatedAt time.Time `xorm:"comment('更新时间') TIMESTAMP"`
	Comment   string    `xorm:"comment('备注信息') VARCHAR(255)"`
}
