package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go-speed/util"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/os/glog"

	"go-speed/global"
	"go-speed/initialize"
	"go-speed/service"
	"go-speed/service/email"
)

type DialType string

const (
	DialTypeAPI  DialType = "API"
	DialTypeNode DialType = "Node"
)

var errCounterAPIServerDNS = make(map[string]uint64)
var errCounterNodeServerDNS = make(map[string]uint64)
var alarmReceiver = []string{
	"pmm73219@gmail.com",
	"hs.alarm@outlook.com",
}

func getConfig(c string) (out []string) {
	for _, v := range strings.Split(c, ",") {
		v = strings.TrimSpace(v)
		if v != "" {
			out = append(out, v)
		}
	}
	return out
}

func main() {
	glog.SetLevel(glog.LEVEL_ALL)
	initialize.InitComponentsV2()

	var (
		apiServerDNSList  = util.MustReadFile("./api_server_dns.list")
		nodeServerDNSList = util.MustReadFile("./node_server_dns.list")
	)

	if len(apiServerDNSList) == 0 || len(nodeServerDNSList) == 0 {
		global.Logger.Fatal().Msgf("apiServerDNSList or nodeServerDNSList is empty")
	}

	for _, dns := range apiServerDNSList {
		errCounterAPIServerDNS[dns] = 0
	}

	for _, dns := range nodeServerDNSList {
		errCounterNodeServerDNS[dns] = 0
	}

	for {
		doDialApiServer(apiServerDNSList)
		doDialNodeServer(nodeServerDNSList)
		time.Sleep(time.Minute * 1)
	}
}

var (
	times        uint64
	successTimes uint64 = 1
)

func doDialApiServer(dnsList []string) {
	times++
	global.Logger.Info().Msgf("--------------------------------------------------------------------------------")
	global.Logger.Info().Msgf("--------------- DialApiServer [times: %d] start -----------------------------", times)

	var (
		ctx        = context.Background()
		err        error
		failedFlag = false
	)
	for _, dns := range dnsList {
		err = dialAPIServer(ctx, fmt.Sprintf("https://%s/app-api/dns_list", dns))
		if err != nil {
			failedFlag = true
			errCounterAPIServerDNS[dns]++
		} else {
			errCounterAPIServerDNS[dns] = 0
		}
		global.Logger.Info().Msgf("dns: %s, errCounterAPIServerDNS: %d", dns, errCounterAPIServerDNS[dns])

		if errCounterAPIServerDNS[dns] > 5 {
			if e := sendAlarm(DialTypeAPI, dns, err); e == nil {
				errCounterAPIServerDNS[dns] = 3
			}
		}
	}
	if !failedFlag {
		successTimes++
	}
	if successTimes%(120) == 0 {
		sendSuccessNotify(DialTypeAPI)
	}
	global.Logger.Info().Msgf("--------------- DialApiServer [times: %d, successTimes: %d] end --------------", times, successTimes)
	global.Logger.Info().Msgf("################################################################################")
}

func sendAlarm(dType DialType, dns string, err error) error {
	var emailSubject string
	if dType == DialTypeAPI {
		emailSubject = "【拨测失败告警】应用服务器-异常！！！"
	} else {
		emailSubject = "【拨测失败告警】节点服务器-异常！！！"
	}
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("拨测对象：%s (%s)<br/>", dns, time.Now().Format(time.DateTime)))
	builder.WriteString(fmt.Sprintf("异常原因：%s<br/>", err.Error()))
	builder.WriteString("<br/>")
	builder.WriteString("赶紧处理!!!<br/>\n")
	return sendEmail(emailSubject, builder.String())
}

func sendEmail(emailSubject, emailBody string) (err error) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	defer ctx.Done()
	var (
		fromAccount = "heronetwork@herovpn.live"
		fromPasswd  = "pingguoqm23"
		fromHost    = "smtpout.secureserver.net:465"
	)

	if emails, e := util.ReadFile("./alarm_receiver_email.list"); e == nil {
		for _, em := range emails {
			if !util.IsInArrayIgnoreCase(em, alarmReceiver) {
				alarmReceiver = append(alarmReceiver, em)
			}
		}
	}

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

func sendSuccessNotify(dType DialType) {
	var emailSubject string
	if dType == DialTypeAPI {
		emailSubject = "<拨测成功> 应用服务器-拨测定时通知"
	} else {
		emailSubject = "<拨测成功> 节点服务器-拨测定时通知"
	}
	emailBody := fmt.Sprintf("应用服务器和nginx服务器拨测成功 (%s)", time.Now().Format(time.DateTime))
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

// /////////////////////////////////// dial node server
var (
	dialNodeTimes        uint64
	dialNodeSuccessTimes uint64 = 1
)

func doDialNodeServer(dnsList []string) {
	dialNodeTimes++
	global.Logger.Info().Msgf("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	global.Logger.Info().Msgf("--------------- doDialNodeServer [dialNodeTimes: %d] start -----------------------------", dialNodeTimes)

	var (
		ctx, _     = gin.CreateTestContext(httptest.NewRecorder())
		err        error
		failedFlag = false
	)
	defer ctx.Done()

	for _, dns := range dnsList {
		_, err = service.GetSysStatsByIp(ctx, dns)
		if err != nil {
			failedFlag = true
			errCounterNodeServerDNS[dns]++
		} else {
			errCounterNodeServerDNS[dns] = 0
		}
		global.Logger.Info().Msgf("@@@@@@@@@@@@@@@@@@@@@@@ dns: %s, errCounterNodeServerDNS: %d", dns, errCounterNodeServerDNS[dns])

		if errCounterNodeServerDNS[dns] > 5 {
			if e := sendAlarm(DialTypeNode, dns, err); e == nil {
				errCounterNodeServerDNS[dns] = 3
			}
		}
	}
	for k, v := range errCounterNodeServerDNS {
		if v > 0 {
			global.Logger.Err(nil).Msgf(">>>>>>>>>>>>>> node(%s), dial failed times(%d)", k, v)
		}
	}

	if !failedFlag {
		dialNodeSuccessTimes++
	}
	if dialNodeSuccessTimes%(120) == 0 {
		sendSuccessNotify(DialTypeNode)
	}
	global.Logger.Info().Msgf("--------------- doDialNodeServer [dialNodeTimes: %d, dialNodeSuccessTimes: %d] end --------------", dialNodeTimes, dialNodeSuccessTimes)
	global.Logger.Info().Msgf("################################################################################")
}

// ps -ef | grep go-monitor | grep -v 'grep' | awk '{print $2}' | xargs kill && cd /wwwroot/go/go-monitor/ && cp -rf backup/go-monitor ./ && ./restart.sh
