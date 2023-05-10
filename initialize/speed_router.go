package initialize

import (
	"github.com/gin-gonic/gin"
	"go-speed/api"
	"go-speed/global"
	"go-speed/router"
	"net/http"
)

func AdminRouters() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	routers := gin.Default()
	routers.Use(Cors()) //解决跨域
	routers.Use(api.PrintRequest)
	routers.Use(api.FilteredSQLInject)
	routers.Use(api.RateMiddleware(api.NewLimiterV2()))
	routers.GET("test", api.Test)
	publicGroup := routers.Group("")
	router.AdminRoute(publicGroup)
	return routers
}

func ApiRouters() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	routers := gin.Default()
	routers.Use(Cors()) //解决跨域
	//routers.Use(api.PrintRequest)
	//routers.Use(api.FilteredSQLInject)
	routers.Use(api.RateMiddleware(api.NewLimiterV2()))
	routers.GET("test", api.Test)
	publicGroup := routers.Group("")
	router.ApiRoute(publicGroup)
	return routers
}

func ExecutorRouters() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	routers := gin.Default()
	routers.Use(Cors()) //解决跨域
	//routers.Use(api.PrintRequest)
	//routers.Use(api.FilteredSQLInject)
	//routers.Use(api.RateMiddleware(api.NewLimiterV2()))
	routers.GET("test", api.Test)
	//publicGroup := routers.Group("")
	//router.ApiRoute(publicGroup)
	return routers
}

func JobRouters() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	routers := gin.Default()
	routers.Use(Cors()) //解决跨域
	//routers.Use(api.PrintRequest)
	//routers.Use(api.FilteredSQLInject)
	//routers.Use(api.RateMiddleware(api.NewLimiterV2()))
	routers.GET("test", api.Test)
	//publicGroup := routers.Group("")
	//router.ApiRoute(publicGroup)
	return routers
}

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,Authorization,Token,Authorization-Token,xx-device-type")
		context.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS")
		context.Header("Access-Control-Expose-Headers", "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Content-Type")
		context.Header("Access-Control-Allow-Credentials", "True")
		//放行索引options
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		//处理请求
		context.Next()
	}
}
