package order

import (
	"fmt"
	"go-speed/api/api/common"
	"go-speed/constant"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/entity"
	"go-speed/model/response"
	"go-speed/service"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

type ConfirmOrderReq struct {
	OrderNo string `form:"order_no" binding:"required" json:"order_no" dc:"订单号"`
}

type ConfirmOrderRes struct {
	Status string `json:"status" dc:"订单状态. success:成功，fail:支付失败, waiting：等待支付中"`
}

func ConfirmOrder(ctx *gin.Context) {
	var (
		err         error
		req         = new(ConfirmOrderReq)
		user        *entity.TUser
		payOrder    *entity.TPayOrder
		orderStatus string
	)
	defer func() {
		if r := recover(); r != nil {
			// 同时打印到日志文件和标准输出中
			global.MyLogger(ctx).Err(err).Msgf("%+v\n%+v", r, string(debug.Stack()))
		}
	}()

	// 绑定请求参数
	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
		response.RespFail(ctx, i18n.RetMsgParamParseErr, nil)
		return
	}
	global.MyLogger(ctx).Info().Msgf("OrderNo: %s", req.OrderNo)

	// validate user
	user, err = common.ValidateClaims(ctx)
	if err != nil {
		return
	}

	// validate order
	payOrder, err = ValidateOrder(ctx, user.Email, req.OrderNo)
	if err != nil {
		return
	}

	// validate proof
	if payOrder.PaymentChannelId == constant.PayChannelBankCardPay && payOrder.PaymentProof == "" {
		err = fmt.Errorf(i18n.RetMsgProofUploadNone)
		global.MyLogger(ctx).Err(err).Msgf("PaymentProof: %s", payOrder.PaymentProof)
		response.RespFail(ctx, i18n.RetMsgProofUploadNone, nil)
		return
	}
	global.MyLogger(ctx).Info().Msgf("OrderNo: %s, ResultStatus: %s", payOrder.OrderNo, payOrder.ResultStatus)

	// sync order status
	orderStatus, err = service.SyncOrderStatus(ctx, req.OrderNo)
	if err != nil {
		response.ResFail(ctx, err.Error())
		return
	}

	response.RespOk(ctx, i18n.RetMsgSuccess, ConfirmOrderRes{Status: orderStatus})
}
