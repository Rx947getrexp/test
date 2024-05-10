package paymentChannel

import (
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/entity"
	"go-speed/model/response"

	"github.com/gin-gonic/gin"
)

type PaymentListReq struct {
}

type PaymentList struct {
	Name          string `json:"name" dc:"支付通道名称"`
	IsActive      int    `json:"is_active" dc:"支付通道是否可用，1表示可用,2表示不可用"`
	FreeTrialDays int    `json:"free_trial_days" dc:"赠送的免费时长（以天为单位）"`
	CreatedAt     string `json:"created_at" dc:"创建时间"`
	UpdatedAt     string `json:"updated_at" dc:"更新时间"`
}

type PaymentListRes struct {
	Items []PaymentList `json:"items" dc:"支付通道列表"`
}

func ChannelList(ctx *gin.Context) {
	var (
		err   error
		items []entity.TPaymentChannels
	)
	err = dao.TPaymentChannels.Ctx(ctx).Scan(&items)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("get payment failed")
		response.ResFail(ctx, err.Error())
		return
	}
	payment := make([]PaymentList, 0)
	for _, item := range items {
		payment = append(payment, PaymentList{
			Name:          item.Name,
			IsActive:      item.IsActive,
			FreeTrialDays: item.FreeTrialDays,
			CreatedAt:     item.CreatedAt.String(),
			UpdatedAt:     item.UpdatedAt.String(),
		})
	}
	response.RespOk(ctx, i18n.RetMsgSuccess, PaymentListRes{
		Items: payment,
	})
}
