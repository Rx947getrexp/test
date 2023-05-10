package lang

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// 英文
func initEn(tag language.Tag) {
	message.SetString(tag, "success", "Success")
	message.SetString(tag, "fail", "Fail")
	message.SetString(tag, "param err", "Param err")
	message.SetString(tag, "server err", "Data request failed")
	message.SetString(tag, "astrict", "Take a break")
	message.SetString(tag, "wrong ip", "Refused")
	message.SetString(tag, "frequently", "Visit too often")
	message.SetString(tag, "file ext error", "File format error")
}
