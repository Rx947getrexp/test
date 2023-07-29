package lang

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// 俄语
func initRussian(tag language.Tag) {
	message.SetString(tag, "success", "успех")
	message.SetString(tag, "fail", "неудача")
	message.SetString(tag, "param err", "Ошибка параметра")
	message.SetString(tag, "server err", "Ошибка сети")
	message.SetString(tag, "astrict", "сделать перерыв")
	message.SetString(tag, "wrong ip", "отклонять")
	message.SetString(tag, "frequently", "навещать слишком часто")
	message.SetString(tag, "file ext error", "неправильный формат файла")
}
