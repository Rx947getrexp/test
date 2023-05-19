package executor

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-speed/constant"
	"go-speed/global"
	"go-speed/model/request"
	"go-speed/model/response"
	"go-speed/util"
	"net/http"
)

func NodeAuth() gin.HandlerFunc {
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

func AddEmail(c *gin.Context) {
	param := new(request.NodeAddEmailRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	response.ResOk(c, "成功")
}

func RemoveEmail(c *gin.Context) {
	param := new(request.NodeRemoveEmailRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	response.ResOk(c, "成功")
}
