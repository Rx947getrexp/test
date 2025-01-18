package internal

import (
	"context"
	"fmt"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
)

const (
	secondsInADay  = 24 * 60 * 60       // 获取一天的秒数
	maxExpiredTime = 30 * secondsInADay // 定义30天的过期时间
	giftTimestamp  = 15 * secondsInADay // 定义赠送15天的时长
)

func UserRecovery(c *gin.Context, userInfo *entity.TUser) {
	// 当前时间
	now := gtime.Now()
	nowUnix := now.Unix()
	// 获取用户是否已经过期maxExpiredTime天
	expiredPlusThirtyDays := userInfo.ExpiredTime + int64(maxExpiredTime)
	isExpiredMoreThanThirtyDays := nowUnix > expiredPlusThirtyDays
	// 活动结束目标时间
	targetTime := gtime.NewFromStr("2025-03-01 00:00:00")

	// 如果用户ExpiredTime过期已经超过maxExpiredTime天，开始赠送时长
	if isExpiredMoreThanThirtyDays && userInfo.Id > 0 && now.Before(targetTime) {
		cid := global.GetClientId(c)
		// 计算出新的过期时间giftTimestamp天，也就是赠送的时长到期时间
		newExpiredTime := nowUnix + int64(giftTimestamp)
		ctx := c
		err := dao.TUser.Ctx(c).Transaction(c, func(c context.Context, tx gdb.TX) error {
			// 加时长
			affected, err := dao.TUser.Ctx(c).Data(do.TUser{
				ExpiredTime: newExpiredTime, // 直接重置过期时间
				UpdatedAt:   now,
				Version:     userInfo.Version + 1,
				Kicked:      0,
			}).Where(do.TUser{
				Id:      userInfo.Id,
				Version: userInfo.Version,
			}).UpdateAndGetAffected()
			if err != nil {
				global.MyLogger(ctx).Err(err).Msgf(`user_recovery update t_user failed`)
				return err
			}
			if affected != 1 {
				err = fmt.Errorf("user_recovery update t_user affected(%d) != 1", affected)
				global.MyLogger(ctx).Err(err).Msgf("user_recovery update t_user failed")
				return err
			}

			global.MyLogger(ctx).Debug().Msgf("reset user ExpiredTime from(%d) to(%d)", userInfo.ExpiredTime, newExpiredTime)

			// 记录操作流水
			OpId, _ := service.GenSnowflake()
			lastInsertId, err := dao.TGift.Ctx(c).Data(do.TGift{
				UserId:    userInfo.Id,
				OpId:      strconv.FormatInt(OpId, 10),
				OpUid:     userInfo.Id,
				Title:     "老用户挽回赠送15天时长",
				GiftSec:   giftTimestamp,
				GType:     3,
				CreatedAt: now,
				UpdatedAt: now,
				Comment:   "",
			}).InsertAndGetId()
			if err != nil {
				global.MyLogger(ctx).Err(err).Msgf(`insert t_gift failed`)
				return err
			}
			global.MyLogger(ctx).Debug().Msgf("insert t_gift, lastInsertId: %d", lastInsertId)
			return nil
		})
		if err != nil {
			global.MyLogger(c).Err(err).Msgf("set expired_time failed")
			return
		}
		// 操作成功日志
		global.MyLogger(c).Info().Msgf("老用户挽回赠送15天时长成功, clientId: %s, userId: %d", cid, userInfo.Id)
	}
}
