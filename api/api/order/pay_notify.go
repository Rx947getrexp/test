package order

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-speed/constant"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/response"
	"go-speed/service"
	"go-speed/util"
	"strings"
)

//type PayNotifyReq struct {
//	MerNo              string `form:"mer_no" json:"mer_no"`                           // y
//	OrderNo            string `form:"order_no" json:"order_no"`                       // y
//	PayTypeCode        string `form:"paytypecode" json:"paytypecode"`                 // y
//	OrderAmount        string `form:"order_amount" json:"order_amount"`               // y
//	OrderRealityAmount string `form:"order_realityamount" json:"order_realityamount"` // y
//	Status             string `form:"status" json:"status"`
//	Sign               string `form:"sign" json:"sign"`
//}

// /pay_notify?
// MERCHANT_ID=52928&
// AMOUNT=500&
// intid=0&
// MERCHANT_ORDER_ID=100718114744251&
// P_EMAIL=zhang@qq.com&
// P_PHONE=&
// CUR_ID=44&
// commission=0&
// SIGN=1418e03619c57a4a1803fd111e725699

type PayNotifyRes struct {
}

func PayNotify(ctx *gin.Context) {
	//var (
	//	err    error
	//	req    = new(PayNotifyReq)
	//	status string
	//)
	//if err = ctx.ShouldBind(req); err != nil {
	//	global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
	//	response.RespFail(ctx, i18n.RetMsgParamParseErr, nil)
	//	return
	//}
	//global.MyLogger(ctx).Info().Msgf("request: %+v", *req)

	//status, err = service.SyncOrderStatus(ctx, req.OrderNo)
	//if err != nil {
	//	response.ResFail(ctx, err.Error())
	//	return
	//}
	//if status == constant.ReturnStatusSuccess {
	//	ctx.Writer.Write([]byte("ok"))
	//} else {
	//	ctx.Writer.Write([]byte(status))
	//}
	//ctx.Done()
	// TODO 目前只有Freekassa支付平台对接了回调通知
	FreekassaNotify(ctx)
}

func FreekassaNotify(ctx *gin.Context) {
	var (
		err    error
		req    = new(service.PayNotifyReq)
		status string
	)
	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
		response.RespFail(ctx, i18n.RetMsgParamParseErr, nil)
		return
	}
	global.MyLogger(ctx).Info().Msgf("request: %+v", *req)

	if err = checkFreekassaNotifyClientIP(ctx); err != nil {
		return
	}

	if err = checkFreekassaNotifySign(ctx, req); err != nil {
		return
	}

	status, err = service.SyncOrderStatus(ctx, req.OrderId, req)
	if err != nil {
		response.ResFail(ctx, err.Error())
		return
	}
	if status == constant.ReturnStatusSuccess {
		ctx.Writer.Write([]byte("yes"))
	} else {
		ctx.Writer.Write([]byte(status))
	}
	ctx.Done()
}

func checkFreekassaNotifySign(ctx *gin.Context, req *service.PayNotifyReq) (err error) {
	amount := req.Amount
	orderID := req.OrderId
	sign := req.Sign
	//content := fmt.Sprintf("%s:%s:%s:%s", "52928", amount, "19c44da8067cdbe4bc7f0b4db68891d3", orderID)
	content := fmt.Sprintf("%s:%s:%s:%s", "52928", amount, "HzMD1ztljV([jzk", orderID)
	_sign := md5.Sum([]byte(content))
	expectedSign := fmt.Sprintf("%x", _sign)
	global.MyLogger(ctx).Info().Msgf(`content: %s`, content)
	global.MyLogger(ctx).Info().Msgf(`expectedSign: %s`, expectedSign)
	global.MyLogger(ctx).Info().Msgf(`sign: %s`, sign)
	if sign != expectedSign {
		err = fmt.Errorf(`wrong sign`)
		global.MyLogger(ctx).Err(err).Msgf("check sign failed, expectedSign: %s", expectedSign)
		response.RespFail(ctx, err.Error(), nil)
		return err
	}
	return nil
}

func checkFreekassaNotifyClientIP(ctx *gin.Context) (err error) {
	clientIps := global.Config.FreekassaConfig.NotifyClientIp
	if strings.TrimSpace(clientIps) == "" {
		clientIps = "168.119.157.136,168.119.60.227,178.154.197.79,51.250.54.238"
	}
	if !util.IsInArrayIgnoreCase(ctx.ClientIP(), strings.Split(clientIps, ",")) {
		err = fmt.Errorf(`clientIP hacking attempt`)
		global.MyLogger(ctx).Err(err).Msgf("clientIp check failed, clientIp: %s, IPS: %s",
			ctx.ClientIP(), clientIps)
		response.RespFail(ctx, "hacking attempt!", nil)
		return err
	}
	return nil
}
