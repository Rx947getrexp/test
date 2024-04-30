package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/response"
)

func CheckUserByUserId(ctx *gin.Context, userId uint64) (userEntity *entity.TUser, err error) {
	err = dao.TUser.Ctx(ctx).Where(do.TUser{Id: userId}).Scan(&userEntity)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("userId %d 查询db失败", userId)
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}
	if userEntity == nil {
		err = fmt.Errorf("userId %d 无效", userId)
		global.MyLogger(ctx).Err(err).Msgf("userId %d 无效", userId)
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
