package russ_pay

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-speed/constant"
	"go-speed/global"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"
)

// 创建支付请求结构体
type CreatePaymentRequest struct {
	MerchantNumber  string `json:"merchant_number"`
	MerchantOrderID string `json:"merchant_order_id"`
	Amount          string `json:"amount"`
	PayWay          string `json:"pay_way"`
	ClientIP        string `json:"client_ip"`
	Description     string `json:"description,omitempty"`
	SuccessURL      string `json:"success_url"`
	FailedURL       string `json:"failed_url"`
	Timestamp       string `json:"timestamp"`
	Sign            string `json:"sign"`
}

// 支付响应结构体
type CreatePaymentResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		OrderID string `json:"order_id"`
		URL     string `json:"url"`
	} `json:"data"`
}

type ResponseCommon struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 生成签名
func generateSign(req *CreatePaymentRequest) string {
	params := make(map[string]string)

	// 收集非空参数（排除sign）
	if req.MerchantNumber != "" {
		params["merchant_number"] = req.MerchantNumber
	}
	if req.MerchantOrderID != "" {
		params["merchant_order_id"] = req.MerchantOrderID
	}
	if req.Amount != "" {
		params["amount"] = req.Amount
	}
	if req.PayWay != "" {
		params["pay_way"] = req.PayWay
	}
	if req.ClientIP != "" {
		params["client_ip"] = req.ClientIP
	}
	if req.Description != "" {
		params["description"] = req.Description
	}
	if req.SuccessURL != "" {
		params["success_url"] = req.SuccessURL
	}
	if req.FailedURL != "" {
		params["failed_url"] = req.FailedURL
	}
	if req.Timestamp != "" {
		params["timestamp"] = req.Timestamp
	}
	return GenSign(params)
}

func GenSign(params map[string]string) string {
	// 按ASCII码排序参数名
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 拼接签名字符串
	var stringA string
	for i, k := range keys {
		if i > 0 {
			stringA += "&"
		}
		stringA += fmt.Sprintf("%s=%s", k, strings.TrimSpace(params[k]))
	}

	// 拼接API Key并生成MD5
	stringSignTemp := stringA + "&key=" + "d413URbYZrsQqpb0fGkqCBgj2fFMzv81"
	hasher := md5.New()
	hasher.Write([]byte(stringSignTemp))
	return strings.ToLower(hex.EncodeToString(hasher.Sum(nil)))
}

// 创建支付订单
func CreatePaymentOrder(ctx *gin.Context, reqParams *CreatePaymentRequest) (string, string, error) {
	reqParams.Sign = generateSign(reqParams)
	requestBody, err := json.Marshal(reqParams)
	if err != nil {
		return "", "", fmt.Errorf("JSON marshal error: %v", err)
	}

	global.MyLogger(ctx).Debug().Msgf("russ-new-pay create order requestBody: %s", string(requestBody))

	fmt.Printf("requestBody: %s\n", string(requestBody))

	// 创建HTTP请求
	client := &http.Client{}
	resp, err := client.Post(
		"https://api.russ-pay.com/api/payments/create",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return "", "", fmt.Errorf("HTTP request failed: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", fmt.Errorf("read response failed: %v", err)
	}

	global.MyLogger(ctx).Debug().Msgf("russ-new-pay create order result: %s", string(body))
	fmt.Printf("result: %s\n", string(body))

	var returnCode ResponseCommon
	if err := json.Unmarshal(body, &returnCode); err != nil {
		return "", "", fmt.Errorf("JSON unmarshal error: %v", err)
	}

	// 检查响应状态码
	if returnCode.Code != 200 {
		return "", "", fmt.Errorf("API error: %s (code %d)", returnCode.Msg, returnCode.Code)
	}

	// 解析响应
	var result CreatePaymentResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return "", "", fmt.Errorf("JSON unmarshal error: %v", err)
	}

	return result.Data.OrderID, result.Data.URL, nil
}

// 使用示例
/*
func main() {
	apiHost := "https://example.com"
	apiKey := "your_api_key_here"

	paymentReq := CreatePaymentRequest{
		MerchantNumber:  "M123456789",
		MerchantOrderID: "ORDER_123",
		Amount:          "100.00",
		PayWay:          "CARD",
		ClientIP:        "127.0.0.1",
		Description:     "Test Payment",
		SuccessURL:      "https://yourdomain.com/success",
		FailedURL:       "https://yourdomain.com/failed",
		Timestamp:       "20230807120000", // 建议使用time.Now().Format("20060102150405")
	}
	fmt.Println("Payment URL:", paymentUrl)
}
*/

type CreateOrderReq struct {
	ChannelId       string
	MerchantOrderID string
	Amount          string
	SuccessURL      string
	FailedURL       string
}

type CreateOrderRes struct {
	OrderID string
	URL     string
}

const (
	payWayCard = "CARD"
	payWaySBP  = "SBP"
)

func CreateOrder(ctx *gin.Context, req CreateOrderReq) (resp *CreateOrderRes, err error) {
	var payWay = payWayCard
	if req.ChannelId == constant.PayChannelRussNewPaySBP {
		payWay = payWaySBP
	}
	//payWay = "card_core_payin_rub[mock_server,create_declined_with_callback]"
	paymentReq := &CreatePaymentRequest{
		//MerchantNumber:  global.Config.RussNewPay.MerchantNumber,
		MerchantNumber:  "M1906978889457909760",
		MerchantOrderID: req.MerchantOrderID,
		Amount:          req.Amount,
		PayWay:          payWay,
		//ClientIP:        "45.251.243.140", //ctx.ClientIP(),
		ClientIP:   ctx.ClientIP(),
		SuccessURL: req.SuccessURL,
		FailedURL:  req.FailedURL,
		Timestamp:  fmt.Sprintf("%d", time.Now().Unix()), // 建议使用 .Format("20060102150405")
	}

	id, paymentUrl, err := CreatePaymentOrder(ctx, paymentReq)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("Payment creation failed: %s", err.Error())
		return
	}

	return &CreateOrderRes{
		OrderID: id,
		URL:     paymentUrl,
	}, nil
}
