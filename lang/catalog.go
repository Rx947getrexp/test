package lang

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func init() {
	initEn(language.AmericanEnglish)
	initZhHans(language.SimplifiedChinese)
}

var (
	cnPrinter = message.NewPrinter(language.SimplifiedChinese)
	enPrinter = message.NewPrinter(language.AmericanEnglish)
)

func getPrinterByLang(lang string) *message.Printer {
	switch lang {
	case "cn":
		return cnPrinter
	default:
		return enPrinter
	}
}

func Translate(lang string, key string, a ...interface{}) string {
	return getPrinterByLang(lang).Sprintf(key, a...)
}
