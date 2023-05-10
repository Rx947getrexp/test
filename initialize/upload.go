package initialize

import (
	"github.com/gin-gonic/gin"
	"go-speed/api"
	"go-speed/global"
	"go-speed/router"
)

func RouterUpload() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	routers := gin.Default()
	routers.Use(api.PrintRequest)
	routers.GET("test", api.Test)
	publicGroup := routers.Group("")
	router.InitUploadRoute(publicGroup)
	return routers
}
