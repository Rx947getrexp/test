package task

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-speed/api/types/api"
	"go-speed/constant"
	"go-speed/global"
	"go-speed/model/request"
	"go-speed/model/response"
	"go-speed/service"
	"go-speed/service/api/speed_api"
	"go-speed/util"
	"net/http/httptest"
	"strconv"
	"sync"
	"time"
)

const (
	leaderLockKeyTraffic = "hs-fly-CollectUserTraffic-leader-lock"

	intervalTraffic    = 10 * time.Minute
	lockTimeoutTraffic = intervalTraffic + 10*time.Minute
)

func CollectUserTraffic() {
	global.Recovery()
	global.Logger.Info().Msg("CollectUserTraffic start...")
	//ctx := context.Background()
	for {
		//isLeader, err := tryAcquireLock(ctx, leaderLockKeyTraffic, lockTimeoutTraffic)
		//if err != nil {
		//	global.Logger.Err(err).Msg("tryAcquireLock failed")
		//} else if isLeader {
		//	global.Logger.Info().Msg("I am the leader")
		//	// 在这里执行主进程的逻辑
		//
		//	releaseLock(ctx, leaderLockKeyTraffic)
		//} else {
		//	global.Logger.Info().Msg("I am a follower")
		//	// 在这里执行从进程的逻辑
		//}
		DoCollectUserTraffic()
		time.Sleep(intervalTraffic)
	}
}

func DoCollectUserTraffic() {
	// 查询v2ray数据节点
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	defer ctx.Done()
	resp, err := speed_api.DescribeNodeList(ctx)
	if err != nil {
		global.Logger.Err(err).Msg("get node ip list failed")
		return
	}
	wg := &sync.WaitGroup{}
	for i, _ := range resp.Items {
		wg.Add(1)
		go CollectNodeTraffic(ctx, wg, resp.Items[i])
	}
	wg.Wait()
}

func CollectNodeTraffic(ctx context.Context, wg *sync.WaitGroup, node api.NodeItem) {
	defer wg.Done()

	nodeIp := node.Ip
	date, _ := strconv.Atoi(time.Now().Format("20060102"))
	dataTime := time.Now().Format(constant.TimeFormat)

	// 从节点获取流量数据
	items, err := GetUserTrafficForOneNode(node.Ip)
	if err != nil {
		return
	}

	// 更新用户统计数据
	for _, item := range items {
		if item.UpLink == 0 && item.DownLink == 0 {
			//global.Logger.Info().Msgf("======== user(%s) traffic is zero, ip: %s", item.Email, nodeIp)
			continue
		}

		if e := service.CreateUserTrafficLog(ctx, item.Email, nodeIp, dataTime, item.UpLink, item.DownLink); e != nil {
			global.Logger.Err(e).Msgf("========xxxxxxxx user(%s) CreateUserTrafficLog failed (UpLink: %d, DownLink: %d) at collectTime(%s), ip: %s",
				item.Email, item.UpLink, item.DownLink, dataTime, nodeIp)
			// 插入流水失败，忽略错误，继续更新用户用量统计信息
		}

		userTraffic, err := service.GetUserTrafficByEmail(ctx, item.Email, nodeIp, date)
		if err != nil {
			continue
		}

		if userTraffic == nil {
			_ = service.CreateUserTraffic(ctx, item.Email, nodeIp, date, item.UpLink, item.DownLink)
		} else {
			_ = service.UpdateUserTraffic(ctx, userTraffic, item.UpLink, item.DownLink)
		}
	}
}

func GetUserTrafficForOneNode(ip string) (items []response.UserTrafficItem, err error) {
	//url := fmt.Sprintf("https://%s/site-api/node/get_user_traffic", server)
	//if strings.Contains(server, "http") {
	//
	//}
	url := fmt.Sprintf("http://%s:15003/node/get_user_traffic", ip)
	req := &request.GetUserTrafficRequest{
		All:   true,
		Reset: true,
		//Emails: []string{
		//	"a1@qq.com",
		//	"zzz@qq.com",
		//	"a101@qq.com",
		//	"303468504456@gmail.com",
		//},
	}
	res := new(response.Response)
	timestamp := fmt.Sprint(time.Now().Unix())
	headerParam := make(map[string]string)
	headerParam["timestamp"] = timestamp
	headerParam["accessToken"] = util.MD5(fmt.Sprint(timestamp, constant.AccessTokenSalt))
	err = util.HttpClientPostV2(url, headerParam, req, res)
	if err != nil {
		global.Logger.Err(err).Msgf("%s http failed: %s", url, err.Error())
		return
	}
	if res.Code != response.Success {
		err = fmt.Errorf("Code: %d, Msg: %s ", res.Code, res.Msg)
		global.Logger.Err(err).Msgf("%s return code is not success: Code: %d, Msg: %s", url, res.Code, res.Msg)
		return
	}
	resp := response.GetUserTrafficResponse{}
	err = json.Unmarshal(res.Data, &resp)
	if err != nil {
		global.Logger.Err(err).Msgf("%s Unmarshal failed, Data: %s, err: %s", url, string(res.Data), err.Error())
		return
	}
	return resp.Items, nil
}
