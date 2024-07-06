package webmoney

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-speed/global"
	"golang.org/x/text/encoding/charmap"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type OperationsRequest struct {
	XMLName       xml.Name      `xml:"w3s.request"`
	Reqn          string        `xml:"reqn"`
	WmId          string        `xml:"wmid"`
	Sign          string        `xml:"sign"`
	GetOperations GetOperations `xml:"getoperations"`
}

type GetOperations struct {
	XMLName    xml.Name `xml:"getoperations"`
	Purse      string   `xml:"purse"`
	WmTranId   string   `xml:"wmtranid"`
	TranId     string   `xml:"tranid"`
	WmInvId    string   `xml:"wminvid"`
	OrderId    string   `xml:"orderid"`
	DateStart  string   `xml:"datestart"`
	DateFinish string   `xml:"datefinish"`
}

type Operation struct {
	Id        string `xml:"id,attr"`
	Ts        string `xml:"ts,attr"`
	Pursesrc  string `xml:"pursesrc"`
	Pursedest string `xml:"pursedest"`
	Amount    string `xml:"amount"`
	Comiss    string `xml:"comiss"`
	OperType  string `xml:"opertype"`
	WmInvId   string `xml:"wminvid"`
	OrderId   string `xml:"orderid"`
	TranId    string `xml:"tranid"`
	Period    string `xml:"period"`
	Desc      string `xml:"desc"`
	DateCrt   string `xml:"datecrt"`
	DateUpd   string `xml:"dateupd"`
	Corrwm    string `xml:"corrwm"`
	Rest      string `xml:"rest"`
	TimeLock  string `xml:"timelock"`
}

type Operations struct {
	XMLName   xml.Name    `xml:"operations"`
	Cnt       string      `xml:"cnt,attr"`
	Operation []Operation `xml:"operation"`
}

type OperationsResponse struct {
	XMLName    xml.Name   `xml:"w3s.response"`
	Reqn       string     `xml:"reqn"`
	Retval     string     `xml:"retval"`
	Retdesc    string     `xml:"retdesc"`
	Operations Operations `xml:"operations"`
}

func PrintXML(ctx *gin.Context, xmlBytes []byte) string {
	var data OperationsResponse
	decoder := xml.NewDecoder(bytes.NewReader(xmlBytes))
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if charset == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}
		return nil, fmt.Errorf("unsupported charset: %s", charset)
	}
	err := decoder.Decode(&data)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("xml.Unmarshal failed")
		return ""
	}

	formattedXML, err := xml.MarshalIndent(data, "", "  ")
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("xml.MarshalIndent failed")
		return ""
	}
	return string(formattedXML)
}

func sendOperationsRequest(ctx *gin.Context, request OperationsRequest) (*OperationsResponse, error) {
	fmt.Println(fmt.Sprintf("------------ >>> request:%#v\n", request))
	global.MyLogger(ctx).Info().Msgf("------------ >>> request:%+v", request)

	data, err := xml.Marshal(request)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("Marshal request failed")
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://w3s.webmoney.ru/asp/XMLOperations.asp", bytes.NewBuffer(data))
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("NewRequest failed")
		return nil, err
	}

	req.Header.Set("Content-Type", "text/xml")
	client := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
	resp, err := client.Do(req)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("client.Do failed")
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	fmt.Println("------------ >>> Response Body String:\n", string(body), "\n-----Body <<<")
	fmt.Println("------------ >>> Response Body XML:\n", PrintXML(ctx, body), "\n-----Body <<<")
	global.MyLogger(ctx).Info().Msgf("------------ >>> Response Body: %s", string(body))
	global.MyLogger(ctx).Info().Msgf("------------ >>> Response Body(XML): %s", PrintXML(ctx, body))

	// 创建一个新的 XML 解码器
	decoder := xml.NewDecoder(bytes.NewReader(body))

	// 设置自定义 CharsetReader
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if charset == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}
		return nil, fmt.Errorf("unsupported charset: %s", charset)
	}

	var response OperationsResponse
	err = decoder.Decode(&response)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("Decode failed")
		return nil, err
	}

	return &response, nil
}

func QueryOrder(ctx *gin.Context, orderId string) (out *Operation, err error) {
	//global.Config.WebMoneyConfig = config.WebMoneyConfig{
	//	WmId:     "283361774557",
	//	Purse:    "Z113494876653",
	//	RandCode: "pingguoqm23",
	//}
	//global.Config.WebMoneyConfig = config.WebMoneyConfig{
	//	WmId:     "601199792702",
	//	Purse:    "Z318054603350",
	//	RandCode: "yrj199007181872",
	//}
	//
	fmt.Println("------------ >>> WmId:", global.Config.WebMoneyConfig.WmId)
	request := OperationsRequest{
		Reqn: fmt.Sprintf("%d", time.Now().UnixNano()&0x1FFFFFFFFFFFFF),
		WmId: global.Config.WebMoneyConfig.WmId,
	}

	request.GetOperations.Purse = global.Config.WebMoneyConfig.Purse
	request.GetOperations.OrderId = orderId
	request.GetOperations.DateStart = time.Now().Add(-10 * time.Hour).Format("20060102 15:04:05")
	request.GetOperations.DateFinish = time.Now().Add(10 * time.Hour).Format("20060102 15:04:05")
	request.Sign, err = genSign(ctx, request.GetOperations.Purse+request.Reqn)
	var resp *OperationsResponse
	resp, err = sendOperationsRequest(ctx, request)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("sendOperationsRequest failed")
		return nil, err
	}

	if resp != nil {
		global.MyLogger(ctx).Info().Msgf(">>>>>>>>> OperationsResponse: %#v", *resp)
	}
	for _, i := range resp.Operations.Operation {
		if i.OrderId == orderId {
			fmt.Println(">>>>>>>> found orderId: ", orderId)
			return &i, nil
		}
	}
	return nil, nil
}
