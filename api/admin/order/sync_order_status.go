package order

import (
	"github.com/gin-gonic/gin"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/response"
	"go-speed/service"
)

type SyncOrderStatusReq struct {
	OrderNo string `form:"order_no" binding:"required" json:"order_no"`
}

type SyncOrderStatusRes struct {
}

func SyncOrderStatus(ctx *gin.Context) {
	var (
		err error
		req = new(SyncOrderStatusReq)
	)
	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
		response.RespFail(ctx, i18n.RetMsgParamParseErr, nil)
		return
	}
	global.MyLogger(ctx).Debug().Msgf("request: %+v", *req)

	err = service.SyncOrderStatus(ctx, req.OrderNo)
	if err != nil {
		response.ResFail(ctx, err.Error())
		return
	}
	response.ResOk(ctx, "成功")
	return
}
