package lang

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// 简体中文
func initZhHans(tag language.Tag) {
	message.SetString(tag, "success", "成功")
	message.SetString(tag, "fail", "失败")
	message.SetString(tag, "param err", "参数错误")
	message.SetString(tag, "server err", "数据请求失败")
	message.SetString(tag, "astrict", "休息一下吧")
	message.SetString(tag, "wrong ip", "拒绝")
	message.SetString(tag, "frequently", "访问太频繁")
	message.SetString(tag, "file ext error", "文件格式错误")
}
