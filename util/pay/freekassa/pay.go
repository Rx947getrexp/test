package freekassa

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"go-speed/global"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type OrderListRequest struct {
	ShopID      int    `json:"shopId"`
	Nonce       int64  `json:"nonce"`
	Signature   string `json:"signature"`
	OrderID     int    `json:"orderId,omitempty"`
	PaymentID   string `json:"paymentId,omitempty"`
	OrderStatus int    `json:"orderStatus,omitempty"`
	DateFrom    string `json:"dateFrom,omitempty"`
	DateTo      string `json:"dateTo,omitempty"`
	Page        int    `json:"page,omitempty"`
}

type OrderListResponse struct {
	Type   string  `json:"type"`
	Pages  int     `json:"pages"`
	Orders []Order `json:"orders"`
}

type Order struct {
	MerchantOrderId string  `json:"merchant_order_id"`
	FkOrderId       int     `json:"fk_order_id"`
	Amount          float64 `json:"amount"`
	Currency        string  `json:"currency"`
	Email           string  `json:"email"`
	Account         string  `json:"account"`
	Date            string  `json:"date"`
	Status          int     `json:"status"`
}

var apiKey = "19c44da8067cdbe4bc7f0b4db68891d3"
var apiShopId = 52928

func QueryOrder(ctx *gin.Context, orderId string) (order *Order, err error) {
	shopID := apiShopId
	nonce := time.Now().Unix()
	mapData := make(map[string]interface{}, 0)
	mapData["shopId"] = shopID
	mapData["nonce"] = nonce
	mapData["paymentId"] = orderId
	//mapData["dateFrom"] = time.Now().Add(-5 * time.Hour * 24).Format("2006-01-02 15:04:05")
	//mapData["dateTo"] = time.Now().Format("2006-01-02 15:04:05")
	//mapData["page"] = 1
	mapData["signature"] = getSignature(mapData)
	fmt.Println("请求地址:", "https://api.freekassa.com/v1/orders")
	fmt.Println("请求参数:", mapData)
	var response *OrderListResponse
	for n := 1; n <= 10; n++ {
		response, err = getOrderList(ctx, mapData)
		if err == nil {
			global.MyLogger(ctx).Info().Msgf("getOrderList failed >>>>>>>>>>>>>>>>>> n = %d", n)
			break
		}
		time.Sleep(time.Duration(n*10) * time.Millisecond)
	}
	if err != nil {
		return
	}

	if response.Type != "success" {
		err = fmt.Errorf("response.Type: (%s) is not success", response.Type)
		global.MyLogger(ctx).Err(err).Msg(">>>>>> response is not ok")
		return
	}
	for _, i := range response.Orders {
		if i.MerchantOrderId == orderId {
			global.MyLogger(ctx).Info().Msgf("############### find MerchantOrder: %#v", i)
			return &i, nil
		}
	}
	return nil, nil
}

func getSignature(data map[string]interface{}) string {
	keys := make([]string, 0, len(data))
	for key := range data {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	msg := ""
	for _, key := range keys {
		msg += fmt.Sprintf("%v|", data[key])
	}
	msg = strings.TrimSuffix(msg, "|")

	hash := hmac.New(sha256.New, []byte(apiKey))
	hash.Write([]byte(msg))
	return hex.EncodeToString(hash.Sum(nil))
}

func createSignature(data, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func getOrderList(ctx *gin.Context, mapData map[string]interface{}) (response *OrderListResponse, err error) {
	apiURL := "https://api.freekassa.com/v1/orders"

	jsonData, err := json.Marshal(mapData)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msg("json.Marshal failed")
		return
	}
	global.MyLogger(ctx).Info().Msgf(">>>>>>>>> jsonData: %s", string(jsonData))

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		global.MyLogger(ctx).Err(err).Msg("http.NewRequest failed")
		return
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
	resp, err := client.Do(req)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msg("http.Do failed")
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	global.MyLogger(ctx).Info().Msgf(">>>>>> response: %s", string(body))

	defer resp.Body.Close()
	if err != nil {
		global.MyLogger(ctx).Err(err).Msg("ioutil.ReadAll failed")
		return
	}
	//
	//resp, err := http.Post(apiURL, "application/json", strings.NewReader(string(jsonData)))
	//if err != nil {
	//	fmt.Println("Error making request:", err)
	//	return
	//}
	//defer resp.Body.Close()
	//
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println("Error reading response body:", err)
	//	return
	//}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("Request failed with status code %d and response: %s\n", resp.StatusCode, string(body))
		global.MyLogger(ctx).Err(err).Msg(">>>>>> http StatusCode is not ok")
		return
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msg("json.Unmarshal failed")
		return
	}

	global.MyLogger(ctx).Info().Msg(">>>>>>>> query success >>>>>>>")
	return
}
