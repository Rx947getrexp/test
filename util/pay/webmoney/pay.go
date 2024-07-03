package webmoney

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-speed/global"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/baibaratsky/go-wmsigner"
	"golang.org/x/text/encoding/charmap"
)

//const (
//	WMID     = "283361774557"
//	PURSE    = "Z113494876653"
//	RandCode = "pingguoqm23"
//)

type OutInvoicesRequest struct {
	XMLName        xml.Name       `xml:"w3s.request"`
	Reqn           int64          `xml:"reqn"`
	WMID           string         `xml:"wmid"`
	Sign           string         `xml:"sign"`
	GetOutInvoices GetOutInvoices `xml:"getoutinvoices"`
}

type OutInvoicesResponse struct {
	XMLName     xml.Name    `xml:"w3s.response"`
	Reqn        int64       `xml:"reqn"`
	Retval      int         `xml:"retval"`
	Retdesc     string      `xml:"retdesc"`
	OutInvoices OutInvoices `xml:"outinvoices"`
}

type GetOutInvoices struct {
	Purse      string `xml:"purse"`
	WMInvid    int64  `xml:"wminvid,omitempty"`
	OrderID    int64  `xml:"orderid,omitempty"`
	DateStart  string `xml:"datestart"`
	DateFinish string `xml:"datefinish"`
}

type OutInvoice struct {
	ID           int64   `xml:"id,attr"`
	Ts           int64   `xml:"ts,attr"`
	OrderID      int64   `xml:"orderid"`
	CustomerWMID string  `xml:"customerwmid"`
	StorePurse   string  `xml:"storepurse"`
	Amount       float64 `xml:"amount"`
	Desc         string  `xml:"desc"`
	Address      string  `xml:"address"`
	Period       int     `xml:"period"`
	Expiration   int     `xml:"expiration"`
	State        int     `xml:"state"`
	Datecrt      string  `xml:"datecrt"`
	Dateupd      string  `xml:"dateupd"`
	WMTranID     int64   `xml:"wmtranid,omitempty"`
}

type OutInvoices struct {
	Cnt        int          `xml:"cnt,attr"`
	OutInvoice []OutInvoice `xml:"outinvoice"`
}

//
//func (this *OutInvoicesResponse) Print() string {
//	return fmt.Sprintf("订单信息：OutInvoicesResponse:\n"+
//		"Reqn:%d\n"+
//		"Retval:%d\n"+
//		"Retdesc:%s\n"+
//		"OutInvoices:{\n"+
//		"  Cnt:%d\n"+
//		"  OutInvoice:[\n"+
//		"               ID:\t\t%d\n"+
//		"               Ts:\t\t%d\n"+
//		"               OrderID:\t%d\n"+
//		"               CustomerWMID:\t%s\n"+
//		"               StorePurse:\t%s\n"+
//		"               Amount:\t\t%f\n"+
//		"               Desc:\t\t%s\n"+
//		"               Address:\t%s\n"+
//		"               Period:\t\t%d\n"+
//		"               Expiration:\t%d\n"+
//		"               State:\t\t%d\n"+
//		"               Datecrt:\t%s\n"+
//		"               Dateupd:\t%s\n"+
//		"               WMTranID:\t%d\n"+
//		"       ]\n"+
//		"}",
//		this.Reqn,
//		this.Retval,
//		this.Retdesc,
//		this.OutInvoices.Cnt,
//		this.OutInvoices.OutInvoice[0].ID,
//		this.OutInvoices.OutInvoice[0].Ts,
//		this.OutInvoices.OutInvoice[0].OrderID,
//		this.OutInvoices.OutInvoice[0].CustomerWMID,
//		this.OutInvoices.OutInvoice[0].StorePurse,
//		this.OutInvoices.OutInvoice[0].Amount,
//		this.OutInvoices.OutInvoice[0].Desc,
//		this.OutInvoices.OutInvoice[0].Address,
//		this.OutInvoices.OutInvoice[0].Period,
//		this.OutInvoices.OutInvoice[0].Expiration,
//		this.OutInvoices.OutInvoice[0].State,
//		this.OutInvoices.OutInvoice[0].Datecrt,
//		this.OutInvoices.OutInvoice[0].Dateupd,
//		this.OutInvoices.OutInvoice[0].WMTranID,
//	)
//}

func sendOutInvoicesRequest(ctx *gin.Context, request OutInvoicesRequest) (*OutInvoicesResponse, error) {
	global.MyLogger(ctx).Info().Msgf("------------ >>> request:%+v", request)
	url := "https://w3s.webmoney.ru/asp/XMLOutInvoices.asp"

	data, err := xml.Marshal(request)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("Marshal request failed")
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "text/xml")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("client.Do failed")
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	fmt.Println("------------ >>> Response Body:\n", string(body), "\n-----Body <<<")
	global.MyLogger(ctx).Info().Msgf("------------ >>> Response Body: %s", string(body))

	// 创建一个新的 XML 解码器
	input := bytes.NewReader(body)
	decoder := xml.NewDecoder(input)

	// 设置自定义 CharsetReader
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if charset == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}
		return nil, fmt.Errorf("unsupported charset: %s", charset)
	}

	var response OutInvoicesResponse
	err = decoder.Decode(&response)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("Decode failed")
		return nil, err
	}

	//var response OutInvoicesResponse
	//err = xml.NewDecoder(resp.Body).Decode(&response)
	//if err != nil {
	//      return nil, err
	//}

	return &response, nil
}

func QueryDeal(ctx *gin.Context, orderId string) (out *OutInvoice, err error) {
	request := OutInvoicesRequest{
		Reqn: time.Now().UnixNano() & 0x1FFFFFFFFFFFFF,
		WMID: global.Config.WebMoneyConfig.WmId,
	}

	orderIdInt, err := strconv.ParseInt(orderId, 10, 64)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("orderId(%s) ParseInt failed", orderId)
		return
	}

	request.GetOutInvoices.Purse = global.Config.WebMoneyConfig.Purse
	//request.GetOutInvoices.OrderID = orderIdInt
	request.GetOutInvoices.DateStart = time.Now().Add(-30 * time.Hour).Format("20060102 15:04:05")
	request.GetOutInvoices.DateFinish = time.Now().Format("20060102 15:04:05")
	request.Sign, err = genSign(ctx, request.GetOutInvoices.Purse+strconv.FormatInt(request.Reqn, 10))
	var resp *OutInvoicesResponse
	resp, err = sendOutInvoicesRequest(ctx, request)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("sendOutInvoicesRequest failed")
		return nil, err
	}

	if resp != nil {
		global.MyLogger(ctx).Info().Msgf(">>>>>>>>> OutInvoicesResponse: %#v", *resp)
	}
	for _, i := range resp.OutInvoices.OutInvoice {
		if i.OrderID == orderIdInt {
			return &i, nil
		}
	}
	return nil, nil
}

func genSign(ctx *gin.Context, data string) (sign string, err error) {
	var fileName string
	fileName = fmt.Sprintf("/wwwroot/go/go-api/config/%s.kwm", global.Config.WebMoneyConfig.WmId)
	//fileName = fmt.Sprintf("/Users/Shared/src/hs/go-speed/manifest/config/%s.kwm", global.Config.WebMoneyConfig.WmId)
	signer, err := wmsigner.New(global.Config.WebMoneyConfig.WmId, fileName, global.Config.WebMoneyConfig.RandCode)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("wmsigner New failed")
		return "", err
	}

	sign, err = signer.Sign(data)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("signer.Sign failed")
		return "", err
	}
	global.MyLogger(ctx).Info().Msgf("sign: %s", sign)
	return
}
