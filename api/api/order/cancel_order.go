package order

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/os/gtime"
	"go-speed/api/api/common"
	"go-speed/constant"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/response"
)

type CancelOrderReq struct {
	OrderNo string `form:"order_no" binding:"required" json:"order_no" dc:"订单号"`
}

type CancelOrderRes struct {
}

func CancelOrder(ctx *gin.Context) {
	var (
		err            error
		req            = new(CancelOrderReq)
		userEntity     *entity.TUser
		payOrderEntity *entity.TPayOrder
		affected       int64
	)
	// 绑定请求参数
	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
		response.RespFail(ctx, i18n.RetMsgParamParseErr, nil)
		return
	}
	global.MyLogger(ctx).Info().Msgf("OrderNo: %s", req.OrderNo)

	// validate user
	userEntity, err = common.ValidateClaims(ctx)
	if err != nil {
		return
	}
	payOrderEntity, err = ValidateOrder(ctx, userEntity.Email, req.OrderNo)
	if err != nil {
		return
	}

	affected, err = dao.TPayOrder.Ctx(ctx).Data(do.TPayOrder{
		Status:    constant.ParOrderStatusClosed,
		UpdatedAt: gtime.Now(),
		Version:   payOrderEntity.Version + 1,
	}).Where(do.TPayOrder{
		OrderNo: payOrderEntity.OrderNo,
		Version: payOrderEntity.Version,
	}).UpdateAndGetAffected()
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("modify order status failed")
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}
	global.MyLogger(ctx).Debug().Msgf("affected: %d", affected)
	response.ResOk(ctx, i18n.RetMsgSuccess)
}

func ValidateOrder(ctx *gin.Context, email, orderNo string) (payOrderEntity *entity.TPayOrder, err error) {
	// 根据订单号查询订单信息
	err = dao.TPayOrder.Ctx(ctx).Where(do.TPayOrder{OrderNo: orderNo}).Scan(&payOrderEntity)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("query pay order failed")
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}

	if payOrderEntity.Email != email {
		err = fmt.Errorf(`order's user is not match`)
		global.MyLogger(ctx).Err(err).Msgf(`%s is not match order's email(%s)'`, email, payOrderEntity.Email)
		response.RespFail(ctx, i18n.RetMsgParamInvalid, nil)
		return
	}

	if payOrderEntity.Status != constant.ParOrderStatusInit &&
		payOrderEntity.Status != constant.ParOrderStatusUnpaid {
		err = fmt.Errorf(`order's status is not match`)
		global.MyLogger(ctx).Err(err).Msgf(`status: %s'`, email, payOrderEntity.Status)
		response.RespFail(ctx, i18n.RetMsgParamInvalid, nil)
		return
	}
	return
}
