package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	statscmd "github.com/v2fly/v2ray-core/v5/app/stats/command"
	"go-speed/constant"
	"go-speed/global"
	"go-speed/model"
	"go-speed/model/entity"
	"go-speed/model/request"
	"go-speed/model/response"
	"go-speed/util"
	"google.golang.org/grpc"
	"time"
)

func InsertUserUuid(user *model.TUser, nodeId int64) (bool, error) {
	return true, nil
}

func InsertNodeUuid() {

}

func GetSysStats(ctx *gin.Context) (resp *statscmd.SysStatsResponse, err error) {
	conn, err := grpc.Dial(global.Config.System.V2rayApiAddress, grpc.WithInsecure())
	if err != nil {
		global.MyLogger(ctx).Err(err).Msg("grpc.Dial failed, err: " + err.Error())
		return nil, err
	}
	defer conn.Close()

	resp, err = statscmd.NewStatsServiceClient(conn).GetSysStats(ctx, &statscmd.SysStatsRequest{})
	if err != nil {
		global.MyLogger(ctx).Err(err).Msg("GetSysStats failed, err: " + err.Error())
		return nil, err
	}
	global.MyLogger(ctx).Info().Msgf("GetSysStats resp is: %s", resp.String())
	return
}

//func GetSysStatsByIp(ctx *gin.Context, nodeIP string) {
//	// 获取参数值
//	var (
//		//ctx  = gctx.New()
//		node = utils.GetNode(ctx, nodeIP)
//	)
//
//	url := fmt.Sprintf("http://%s:15003/node/get_v2ray_sys_stats", node.Ip)
//	fmt.Println(url)
//
//	timestamp := fmt.Sprint(time.Now().Unix())
//	headerParam := make(map[string]string)
//	res := new(response.Response)
//	headerParam["timestamp"] = timestamp
//	headerParam["accessToken"] = util.MD5(fmt.Sprint(timestamp, constant.AccessTokenSalt))
//	//fmt.Println(headerParam)
//	err := util.HttpClientPostV2(url, headerParam, &request.Request{}, res)
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//	fmt.Println("err：", err)
//	res.Print()
//}

func GetSysStatsByIp(ctx *gin.Context, nodeIP string) (resp *response.GetV2raySysStatsResponse, err error) {
	url := fmt.Sprintf("http://%s:15003/node/get_v2ray_sys_stats", nodeIP)
	global.Logger.Info().Msg(url)

	timestamp := fmt.Sprint(time.Now().Unix())
	headerParam := make(map[string]string)
	res := new(response.Response)
	headerParam["timestamp"] = timestamp
	headerParam["accessToken"] = util.MD5(fmt.Sprint(timestamp, constant.AccessTokenSalt))
	err = util.HttpClientPostV2(url, headerParam, &request.Request{}, res)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msg("HttpClientPostV2 failed")
		return
	}
	if res.Code != response.Success {
		err = fmt.Errorf("Code: %d, Msg: %s ", res.Code, res.Msg)
		global.MyLogger(ctx).Err(err).Msgf("%s return code is not success: Code: %d, Msg: %s", url, res.Code, res.Msg)
		return
	}
	res.Print()
	global.MyLogger(ctx).Info().Msgf("get_v2ray_sys_stats >>>>> ip: %s, resp: %s", nodeIP, res.Dump())
	resp = &response.GetV2raySysStatsResponse{}
	err = json.Unmarshal(res.Data, &resp)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("%s Unmarshal failed, Data: %s, err: %s", url, string(res.Data), err.Error())
		return
	}
	global.MyLogger(ctx).Info().Msgf("resp: %+v", *resp)
	return resp, nil
}

func GetMinLoadNode(ctx *gin.Context, nodes []entity.TNode) (nodeId int64, err error) {
	var numGoroutine uint32 = 200000000
	for _, node := range nodes {
		res, err := GetSysStatsByIp(ctx, node.Ip)
		if err != nil {
			global.MyLogger(ctx).Err(err).Msgf(">>>>>>>>>>>>>>>>>>>> numGoroutine: %d, %s skip, CountryEn: %s", numGoroutine, node.Ip, node.CountryEn)
			continue
		}
		if res.NumGoroutine < numGoroutine {
			numGoroutine = res.NumGoroutine
			nodeId = node.Id
			global.MyLogger(ctx).Info().Msgf(">>>>>>>>>> numGoroutine: %d, id: %d, ip: %s, CountryEn: %s", numGoroutine, node.Id, node.Ip, node.CountryEn)
		}
	}
	return
}
