package speed_api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/errors/gerror"
	"go-speed/api/types/api"
	"go-speed/global"
	"go-speed/model/response"
	"go-speed/util"
)

const (
	Token = "J7RtY3DvV2pK0fM5rW4aU1cL8yB9eQ6sI8gH2kZ5xT7uF1oP6vN8jA4lR9mG3bE0wH7nY6tS5zC8iQ1fX9rV6hO5lJ4dU3pV8aB2e(*&12"

	HttpCodeSuccess     = 200
	APIDescribeUserInfo = "internal/describe_user_info"
	APIDescribeNodeList = "internal/describe_node_list"
)

func genHeaders() map[string]string {
	headerParam := make(map[string]string)
	headerParam["Authorization-Token"] = Token
	return headerParam
}

func genUrl(apiName string) string {
	return fmt.Sprintf("http://%s/%s", "31.128.41.86:13002", apiName)
}

func DescribeUserInfo(ctx *gin.Context, req *api.DescribeUserInfoReq) (resp *api.DescribeUserInfoRes, err error) {
	var (
		apiName = APIDescribeUserInfo
		url     = genUrl(apiName)
		res     = new(response.Response)
	)
	global.MyLogger(ctx).Info().Msgf("call %s, request: %+v", url, *req)
	err = util.HttpClientPostV2(url, genHeaders(), req, res)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("call api failed")
		return
	}
	global.MyLogger(ctx).Info().Msgf("code: %d, msg: %s, data: %s", res.Code, res.Msg, string(res.Data))
	if res.Code != HttpCodeSuccess {
		return nil, gerror.New(res.Msg)
	}
	err = json.Unmarshal(res.Data, &resp)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("json.Unmarshal")
		return
	}
	return resp, nil
}

func DescribeNodeList(ctx *gin.Context) (resp *api.DescribeNodeListRes, err error) {
	var (
		apiName = APIDescribeNodeList
		url     = genUrl(apiName)
	)
	global.MyLogger(ctx).Info().Msgf("call %s", url)
	res := new(response.Response)
	err = util.HttpClientPostV2(url, genHeaders(), api.DescribeNodeListReq{}, res)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("call api failed")
		return
	}
	global.MyLogger(ctx).Info().Msgf("code: %d, msg: %s, data: %s", res.Code, res.Msg, string(res.Data))
	if res.Code != HttpCodeSuccess {
		return nil, gerror.New(res.Msg)
	}
	err = json.Unmarshal(res.Data, &resp)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("json.Unmarshal")
		return
	}
	return resp, nil
}
