package vip

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"go-speed/api/api/common"
	"go-speed/constant"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/response"
	"go-speed/util"
	"time"
)

type EditMemberExpiredTimeReq struct {
	Id          uint64 `form:"id" binding:"required" json:"id"`  //用户ID
	ExpiredTime int64  `form:"expired_time" json:"expired_time"` //会员过期时间，时间戳秒，最大支持当前时间+5年
}

func EditMemberExpiredTime(c *gin.Context) {
	var (
		err          error
		req          = new(EditMemberExpiredTimeReq)
		userEntity   *entity.TUser
		affected     int64
		lastInsertId int64
	)
	if err = c.ShouldBind(req); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	userEntity, err = common.CheckUserByUserId(c, req.Id)
	if err != nil {
		return
	}
	global.Logger.Err(err).Msgf("uname: %s", userEntity.Uname)
	if !isValidTimestamp(req.ExpiredTime) {
		global.Logger.Err(nil).Msgf("ExpiredTime: (%d) 参数无效！不能超过20年时间长度。", req.ExpiredTime)
		response.ResFail(c, "用户过期参数无效！不能超过20年时间长度。")
		return
	}

	ctx := c
	err = dao.TUser.Ctx(c).Transaction(c, func(c context.Context, tx gdb.TX) error {
		// 加时长
		affected, err = dao.TUser.Ctx(c).Data(do.TUser{
			ExpiredTime: req.ExpiredTime, // 管理员直接重置过期时间
			UpdatedAt:   gtime.Now(),
			Version:     userEntity.Version + 1,
		}).Where(do.TUser{
			Id:      userEntity.Id,
			Version: userEntity.Version,
		}).UpdateAndGetAffected()
		if err != nil {
			global.MyLogger(ctx).Err(err).Msgf(`update t_user failed`)
			return err
		}
		if affected != 1 {
			err = fmt.Errorf("update t_user affected(%d) != 1", affected)
			global.MyLogger(ctx).Err(err).Msgf("update t_user failed")
			return err
		}

		global.MyLogger(ctx).Debug().Msgf("reset user ExpiredTime from(%d) to(%d)", userEntity.ExpiredTime, req.ExpiredTime)

		// 记录操作流水
		lastInsertId, err = dao.TUserVipAttrRecord.Ctx(c).Data(do.TUserVipAttrRecord{
			Email:           userEntity.Email,
			Source:          constant.UserVipAttrOpSourceAdminSet,
			OrderNo:         gtime.Datetime(),
			ExpiredTimeFrom: userEntity.ExpiredTime,
			ExpiredTimeTo:   req.ExpiredTime,
			Desc:            fmt.Sprintf("ExpiredTime set to(%s)", util.TimeFormat(req.ExpiredTime)),
			CreatedAt:       gtime.Now(),
		}).InsertAndGetId()
		if err != nil {
			global.MyLogger(ctx).Err(err).Msgf(`insert TUserVipAttrRecords failed`)
			return err
		}
		global.MyLogger(ctx).Debug().Msgf("insert TUserVipAttrRecords, lastInsertId: %d", lastInsertId)
		return nil
	})
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("set expired_time failed")
		response.ResFail(c, err.Error())
		return
	}
	response.ResOk(c, "成功")
}

func isValidTimestamp(expiredTime int64) bool {
	// 将 ExpiredTime 转换为 time.Time 类型
	expiredTimeInTime := time.Unix(expiredTime, 0)

	// 获取当前时间
	now := time.Now()

	// 获取1年前的时间
	oneYearAgo := now.AddDate(-1, 0, 0)

	// 获取5年后的时间
	fiveYearsLater := now.AddDate(20, 0, 0)

	// 判断 ExpiredTime 是否在有效范围内
	if expiredTimeInTime.After(oneYearAgo) && expiredTimeInTime.Before(fiveYearsLater) {
		return true
	}
	return false
}
