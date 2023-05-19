package router

import (
	"github.com/gin-gonic/gin"
	"go-speed/api/executor"
)

func ExecutorRoute(group *gin.RouterGroup) {
	nodeGroup := group.Group("node")
	nodeGroup.Use(executor.NodeAuth())
	{
		nodeGroup.POST("add_email", executor.AddEmail)
		nodeGroup.POST("remove_email", executor.RemoveEmail)
	}
}
