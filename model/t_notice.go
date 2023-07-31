package model

import (
	"time"
)

type TNotice struct {
	Id         int64     `xorm:"pk autoincr comment('自增id') BIGINT"`
	Title      string    `xorm:"not null comment('标题') VARCHAR(128)"`
	TitleEn    string    `xorm:"comment('标题（英文）') VARCHAR(128)"`
	TitleRus   string    `xorm:"comment('标题（俄文）') VARCHAR(128)"`
	Tag        string    `xorm:"comment('标签') VARCHAR(255)"`
	TagEn      string    `xorm:"comment('标签（英文）') VARCHAR(255)"`
	TagRus     string    `xorm:"comment('标签（俄文）') VARCHAR(255)"`
	Content    string    `xorm:"comment('正文内容') TEXT"`
	ContentEn  string    `xorm:"comment('正文内容（英文)') TEXT"`
	ContentRus string    `xorm:"comment('正文内容（俄文)') TEXT"`
	Author     string    `xorm:"not null comment('作者') VARCHAR(128)"`
	CreatedAt  time.Time `xorm:"comment('创建时间') TIMESTAMP"`
	UpdatedAt  time.Time `xorm:"comment('更新时间') TIMESTAMP"`
	Status     int       `xorm:"comment('状态:1-发布；2-软删') INT"`
	Comment    string    `xorm:"comment('备注信息') VARCHAR(255)"`
}
