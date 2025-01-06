package ad

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"go-speed/api/api/common"
	"go-speed/constant"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/response"
	"go-speed/service"
	"runtime/debug"
	"time"
)

type ADCompletionNotifyReq struct {
	Name string `form:"name" binding:"required" json:"name" dc:"广告名称"`
}

type ADCompletionNotifyRes struct{}

func ADCompletionNotify(ctx *gin.Context) {
	var (
		err        error
		req        = new(ADCompletionNotifyReq)
		adInfo     *entity.TAd
		user       *entity.TUser
		lastADGift *entity.TAdGift
	)
	defer func() {
		if r := recover(); r != nil {
			// 同时打印到日志文件和标准输出中
			global.MyLogger(ctx).Err(err).Msgf("%+v\n%+v", r, string(debug.Stack()))
		}
	}()

	// 绑定请求参数
	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
		response.RespFail(ctx, i18n.RetMsgParamParseErr, nil)
		return
	}
	global.MyLogger(ctx).Info().Msgf("req: %+v", *req)

	// validate user
	user, err = common.ValidateClaims(ctx)
	if err != nil {
		return
	}

	// validate order
	err = dao.TAd.Ctx(ctx).Where(do.TAd{Name: req.Name}).Scan(&adInfo)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("query ad failed")
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}
	if adInfo == nil {
		global.MyLogger(ctx).Warn().Msgf("ad \"%s\" is not exist", req.Name)
		response.RespFail(ctx, i18n.RetMsgParamInvalid, nil)
		return
	}

	err = dao.TAdGift.Ctx(ctx).
		Where(do.TAdGift{UserId: user.Id}).
		Order(dao.TAdGift.Columns().Id, constant.OrderTypeDesc).Limit(1).
		Scan(&lastADGift)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("query ad gift failed")
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}
	now := gtime.Now()
	if lastADGift != nil && lastADGift.CreatedAt.Add(time.Duration(adInfo.ExposureTime)*time.Second).After(now) {
		err = gerror.Newf(`last ad gift at "%s" + "%d"s is after now(%s), that is limited, userId: %d`,
			lastADGift.CreatedAt.String(), adInfo.ExposureTime, now.String(), user.Id)
		global.MyLogger(ctx).Err(err).Msgf("query ad gift failed")
		response.RespFail(ctx, i18n.RetMsgInternalErr, nil)
		return
	}

	// 加时长
	var (
		userUpdate = do.TUser{
			Kicked:    0,
			UpdatedAt: gtime.Now(),
			Version:   user.Version + 1,
		}
		_ctx     = ctx
		affected int64
	)
	if service.IsVIPExpired(user) {
		userUpdate.ExpiredTime = time.Now().Unix() + int64(adInfo.GiftDuration)
	} else {
		userUpdate.ExpiredTime = user.ExpiredTime + int64(adInfo.GiftDuration)
	}
	err = dao.TPayOrder.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		affected, err = dao.TUser.Ctx(ctx).Data(userUpdate).Where(do.TUser{
			Id:      user.Id,
			Version: user.Version,
		}).UpdateAndGetAffected()
		if err != nil {
			global.MyLogger(_ctx).Err(err).Msgf(`update t_user failed`)
			return err
		}
		if affected != 1 {
			err = fmt.Errorf("update t_user affected(%d) != 1", affected)
			global.MyLogger(_ctx).Err(err).Msgf("update t_user failed")
			return err
		}

		// 记录流水
		_, err = dao.TAdGift.Ctx(ctx).Data(do.TAdGift{
			UserId:       user.Id,
			AdId:         adInfo.Id,
			AdName:       adInfo.Name,
			ExposureTime: adInfo.ExposureTime,
			GiftDuration: adInfo.GiftDuration,
			CreatedAt:    gtime.Now(),
		}).Insert()
		if err != nil {
			global.MyLogger(_ctx).Err(err).Msgf(`insert TAdGift failed`)
			return err
		}
		return nil
	})
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("add ad gift failed")
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}
	global.MyLogger(ctx).Info().Msgf("ADName: %s, notify success, userId: %d", req.Name, user.Id)
	response.RespOk(ctx, i18n.RetMsgSuccess, ADCompletionNotifyRes{})
}
