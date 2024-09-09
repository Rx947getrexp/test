package api

import (
	"github.com/gin-gonic/gin"
	"go-speed/i18n"
	"go-speed/model/response"
	"net/http"
	"time"
)

// Test 测试api
func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"now": time.Now().Unix(),
	})
}

func APIDeprecated(ctx *gin.Context) {
	response.RespOk(ctx, i18n.RetMsgSuccess, nil)
	return
}
