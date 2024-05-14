package upay

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"go-speed/global"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Deposit struct {
	ID            string `json:"id"`
	Amount        string `json:"amount"`
	Coin          string `json:"coin"`
	Network       string `json:"network"`
	Status        int    `json:"status"`
	Address       string `json:"address"`
	AddressTag    string `json:"addressTag"`
	TxId          string `json:"txId"`
	InsertTime    int64  `json:"insertTime"`
	TransferType  int    `json:"transferType"`
	ConfirmTimes  string `json:"confirmTimes"`
	UnlockConfirm int    `json:"unlockConfirm"`
	WalletType    int    `json:"walletType"`
}

func QueryPayOrder(ctx *gin.Context, minutesAgo time.Duration) (deposits []Deposit, err error) {
	// 填写你的API密钥
	apiKey := "nDjKls9WZvFlrw6GHNptvEJ8ipCOmd9aLvCf6MpjzrWQskH6mnH3BLZg1YVGtOgN"
	apiSecret := "mTwKB9rsRzBvsMHz9MUPM9hKu8KTeDe0mRstjIXY9fhvMDy9HHtkCwUqDORzlqqJ"

	// 构造请求参数
	params := map[string]string{
		"coin":      "",                                                                  // 可选，填写币种，不填则返回所有币种的充值历史
		"status":    "",                                                                  // 可选，0表示pending，1表示success，不填则返回所有状态的充值历史
		"startTime": fmt.Sprintf("%d", time.Now().Add(-1*time.Minute*minutesAgo).Unix()), // 默认当前时间90天的时间戳
		"endTime":   "",
		"limit":     "1000",                                                               // 默认1000，最大1000
		"timestamp": strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10), // 必填，以毫秒为单位的UNIX时间戳
	}

	// 生成签名
	queryString := ""
	for k, v := range params {
		queryString += k + "=" + v + "&"
	}
	queryString = queryString[:len(queryString)-1] // 移除最后一个'&'

	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write([]byte(queryString))
	signature := hex.EncodeToString(h.Sum(nil))

	// 发送请求
	url := fmt.Sprintf("https://api.binance.com/sapi/v1/capital/deposit/hisrec?%s&signature=%s", queryString, signature)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-MBX-APIKEY", apiKey)
	client := &http.Client{}

	global.MyLogger(ctx).Debug().Msgf(">>>>>>>>>>>> binance params: %s", gjson.MustEncode(params))

	var resp *http.Response
	resp, err = client.Do(req)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("request binance failed")
		return
	}
	defer resp.Body.Close()
	global.MyLogger(ctx).Debug().Msgf(">>>>>>>>>>>> binance response deposits: %s", gjson.MustEncode(resp))

	// 处理响应
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		err = gerror.Newf("StatusCode: %d != 200", resp.StatusCode)
		global.MyLogger(ctx).Err(err).Msgf("binance query order failed, response: %+v", gjson.MustEncode(resp))
		return
	}

	deposits = make([]Deposit, 0)
	err = json.Unmarshal(body, &deposits)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("binance Unmarshal failed")
		return
	}
	return
}

func CheckBinanceOrder(ctx *gin.Context, minutesAgo time.Duration, amount string) (found bool, err error) {
	var (
		deposits    []Deposit
		amountFloat float64
	)
	amountFloat, err = strconv.ParseFloat(amount, 64)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("amount ParseFloat failed, (%s)", amount)
		return
	}
	deposits, err = QueryPayOrder(ctx, minutesAgo)
	if err != nil {
		return
	}
	for _, deposit := range deposits {
		depositTime := time.Unix(deposit.InsertTime/1000, 0).Format("2006-01-02 15:04:05")
		global.MyLogger(ctx).Debug().Msgf("充值时间: %s, 充值id: %s, 充值金额: %s %s, 充值网络地址: %s, 充值状态: %d",
			depositTime, deposit.ID, deposit.Amount, deposit.Coin, deposit.Network, deposit.Status)

		value, _err := strconv.ParseFloat(deposit.Amount, 64)
		if _err != nil {
			global.MyLogger(ctx).Err(err).Msgf("binance ParseFloat failed, (%s)", deposit.Amount)
			continue
		}
		if value == amountFloat {
			return true, nil
		}
	}
	global.MyLogger(ctx).Debug().Msgf("can not found Amount(%f) order", amount)
	return false, nil
}
