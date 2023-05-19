package service

import (
	"fmt"
	"go-speed/constant"
	"go-speed/global"
	"go-speed/model/request"
	"go-speed/model/response"
	"go-speed/util"
	"net/http"
	"time"
)

func CheckUrlDelay(url string) int64 {
	// 测试目标URL
	targetURL := url + "/test"

	// 发起请求并计算延迟
	startTime := time.Now()
	resp, err := http.Get(targetURL)
	if err != nil {
		global.Logger.Err(err).Msg("访问出错!")
		return 999999
	}
	defer resp.Body.Close()

	duration := time.Since(startTime)
	delay := duration.Milliseconds()
	return delay
}

func Heartbeat() {
	req := &request.HeartbeatAdminRequest{
		NodeVersion: constant.NodeVersion,
	}
	url := global.Viper.GetString("system.admin_address") + "/node_report/heartbeat"
	res := new(response.Response)
	headerParam := make(map[string]string)
	timestamp := fmt.Sprint(time.Now().Unix())
	headerParam["timestamp"] = timestamp
	headerParam["accessToken"] = util.MD5(fmt.Sprint(timestamp, constant.AccessTokenSalt))
	err := util.HttpClientPostV2(url, headerParam, req, res)
	if err != nil {
		global.Logger.Err(err).Msg("发送心跳包失败...")
		return
	}
	if res.Code == 401 {
		global.Logger.Err(err).Msg("发送心跳包鉴权失败...")
		return
	}
}
