package order

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-speed/api/api/common"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/response"
	"go-speed/service"
)

type QueryOrderReq struct {
	OrderNo string `form:"order_no" binding:"required" json:"order_no" dc:"订单号"`
}

type QueryOrderRes struct {
	OrderNo            string `json:"order_no"             dc:"订单号"`
	PaymentChannelId   string `json:"payment_channel_id"   dc:"支付通道ID"`
	GoodsId            int    `json:"goods_id"             dc:"套餐ID"`
	OrderAmount        string `json:"order_amount"         dc:"交易金额"`
	Currency           string `json:"currency"             dc:"交易币种"`
	Status             string `json:"status"               dc:"状态"`
	OrderRealityAmount string `json:"order_reality_amount" dc:"实际交易金额"`
	PaymentProof       string `json:"payment_proof"        dc:"支付凭证地址"`
	CreatedAt          string `json:"created_at"           dc:"创建时间"`
	UpdatedAt          string `json:"updated_at"           dc:"更新时间"`
}

func QueryOrder(ctx *gin.Context) {
	var (
		err      error
		req      = new(QueryOrderReq)
		payOrder *entity.TPayOrder
		user     *entity.TUser
	)
	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
		response.ResFail(ctx, i18n.RetMsgParamParseErr)
		return
	}
	user, err = common.ValidateClaims(ctx)
	if err != nil {
		return
	}

	// sync order status
	_, err = service.SyncOrderStatus(ctx, req.OrderNo)
	if err != nil {
		response.ResFail(ctx, err.Error())
		return
	}

	err = dao.TPayOrder.Ctx(ctx).Where(do.TPayOrder{
		UserId:  user.Id,
		OrderNo: req.OrderNo,
	}).Scan(&payOrder)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("query pay order failed")
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}
	if payOrder == nil {
		err = fmt.Errorf("param OrderNo invalid")
		global.MyLogger(ctx).Err(err).Msgf("%s", req.OrderNo)
		response.ResFail(ctx, i18n.RetMsgParamInvalid)
		return
	}
	resp := QueryOrderRes{
		OrderNo:            payOrder.OrderNo,
		PaymentChannelId:   payOrder.PaymentChannelId,
		GoodsId:            payOrder.GoodsId,
		OrderAmount:        payOrder.OrderAmount,
		Currency:           payOrder.Currency,
		Status:             payOrder.Status,
		OrderRealityAmount: payOrder.OrderRealityAmount,
		PaymentProof:       payOrder.PaymentProof,
		CreatedAt:          payOrder.CreatedAt.String(),
		UpdatedAt:          payOrder.UpdatedAt.String(),
	}
	response.RespOk(ctx, i18n.RetMsgSuccess, resp)
}
