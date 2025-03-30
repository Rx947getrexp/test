package app

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"

	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/response"
)

type AppVersionReq struct {
}

type AppVersionRes struct {
	Type        int    `json:"type" dc:"0:无需升级，1：建议升级，2：需要升级"`
	Description string `json:"description" dc:"描述"`
}

func AppVersion(ctx *gin.Context) {
	var (
		lang  = strings.ToLower(global.GetLang(ctx))
		level = calcLevel(ctx)
	)

	global.MyLogger(ctx).Debug().Msgf("lang: %s, level: %d", lang, level)
	if lang != i18n.LangCN && lang != i18n.LangEN && lang != i18n.LangRU {
		lang = i18n.LangRU
	}

	if level != 0 && level != 1 && level != 2 {
		level = -1
	}

	res := AppVersionRes{
		Type:        level,
		Description: messages[level][lang],
	}
	response.RespOk(ctx, i18n.RetMsgSuccess, res)
}

func calcLevel(ctx *gin.Context) int {
	var (
		l1 = global.Config.System.AppVersionL1
		l2 = global.Config.System.AppVersionL2

		appVersion = global.GetAppVersion(ctx)
	)
	global.MyLogger(ctx).Debug().Msgf("l1: %s, l2: %s, appVersion: %s", l1, l2, appVersion)
	// 转换字符串为浮点数
	appVer, err := strconv.ParseFloat(appVersion, 64)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("invalid app_version: %s", appVersion)
		return -1
	}

	l1Float, err := strconv.ParseFloat(l1, 64)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("invalid l1: %s", l1)
		return -1
	}

	l2Float, err := strconv.ParseFloat(l2, 64)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("invalid l2: %s", l2)
		return -1
	}

	// 确定级别
	level := 0
	switch {
	case appVer >= l2Float:
		level = 0

	case appVer >= l1Float:
		level = 1

	default:
		level = 2
	}
	return level
}

// 多语言消息映射表
var messages = map[int]map[string]string{
	-1: {
		"zh": "Unknown version",
		"en": "Unknown version",
		"ru": "Неизвестная версия",
	},
	0: {
		"zh": "APP版本是当前最新版本",
		"en": "Your app is on the latest version",
		"ru": "Ваше приложение использует актуальную версию",
	},
	1: {
		"zh": "APP版本不是当前最新版本，建议升级到最新版本",
		"en": "Your app is not up-to-date, recommended to upgrade",
		"ru": "Ваша версия приложения устарела, рекомендуется обновление",
	},
	2: {
		"zh": "APP版本过低，需要升级到最新版本",
		"en": "App version is too low, immediate upgrade required",
		"ru": "Версия приложения слишком старая, требуется срочное обновление",
	},
}
