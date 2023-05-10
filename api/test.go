package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Test 测试api
func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"now": time.Now().Unix(),
	})
}
