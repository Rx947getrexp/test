package russpay

import (
	"bytes"
	"crypto/md5"
	"crypto/tls"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"go-speed/constant"
	"go-speed/global"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	AppKey    = "81001607"
	SecretKey = "5550be23ed414cf0a9b011175504105a"
)

type CreateOrderReq struct {
	ChannelId   string
	Amount      string `json:"amount" dc:"支付余额"`
	OrderNumber string `json:"orderNumber" dc:"订单号"`
	CompanyPage string `json:"companyPage" dc:"支付页面完成以后，跳转到CP 的结果页地址"`
	DeviceType  string `json:"deviceType" dc:"设备类型,MOBILE 或者 DESKTOP"`
}

type RussPayReq struct {
	AppKey  string `json:"appKey"`
	Sign    string `json:"sign"`
	Content string `json:"content"`
}

type Content struct {
	Amount           string `json:"amount" dc:"支付余额"`
	OrderNumber      string `json:"orderNumber" dc:"订单号"`
	CompanyPage      string `json:"companyPage" dc:"支付页面完成以后，跳转到CP 的结果页地址"`
	CountryCode      string `json:"countryCode" dc:"国家码"`
	Currency         string `json:"currency" dc:"货币单位"`
	OrderDescription string `json:"orderDescription" dc:"订单描述"`
}

type requestRussPayRes struct {
	Status    string   `json:"status"`
	Message   string   `json:"message"`
	ErrorCode string   `json:"errorCode"`
	Content   Content1 `json:"content"`
}

type Content1 struct {
	Content string `json:"content"`
	Sign    string `json:"sign"`
}

type CreateOrderResponse struct {
	PayUrl        string `json:"payUrl"`
	OrderNumber   string `json:"orderNumber"`
	BillingNumber string `json:"billingNumber"`
	InvoiceId     string `json:"invoiceId"`
}

func requestRussPay(ctx *gin.Context, createOrderUrl string, mapData map[string]interface{}) (res *requestRussPayRes, err error) {
	global.MyLogger(ctx).Info().Msgf(">>>>>> 请求地址: %#v", createOrderUrl)
	fmt.Println(createOrderUrl)
	global.MyLogger(ctx).Info().Msgf(">>>>>> 请求参数: %#v", mapData)
	fmt.Println(mapData)
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

	if strings.ToLower(res.Status) != "success" {
		err = fmt.Errorf("russpay return failed, %s, %s", res.ErrorCode, res.Message)
		global.MyLogger(ctx).Err(err).Msgf("status is not success")
		return
	}
	return
}

func getOSFromUserAgent(ctx *gin.Context, deviceType string) string {
	deviceType = strings.ToLower(strings.TrimSpace(deviceType))
	if deviceType == "" {
		userAgent := strings.ToLower(global.GetUserAgent(ctx))
		if strings.Contains(userAgent, "iphone") || strings.Contains(userAgent, "ipad") {
			deviceType = "iphone"
		} else if strings.Contains(userAgent, "android") {
			deviceType = "android"
		} else if strings.Contains(userAgent, "windows") {
			deviceType = "windows"
		} else if strings.Contains(userAgent, "mac os") {
			deviceType = "mac"
		} else {
			deviceType = "unknown"
		}
	}

	if deviceType == "windows" || deviceType == "mac" {
		return "DESKTOP"
	} else {
		return "MOBILE"
	}
}

func CreateOrder(ctx *gin.Context, req CreateOrderReq) (resp *CreateOrderResponse, err error) {
	mapReq := make(map[string]interface{})
	mapReq["appKey"] = AppKey

	mapContent := make(map[string]interface{})
	mapContent["amount"] = req.Amount
	mapContent["orderNumber"] = req.OrderNumber
	mapContent["companyPage"] = req.CompanyPage
	mapContent["countryCode"] = "RU"
	mapContent["currency"] = "RUB"
	if req.ChannelId == constant.PayChannelRussPaySBER {
		mapContent["deviceType"] = getOSFromUserAgent(ctx, req.DeviceType)
	}

	contentBase64, err := Base64Encode(mapContent)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("Base64Encode failed")
		return
	}
	mapReq["content"] = contentBase64
	mapReq["sign"] = MD5(contentBase64, SecretKey)

	var url string
	switch req.ChannelId {
	case constant.PayChannelRussPayBankCard:
		url = "https://api.yolesdk.com/v2/api/ruBankCard/transaction/pay"
	case constant.PayChannelRussPaySBP:
		url = "https://api.yolesdk.com/v2/api/ruBankCard/transaction/payBySbp"
	case constant.PayChannelRussPaySBER:
		url = "https://api.yolesdk.com/v2/api/ruBankCard/transaction/payBySber"
	}

	for n := 1; n <= 1; n++ {
		var response *requestRussPayRes
		response, err = requestRussPay(ctx, url, mapReq)
		if err != nil {
			fmt.Println("-----> err: ", err.Error())
			global.MyLogger(ctx).Info().Msgf("createOrder failed >>>>>>>>>>>>>>>>>> n = %d", n)
			time.Sleep(time.Duration(n*10) * time.Millisecond)
		}
		if err == nil {
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

			global.MyLogger(ctx).Info().Msgf(">>>>>> createOrder success, orderResponse: %#v", *resp)
			return
		}
	}
	return
}

func Base64Decode(content string) (out []byte, err error) {
	out, err = base64.StdEncoding.DecodeString(content)
	if err != nil {
		return
	}
	return
}

func Base64Encode(data map[string]interface{}) (out string, err error) {
	contentBytes, err := json.Marshal(data)
	if err != nil {
		return
	}

	out = base64.StdEncoding.EncodeToString(contentBytes)
	return
}

func MD5(content, secretKey string) string {
	data := content + secretKey
	hash := md5.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}
