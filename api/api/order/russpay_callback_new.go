package order

import (
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
	russ_pay "go-speed/util/pay/russ-pay"
	"strings"
)

// RussPayNewCallbackReq 表示新版回调参数结构体
type RussPayNewCallbackReq struct {
	OrderID         string `json:"order_id" binding:"required"`          // 平台订单号
	Status          string `json:"status" binding:"required"`            // 订单状态
	MerchantNumber  string `json:"merchant_number" binding:"required"`   // 商户编号
	MerchantOrderID string `json:"merchant_order_id" binding:"required"` // 商户订单号
	Amount          string `json:"amount" binding:"required"`            // 实际支付金额
	DepositAmount   string `json:"deposit_amount" binding:"required"`    // 到账金额
	Fee             string `json:"fee" binding:"required"`               // 手续费
	ReceivedAmount  string `json:"received_amount" binding:"required"`   // 实收金额
	Description     string `json:"description,omitempty"`                // 描述
	Sign            string `json:"sign" binding:"required"`              // 签名
}

func RussPayNewCallback(ctx *gin.Context) {
	var (
		err            error
		req            = new(RussPayNewCallbackReq)
		payOrderEntity *entity.TPayOrder
	)

	// 绑定请求参数
	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msg("绑定参数失败")
		response.RespFail(ctx, i18n.RetMsgParamParseErr, nil)
		return
	}
	global.MyLogger(ctx).Info().Msgf("req: %#v", *req)

	// 验证签名
	sign := GenerateCallbackSign(req)
	if !strings.EqualFold(sign, req.Sign) {
		err = fmt.Errorf("signature mismatch, generated: %s, received: %s", sign, req.Sign)
		global.MyLogger(ctx).Err(err).Msg("signature verification failed")
		response.RespFail(ctx, i18n.RetMsgParamInvalid, nil)
		return
	}

	// 查询订单信息
	err = dao.TPayOrder.Ctx(ctx).Where(do.TPayOrder{OrderNo: req.MerchantOrderID}).Scan(&payOrderEntity)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msg("query pay order failed")
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}
	if payOrderEntity == nil {
		err = fmt.Errorf("order %s does not exist", req.MerchantOrderID)
		global.MyLogger(ctx).Err(err).Msg("order record not found")
		response.RespFail(ctx, i18n.RetMsgParamInvalid, nil)
		return
	}

	global.MyLogger(ctx).Info().Msgf("order entity: %+v", payOrderEntity)

	// 校验支付通道是否匹配
	if !util.IsInArrayIgnoreCase(
		payOrderEntity.PaymentChannelId,
		[]string{constant.PayChannelRussNewPayCard, constant.PayChannelRussNewPaySBP}) {
		err = fmt.Errorf("payment channel mismatch for order %s: expected channel %s", req.MerchantOrderID, payOrderEntity.PaymentChannelId)
		global.MyLogger(ctx).Err(err).Msg("Payment channel does not match")
		response.RespFail(ctx, i18n.RetMsgParamInvalid, nil)
		return
	}

	// 校验商户订单号是否匹配
	if payOrderEntity.OrderData != req.OrderID {
		err = fmt.Errorf("merchant order ID mismatch: expected %s, got %s", payOrderEntity.OrderData, req.OrderID)
		global.MyLogger(ctx).Err(err).Msg("Merchant order ID does not match")
		response.RespFail(ctx, i18n.RetMsgParamInvalid, nil)
		return
	}

	// 校验金额
	pass, err := service.CheckAmount(ctx, req.Amount, payOrderEntity.OrderAmount)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msg("BillingNumber not match")
		response.RespFail(ctx, i18n.RetMsgParamInvalid, nil)
		return
	}
	if !pass {
		err = fmt.Errorf("%s Amount(%s) is not match", req.Amount, payOrderEntity.OrderAmount)
		global.MyLogger(ctx).Err(err).Msg("Amount not match")
		response.RespFail(ctx, i18n.RetMsgParamInvalid, nil)
		return
	}

	// 同步订单状态
	status, err := service.SyncOrderStatus(ctx, req.MerchantOrderID, nil)
	if err != nil {
		response.ResFail(ctx, err.Error())
		return
	}

	// 返回响应
	if status == constant.ReturnStatusSuccess {
		ctx.Writer.Write([]byte("success_" + payOrderEntity.OrderData))
	} else {
		ctx.Writer.Write([]byte(status))
	}

	global.MyLogger(ctx).Info().Msgf("Callback processing completed for order ID: %s", req.OrderID)
}

// GenerateCallbackSign 生成回调签名
func GenerateCallbackSign(req *RussPayNewCallbackReq) string {
	params := make(map[string]string)

	// 收集非空参数（严格排除 sign）
	if req.OrderID != "" {
		params["order_id"] = req.OrderID
	}
	if req.Status != "" {
		params["status"] = req.Status
	}
	if req.MerchantNumber != "" {
		params["merchant_number"] = req.MerchantNumber
	}
	if req.MerchantOrderID != "" {
		params["merchant_order_id"] = req.MerchantOrderID
	}
	if req.Amount != "" {
		params["amount"] = req.Amount
	}
	if req.DepositAmount != "" {
		params["deposit_amount"] = req.DepositAmount
	}
	if req.Fee != "" {
		params["fee"] = req.Fee
	}
	if req.ReceivedAmount != "" {
		params["received_amount"] = req.ReceivedAmount
	}
	if req.Description != "" {
		params["description"] = req.Description
	}

	return russ_pay.GenSign(params)
}
