package report

import (
	"github.com/gin-gonic/gin"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/response"
)

type GetClientIPResponse struct {
	ClientIP string `json:"client_ip"`
}

func GetClientIP(c *gin.Context) {
	global.MyLogger(c).Info().Msgf("ClientIP: %+v", c.ClientIP())
	response.RespOk(c, i18n.RetMsgSuccess, GetClientIPResponse{ClientIP: c.ClientIP()})
}
