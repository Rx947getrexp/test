package api

import (
	"github.com/gin-gonic/gin"
	"go-speed/service"
	"net/http"
)

// JWTAuth 验证token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization-Token")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"code":    301,
				"message": "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}

		// parseToken 解析token包含的信息
		claims, err := service.ParseTokenByUser(token, service.CommonUserType)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    301,
				"message": "授权已过期",
			})
			c.Abort()
			return
		}

		err = service.AddLog(c, claims.UserId)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    100,
				"message": "网络错误",
			})
			c.Abort()
			return
		}
		//isBool := RedisInsert(strconv.FormatInt(claims.UserId, 10))
		//if !isBool {
		//	c.JSON(http.StatusOK, gin.H{
		//		"code": 301,
		//		"msg":  "授权已过期",
		//	})
		//	c.Abort()
		//	return
		//}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)
		//uu := c.MustGet("claims").(*service.CustomClaims)
		//fmt.Println("claims...", uu.UserId)
	}
}
