package freekassa

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"go-speed/global"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateOrderReq struct {
	//ShopId    int
	//Nonce     int
	//Signature string
	PaymentId string
	I         string
	Email     string
	Ip        string
	Amount    float64
	Currency  string
	//Tel       string
	SuccessUrl string
	//FailureUrl      string
	//NotificationUrl string
}

type OrderResponse struct {
	Type      string
	OrderId   int
	OrderHash string
	Location  string
}

var createOrderUrl = "https://api.freekassa.com/v1/orders/create"

func createOrder(ctx *gin.Context, mapData map[string]interface{}) (res *OrderResponse, err error) {
	global.MyLogger(ctx).Info().Msgf(">>>>>> 请求地址: %#v", createOrderUrl)
	global.MyLogger(ctx).Info().Msgf(">>>>>> 请求参数: %#v", mapData)

	jsonData, err := json.Marshal(mapData)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msg(">>>>>>  json.Marshal failed")
		return
	}

	req, err := http.NewRequest("POST", createOrderUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		global.MyLogger(ctx).Err(err).Msg(">>>>>>  NewRequest failed")
		return
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
	resp, err := client.Do(req)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msg(">>>>>>  client.Do(req) failed")
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("返回结果:", string(body))
	global.MyLogger(ctx).Info().Msgf(">>>>>>  返回结果: %q", string(body))
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("StatusCode: %d is not ok", resp.StatusCode)
		global.MyLogger(ctx).Err(err).Msgf("Request failed with status code %d and response: %s\n", resp.StatusCode, string(body))
		return
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msg(">>>>>> Unmarshal failed")
		return
	}
	return
}

func CreateOrder(ctx *gin.Context, req CreateOrderReq) (orderResponse *OrderResponse, err error) {
	i, err := strconv.Atoi(req.I)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf(">>>>>> strconv.Atoi is failed. I: %s", req.I)
		return
	}

	shopID := apiShopId
	nonce := time.Now().Unix()
	mapData := make(map[string]interface{}, 0)
	mapData["shopId"] = shopID
	mapData["nonce"] = nonce
	mapData["paymentId"] = req.PaymentId
	mapData["i"] = i
	mapData["email"] = req.Email
	mapData["ip"] = req.Ip
	mapData["amount"] = req.Amount
	mapData["currency"] = req.Currency
	mapData["successUrl"] = global.Config.PNSafePay.CallBackUrl
	mapData["signature"] = getSignature(mapData)
	for n := 1; n <= 10; n++ {
		orderResponse, err = createOrder(ctx, mapData)
		if err != nil {
			global.MyLogger(ctx).Info().Msgf("createOrder failed >>>>>>>>>>>>>>>>>> n = %d", n)
			time.Sleep(time.Duration(n*10) * time.Millisecond)
		}
		if err == nil {
			global.MyLogger(ctx).Info().Msgf(">>>>>> createOrder success, orderResponse: %#v", *orderResponse)
			return
		}
	}
	return
}
