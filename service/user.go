package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/os/gtime"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/model/do"
)

func UpdateLoginInfo(c *gin.Context, uid int64) {
	ip := c.ClientIP()
	global.MyLogger(c).Info().Msgf("clientIp is %s, uid: %d", ip, uid)
	if ip != "" {
		_, err := dao.TUser.Ctx(c).Data(do.TUser{
			LastLoginIp: ip, UpdatedAt: gtime.Now(),
		}).Where(do.TUser{
			Id: uid,
		}).Update()
		if err != nil {
			global.MyLogger(c).Err(err).Msgf("update LastLoginIp failed, uid: %d", uid)
			return
		}
	} else {
		global.MyLogger(c).Info().Msgf("clientIp is empty, uid: %d", uid)
	}
}
