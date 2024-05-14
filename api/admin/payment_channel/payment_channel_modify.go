package payment_channel

import (
	"encoding/json"
	"go-speed/constant"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/response"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/os/gtime"
)

type PaymentChannelModifyReq struct {
	ChannelId           string               `form:"channel_id" binding:"required" dc:"支付通道ID，前后端交互时使用.（不可以修改）"`
	ChannelName         string               `form:"channel_name" dc:"支付通道名称，展示给用户"`
	IsActive            int                  `form:"is_active" dc:"支付通道是否可用，1：可用，2：不可用"`
	FreeTrialDays       int                  `form:"free_trial_days" dc:"赠送的免费时长（以天为单位）"`
	TimeoutDuration     int                  `form:"timeout_duration" dc:"订单未支付超时自动关闭时间（单位分钟）"`
	PaymentQRCode       *string              `form:"payment_qr_code" dc:"支付收款码. eg: U支付收款码"`
	BankCardInfo        []BankCardInfo       `form:"bank_card_info" dc:"银行卡信息"`
	CustomerServiceInfo *CustomerServiceInfo `form:"customer_service_info" dc:"客服信息"`
	Weight              int                  `form:"weight" dc:"权重，根据权重排序"`
}

type PaymentChannelModifyRes struct {
}

func PaymentChannelModify(ctx *gin.Context) {
	var (
		err           error
		req           = new(PaymentChannelModifyReq)
		paymentEntity *entity.TPaymentChannel
		affected      int64
	)
	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
		response.ResFail(ctx, i18n.RetMsgParamParseErr)
		return
	}
	global.MyLogger(ctx).Info().Msgf("req: %+v", *req)
	if req.IsActive != constant.PaymentChannelIsActiveYes && req.IsActive != constant.PaymentChannelIsActiveNo {
		response.ResFail(ctx, `param "IsActive" invalid`)
		return
	}

	if req.FreeTrialDays > global.Config.PayConfig.MaxFreeTrialDays {
		response.ResFail(ctx, `param "FreeTrialDays" invalid`)
		return
	}

	err = dao.TPaymentChannel.Ctx(ctx).Where(do.TPaymentChannel{ChannelId: req.ChannelId}).Scan(&paymentEntity)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("query TPaymentChannels failed")
		response.ResFail(ctx, err.Error())
		return
	}
	if paymentEntity == nil {
		global.MyLogger(ctx).Err(err).Msgf("param `ChannelId` invalid")
		response.ResFail(ctx, "param `ChannelId` invalid")
		return
	}
	updateData := do.TPaymentChannel{UpdatedAt: gtime.Now()}
	if req.IsActive > 0 {
		updateData.IsActive = req.IsActive
	}
	if req.FreeTrialDays > 0 {
		updateData.FreeTrialDays = req.FreeTrialDays
	}
	if req.TimeoutDuration > 0 {
		updateData.TimeoutDuration = req.TimeoutDuration
	}
	if req.PaymentQRCode != nil {
		updateData.PaymentQrCode = *req.PaymentQRCode
	}
	if req.Weight > 0 {
		updateData.Weight = req.Weight
	}
	if len(req.BankCardInfo) > 0 {
		bytes, err := json.Marshal(req.BankCardInfo)
		if err != nil {
			global.MyLogger(ctx).Err(err).Msgf("param `BankCardInfo` invalid")
			response.ResFail(ctx, "param `BankCardInfo` invalid")
			return
		}
		updateData.BankCardInfo = string(bytes)
	}
	if req.CustomerServiceInfo != nil {
		bytes, err := json.Marshal(req.CustomerServiceInfo)
		if err != nil {
			global.MyLogger(ctx).Err(err).Msgf("param `CustomerServiceInfo` invalid")
			response.ResFail(ctx, "param `CustomerServiceInfo` invalid")
			return
		}
		updateData.CustomerServiceInfo = string(bytes)
	}
	affected, err = dao.TNode.Ctx(ctx).Data(updateData).
		Where(do.TPaymentChannel{
			ChannelId: req.ChannelId,
		}).UpdateAndGetAffected()
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("modify Name failed")
		response.ResFail(ctx, err.Error())
		return
	}
	global.MyLogger(ctx).Debug().Msgf("affected: %d", affected)
	response.ResOk(ctx, i18n.RetMsgSuccess)
}
