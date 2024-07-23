package order

import (
	"encoding/json"
	"go-speed/api/api/common"
	"go-speed/constant"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/response"
	"golang.org/x/exp/rand"
	"time"

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
	PaymentQRCode       string              `json:"payment_qr_code" dc:"支付收款码. eg: U支付收款码"`
	PaymentQRUrl        string              `json:"payment_qr_url" dc:"支付收款码链接"`
	UsdNetwork          string              `json:"usd_network" dc:"USD支付网络"`
	BankCardInfo        BankCardInfo        `json:"bank_card_info" dc:"银行卡信息"`
	CustomerServiceInfo CustomerServiceInfo `json:"customer_service_info" dc:"客服信息"`
	Weight              int                 `json:"weight" dc:"权重，根据权重排序"`
	WmId                string              `json:"wmid" dc:"wmid"`
	Purse               string              `json:"purse" dc:"purse"`
	CurrencyType        string              `json:"currency_type" dc:"支付渠道币种"`
	FreekassaCode       string              `json:"freekassa_code" dc:"freekassa支付通道"`
	CommissionRate      float64             `json:"commission_rate" dc:"手续费比例"`
	Commission          float64             `json:"commission" dc:"手续费"`
	MinPayAmount        float64             `json:"min_pay_amount" dc:"最低支付金额"`
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
	_, err = common.ValidateClaims(ctx)
	if err != nil {
		return
	}
	err = dao.TPaymentChannel.Ctx(ctx).
		Where(do.TPaymentChannel{IsActive: constant.PaymentChannelIsActiveYes}).
		Order(dao.TPaymentChannel.Columns().Weight, constant.OrderTypeDesc).
		Scan(&entityItems)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("query payment channel failed")
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}

	items := make([]PaymentChannel, 0)
	for _, item := range entityItems {
		bankCardInfo := make([]BankCardInfo, 0)
		customerServiceInfo := CustomerServiceInfo{}
		_ = json.Unmarshal([]byte(item.BankCardInfo), &bankCardInfo)
		_ = json.Unmarshal([]byte(item.CustomerServiceInfo), &customerServiceInfo)

		var card BankCardInfo
		if len(bankCardInfo) > 0 {
			randNum := genRandForBankCard(len(bankCardInfo))
			if randNum < len(bankCardInfo) {
				card = bankCardInfo[randNum]
			} else {
				card = bankCardInfo[0]
			}
		}

		var wmId, purse string
		if item.ChannelId == constant.PayChannelWebMoneyPay {
			wmId, purse = global.Config.WebMoneyConfig.WmId, global.Config.WebMoneyConfig.Purse
		}

		items = append(items, PaymentChannel{
			ChannelId:           item.ChannelId,
			ChannelName:         item.ChannelName,
			PaymentQRCode:       item.PaymentQrCode,
			PaymentQRUrl:        item.PaymentQrUrl,
			UsdNetwork:          item.UsdNetwork,
			BankCardInfo:        card,
			CustomerServiceInfo: customerServiceInfo,
			Weight:              item.Weight,
			WmId:                wmId,
			Purse:               purse,
			CurrencyType:        item.CurrencyType,
			FreekassaCode:       item.FreekassaCode,
			CommissionRate:      item.CommissionRate,
			Commission:          item.Commission,
			MinPayAmount:        item.MinPayAmount,
		})
	}
	response.RespOk(ctx, i18n.RetMsgSuccess, PaymentChannelListRes{Items: items})
}

func genRandForBankCard(n int) int {
	rand.Seed(uint64(time.Now().UnixNano()))
	return rand.Intn(n) // 生成一个0到9999之间的随机数
}
