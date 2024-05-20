package payment_channel

import (
	"encoding/json"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/entity"
	"go-speed/model/response"

	"github.com/gin-gonic/gin"
)

type PaymentChannelListReq struct {
}

type PaymentChannelListRes struct {
	Items []PaymentChannel `json:"items" dc:"支付通道列表"`
}

type PaymentChannel struct {
	ChannelId           string              `json:"channel_id" dc:"支付通道ID，前后端交互时使用"`
	ChannelName         string              `json:"channel_name" dc:"支付通道名称，展示给用户"`
	IsActive            int                 `json:"is_active" dc:"支付通道是否可用，1：可用，2：不可用"`
	FreeTrialDays       int                 `json:"free_trial_days" dc:"赠送的免费时长（以天为单位）"`
	TimeoutDuration     int                 `json:"timeout_duration" dc:"订单未支付超时自动关闭时间（单位分钟）"`
	PaymentQRCode       string              `json:"payment_qr_code" dc:"支付收款码. eg: U支付收款码"`
	PaymentQRUrl        string              `json:"payment_qr_url" dc:"支付收款码链接"`
	UsdNetwork          string              `json:"usd_network" dc:"USD支付网络"`
	BankCardInfo        []BankCardInfo      `json:"bank_card_info" dc:"银行卡信息"`
	CustomerServiceInfo CustomerServiceInfo `json:"customer_service_info" dc:"客服信息"`
	Weight              int                 `json:"weight" dc:"权重，根据权重排序"`
	CreatedAt           string              `json:"created_at" dc:"创建时间"`
	UpdatedAt           string              `json:"updated_at" dc:"更新时间"`
}

type BankCardInfo struct {
	Cardholder     string `json:"cardholder" dc:"持卡人"`
	BankCardNumber string `json:"bank_card_number" dc:"银行卡号"`
	IssuingBank    string `json:"issuing_bank" dc:"开户银行"`
}

type CustomerServiceInfo struct {
	Phone        string `json:"phone" dc:"持卡人姓名"`
	Email        string `json:"email" dc:"银行卡号"`
	WorkingHours string `json:"working_hours" dc:"工作时间"`
}

func PaymentChannelList(ctx *gin.Context) {
	var (
		err         error
		entityItems []entity.TPaymentChannel
	)
	err = dao.TPaymentChannel.Ctx(ctx).Order(dao.TPaymentChannel.Columns().Weight, "desc").Scan(&entityItems)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("get payment failed")
		response.ResFail(ctx, err.Error())
		return
	}
	items := make([]PaymentChannel, 0)
	for _, item := range entityItems {
		bankCardInfo := make([]BankCardInfo, 0)
		customerServiceInfo := CustomerServiceInfo{}
		_ = json.Unmarshal([]byte(item.BankCardInfo), &bankCardInfo)
		_ = json.Unmarshal([]byte(item.CustomerServiceInfo), &customerServiceInfo)
		items = append(items, PaymentChannel{
			ChannelId:           item.ChannelId,
			ChannelName:         item.ChannelName,
			IsActive:            item.IsActive,
			FreeTrialDays:       item.FreeTrialDays,
			TimeoutDuration:     item.TimeoutDuration,
			PaymentQRCode:       item.PaymentQrCode,
			PaymentQRUrl:        item.PaymentQrUrl,
			UsdNetwork:          item.UsdNetwork,
			BankCardInfo:        bankCardInfo,
			CustomerServiceInfo: customerServiceInfo,
			Weight:              item.Weight,
			CreatedAt:           item.CreatedAt.String(),
			UpdatedAt:           item.UpdatedAt.String(),
		})
	}
	response.RespOk(ctx, i18n.RetMsgSuccess, PaymentChannelListRes{Items: items})
}
