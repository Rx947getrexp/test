package admin

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-speed/constant"
	"go-speed/global"
	"go-speed/model/request"
	"go-speed/model/response"
	"go-speed/util"
	"net/http"
	"time"
)

func NodeReportAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader("accessToken")
		timestamp := c.GetHeader("timestamp")
		md5Str := util.MD5(fmt.Sprint(timestamp, constant.AccessTokenSalt))
		if accessToken != md5Str {
			c.JSON(http.StatusOK, gin.H{
				"code":    401,
				"message": "token鉴权失败，无权限访问",
			})
			c.Abort()
			return
		}
	}
}

func Heartbeat(c *gin.Context) {
	param := new(request.HeartbeatAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	reqIP := c.ClientIP()
	fmt.Println("client ip :", reqIP, param.NodeVersion)
	lastTime := time.Now().Unix()
	version := param.NodeVersion
	nodeKey := "node_" + reqIP
	global.Redis.HSet(context.Background(), nodeKey, "lastTime", lastTime).Err()
	global.Redis.HSet(context.Background(), nodeKey, "version", version).Err()
	//a, _ := global.Redis.HGet(context.Background(), nodeKey, "lastTime").Int64()
	//b, _ := global.Redis.HGet(context.Background(), nodeKey, "version").Result()
	//fmt.Println(a, b)
	//hashMap, _ := global.Redis.HGetAll(context.Background(), nodeKey).Result()
	//for k,v := range hashMap {
	//	fmt.Println(k, v)
	//}
	response.ResOk(c, "成功")
}

func Data(c *gin.Context) {
	param := new(request.ReportDataAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	reqIP := c.ClientIP()
	fmt.Println("client ip :", reqIP)
	response.ResOk(c, "成功")
}
