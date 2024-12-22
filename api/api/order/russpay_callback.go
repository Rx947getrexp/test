package order

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-speed/constant"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/response"
	"go-speed/service"
	"go-speed/util"
	"go-speed/util/pay/russpay"
)

type RussPayCallbackReq struct {
	AppKey  string `form:"appKey" binding:"required" json:"appKey" dc:"appKey"`
	Sign    string `form:"sign" binding:"required" json:"sign" dc:"sign"`
	Content string `form:"content" binding:"required" json:"content" dc:"content"`
}

type Content struct {
	BillingNumber string `form:"billingNumber" binding:"required" json:"billingNumber" dc:"billingNumber"`
	OrderNumber   string `form:"orderNumber" binding:"required" json:"orderNumber" dc:"orderNumber"`
	Amount        string `form:"amount" binding:"required" json:"amount" dc:"amount"`
	Currency      string `form:"currency" binding:"required" json:"currency" dc:"currency"`
	Status        string `form:"status" binding:"required" json:"status" dc:"status"`
	Message       string `form:"message" json:"message" dc:"message"`
	PaymentMethod string `form:"paymentMethod" json:"paymentMethod" dc:"paymentMethod"`
}

func RussPayCallback(ctx *gin.Context) {
	var (
		err            error
		req            = new(RussPayCallbackReq)
		payOrderEntity *entity.TPayOrder
	)

	// 绑定请求参数
	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
		response.RespFail(ctx, i18n.RetMsgParamParseErr, nil)
		return
	}
	global.MyLogger(ctx).Info().Msgf("req: %#v", *req)

	if req.AppKey != russpay.AppKey {
		err = fmt.Errorf("appKey is invalid")
		global.MyLogger(ctx).Err(err).Msgf("invalid notify")
		response.RespFail(ctx, i18n.RetMsgParamInvalid, nil)
		return
	}

	content, err := russpay.Base64Decode(req.Content)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("Content Base64Decode failed")
		response.RespFail(ctx, i18n.RetMsgParamInvalid, nil)
		return
	}
	global.MyLogger(ctx).Info().Msgf("content: %#v", string(content))

	var c Content
	err = json.Unmarshal(content, &c)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msg(">>>>>> Content Unmarshal failed")
		response.RespFail(ctx, i18n.RetMsgParamInvalid, nil)
		return
	}

	// 根据订单号查询订单信息
	err = dao.TPayOrder.Ctx(ctx).Where(do.TPayOrder{OrderNo: c.OrderNumber}).Scan(&payOrderEntity)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("query pay order failed")
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}
	if payOrderEntity == nil {
		err = fmt.Errorf("%s can not found pay order entity", c.OrderNumber)
		global.MyLogger(ctx).Err(err).Msg("record is nil")
		response.RespFail(ctx, i18n.RetMsgParamInvalid, nil)
		return
	}
	if !util.IsInArrayIgnoreCase(
		payOrderEntity.PaymentChannelId,
		[]string{constant.PayChannelRussPayBankCard, constant.PayChannelRussPaySBP, constant.PayChannelRussPaySBER}) {
		err = fmt.Errorf("%s PaymentChannelId(%s) is not match", c.OrderNumber, payOrderEntity.PaymentChannelId)
		global.MyLogger(ctx).Err(err).Msg("PaymentChannelId not match")
		response.RespFail(ctx, i18n.RetMsgParamInvalid, nil)
		return
	}

	if payOrderEntity.OrderData != c.BillingNumber {
		err = fmt.Errorf("%s BillingNumber(%s) is not match", c.BillingNumber, payOrderEntity.OrderData)
		global.MyLogger(ctx).Err(err).Msg("BillingNumber not match")
		response.RespFail(ctx, i18n.RetMsgParamInvalid, nil)
		return
	}

	pass, err := service.CheckAmount(ctx, c.Amount, payOrderEntity.OrderAmount)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msg("BillingNumber not match")
		response.RespFail(ctx, i18n.RetMsgParamInvalid, nil)
		return
	}
	if !pass {
		err = fmt.Errorf("%s Amount(%s) is not match", c.Amount, payOrderEntity.OrderAmount)
		global.MyLogger(ctx).Err(err).Msg("Amount not match")
		response.RespFail(ctx, i18n.RetMsgParamInvalid, nil)
		return
	}

	status, err := service.SyncOrderStatus(ctx, c.OrderNumber, nil)
	if err != nil {
		response.ResFail(ctx, err.Error())
		return
	}
	if status == constant.ReturnStatusSuccess {
		ctx.Writer.Write([]byte("success_" + payOrderEntity.OrderData))
	} else {
		ctx.Writer.Write([]byte(status))
	}
	ctx.Done()

	global.MyLogger(ctx).Info().Msgf("OrderNo: %s, callback end", c.OrderNumber)
}
