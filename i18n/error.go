package i18n

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-speed/global"
)

func I18nTrans(c *gin.Context, msg string) string {
	lang := c.GetHeader("Lang")
	if lang == "rus" {
		lang = "ru"
	}
	claims, claimsExist := c.Get("claims")
	global.MyLogger(c).Err(fmt.Errorf(msg)).Msgf("lang: %s, msg: %s, Client-Id: %+v, Dev-Id: %+v, claims: %+v, claimsExist: %+v",
		lang, msg, c.GetHeader("Client-Id"), c.GetHeader("Dev-Id"), claims, claimsExist)
	m, ok := ReturnMsgMap[msg]
	if !ok {
		return unknownError(lang)
	}
	if lang == LangCN {
		return msg
	}
	ret, ok := m[lang]
	global.MyLogger(c).Info().Msgf("ok: %v, ret: %v", ok, ret)
	if !ok {
		return unknownError(lang)
	}
	return ret
}

func unknownError(lang string) string {
	global.Logger.Info().Msgf(">>>>>>>>>>> lang: %s", lang)
	if lang == LangRU {
		return "Запрос не удался..."
	}
	if lang == LangCN {
		return "查询失败..."
	}
	return "Query failed..."
}

type I18nMsgMap map[string]map[string]string

var ReturnMsgMap I18nMsgMap

const (
	LangCN = "cn" // 中文
	LangEN = "en" // 英语
	LangRU = "ru" // 俄语

	RetMsgSuccess                   = "成功"
	RetMsgDBErr                     = "查询失败" //"数据库访问失败。"
	RetMsgDevIdNotExitsErr          = "设备号不存在。"
	RetMsgParamParseErr             = "参数解析失败。"
	RetMsgParamInvalid              = "参数错误。"
	RetMsgEmailNotReg               = "邮箱地址未注册。"
	RetMsgVerifyCodeSendFail        = "发送验证码失败,请稍后再试。"
	RetMsgSendSuccess               = "发送成功。"
	RetMsgParamInputInvalid         = "请检查输入参数。"
	RetMsgTwoPasswordNotMatch       = "两次输入的密码不一致。"
	RetMsgEmailHasRegErr            = "该邮箱已注册。"
	RetMsgRegFailed                 = "注册失败。"
	RetMsgReferrerIDIncorrect       = "推荐人ID不正确。"
	RetMsgRegSuccess                = "注册成功。"
	RetMsgAccountNotExist           = "账号不存在。"
	RetMsgAccountPasswordEmptyErr   = "账号、密码都不可以为空。"
	RetMsgAccountPasswordIncorrect  = "用户名或密码不正确。"
	RetMsgReachedDevicesLimit       = "达到登录设备上限。"
	RetMsgLoginError                = "登录出错。"
	RetMsgAuthFailed                = "用户鉴权失败，请重新登陆！"
	RetMsgOperateFailed             = "操作失败。"
	RetMsgVerificationCodeErr       = "验证码错误。"
	RetMsgQueryResultIsEmpty        = "查询结果为空。"
	RetMsgActivity3TimesLimits      = "每天参与活动限制3次。"
	RetMsgDeviceAuthFailed          = "设备鉴权失败。"
	RetMsgUploadLogFailed           = "上传日志失败。"
	RetMsgDealCreateFailed          = "创建订单失败。"
	RetMsgRemoveDevFailed           = "踢除设备失败。"
	RetMsgAccountExpired            = "账户已过期。"
	RetMsgLogoutFailed              = "注销失败。"
	RetMsgAuthExpired               = "授权已过期，请重新登陆！"
	RetMsgDevIdInvalid              = "Dev-Id无效"
	RetMsgUserIdInvalid             = "User-Id无效"
	RetMsgAuthorizationTokenInvalid = "Token无效"
)

func Init() {
	ReturnMsgMap = make(I18nMsgMap)
	ReturnMsgMap[RetMsgDBErr] = map[string]string{
		LangEN: "Query failed.",
		LangRU: "Запрос не удался.",
	}
	ReturnMsgMap[RetMsgDevIdNotExitsErr] = map[string]string{
		LangEN: "DeviceId not exist.",
		LangRU: "Номер устройства не существует.",
	}
	ReturnMsgMap[RetMsgParamParseErr] = map[string]string{
		LangEN: "Parameter parsing failed.",
		LangRU: "Параметрический анализ не работает.",
	}
	ReturnMsgMap[RetMsgParamInvalid] = map[string]string{
		LangEN: "Parameter error.",
		LangRU: "Ошибка параметра.",
	}
	ReturnMsgMap[RetMsgEmailNotReg] = map[string]string{
		LangEN: "Email is not registered.",
		LangRU: "Почтовый адрес не зарегистрирован.",
	}
	ReturnMsgMap[RetMsgVerifyCodeSendFail] = map[string]string{
		LangEN: "Failed to send the verification code, please try again later.",
		LangRU: "Отправка кода не удалась, пожалуйста, попробуйте позже.",
	}
	ReturnMsgMap[RetMsgSendSuccess] = map[string]string{
		LangEN: "Send succeeded.",
		LangRU: "Отправить успешно.",
	}
	ReturnMsgMap[RetMsgParamInputInvalid] = map[string]string{
		LangEN: "Please check the input parameters.",
		LangRU: "Пожалуйста, проверьте входные параметры.",
	}
	ReturnMsgMap[RetMsgTwoPasswordNotMatch] = map[string]string{
		LangEN: "The two passwords entered do not match.",
		LangRU: "Введенные пароли не совпадают.",
	}
	ReturnMsgMap[RetMsgEmailHasRegErr] = map[string]string{
		LangEN: "This email has already been registered.",
		LangRU: "Этот адрес электронной почты уже зарегистрирован.",
	}
	ReturnMsgMap[RetMsgRegFailed] = map[string]string{
		LangEN: "Registration failed.",
		LangRU: "Регистрация не удалась.",
	}
	ReturnMsgMap[RetMsgReferrerIDIncorrect] = map[string]string{
		LangEN: "The referrer ID is incorrect.",
		LangRU: "ID рекомендателя неверен.",
	}
	ReturnMsgMap[RetMsgRegSuccess] = map[string]string{
		LangEN: "Registration successful.",
		LangRU: "Регистрация успешно завершена.",
	}
	ReturnMsgMap[RetMsgAccountNotExist] = map[string]string{
		LangEN: "Account does not exist.",
		LangRU: "Учетная запись не существует.",
	}
	ReturnMsgMap[RetMsgAccountPasswordEmptyErr] = map[string]string{
		LangEN: "Both account and password cannot be empty.",
		LangRU: "Учетная запись и пароль не могут быть пустыми.",
	}
	ReturnMsgMap[RetMsgAccountPasswordIncorrect] = map[string]string{
		LangEN: "Incorrect username or password.",
		LangRU: "Неправильное имя пользователя или пароль.",
	}
	ReturnMsgMap[RetMsgReachedDevicesLimit] = map[string]string{
		LangEN: "Reached the limit of login devices.",
		LangRU: "Достигнут предел количества устройств для входа.",
	}
	ReturnMsgMap[RetMsgLoginError] = map[string]string{
		LangEN: "Login error.",
		LangRU: "Ошибка входа.",
	}
	ReturnMsgMap[RetMsgAuthFailed] = map[string]string{
		LangEN: "User authentication failed, please login again.",
		LangRU: "Аутентификация пользователя не удалась, пожалуйста, войдите снова.",
	}
	ReturnMsgMap[RetMsgOperateFailed] = map[string]string{
		LangEN: "Operation failed.",
		LangRU: "Операция не удалась.",
	}
	ReturnMsgMap[RetMsgVerificationCodeErr] = map[string]string{
		LangEN: "Verification code error.",
		LangRU: "Ошибка проверочного кода.",
	}
	ReturnMsgMap[RetMsgQueryResultIsEmpty] = map[string]string{
		LangEN: "The query result is empty.",
		LangRU: "Результат запроса пуст.",
	}
	ReturnMsgMap[RetMsgActivity3TimesLimits] = map[string]string{
		LangEN: "You can participate in the activity up to 3 times per day.",
		LangRU: "Вы можете принимать участие в мероприятии не более 3 раз в день.",
	}
	ReturnMsgMap[RetMsgDeviceAuthFailed] = map[string]string{
		LangEN: "Device authentication failed.",
		LangRU: "Ошибка аутентификации устройства.",
	}
	ReturnMsgMap[RetMsgUploadLogFailed] = map[string]string{
		LangEN: "Failed to upload log.",
		LangRU: "Не удалось загрузить журнал.",
	}
	ReturnMsgMap[RetMsgDealCreateFailed] = map[string]string{
		LangEN: "Order creation failed.",
		LangRU: "Не удалось создать заказ.",
	}
	ReturnMsgMap[RetMsgRemoveDevFailed] = map[string]string{
		LangEN: "Failed to remove the device.",
		LangRU: "Не удалось удалить устройство.",
	}
	ReturnMsgMap[RetMsgAccountExpired] = map[string]string{
		LangEN: "The account has expired.",
		LangRU: "Срок действия учетной записи истек.",
	}
	ReturnMsgMap[RetMsgLogoutFailed] = map[string]string{
		LangEN: "Logout failed.",
		LangRU: "Не удалось выйти из системы.",
	}
	ReturnMsgMap[RetMsgAuthExpired] = map[string]string{
		LangEN: "Auth expired. please login again.",
		LangRU: "Авторизация истекла. пожалуйста, войдите снова.",
	}
	ReturnMsgMap[RetMsgDevIdInvalid] = map[string]string{
		LangEN: "Dev-Id invalid.",
		LangRU: "Dev-Id недействителен.",
	}
	ReturnMsgMap[RetMsgUserIdInvalid] = map[string]string{
		LangEN: "User-Id invalid.",
		LangRU: "User-Id недействителен.",
	}
	ReturnMsgMap[RetMsgAuthorizationTokenInvalid] = map[string]string{
		LangEN: "Token invalid.",
		LangRU: "Token недействителен.",
	}

}
