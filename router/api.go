package router

import (
	"github.com/gin-gonic/gin"
)

func ApiRoute(group *gin.RouterGroup) {
	//group.POST("upload", api.Upload)

	//v1 := group.Group("v1")
	//v1.Use() //中间件
	//v1.POST("createCollectOrder", api.CreateCollectOrder)   //创建代收订单
	//v1.POST("queryCollectOrder", api.QueryCollectOrder)     //查询代收订单
	//v1.POST("cancelCollectOrder", api.CancelCollectOrder)   //取消代收订单
	//v1.POST("confirmCollectOrder", api.ConfirmCollectOrder) //确认代收订单（含电子账单)
	//v1.POST("createPaymentOrder", api.CreatePaymentOrder)   //创建代付订单
	//v1.POST("queryPaymentOrder", api.QueryPaymentOrder)     //查询代付订单
	//v1.POST("cancelPaymentOrder", api.CancelPaymentOrder)   //取消代付订单
}
