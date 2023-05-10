package router

import (
	"github.com/gin-gonic/gin"
	"go-speed/api/upload"
)

func InitUploadRoute(group *gin.RouterGroup) {
	group.Static("public", "./public")
	group.Static("uploads", "./uploads")
	group.POST("upload", upload.PathUpload)
}
