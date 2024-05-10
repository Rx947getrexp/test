package order

import (
	"fmt"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/response"

	"github.com/gin-gonic/gin"
)

type CheckOrderReq struct {
	UserId  int    `form:"user_id" binding:"required" json:"user_id" dc:"用户ID"`
	OrderNo string `form:"order_no" binding:"required" json:"order_no" dc:"订单号"`
}

type CheckOrderRes struct {
	UserId             string `json:"user_id" dc:"用户ID"`
	OrderNo            string `json:"order_no" dc:"订单号"`
	PaymentChannelName string `json:"payment_channel_name" dc:"支付通道名称"`
}

// 检查订单结果并执行相应逻辑
func CheckOrder(ctx *gin.Context) {
	var (
		err error
		req = new(CheckOrderReq)
	)
	// 绑定请求参数
	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
		response.RespFail(ctx, i18n.RetMsgParamParseErr, nil)
		return
	}
	global.MyLogger(ctx).Info().Msgf("request: %+v", *req)
	// 根据订单号查询订单信息
	order := new(do.TPayOrder)
	_, err = dao.TPayOrder.Ctx(ctx).Where("user_id = ? and order_no = ? and", req.UserId, req.OrderNo).One(&order)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("数据库查询出错, userId: %d order_no: %d", req.UserId, req.OrderNo)
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}
	// 判断订单状态，如果订单支付失败，执行赠送时长逻辑
	PaymentChannelName := order.PaymentChannelName.(string)
	if order.Status == "fail" {
		global.MyLogger(ctx).Err(err).Msgf("检测该笔订单充值失败,触发免费赠送时长 OrderNo: %d", req.OrderNo)
		giveFreeToUser(ctx, req.UserId, PaymentChannelName)
		return
	}
	res := CheckOrderRes{
		UserId:             fmt.Sprintf("%d", req.UserId),
		OrderNo:            req.OrderNo,
		PaymentChannelName: PaymentChannelName,
	}
	response.RespOk(ctx, i18n.RetMsgSuccess, res)
}
