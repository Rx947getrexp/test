package main

//import (
//	"fmt"
//	"go-speed/constant"
//	"go-speed/global"
//	"go-speed/model/response"
//	"go-speed/util"
//	"strings"
//	"time"
//)
//
//func main() {
//	url := fmt.Sprintf("https://110.42.42.229/site-api/node/add_sub")
//	if strings.Contains(item.Server, "http") {
//		url = fmt.Sprintf("http://81.70.92.211:15003/node/add_sub", item.Server)
//	}
//	global.MyLogger(ctx).Info().Msgf(">>>>>>>>> url: %s", url)
//	timestamp := fmt.Sprint(time.Now().Unix())
//	headerParam := make(map[string]string)
//	res := new(response.Response)
//	headerParam["timestamp"] = timestamp
//	headerParam["accessToken"] = util.MD5(fmt.Sprint(timestamp, constant.AccessTokenSalt))
//	err = util.HttpClientPostV2(url, headerParam, nodeAddSubRequest, res)
//	if res != nil {
//		global.MyLogger(ctx).Info().Msgf(">>>>>>>>> nodeAddSubRequest: %+v, res: %+v", *nodeAddSubRequest, *res)
//	} else {
//		global.MyLogger(ctx).Info().Msgf(">>>>>>>>> nodeAddSubRequest: %+v, res: is nil", *nodeAddSubRequest)
//	}
//	if err != nil {
//		global.MyLogger(ctx).Err(err).Msgf("email: %s, add_sub 发送失败", userEntity.Email)
//		continue
//	}
//}
