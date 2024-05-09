package order

import (
	"github.com/gin-gonic/gin"
	"go-speed/constant"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/response"
	"go-speed/service"
)

type PayNotifyReq struct {
	MerNo              string `form:"mer_no" json:"mer_no"`                           // y
	OrderNo            string `form:"order_no" json:"order_no"`                       // y
	PayTypeCode        string `form:"paytypecode" json:"paytypecode"`                 // y
	OrderAmount        string `form:"order_amount" json:"order_amount"`               // y
	OrderRealityAmount string `form:"order_realityamount" json:"order_realityamount"` // y
	Status             string `form:"status" json:"status"`
	Sign               string `form:"sign" json:"sign"`
}

type PayNotifyRes struct {
}

func PayNotify(ctx *gin.Context) {
	var (
		err    error
		req    = new(PayNotifyReq)
		status string
	)
	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
		response.RespFail(ctx, i18n.RetMsgParamParseErr, nil)
		return
	}
	global.MyLogger(ctx).Info().Msgf("request: %+v", *req)

	status, err = service.SyncOrderStatus(ctx, req.OrderNo)
	if err != nil {
		response.ResFail(ctx, err.Error())
		return
	}
	if status == constant.ReturnStatusSuccess {
		ctx.Writer.Write([]byte("ok"))
	} else {
		ctx.Writer.Write([]byte(status))
	}
	ctx.Done()
	return
}
