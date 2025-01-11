package order

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"go-speed/constant"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/response"
)

const (
	ConfirmResultPass   = "PASS"   // 审核通过
	ConfirmResultRevert = "REVERT" // 审核不通过，需要撤销时间
)

type ConfirmOrderStatusReq struct {
	OrderNo       string `form:"order_no" binding:"required" json:"order_no"`
	ConfirmResult string `form:"confirm_result" binding:"required" json:"confirm_result"`
}

type ConfirmOrderStatusRes struct {
}

func ConfirmOrderStatus(ctx *gin.Context) {
	var (
		err              error
		req              = new(ConfirmOrderStatusReq)
		payOrderEntity   *entity.TPayOrder
		attrRecordEntity *entity.TUserVipAttrRecord
		targetStatus     string
	)
	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
		response.ResFail(ctx, err.Error())
		return
	}
	global.MyLogger(ctx).Info().Msgf("request: %+v", *req)

	// validate params
	switch req.ConfirmResult {
	case ConfirmResultPass:
		targetStatus = constant.ParOrderStatusAdminConfirmPassed
	case ConfirmResultRevert:
		targetStatus = constant.ParOrderStatusAdminConfirmClosed
	default:
		err = fmt.Errorf(`param "ConfirmResult" invalid`)
		global.MyLogger(ctx).Err(err).Msgf(`ConfirmResult: %s`, req.ConfirmResult)
		response.ResFail(ctx, err.Error())
		return
	}
	payOrderEntity, err = ValidateOrder(ctx, req.OrderNo)
	if err != nil {
		return
	}
	global.MyLogger(ctx).Info().Msgf("payOrderEntity: %+v", *payOrderEntity)
	attrRecordEntity, err = ValidateUserVipAttrRecord(ctx, payOrderEntity.Email, req.OrderNo)
	if err != nil {
		return
	}

	global.MyLogger(ctx).Info().Msgf("attrRecordEntity: %+v", *attrRecordEntity)
	// validate order status
	switch payOrderEntity.Status {
	case targetStatus:
		response.ResFail(ctx, "不需要重复确认，请刷新页面查看订单最新状态")
		return

	case constant.ParOrderStatusPaid:
		if payOrderEntity.PaymentChannelId != constant.PayChannelBankCardPay {
			response.ResFail(ctx, "只有银行卡支付的订单才需要审核")
			return
		}
	default:
		err = fmt.Errorf(`只有用户确认支付完成的订单才可以发起审核，当前订单状态为 "%s"，`, payOrderEntity.Status)
		response.ResFail(ctx, err.Error())
		return
	}

	// confirm order process
	switch targetStatus {
	case constant.ParOrderStatusAdminConfirmPassed:
		err = AdminConfirmPassedProcess(ctx, payOrderEntity)
	case constant.ParOrderStatusAdminConfirmClosed:
		err = AdminConfirmRevertProcess(ctx, payOrderEntity, attrRecordEntity)
	}
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("confirm failed")
		return
	}
	response.ResOk(ctx, "成功")
	return
}

func ValidateUserVipAttrRecord(ctx *gin.Context, email, orderNo string) (record *entity.TUserVipAttrRecord, err error) {
	err = dao.TUserVipAttrRecord.Ctx(ctx).Where(do.TUserVipAttrRecord{Email: email, OrderNo: orderNo}).Scan(&record)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("query TUserVipAttrRecord failed")
		response.ResFail(ctx, err.Error())
		return
	}

	if record == nil {
		err = fmt.Errorf(`TUserVipAttrRecord is not exist`)
		global.MyLogger(ctx).Err(err).Msgf(`order: %s`, orderNo)
		response.ResFail(ctx, err.Error())
	}
	return
}

func ValidateOrder(ctx *gin.Context, orderNo string) (payOrderEntity *entity.TPayOrder, err error) {
	err = dao.TPayOrder.Ctx(ctx).Where(do.TPayOrder{OrderNo: orderNo}).Scan(&payOrderEntity)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("query pay order failed")
		response.ResFail(ctx, err.Error())
		return
	}

	if payOrderEntity == nil {
		err = fmt.Errorf(`order is not exist`)
		global.MyLogger(ctx).Err(err).Msgf(`order: %s`, orderNo)
		response.ResFail(ctx, err.Error())
	}
	return
}

func AdminConfirmPassedProcess(ctx *gin.Context, payOrder *entity.TPayOrder) (err error) {
	_ctx := ctx
	err = dao.TPayOrder.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		var affected int64
		affected, err = dao.TPayOrder.Ctx(ctx).Data(do.TPayOrder{
			Status:    constant.ParOrderStatusAdminConfirmPassed,
			UpdatedAt: gtime.Now(),
			Version:   payOrder.Version + 1,
		}).Where(do.TPayOrder{Id: payOrder.Id, Version: payOrder.Version}).UpdateAndGetAffected()
		if err != nil {
			global.MyLogger(_ctx).Err(err).Msgf(`confirm order failed`)
			response.ResFail(_ctx, err.Error())
			return err
		}
		if affected != 1 {
			err = fmt.Errorf(`您看到的页面数据可能不是最新的，请刷新页面后再重试。`)
			response.ResFail(_ctx, err.Error())
			return err
		}
		return nil
	})
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("confirm order failed")
		return
	}
	return err
}

func ValidateUser(ctx *gin.Context, userId uint64) (user *entity.TUser, err error) {
	err = dao.TUser.Ctx(ctx).Where(do.TUser{Id: userId}).Scan(&user)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("query user failed")
		response.ResFail(ctx, err.Error())
		return
	}

	if user == nil {
		err = fmt.Errorf(`user(%d) is not exist`, userId)
		global.MyLogger(ctx).Err(err).Msgf(`TUser is nil`)
		response.ResFail(ctx, err.Error())
	}
	return
}

func AdminConfirmRevertProcess(ctx *gin.Context, payOrder *entity.TPayOrder, attrRecord *entity.TUserVipAttrRecord) (err error) {
	var (
		affected         int64
		lastInsertId     int64
		user             *entity.TUser
		direct           *entity.TUser
		directAttrRecord *entity.TUserVipAttrRecord
	)
	user, err = ValidateUser(ctx, payOrder.UserId)
	if err != nil {
		return
	}
	global.MyLogger(ctx).Info().Msgf("user: %+v", *user)
	direct, directAttrRecord, err = GetDirectAttrRecord(ctx, user.Id, payOrder.OrderNo)
	if err != nil {
		return
	}
	_ctx := ctx
	err = dao.TPayOrder.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// update order status
		affected, err = dao.TPayOrder.Ctx(ctx).Data(do.TPayOrder{
			Status:    constant.ParOrderStatusAdminConfirmClosed,
			UpdatedAt: gtime.Now(),
			Version:   payOrder.Version + 1,
		}).Where(do.TPayOrder{Id: payOrder.Id, Version: payOrder.Version}).UpdateAndGetAffected()
		if err != nil {
			global.MyLogger(_ctx).Err(err).Msgf(`confirm order failed`)
			response.ResFail(_ctx, err.Error())
			return err
		}
		if affected != 1 {
			err = fmt.Errorf(`您看到的页面数据可能不是最新的，请刷新页面后再重试。`)
			response.ResFail(_ctx, err.Error())
			return err
		}
		global.MyLogger(_ctx).Info().Msgf("1 affected: %+v", affected)

		// attr mark revert tag
		affected, err = dao.TUserVipAttrRecord.Ctx(ctx).Data(do.TUserVipAttrRecord{
			IsRevert: constant.OrderRevertFlag,
		}).Where(do.TUserVipAttrRecord{Id: attrRecord.Id}).UpdateAndGetAffected()
		if err != nil {
			global.MyLogger(_ctx).Err(err).Msgf(`attr mark revert tag failed`)
			response.ResFail(_ctx, err.Error())
			return err
		}
		global.MyLogger(_ctx).Info().Msgf("2 affected: %+v", affected)

		// revert user ExpiredTime
		addTime := -1 * int64(attrRecord.ExpiredTimeTo-attrRecord.ExpiredTimeFrom)
		newExpiredTime := user.ExpiredTime + addTime
		affected, err = dao.TUser.Ctx(ctx).Data(do.TUser{
			ExpiredTime: newExpiredTime,
			UpdatedAt:   gtime.Now(),
			Version:     user.Version + 1,
			Kicked:      0,
		}).Where(do.TUser{
			Id:      user.Id,
			Version: user.Version,
		}).UpdateAndGetAffected()
		if err != nil {
			global.MyLogger(_ctx).Err(err).Msgf(`update t_user failed`)
			response.ResFail(_ctx, err.Error())
			return err
		}
		if affected != 1 {
			err = fmt.Errorf("update t_user affected(%d) != 1", affected)
			global.MyLogger(_ctx).Err(err).Msgf("update t_user failed")
			response.ResFail(_ctx, err.Error())
			return err
		}
		global.MyLogger(_ctx).Info().Msgf("3 affected: %+v", affected)

		// add attr record
		lastInsertId, err = dao.TUserVipAttrRecord.Ctx(ctx).Data(do.TUserVipAttrRecord{
			Email:           user.Email,
			Source:          constant.UserVipAttrOpSourceAdminRevertOrder,
			OrderNo:         payOrder.OrderNo + "-revert",
			ExpiredTimeFrom: user.ExpiredTime,
			ExpiredTimeTo:   newExpiredTime,
			Desc:            fmt.Sprintf("ExpiredTime add[%d]", addTime),
			CreatedAt:       gtime.Now(),
		}).InsertAndGetId()
		if err != nil {
			global.MyLogger(_ctx).Err(err).Msgf(`insert TUserVipAttrRecords failed`)
			response.ResFail(_ctx, err.Error())
			return err
		}
		global.MyLogger(_ctx).Info().Msgf(">>>> insert TUserVipAttrRecords, lastInsertId: %d", lastInsertId)
		// 推荐人也要扣减时间
		if direct != nil && directAttrRecord != nil {
			addTimeDirect := -1 * int64(directAttrRecord.ExpiredTimeTo-directAttrRecord.ExpiredTimeFrom)
			newExpiredTimeDirect := direct.ExpiredTime + addTimeDirect

			affected, err = dao.TUser.Ctx(ctx).Data(do.TUser{
				ExpiredTime: newExpiredTimeDirect,
				UpdatedAt:   gtime.Now(),
				Version:     direct.Version + 1,
				Kicked:      0,
			}).Where(do.TUser{
				Id:      direct.Id,
				Version: direct.Version,
			}).UpdateAndGetAffected()
			if err != nil {
				global.MyLogger(_ctx).Err(err).Msgf(`update direct t_user failed`)
				response.ResFail(_ctx, err.Error())
				return err
			}
			if affected != 1 {
				err = fmt.Errorf("update direct t_user affected(%d) != 1", affected)
				global.MyLogger(_ctx).Err(err).Msgf("update direct t_user failed")
				response.ResFail(_ctx, err.Error())
				return err
			}

			// add attr record
			lastInsertId, err = dao.TUserVipAttrRecord.Ctx(ctx).Data(do.TUserVipAttrRecord{
				Email:           direct.Email,
				Source:          constant.UserVipAttrOpSourceAdminRevertOrder,
				OrderNo:         payOrder.OrderNo + "-revert-direct",
				ExpiredTimeFrom: direct.ExpiredTime,
				ExpiredTimeTo:   newExpiredTimeDirect,
				Desc:            fmt.Sprintf("expiredTime add[%d]", addTimeDirect),
				CreatedAt:       gtime.Now(),
			}).InsertAndGetId()
			if err != nil {
				global.MyLogger(_ctx).Err(err).Msgf(`insert TUserVipAttrRecords failed`)
				response.ResFail(_ctx, err.Error())
				return err
			}
			global.MyLogger(_ctx).Info().Msgf(">>>> insert TUserVipAttrRecords, lastInsertId: %d", lastInsertId)
		}
		return nil
	})
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("sync order status failed")
		return
	}
	return err
}

func GetDirectAttrRecord(ctx *gin.Context, userId int64, orderNo string) (direct *entity.TUser, record *entity.TUserVipAttrRecord, err error) {
	var (
		userTeam *entity.TUserTeam
	)
	err = dao.TUserTeam.Ctx(ctx).Where(do.TUserTeam{UserId: userId}).Scan(&userTeam)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf(`query user team failed`)
		response.ResFail(ctx, err.Error())
		return
	}
	if userTeam == nil {
		global.MyLogger(ctx).Info().Msgf("userId: %d", userId)
		return
	}
	global.MyLogger(ctx).Info().Msgf(">>>> DirectId: %d", userTeam.DirectId)
	if userTeam.DirectId <= 0 {
		global.MyLogger(ctx).Info().Msgf("direct user is not exist")
		return
	}

	// 查询推荐人用户信息
	direct, err = ValidateUser(ctx, uint64(userTeam.DirectId))
	if err != nil {
		return
	}
	global.MyLogger(ctx).Info().Msgf(">>>> Direct Email: %s", direct.Email)

	// 推荐人记录
	record, err = ValidateUserVipAttrRecord(ctx, direct.Email, orderNo)
	if err != nil {
		return
	}
	return
}
