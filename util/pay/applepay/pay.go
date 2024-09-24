package applepay

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/errors/gerror"
	"go-speed/global"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	EnvironmentSandbox = "Sandbox"
)

type VerifyReceiptResponse struct {
	Receipt     Receipt `json:"receipt"`
	Environment string  `json:"environment"`
	Status      int     `json:"status"`
}

type Receipt struct {
	ReceiptType                string      `json:"receipt_type"`
	AdamId                     int         `json:"adam_id"`
	AppItemId                  int         `json:"app_item_id"`
	BundleId                   string      `json:"bundle_id"`
	ApplicationVersion         string      `json:"application_version"`
	DownloadId                 int         `json:"download_id"`
	VersionExternalIdentifier  int         `json:"version_external_identifier"`
	ReceiptCreationDate        string      `json:"receipt_creation_date"`
	ReceiptCreationDateMs      string      `json:"receipt_creation_date_ms"`
	ReceiptCreationDatePst     string      `json:"receipt_creation_date_pst"`
	RequestDate                string      `json:"request_date"`
	RequestDateMs              string      `json:"request_date_ms"`
	RequestDatePst             string      `json:"request_date_pst"`
	OriginalPurchaseDate       string      `json:"original_purchase_date"`
	OriginalPurchaseDateMs     string      `json:"original_purchase_date_ms"`
	OriginalPurchaseDatePst    string      `json:"original_purchase_date_pst"`
	OriginalApplicationVersion string      `json:"original_application_version"`
	InApp                      []InAppItem `json:"in_app"`
}

type InAppItem struct {
	Quantity                string `json:"quantity"`
	ProductId               string `json:"product_id"`
	TransactionId           string `json:"transaction_id"`
	OriginalTransactionId   string `json:"original_transaction_id"`
	PurchaseDate            string `json:"purchase_date"`
	PurchaseDateMs          string `json:"purchase_date_ms"`
	PurchaseDatePst         string `json:"purchase_date_pst"`
	OriginalPurchaseDate    string `json:"original_purchase_date"`
	OriginalPurchaseDateMs  string `json:"original_purchase_date_ms"`
	OriginalPurchaseDatePst string `json:"original_purchase_date_pst"`
	IsTrialPeriod           string `json:"is_trial_period"`
	InAppOwnershipType      string `json:"in_app_ownership_type"`
}

type VerifyReceiptRequest struct {
	ReceiptData string `json:"receipt-data"`
}

func AppleVerify(ctx *gin.Context, transactionId, transactionReceipt string) (status int, err error) {
	var (
		url   string
		param VerifyReceiptRequest
	)
	if global.Config.ApplePayConfig.Environment == EnvironmentSandbox {
		url = "https://sandbox.itunes.apple.com/verifyReceipt"
	} else {
		url = "https://buy.itunes.apple.com/verifyReceipt"
	}
	global.MyLogger(ctx).Info().Msgf("apple pay request url: %s", url)

	// 苹果验证
	param.ReceiptData = transactionReceipt
	client := &http.Client{Timeout: 15 * time.Second}
	jsonStr, _ := json.Marshal(param)
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	//·21000 App Store 无法读取你提供的JSON数据
	//·21002 收据数据不符合格式
	//·21003 收据无法被验证
	//·21004 你提供的共享密钥和账户的共享密钥不一致
	//·21005 收据服务器当前不可用
	//·21006 收据是有效的，但订阅服务已经过期。当收到这个信息时，解码后的收据信息也包含在返回内容中
	//·21007 收据信息是测试用（sandbox），但却被发送到产品环境中验证
	//·21008 收据信息是产品环境中使用，但却被发送到测试环境中验证
	if err != nil {
		err = gerror.Wrap(err, `http.Client.Post failed`)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	global.MyLogger(ctx).Info().Msgf("apple pay response: %s", string(body))

	var response VerifyReceiptResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		err = gerror.Wrap(err, `response json.Unmarshal failed`)
		return
	}

	if len(response.Receipt.InApp) <= 0 {
		err = gerror.New(`response empty Receipt`)
		return
	}

	global.MyLogger(ctx).Info().Msgf("apple pay response BundleId(%s), Config(%s)", response.Receipt.BundleId, global.Config.ApplePayConfig.BundleId)
	if response.Receipt.BundleId != global.Config.ApplePayConfig.BundleId {
		err = gerror.New(`response "BundleId" invalid`)
		return
	}

	global.MyLogger(ctx).Info().Msgf("apple pay response Environment(%s), Config(%s)", response.Environment, global.Config.ApplePayConfig.Environment)
	if response.Environment != global.Config.ApplePayConfig.Environment {
		err = gerror.New(`response "Environment" invalid`)
		return
	}

	var appItem *InAppItem
	for _, item := range response.Receipt.InApp {
		if item.TransactionId == transactionId {
			appItem = &item
			break
		}
	}
	if appItem == nil {
		err = gerror.New(`transactionId invalid`)
		return
	}

	global.MyLogger(ctx).Info().Msgf("apple pay response Status(%d)", response.Status)
	return response.Status, nil
}
