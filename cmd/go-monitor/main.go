package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/os/glog"
	"go-speed/global"
	"go-speed/initialize"
	"go-speed/service/email"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"
)

var dnsList = []string{
	"eigrrht.xyz",
	"siaax.xyz",
	"beiyo.xyz",
	"thertee.xyz",
	"weechat.xyz",
	"2yiny.xyz",
	"yinyong.xyz",
}

func main() {
	glog.SetLevel(glog.LEVEL_ALL)
	initialize.InitComponentsV2()
	errCounter := make(map[string]uint64)
	for _, dns := range dnsList {
		errCounter[dns] = 0
	}
	var (
		times        uint64
		failedFlag   bool
		successTimes uint64 = 1
	)
	for {
		global.Logger.Info().Msgf("------------------------------ [%d] ----------------------------------------", times)
		times++
		failedFlag = false
		global.Logger.Info().Msgf("-------- begin-to-dial-dns -------- [times: %d]", times)
		for _, dns := range dnsList {
			ctx := context.Background()
			err := dialAPIServer(ctx, fmt.Sprintf("https://%s/app-api/dns_list", dns))
			if err != nil {
				failedFlag = true
				errCounter[dns]++
			} else {
				errCounter[dns] = 0
			}
			global.Logger.Info().Msgf("dns: %s, errCounter: %d", dns, errCounter[dns])
			if errCounter[dns] > 5 {
				sendAlarm(dns, err)
			}
		}
		if !failedFlag {
			successTimes++
		}
		if successTimes%60 == 0 {
			sendSuccessNotify()
		}
		global.Logger.Info().Msgf("-------- finished-once-dial -------- [times: %d, successTimes: %d]", times, successTimes)
		global.Logger.Info().Msgf("begin to sleep, 1 minute ...")
		global.Logger.Info().Msgf("----------------------------------------------------------------------------")
		time.Sleep(time.Minute * 1)
	}
}

func sendAlarm(dns string, err error) {
	var builder strings.Builder
	// 添加邮件内容
	builder.WriteString(fmt.Sprintf("<h1>拨测对象：%s (%s)</h1>\n", dns, time.Now().Format(time.DateTime)))
	builder.WriteString(fmt.Sprintf("<h1>报错信息：%s</h1>\n", err.Error()))
	builder.WriteString("\n")
	builder.WriteString("<h1>请赶紧处理！</h1>\n")

	emailSubject := "应用服务器拨测失败告警通知"
	fmt.Println(builder.String())
	_ = sendEmail(emailSubject, builder.String())
}

func sendEmail(emailSubject, emailBody string) (err error) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	defer ctx.Done()
	var (
		fromAccount   = "heronetwork@herovpn.live"
		fromPasswd    = "pingguoqm23"
		fromHost      = "smtpout.secureserver.net:465"
		alarmReceiver = []string{
			"xiaomingchuan1990@gmail.com",
		}
	)
	email.SetSendAccount(fromAccount, fromPasswd, fromHost)
	for i := 0; i < 10; i++ {
		err = email.SendEmailTLS(ctx, emailSubject, emailBody, alarmReceiver)
		if err == nil {
			return
		}
		time.Sleep(time.Second * 2)
	}
	return
}

func sendSuccessNotify() {
	emailSubject := "应用服务器拨测成功定时通知"
	emailBody := fmt.Sprintf("<h1>应用服务器和nginx服务器拨测成功 (%s)</h1>", time.Now().Format(time.DateTime))
	_ = sendEmail(emailSubject, emailBody)
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type Data struct {
	List []DNSItem `json:"list"`
}

type DNSItem struct {
	DNS      string `json:"dns"`
	ID       int    `json:"id"`
	SiteType int    `json:"site_type"`
}

func dialAPIServer(c context.Context, url string) (err error) {
	ctx, cancel := context.WithTimeout(c, 30*time.Second)
	defer cancel()

	global.Logger.Info().Msgf("url: %s", url)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		glog.Errorf(ctx, "NewRequest failed, err: %s", err.Error())
		return
	}

	// 设置请求头部
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Lang", "cn")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		global.Logger.Error().Msgf("Do failed, err: %s", err.Error())
		return
	}
	defer resp.Body.Close()

	// 检查HTTP响应状态码
	if resp.StatusCode != http.StatusOK {
		global.Logger.Error().Msgf("HTTP请求失败，状态码: %d\n", resp.StatusCode)
		return
	}

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		global.Logger.Error().Msgf("ReadAll failed, err: %s", err.Error())
		return
	}

	// 打印响应内容
	//fmt.Println(string(body))
	// 解析JSON响应
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		global.Logger.Error().Msgf("json.Unmarshal failed, err: %s", err.Error())
		return
	}

	// 检查业务code和message
	if response.Code != 200 || response.Message != "成功" {
		err = fmt.Errorf("业务处理失败，code:%d, message:%s", response.Code, response.Message)
		global.Logger.Error().Msgf(err.Error())
		return
	}

	// 检查data.list是否有值
	if len(response.Data.List) == 0 {
		err = fmt.Errorf("data.list为空")
		global.Logger.Error().Msgf(err.Error())
		return
	}
	global.Logger.Info().Msgf("%s 拨测成功!", url)

	// 打印解析后的数据
	//for _, item := range response.Data.List {
	//	fmt.Printf("DNS: %s, ID: %d, SiteType: %d\n", item.DNS, item.ID, item.SiteType)
	//}
	return
}
