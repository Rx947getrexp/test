package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mssola/user_agent"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/response"
	"go-speed/service"
)

func ValidateClaims(ctx *gin.Context) (userEntity *entity.TUser, err error) {
	claims := ctx.MustGet("claims").(*service.CustomClaims)
	//claims := service.CustomClaims{UserId: 219122692}
	return CheckUserByUserId(ctx, uint64(claims.UserId))
}

func CheckUserByUserId(ctx *gin.Context, userId uint64) (userEntity *entity.TUser, err error) {
	err = dao.TUser.Ctx(ctx).Where(do.TUser{Id: userId}).Scan(&userEntity)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("userId %d 查询db失败", userId)
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}
	if userEntity == nil {
		err = fmt.Errorf("userId无效: %d", userId)
		global.MyLogger(ctx).Warn().Msgf("userId无效: %d", userId)
		response.RespFail(ctx, i18n.RetMsgAccountNotExist, nil)
		return
	}
	if userEntity.Status != 0 {
		err = fmt.Errorf("user(%d) 用户状态(%d)无效", userId, userEntity.Status)
		global.MyLogger(ctx).Err(err).Msgf("user(%d) 用户状态(%d)无效", userId, userEntity.Status)
		response.RespFail(ctx, i18n.RetMsgUserIdInvalid, nil)
		return
	}
	return userEntity, nil
}

func GetUserByEmail(ctx *gin.Context, email string) (userEntity *entity.TUser, err error) {
	err = dao.TUser.Ctx(ctx).Where(do.TUser{Email: email}).Scan(&userEntity)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("email %s 查询db失败", email)
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}
	if userEntity == nil {
		err = fmt.Errorf("email %s 无效", email)
		global.MyLogger(ctx).Err(err).Msgf("email not exist")
		response.RespFail(ctx, i18n.RetMsgAccountNotExist, nil)
		return
	}
	return userEntity, nil
}

func CheckDevId(ctx *gin.Context, devId string) (devEntity *entity.TDev, err error) {
	err = dao.TDev.Ctx(ctx).Where(do.TDev{
		Id: devId,
	}).Scan(&devEntity)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("devId %d 查询db失败", devId)
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}
	if devEntity == nil {
		err = fmt.Errorf("%s", i18n.RetMsgDevIdNotExitsErr)
		global.MyLogger(ctx).Err(err).Msgf("devId %d 无效", devId)
		response.RespFail(ctx, i18n.RetMsgDevIdNotExitsErr, nil)
		return
	}
	return
}

func SaveDeviceID(ctx *gin.Context, uid int64) {
	deviceID := global.GetClientId(ctx)
	if deviceID == "" {
		global.MyLogger(ctx).Warn().Msgf("deviceID is empty, userId: %d", uid)
		return
	}
	userAgent := global.GetUserAgent(ctx)
	ua := user_agent.New(userAgent)
	os := ua.OS()
	if os == "" {
		os = userAgent
	}

	lastInsertId, err := dao.TUserDevice.Ctx(ctx).Data(do.TUserDevice{
		UserId:    uid,
		ClientId:  deviceID,
		Os:        os,
		CreatedAt: gtime.Now(),
	}).InsertIgnore()
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("add TUserDevice failed")
		return
	}
	global.MyLogger(ctx).Debug().Msgf("add TUserDevice success, lastInsertId: %d", lastInsertId)
	return
}
