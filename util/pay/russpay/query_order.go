package russpay

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-speed/global"
)

const (
	queryOrderUrl = "https://api.yolesdk.com/v2/api/payment/transaction/query"
)

type QueryOrderReq struct {
	BillingNumber string `json:"billingNumber" dc:"交易记录号"`
}

type QueryOrderResponse struct {
	PaymentStatus   string `json:"paymentStatus"`
	PaymentDatetime string `json:"paymentDatetime"`
	BillingNumber   string `json:"billingNumber"`
	ErrorMassage    string `json:"errorMassage"`
	PaymentMethod   string `json:"paymentMethod"`
}

func QueryOrder(ctx *gin.Context, req QueryOrderReq) (resp *QueryOrderResponse, err error) {
	mapReq := make(map[string]interface{})
	mapReq["appKey"] = AppKey

	mapContent := make(map[string]interface{})
	mapContent["billingNumber"] = req.BillingNumber

	contentBase64, err := Base64Encode(mapContent)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("Base64Encode failed")
		return
	}
	mapReq["content"] = contentBase64
	mapReq["sign"] = MD5(contentBase64, SecretKey)

	response, err := requestRussPay(ctx, queryOrderUrl, mapReq)
	if err != nil {
		fmt.Println("-----> err: ", err.Error())
		global.MyLogger(ctx).Err(err).Msgf("requestRussPay failed")
		return
	}

	var out []byte
	out, err = Base64Decode(response.Content.Content)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("Base64Decode failed")
		return
	}

	err = json.Unmarshal(out, &resp)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("Unmarshal failed")
		return
	}

	global.MyLogger(ctx).Info().Msgf(">>>>>> queryOrder success, orderResponse: %#v", *resp)
	return
}
