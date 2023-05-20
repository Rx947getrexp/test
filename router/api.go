package router

import (
	"github.com/gin-gonic/gin"
	"go-speed/api/api"
)

func ApiRoute(group *gin.RouterGroup) {
	group.POST("generate_dev_id", api.GenerateDevId)
	group.POST("send_email", api.SendEmail)
	group.POST("reg", api.Reg)
	group.POST("login", api.Login)
	group.POST("forget_passwd", api.ForgetPasswd)

	group.Use(api.JWTAuth())
	{
		group.POST("change_passwd", api.ChangePasswd)
	}
}
