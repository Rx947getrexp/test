package paymentChannel

import (
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/response"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/os/gtime"
)

type PaymentModifyReq struct {
	Name          string `binding:"required" form:"name" dc:"支付通道名称"`
	IsActive      int    `form:"is_active" dc:"支付通道是否可用，1表示可用,2表示不可用"`
	FreeTrialDays int    `form:"free_trial_days" dc:"赠送的免费时长（以天为单位）"`
	CreatedAt     string `form:"created_at" dc:"创建时间"`
	UpdatedAt     string `form:"updated_at" dc:"更新时间"`
}

type PaymentModifyRes struct {
}

// PaymentModify 修改支付通道
func PaymentModify(ctx *gin.Context) {
	var (
		err           error
		req           = new(PaymentModifyReq)
		paymentEntity *entity.TPaymentChannels
		affected      int64
	)
	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
		response.ResFail(ctx, i18n.RetMsgParamParseErr)
		return
	}
	global.MyLogger(ctx).Info().Msgf("req: %+v", *req)
	err = dao.TPaymentChannels.Ctx(ctx).Where(do.TPaymentChannels{Name: req.Name}).Scan(&paymentEntity)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("query payment failed")
		response.ResFail(ctx, err.Error())
		return
	}
	if paymentEntity == nil {
		global.MyLogger(ctx).Err(err).Msgf("param `payment` invalid")
		response.ResFail(ctx, i18n.RetMsgParamInvalid)
		return
	}
	updateData := do.TPaymentChannels{UpdatedAt: gtime.Now()}
	if req.Name != "" {
		updateData.Name = req.Name
	}
	if req.IsActive > 0 {
		updateData.IsActive = req.IsActive
	}
	if req.FreeTrialDays > 0 {
		updateData.FreeTrialDays = req.FreeTrialDays
	}
	affected, err = dao.TNode.Ctx(ctx).Data(updateData).Where(do.TPaymentChannels{
		Name: req.Name,
	}).UpdateAndGetAffected()
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("modify Name failed")
		response.ResFail(ctx, err.Error())
		return
	}
	global.MyLogger(ctx).Debug().Msgf("affected: %d", affected)
	response.ResOk(ctx, i18n.RetMsgSuccess)
}
