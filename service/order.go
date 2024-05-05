package service

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
	"go-speed/util/pay/pnsafepay"
	"strconv"
	"strings"
)

const (
	ReturnStatusSuccess = "success"
)

func SyncOrderStatus(ctx *gin.Context, orderNo string) (status string, err error) {
	var (
		affected     int64
		lastInsertId int64
		payOrder     *entity.TPayOrder
		payResponse  *pnsafepay.QueryOrderResponse
		userEntity   *entity.TUser
	)
	global.MyLogger(ctx).Info().Msgf("$$$$$$$$$$$$$$ orderNo: %s", orderNo)

	err = dao.TPayOrder.Ctx(ctx).Where(do.TPayOrder{
		OrderNo: orderNo,
	}).Scan(&payOrder)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("query pay order failed")
		return
	}
	if payOrder == nil {
		err = fmt.Errorf(`order not exists`)
		global.MyLogger(ctx).Err(err).Msgf("order is nil")
		return
	}
	if payOrder.Status == constant.ParOrderStatusPaid {
		global.MyLogger(ctx).Info().Msgf("$$$$$$$$$$$$$$ orderNo: %s, status: %s, has notified success", orderNo, payOrder.Status)
		return ReturnStatusSuccess, nil
	}

	payResponse, err = pnsafepay.QueryPayOrder(ctx, payOrder.OrderNo)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("QueryPayOrder failed")
		return
	}

	if strings.ToLower(payResponse.CheckStatus) != ReturnStatusSuccess {
		err = fmt.Errorf(`"checkstatus"(%s) is not success`, payResponse.CheckStatus)
		global.MyLogger(ctx).Err(err).Msgf("CheckStatus failed")
		return
	}
	orderAmount, _ := strconv.ParseFloat(payResponse.OrderAmount, 64)
	payOrderAmount, _ := strconv.ParseFloat(payOrder.OrderAmount, 64)
	if payResponse.OrderNo != payOrder.OrderNo ||
		orderAmount != payOrderAmount ||
		payResponse.MerNo != global.Config.PNSafePay.MerNo {
		err = fmt.Errorf(`QueryPayOrder response order info is invalid, [order_no,order_amount,mer_no] is not right`)
		global.MyLogger(ctx).Err(err).Msgf(`payResponse: %+v, payOrder: %+v`, *payResponse, *payOrder)
		return
	}

	err = dao.TUser.Ctx(ctx).Where(do.TUser{Id: payOrder.UserId}).Scan(&userEntity)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("query user info failed")
		return
	}
	if userEntity == nil {
		err = fmt.Errorf(`user not exist, userId: %d, emal: %s`, payOrder.UserId, payOrder.Email)
		global.MyLogger(ctx).Err(err).Msgf(`user not exist`)
		return
	}

	err = dao.TPayOrder.Ctx(ctx).Transaction(ctx, func(_ctx context.Context, tx gdb.TX) error {
		updateDo := do.TPayOrder{
			ResultStatus:       payResponse.ResultStatus,
			OrderRealityAmount: payResponse.OrderRealityAmount,
			Version:            payOrder.Version + 1,
			UpdatedAt:          gtime.Now(),
		}

		// 订单支付成功时，需要执行相关操作
		if payResponse.ResultStatus == ReturnStatusSuccess {
			// 修改订单状态
			updateDo.Status = constant.ParOrderStatusPaid
			global.MyLogger(ctx).Info().Msgf(">>>> 1 order status from(%s) to(%s)", payOrder.Status, constant.ParOrderStatusPaid)

			// 加时长
			affected, err = dao.TUser.Ctx(ctx).Data(do.TUser{
				ExpiredTime: userEntity.ExpiredTime + 30*24*60*60, // 加30天
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

			global.MyLogger(ctx).Info().Msgf(">>>> 2 add(%d) user ExpiredTime from(%d) to(%d)",
				30*24*60*60, userEntity.ExpiredTime, userEntity.ExpiredTime+30*24*60*60)

			// 记录操作流水
			lastInsertId, err = dao.TUserVipAttrRecords.Ctx(ctx).Data(do.TUserVipAttrRecords{
				Email:       userEntity.Email,
				Source:      constant.UserVipAttrOpSourcePayOrder,
				OrderNo:     payOrder.OrderNo,
				ExpiredTime: userEntity.ExpiredTime + 30*24*60*60,
				Desc:        fmt.Sprintf("expired_time add [%d]", 30*24*60*60),
				CreatedAt:   gtime.Now(),
			}).InsertAndGetId()
			if err != nil {
				global.MyLogger(ctx).Err(err).Msgf(`insert TUserVipAttrRecords failed`)
				return err
			}
			global.MyLogger(ctx).Info().Msgf(">>>> 3 insert TUserVipAttrRecords, lastInsertId: %d", lastInsertId)
		}

		affected, err = dao.TPayOrder.Ctx(ctx).Data(updateDo).Where(do.TPayOrder{
			Id:      payOrder.Id,
			Version: payOrder.Version,
		}).UpdateAndGetAffected()
		if err != nil {
			global.MyLogger(ctx).Err(err).Msgf("update TPayOrder failed")
			return err
		}

		if affected != 1 {
			err = fmt.Errorf("update TPayOrder affected(%d) != 1", affected)
			global.MyLogger(ctx).Err(err).Msgf("update TPayOrder failed")
			return err
		}
		return nil
	})
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("sync order status failed")
		return
	}
	global.MyLogger(ctx).Info().Msgf("sync order status success, orderNo: %s", orderNo)
	return payResponse.ResultStatus, nil
}
