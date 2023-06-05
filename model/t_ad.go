package model

import (
	"time"
)

type TAd struct {
	Id        int64     `xorm:"pk autoincr comment('自增id') BIGINT"`
	Status    int       `xorm:"not null comment('状态:1-上架；2-下架') INT"`
	Sort      int       `xorm:"not null comment('排序') INT"`
	Name      string    `xorm:"comment('广告名称') VARCHAR(255)"`
	Logo      string    `xorm:"comment('广告logo') VARCHAR(255)"`
	Link      string    `xorm:"comment('广告链接') VARCHAR(255)"`
	AdType    int       `xorm:"not null comment('广告分类：1-社交；2-游戏；3-漫画；4-视频...') INT"`
	Tag       string    `xorm:"comment('标签标题') VARCHAR(255)"`
	Content   string    `xorm:"comment('正文介绍') TEXT"`
	Author    string    `xorm:"not null comment('作者') VARCHAR(128)"`
	CreatedAt time.Time `xorm:"comment('创建时间') TIMESTAMP"`
	UpdatedAt time.Time `xorm:"comment('更新时间') TIMESTAMP"`
	Comment   string    `xorm:"comment('备注信息') VARCHAR(255)"`
}
