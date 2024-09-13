package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/os/glog"
	"go-speed/service/email"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"time"
)

var dnsList = []string{
	"eigrrht.xyz",
	"siaax.xyz",
	"beiyo.xyz",
	"beiy1o.xyz",
}

func main() {
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
		times++
		failedFlag = false
		glog.Infof(context.Background(), ">>>>>>>> begin to dial, times: %d", times)
		for _, dns := range dnsList {
			ctx := context.Background()
			err := dialAPIServer(ctx, fmt.Sprintf("https://%s/app-api/dns_list", dns))
			if err != nil {
				failedFlag = true
				errCounter[dns]++
			} else {
				errCounter[dns] = 0
			}
			glog.Infof(ctx, "dns: %s, errCounter: %d", dns, errCounter[dns])
			if errCounter[dns] > 0 {
				sendAlarm(dns, err)
			}
		}
		if !failedFlag {
			successTimes++
		}
		if successTimes%12 == 0 {
			sendSuccessNotify()
		}
		glog.Infof(context.Background(), ">>>>>>>> finished one time dial, times: %d, successTimes: %d, going to sleep...\n\n", times, successTimes)
		time.Sleep(time.Second * 5)
	}
}

var alarmReceiver = []string{}

func sendAlarm(dns string, err error) {
	emailSubject := "应用服务器拨测失败告警通知"
	emailBody := fmt.Sprintf("拨测对象：%s\n报错信息：%s", dns, err.Error())
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	defer ctx.Done()
	account := map[string]map[string]string{
		"vpnheroes@outlook.com":  {"pw": "pingguoqm23", "host": "smtp-mail.outlook.com:587"},
		"heroesvpnn@outlook.com": {"pw": "pingguoqm23", "host": "smtp-mail.outlook.com:587"},
		"VPNHERO@outlook.com":    {"pw": "pingguoqm23", "host": "smtp-mail.outlook.com:587"},
	}
	accounts := []string{"vpnheroes@outlook.com", "heroesvpnn@outlook.com", "VPNHERO@outlook.com"}
	for _, userName := range accounts {
		email.SetSendAccount(userName, account[userName]["pw"], account[userName]["host"])
		err = email.SendEmail(ctx, emailSubject, emailBody, alarmReceiver)
		if err == nil {
			// 发送成功，记录
			return
		}
	}
}

func sendSuccessNotify() {
	emailSubject := "应用服务器拨测成功定时通知"
	emailBody := fmt.Sprintf("拨测成功")
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	defer ctx.Done()
	account := map[string]map[string]string{
		"vpnheroes@outlook.com":  {"pw": "pingguoqm23", "host": "smtp-mail.outlook.com:587"},
		"heroesvpnn@outlook.com": {"pw": "pingguoqm23", "host": "smtp-mail.outlook.com:587"},
		"VPNHERO@outlook.com":    {"pw": "pingguoqm23", "host": "smtp-mail.outlook.com:587"},
	}
	accounts := []string{"vpnheroes@outlook.com", "heroesvpnn@outlook.com", "VPNHERO@outlook.com"}
	for _, userName := range accounts {
		email.SetSendAccount(userName, account[userName]["pw"], account[userName]["host"])
		err := email.SendEmail(ctx, emailSubject, emailBody, alarmReceiver)
		if err == nil {
			// 发送成功，记录
			return
		}
	}
}

// 定义与JSON响应对应的结构体
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

func dialAPIServer(ctx context.Context, url string) (err error) {
	// 创建一个请求
	glog.Infof(ctx, "url: %s", url)
	req, err := http.NewRequest("GET", url, nil)
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
		glog.Errorf(ctx, "Do failed, err: %s", err.Error())
		return
	}
	defer resp.Body.Close()

	// 检查HTTP响应状态码
	if resp.StatusCode != http.StatusOK {
		glog.Errorf(ctx, "HTTP请求失败，状态码: %d\n", resp.StatusCode)
		return
	}

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		glog.Errorf(ctx, "ReadAll failed, err: %s", err.Error())
		return
	}

	// 打印响应内容
	//fmt.Println(string(body))
	// 解析JSON响应
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		glog.Errorf(ctx, "json.Unmarshal failed, err: %s", err.Error())
		return
	}

	// 检查业务code和message
	if response.Code != 200 || response.Message != "成功" {
		err = fmt.Errorf("业务处理失败，code:%d, message:%s", response.Code, response.Message)
		glog.Errorf(ctx, err.Error())
		return
	}

	// 检查data.list是否有值
	if len(response.Data.List) == 0 {
		err = fmt.Errorf("data.list为空")
		glog.Errorf(ctx, err.Error())
		return
	}
	glog.Info(ctx, "%s 拨测成功!", url)

	// 打印解析后的数据
	//for _, item := range response.Data.List {
	//	fmt.Printf("DNS: %s, ID: %d, SiteType: %d\n", item.DNS, item.ID, item.SiteType)
	//}
	return
}
