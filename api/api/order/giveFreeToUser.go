package order

import (
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/response"

	"github.com/gin-gonic/gin"
)

// giveFreeToUser 根据用户ID和支付通道名称赠送免费时长
func giveFreeToUser(ctx *gin.Context, userId int, paymentChannelName string) {
	// 根据支付通道名称查询免费时长
	channel := new(do.TPaymentChannels)
	_, err := dao.TPaymentChannels.Ctx(ctx).Where("name = ?", paymentChannelName).One(&channel)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("params 'payment_channel_name'(%s) invalid", paymentChannelName)
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}
	freeTrialDay := channel.FreeTrialDays.(int)
	freeTrialSeconds := freeTrialDay * 24 * 60 * 60
	// 查询用户信息
	user := new(do.TUser)
	_, err = dao.TUser.Ctx(ctx).Where("id = ?", userId).One(&user)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("userId %d 查询db失败", userId)
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}
	// 计算到期时间
	expiredTime := user.ExpiredTime.(int)
	expiredTime += freeTrialSeconds
	// 更新用户到期时间
	_, err = dao.TUser.Ctx(ctx).Where("id = ?", userId).Data(do.TUser{
		ExpiredTime: expiredTime,
	}).UpdateAndGetAffected()
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("赠送免费时长时间失败 email: %s", user.Email)
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}
	global.MyLogger(ctx).Info().Msgf("已为用户补偿免费时长 email: %s", user.Email)
	response.RespOk(ctx, i18n.RetMsgSuccess, nil)
	return
}
