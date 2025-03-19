package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/errors/gerror"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/response"
	"strconv"
	"strings"
	"time"
)

func CheckClientIdNumLimitsWithErrRsp(ctx *gin.Context, userInfo *entity.TUser) (err error) {
	clientId := global.GetClientId(ctx)
	if clientId == "" {
		global.MyLogger(ctx).Warn().Msgf("%s Client-Id is empty. [%s]", i18n.ErrLabelHeader, userInfo.Email)
		return nil
	}

	timeWindow := time.Now().Add(-1 * time.Hour * 30 * 24).Format("2006-01-02 15:04:05")
	var items []entity.TUserDevice
	err = dao.TUserDevice.Ctx(ctx).
		Where(do.TUserDevice{UserId: userInfo.Id, Kicked: 0}).
		WhereGTE(dao.TUserDevice.Columns().UpdatedAt, timeWindow).
		Scan(&items)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("%s query TUserDevice failed. [%s]", i18n.ErrLabelDB, userInfo.Email)
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return err
	}

	clientIdNums := 0
	for _, item := range items {
		if item.ClientId == clientId {
			global.MyLogger(ctx).Info().Msgf("[TUserDeviceAlreadyRecord] %s, level: %d, ClientId: %s, UpdatedAt: %s",
				userInfo.Email, userInfo.Level, item.ClientId, item.UpdatedAt.Format("2006-01-02 15:04:05"))
			return nil
		} else {
			clientIdNums++
		}
	}
	// TODO：先简单处理，目前还没有产品形态定义
	limits := 3
	switch userInfo.Level {
	case 0:
		limits = 3
	case 1:
		limits = 3
	case 2:
		limits = 6
	case 3:
		limits = 10
	default:
		limits = 5
	}

	global.MyLogger(ctx).Info().Msgf("%s, level: %d, clientIdNums: %d, limits: %d",
		userInfo.Email, userInfo.Level, clientIdNums, limits)

	if clientIdNums >= limits {
		err = gerror.Newf("[CheckClientIdNumLimits] 登录出错，设备数量超过限制. account: %s, level: %d, clientIdNums: %d, limits: %d",
			userInfo.Email, userInfo.Level, clientIdNums, limits)
		global.MyLogger(ctx).Warn().Msgf(err.Error())
		response.RespFail(ctx, i18n.RetMsgReachedDevicesLimit, nil)
		return err
	}
	return nil
}

func GetInviteUserWithErrRsp(ctx *gin.Context, inviteCode string) (
	inviteUserId uint64, user *entity.TUser, userTeam *entity.TUserTeam, err error) {
	inviteCode = strings.TrimSpace(inviteCode)
	if inviteCode == "" {
		return
	}

	inviteUserId, err = strconv.ParseUint(inviteCode, 10, 64)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("%s InviteCode: (%s)", i18n.ErrLabelParams, inviteCode)
		response.RespFail(ctx, i18n.RetMsgReferrerIDIncorrect, nil)
		return
	}

	err = dao.TUser.Ctx(ctx).Where(do.TUser{Id: inviteUserId}).Scan(&user)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("%s query InviteUser failed, inviteCode: %s", i18n.ErrLabelDB, inviteCode)
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}
	if user == nil {
		err = dao.TUserCancelled.Ctx(ctx).Where(do.TUserCancelled{Id: inviteUserId}).Scan(&user)
		if err != nil {
			global.MyLogger(ctx).Err(err).Msgf(
				"%s query InviteUser from TUserCancelled failed, inviteCode: %s", i18n.ErrLabelDB, inviteCode)
			response.RespFail(ctx, i18n.RetMsgDBErr, nil)
			return
		}
	}
	err = dao.TUserTeam.Ctx(ctx).Where(do.TUserTeam{UserId: inviteUserId}).Scan(&userTeam)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("%s query TUserTeam failed, inviteCode: %d", i18n.ErrLabelDB, inviteUserId)
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}
	return
}
