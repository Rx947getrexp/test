package service

import (
	"context"
	"errors"
	"fmt"
	"go-speed/constant"
	"go-speed/global"
	"go-speed/util"
	"time"
)

func SendTelSms(mobile, ip string) error {
	var err error
	err = smsIpLimit(ip)
	if err != nil {
		global.Logger.Err(err).Msgf("发送短信失败:%s[%s]", mobile, ip)
		return err
	}
	return sendSms(mobile)
}

func smsIpLimit(ip string) error {
	dateStr := time.Now().Format("2006-01-02")
	ipDayKey := fmt.Sprintf(constant.IpDayMsgKey, ip, dateStr)
	ipCount, _ := global.Redis.Get(context.Background(), ipDayKey).Int64()
	if ipCount >= constant.IpMaxCountMsg {
		return errors.New("IP短信限制，请稍后再试！")
	}
	return global.Redis.Set(context.Background(), ipDayKey, ipCount+1, time.Hour*24).Err()
}

func sendSms(mobile string) error {
	dateStr := time.Now().Format("2006-01-02")
	telKey := fmt.Sprintf(constant.TelMsgKey, mobile)
	telDayKey := fmt.Sprintf(constant.TelDayMsgKey, mobile, dateStr)
	telCount, _ := global.Redis.Get(context.Background(), telDayKey).Int64()
	if telCount >= constant.MaxCountMsg {
		return errors.New("短信限制，请稍后再试！")
	}
	msgCode := util.EncodeToString(6)
	content := fmt.Sprintf(constant.SmsMsg, msgCode)
	err := global.Redis.Set(context.Background(), telKey, msgCode, time.Minute*5).Err()
	if err != nil {
		global.Logger.Err(err).Msg("redis错误")
		return errors.New("短信发送失败")
	}
	err = global.Redis.Set(context.Background(), telDayKey, telCount+1, time.Hour*12).Err()
	if err != nil {
		global.Logger.Err(err).Msg("redis错误")
		return errors.New("短信发送失败")
	}
	fmt.Sprint(content)
	return nil
	//return smsService.SendMsgByKeTong(content, mobile)
}

func VerifyMsg(mobile, code string) error {
	verifyKey := fmt.Sprintf(constant.VerifySmsKey, mobile)
	count, _ := global.Redis.Get(context.Background(), verifyKey).Int64()
	if count >= constant.VerifyCountByHour {
		return errors.New("验证次数受限制，请稍后再试")
	}
	err := global.Redis.Set(context.Background(), verifyKey, count+1, time.Hour).Err()
	if err != nil {
		global.Logger.Err(err).Msg("redis连接出错")
		return errors.New("验证失败，请稍后再试")
	}
	telKey := fmt.Sprintf(constant.TelMsgKey, mobile)
	msgCode, err := global.Redis.Get(context.Background(), telKey).Result()
	if err != nil {
		global.Logger.Err(err).Msg("redis连接出错")
		return errors.New("验证失败，请稍后再试")
	}
	if code != msgCode {
		global.Logger.Err(err).Msgf("code:%s,msgCode:%s", code, msgCode)
		return errors.New("验证码不正确")
	}
	return nil
}
