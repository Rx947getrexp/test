package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mssola/user_agent"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/entity"
)

func UpdateUserDeviceByClientId(ctx *gin.Context, uid int64, email string) (err error) {
	clientId := global.GetClientId(ctx)
	if clientId == "" {
		global.MyLogger(ctx).Warn().Msgf("%s Client-Id is empty. [%s]", i18n.ErrLabelHeader, email)
		return nil
	}
	var row *entity.TUserDevice
	err = dao.TUserDevice.Ctx(ctx).
		Where(do.TUserDevice{UserId: uid, ClientId: clientId}).
		Scan(&row)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("%s query TUserDevice failed. [%s]", i18n.ErrLabelDB, email)
		return err
	}

	if row == nil {
		userAgent := global.GetUserAgent(ctx)
		ua := user_agent.New(userAgent)
		os := ua.OS()
		if os == "" {
			os = userAgent
		}
		var lastId int64
		lastId, err = dao.TUserDevice.Ctx(ctx).Data(do.TUserDevice{
			UserId:    uid,
			ClientId:  clientId,
			Os:        os,
			CreatedAt: gtime.Now(),
			UpdatedAt: gtime.Now(),
		}).InsertAndGetId()
		if err != nil {
			global.MyLogger(ctx).Err(err).Msgf("%s add TUserDevice failed. [%s]", i18n.ErrLabelDB, email)
			return err
		}
		global.MyLogger(ctx).Debug().Msgf("add TUserDevice lastId: %d", lastId)
	} else {
		var affected int64
		affected, err = dao.TUserDevice.Ctx(ctx).
			Where(do.TUserDevice{UserId: uid, ClientId: clientId}).
			Data(do.TUserDevice{Kicked: 0, UpdatedAt: gtime.Now()}).UpdateAndGetAffected()
		if err != nil {
			global.MyLogger(ctx).Err(err).Msgf("%s update TUserDevice failed. [%s]", i18n.ErrLabelDB, email)
			return err
		}
		global.MyLogger(ctx).Debug().Msgf("add TUserDevice affected: %d", affected)
	}
	return nil
}
