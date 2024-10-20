package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mssola/user_agent"
	"go-speed/constant"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/model/do"
	"go-speed/model/entity"
)

func GetDevIdByClientId(c *gin.Context) (devId int64, err error) {
	clientId := global.GetClientId(c)
	if clientId == "" {
		global.MyLogger(c).Warn().Msg("clientId is empty")
		return devId, nil
	}

	devEntity, err := IsClientIdExistInTDev(c, clientId)
	if err != nil {
		return
	}

	if devEntity == nil {
		// 不存在就插入
		devId, _ = GenSnowflake()
		userAgent := global.GetUserAgent(c)
		ua := user_agent.New(userAgent)
		os := ua.OS()
		if os == "" {
			os = userAgent
		}
		_, err = dao.TDev.Ctx(c).Data(do.TDev{
			Id:        devId,
			Os:        os,
			ClientId:  clientId,
			Network:   constant.NetworkAutoMode,
			CreatedAt: gtime.Now(),
			UpdatedAt: gtime.Now(),
			IsSend:    2,
			Comment:   "",
		}).Insert()
		if err != nil {
			return devId, gerror.Wrap(err, "insert TDev failed")
		}
	} else {
		devId = devEntity.Id
	}
	return
}

func IsClientIdExistInTDev(ctx *gin.Context, clientId string) (devEntity *entity.TDev, err error) {
	err = dao.TDev.Ctx(ctx).Where(do.TDev{ClientId: clientId}).Scan(&devEntity)
	if err != nil {
		return nil, gerror.Wrap(err, "get TDev by clientId failed")
	}
	return
}
