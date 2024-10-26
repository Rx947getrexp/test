package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-speed/constant"
	"go-speed/global"
	"go-speed/service/email"
	"go-speed/util"
	"strings"
	"time"
)

func isRedisKeyNil(err error) bool {
	return strings.Contains(err.Error(), "redis: nil")
}

func SendTelSms(ctx *gin.Context, mobile, ip string) error {
	var err error
	err = smsIpLimit(ctx, ip)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("发送短信失败:%s[%s]", mobile, ip)
		return err
	}
	return sendSms(ctx, mobile)
}

func smsIpLimit(ctx *gin.Context, ip string) (err error) {
	dateStr := time.Now().Format("2006-01-02")
	ipDayKey := fmt.Sprintf(constant.IpDayMsgKey, ip, dateStr)
	ipCount, err := global.Redis.Get(ctx, ipDayKey).Int64()
	if err != nil && !isRedisKeyNil(err) {
		global.MyLogger(ctx).Err(err).Msgf("redis get key[%s] failed", ipDayKey)
		return err
	}
	if ipCount >= 200 {
		err = fmt.Errorf(`IP频率限制：ipCount[%d] > IpMaxCountMsg[%d]`, ipCount, 200)
		global.MyLogger(ctx).Err(err).Msgf(`key[%s]`, ipDayKey)
		return err
	}
	return global.Redis.Set(ctx, ipDayKey, ipCount+1, time.Hour*24).Err()
}

func sendSms(ctx *gin.Context, mobile string) error {
	dateStr := time.Now().Format("2006-01-02")
	telKey := fmt.Sprintf(constant.TelMsgKey, mobile)
	telDayKey := fmt.Sprintf(constant.TelDayMsgKey, mobile, dateStr)
	telCount, err := global.Redis.Get(ctx, telDayKey).Int64()
	if err != nil && !isRedisKeyNil(err) {
		global.MyLogger(ctx).Err(err).Msgf("redis get key[%s] failed", telDayKey)
		return err
	}
	if telCount >= 100 {
		err = fmt.Errorf(`account 频率限制：Count[%d] > 100`, telCount)
		global.MyLogger(ctx).Err(err).Msgf(`key[%s]`, telDayKey)
		return err
	}
	msgCode := util.EncodeToString(6)
	//content := fmt.Sprintf(constant.SmsMsg, msgCode)
	err = global.Redis.Set(ctx, telKey, msgCode, time.Minute*6).Err()
	if err != nil {
		global.MyLogger(ctx).Err(err).Msg("redis错误")
		return err
	}
	err = global.Redis.Set(ctx, telDayKey, telCount+1, time.Hour*12).Err()
	if err != nil {
		global.MyLogger(ctx).Err(err).Msg("redis错误")
		return err
	}
	var (
		emailSubject = `VPN Короля сброс пароля`
		emailContent = "<br>【VPN Короля】</br>Код подтверждения: <font color='red'>%s</font>, действителен в течение 5 минут. Для обеспечения безопасности вашего аккаунта, пожалуйста, не раскрывайте эту информацию другим людям."
	)
	email.SetSendAccount("heronetwork@herovpn.live", "pingguoqm23", "smtpout.secureserver.net:465")
	err = email.SendEmailTLS(ctx, emailSubject, fmt.Sprintf(emailContent, msgCode), []string{mobile})
	if err == nil {
		// 发送成功，记录
		markEmailSend(ctx, mobile)
		return nil
	}

	// nc -vz smtp-mail.outlook.com 587
	// nc -vz smtp-mail.outlook.com 587
	account := map[string]map[string]string{
		//"heroesvpn@yandex.com":   {"pw": "gcdgcvwulvzftvan", "host": "smtp.yandex.com:587"}, // heroesvpn gcdgcvwulvzftvan
		"vpnheroes@outlook.com":  {"pw": "pingguoqm23", "host": "smtp-mail.outlook.com:587"},
		"heroesvpnn@outlook.com": {"pw": "pingguoqm23", "host": "smtp-mail.outlook.com:587"},
		"VPNHERO@outlook.com":    {"pw": "pingguoqm23", "host": "smtp-mail.outlook.com:587"},
	}

	accounts := []string{"vpnheroes@outlook.com", "heroesvpnn@outlook.com", "VPNHERO@outlook.com"}
	for _, userName := range accounts {
		email.SetSendAccount(userName, account[userName]["pw"], account[userName]["host"])
		err = email.SendEmail(ctx, emailSubject, fmt.Sprintf(emailContent, msgCode), []string{mobile})
		if err == nil {
			// 发送成功，记录
			markEmailSend(ctx, mobile)
			return nil
		}
	}
	return err
	//return email.SendEmail(constant.ForgetSubject, fmt.Sprintf(constant.ForgetBody, msgCode), []string{mobile})
	//return nil
	//return email.SendEmail(constant.ForgetSubject, fmt.Sprintf(constant.ForgetBody, msgCode), []string{mobile})
	//return smsService.SendMsgByKeTong(content, mobile)
}

func markEmailSend(ctx *gin.Context, email string) {
	telKey := fmt.Sprintf(constant.TelMsgKey, email+"send-flag")
	err := global.Redis.Set(ctx, telKey, "success", time.Minute*5).Err()
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("mark email(%s) send set redis failed", email)
		return
	}
}

func CheckEmailSendFlag(ctx *gin.Context, email string) bool {
	key := fmt.Sprintf(constant.TelMsgKey, email+"send-flag")
	flag, err := global.Redis.Get(ctx, key).Result()
	if err != nil {
		if isRedisKeyNil(err) {
			return false
		}
		global.MyLogger(ctx).Err(err).Msgf("redis get key failed, key: %s", key)
		return false
	}

	global.MyLogger(ctx).Debug().Msgf("%s -> %s", email, flag)
	return flag == "success"
}

func VerifyMsg(ctx *gin.Context, mobile, code string) error {
	//verifyKey := fmt.Sprintf(constant.VerifySmsKey, mobile)
	//count, _ := global.Redis.Get(ctx, verifyKey).Int64()
	//if count >= constant.VerifyCountByHour {
	//	return errors.New("验证次数受限制，请稍后再试")
	//}
	//err := global.Redis.Set(ctx, verifyKey, count+1, time.Hour).Err()
	//if err != nil {
	//	global.MyLogger(ctx).Err(err).Msg("redis连接出错")
	//	return errors.New("验证失败，请稍后再试")
	//}
	telKey := fmt.Sprintf(constant.TelMsgKey, mobile)
	msgCode, err := global.Redis.Get(ctx, telKey).Result()
	if err != nil {
		if isRedisKeyNil(err) {
			global.MyLogger(ctx).Warn().Msgf("验证码不存在，需要重新发送验证码, code:%s, msgCode:%s", code, msgCode)
			return errors.New("验证码不存在，需要重新发送验证码")
		}
		global.MyLogger(ctx).Err(err).Msg("redis连接出错")
		return errors.New("验证失败，请稍后再试")
	}
	global.MyLogger(ctx).Info().Msgf("%s [%s] [%s]", mobile, code, msgCode)
	if code != msgCode {
		global.MyLogger(ctx).Warn().Msgf("验证码不正确, code:%s, msgCode:%s", code, msgCode)
		return errors.New("验证码不正确")
	}
	return nil
}
