package pnsafepay

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"go-speed/global"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

const (
	PaySecretKey      = "97a8df4fadfb613f0b4f0611c7dfc826"
	MethodTradeCreate = "trade.create"
	MethodTradeCheck  = "trade.check"
	PnSafePayUrl      = "http://api.pnsafepay.com/gateway.aspx"
)

type PayRequest struct {
	MerNo       string `json:"mer_no"`
	OrderNo     string `json:"order_no"`
	OrderAmount string `json:"order_amount"`
	PayName     string `json:"payname"`
	PayEmail    string `json:"payemail"`
	PayPhone    string `json:"payphone"`
	Currency    string `json:"currency"`
	PayTypeCode string `json:"paytypecode"`
	Method      string `json:"method"`
	ReturnUrl   string `json:"returnurl"`
	Sign        string `json:"sign"`
}

type PayResponse struct {
	MerNo           string `json:"mer_no"`
	OrderNo         string `json:"order_no"`
	OrderAmount     string `json:"order_amount"`
	Status          string `json:"status"`
	StatusMes       string `json:"status_mes"`
	OrderData       string `json:"order_data"`
	InterIsNeedGift bool   `json:"-"`
	InterMsg        string `json:"-"`
}

func CreatePayOrder(ctx *gin.Context, req *PayRequest) (res *PayResponse, err error) {
	req.Method = MethodTradeCreate
	req.ReturnUrl = global.Config.PNSafePay.CallBackUrl
	requestParams := genRequestSignature(ctx, req)
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	jsonData, err := json.Marshal(requestParams)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("JSON编码失败")
		return
	}

	global.MyLogger(ctx).Info().Msgf(">>>>>>>>>>>> pnsafepay request: %s", gjson.MustEncode(requestParams))
	response, err := g.Client().Post(ctx, PnSafePayUrl, jsonData, headers)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("请求pnsafepay失败")
		return nil, err
	}
	defer response.Close()
	global.MyLogger(ctx).Debug().Msgf("pnsafepay trade.create response: %+v", *response)

	if response.StatusCode != 200 {
		global.MyLogger(ctx).Err(err).Msgf("pnsafepay trade.create failed, StatusCode: %d, response: %+v",
			response.StatusCode, *response)
		var payResponse *PayResponse
		if response.StatusCode == 502 || response.StatusCode == 504 {
			payResponse = &PayResponse{
				InterIsNeedGift: true,
				InterMsg:        fmt.Sprintf("StatusCode is %d", response.StatusCode),
			}
		}
		return payResponse, gerror.Newf("StatusCode: %d != 200", response.StatusCode)
	}

	content := response.ReadAll()
	global.MyLogger(ctx).Info().Msgf(">>>>>>>>>>>> pnsafepay response content: %s", string(content))

	payResponse := &PayResponse{}
	err = json.Unmarshal(content, payResponse)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("pnsafepay Unmarshal failed")
		return nil, err
	}
	return payResponse, nil
}

func genRequestSignature(ctx *gin.Context, req *PayRequest) map[string]string {
	params := map[string]string{
		"currency":     req.Currency,
		"mer_no":       global.Config.PNSafePay.MerNo,
		"method":       req.Method,
		"order_amount": req.OrderAmount,
		"order_no":     req.OrderNo,
		"payemail":     req.PayEmail,
		"payname":      req.PayName,
		"payphone":     req.PayPhone,
		"paytypecode":  req.PayTypeCode,
		"returnurl":    req.ReturnUrl,
	}
	secretKey := PaySecretKey
	signature := generateSignature(ctx, params, secretKey)
	fmt.Println("Signature:", signature)
	params["sign"] = signature
	return params
}

func generateSignature(ctx *gin.Context, params map[string]string, secretKey string) string {
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var buf strings.Builder
	for i, k := range keys {
		v := params[k]
		if v == "" {
			continue
		}
		buf.WriteString(k)
		buf.WriteString("=")
		//buf.WriteString(url.QueryEscape(v))
		buf.WriteString(v)
		if i < len(keys)-1 {
			buf.WriteString("&")
		}
	}
	buf.WriteString(secretKey)
	global.MyLogger(ctx).Info().Msgf(">>>>>>>>>>>> md5 input content: %s", buf.String())

	hasher := md5.New()
	hasher.Write([]byte(buf.String()))
	return hex.EncodeToString(hasher.Sum(nil))
}

type QueryOrderResponse struct {
	MerNo              string `json:"mer_no"`
	OrderNo            string `json:"order_no"`
	CheckStatus        string `json:"checkstatus"`
	ResultStatus       string `json:"resultstatus"`
	OrderAmount        string `json:"order_amount"`
	OrderRealityAmount string `json:"order_realityamount"`
}

func QueryPayOrder(ctx *gin.Context, orderNo string) (res *QueryOrderResponse, err error) {
	params := map[string]string{
		"mer_no":   global.Config.PNSafePay.MerNo,
		"method":   MethodTradeCheck,
		"order_no": orderNo,
	}
	params["sign"] = generateSignature(ctx, params, PaySecretKey)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	jsonData, err := json.Marshal(params)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("JSON编码失败")
		return
	}

	global.MyLogger(ctx).Info().Msgf(">>>>>>>>>>>> pnsafepay params: %s", gjson.MustEncode(params))
	response, err := g.Client().Post(ctx, PnSafePayUrl, jsonData, headers)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("请求pnsafepay失败")
		return nil, err
	}
	defer response.Close()

	if response.StatusCode != 200 {
		global.MyLogger(ctx).Err(err).Msgf("pnsafepay trade.create failed, response: %+v", *response)
		return nil, gerror.Newf("StatusCode: %d != 200", response.StatusCode)
	}

	content := response.ReadAll()
	global.MyLogger(ctx).Info().Msgf(">>>>>>>>>>>> pnsafepay response content: %s", string(content))

	payResponse := &QueryOrderResponse{}
	err = json.Unmarshal(content, payResponse)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("pnsafepay Unmarshal failed")
		return nil, err
	}
	return payResponse, nil
}
