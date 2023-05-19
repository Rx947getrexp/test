package model

import (
	"time"
)

type TNodeUuid struct {
	Id        int64     `xorm:"pk autoincr comment('自增id') BIGINT"`
	UserId    int64     `xorm:"not null comment('用户id') BIGINT"`
	NodeId    int64     `xorm:"not null comment('节点id') BIGINT"`
	Email     string    `xorm:"comment('节点邮箱，用于区分流量') VARCHAR(64)"`
	V2rayUuid string    `xorm:"not null comment('节点UUID') VARCHAR(128)"`
	Server    string    `xorm:"not null comment('公网域名') VARCHAR(64)"`
	Port      int       `xorm:"not null comment('公网端口') INT"`
	UsedFlow  int64     `xorm:"not null comment('已使用流量（单位B）') BIGINT"`
	Status    int       `xorm:"not null comment('状态:1-正常；2-已软删') INT"`
	CreatedAt time.Time `xorm:"comment('创建时间') TIMESTAMP"`
	UpdatedAt time.Time `xorm:"comment('更新时间') TIMESTAMP"`
	Comment   string    `xorm:"comment('备注信息') VARCHAR(255)"`
}
