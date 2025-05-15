package russ_pay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-speed/global"
	"io"
	"net/http"
	"time"
)

type QueryPaymentOrderReq struct {
	MerchantNumber  string `json:"merchant_number"`
	MerchantOrderID string `json:"merchant_order_id"`
	Timestamp       string `json:"timestamp"`
	Sign            string `json:"sign"`
}

type QueryPaymentOrderRes struct {
	Code int                   `json:"code"`
	Msg  string                `json:"msg"`
	Data QueryPaymentOrderData `json:"data"`
}

type QueryPaymentOrderData struct {
	OrderId         string `json:"order_id"`
	Status          string `json:"status"`
	MerchantNumber  string `json:"merchant_number"`
	MerchantOrderId string `json:"merchant_order_id"`
	Amount          string `json:"amount"`
	DepositAmount   string `json:"deposit_amount"`
	Fee             string `json:"fee"`
	ReceivedAmount  string `json:"received_amount"`
	Description     string `json:"description"`
	Sign            string `json:"sign"`
}

type QueryOrderReq struct {
	MerchantOrderID string `json:"merchant_order_id"`
}

type QueryOrderRes struct {
	OrderID string
	URL     string
}

func QueryPaymentOrder(ctx *gin.Context, reqParams *QueryPaymentOrderReq) (*QueryPaymentOrderRes, error) {
	params := make(map[string]string)
	params["merchant_number"] = reqParams.MerchantNumber
	params["merchant_order_id"] = reqParams.MerchantOrderID
	params["timestamp"] = reqParams.Timestamp
	reqParams.Sign = GenSign(params)
	requestBody, err := json.Marshal(reqParams)
	if err != nil {
		return nil, fmt.Errorf("JSON marshal error: %v", err)
	}
	global.MyLogger(ctx).Debug().Msgf("russ-new-pay query requestBody: %s", string(requestBody))
	fmt.Printf("requestBody: %s\n", string(requestBody))
	// 创建HTTP请求
	client := &http.Client{}
	resp, err := client.Post(
		"https://api.russ-pay.com/api/payments/query",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response failed: %v", err)
	}
	global.MyLogger(ctx).Debug().Msgf("russ-new-pay query result: %s", string(body))
	fmt.Printf("response: %s\n", string(body))

	var returnCode ResponseCommon
	if err = json.Unmarshal(body, &returnCode); err != nil {
		return nil, fmt.Errorf("JSON unmarshal error: %v", err)
	}

	// 检查响应状态码
	if returnCode.Code != 200 {
		return nil, fmt.Errorf("API error: %s (code %d)", returnCode.Msg, returnCode.Code)
	}

	// 解析响应
	var result QueryPaymentOrderRes
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("JSON unmarshal error: %v", err)
	}

	return &result, nil
}

func QueryOrder(ctx *gin.Context, req QueryOrderReq) (resp *QueryPaymentOrderRes, err error) {
	requestParam := &QueryPaymentOrderReq{
		//MerchantNumber:  global.Config.RussNewPay.MerchantNumber,
		MerchantNumber:  "M1906978889457909760",
		MerchantOrderID: req.MerchantOrderID,
		Timestamp:       fmt.Sprintf("%d", time.Now().Unix()),
	}

	return QueryPaymentOrder(ctx, requestParam)
}
