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
	"go-speed/util/pay/applepay"
	"go-speed/util/pay/freekassa"
	"go-speed/util/pay/pnsafepay"
	russ_pay "go-speed/util/pay/russ-pay"
	"go-speed/util/pay/russpay"
	"go-speed/util/pay/upay"
	"go-speed/util/pay/webmoney"
	"golang.org/x/exp/rand"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

type PayNotifyReq struct {
	MerchantId string `form:"MERCHANT_ID" json:"MERCHANT_ID"`
	Amount     string `form:"AMOUNT" json:"AMOUNT"`
	IntId      string `form:"intid" json:"intid"`
	OrderId    string `form:"MERCHANT_ORDER_ID" json:"MERCHANT_ORDER_ID"`
	PEmail     string `form:"P_EMAIL" json:"P_EMAIL"`
	CurId      string `form:"CUR_ID" json:"CUR_ID"`
	Commission string `form:"commission" json:"commission"`
	Sign       string `form:"SIGN" json:"SIGN"`
}

type ApplePayNotifyReq struct {
	TransactionId string `form:"transaction_id" json:"transaction_id" dc:"苹果支付平台订单号"`
	ReceiptData   string `form:"receipt_data" json:"receipt_data" dc:"苹果支付平台凭证"`
}

func SyncOrderStatus(ctx *gin.Context, orderNo string, notifyData interface{}) (status string, err error) {
	var (
		affected             int64
		lastInsertId         int64
		payOrder             *entity.TPayOrder
		userEntity           *entity.TUser
		goodsEntity          *entity.TGoods
		resultStatus         string
		orderRealityAmount   string
		directUserEntity     *entity.TUser
		directNewExpiredTime int64
		directAddExpiredTime int64
		paymentProof         string
	)
	global.MyLogger(ctx).Info().Msgf("$$$$$$$$$$$$$$ orderNo: %s", orderNo)
	defer func() {
		if r := recover(); r != nil {
			// 同时打印到日志文件和标准输出中
			global.MyLogger(ctx).Err(err).Msgf("%+v\n%+v", r, string(debug.Stack()))
		}
	}()

	// 查找订单
	err = dao.TPayOrder.Ctx(ctx).Where(do.TPayOrder{OrderNo: orderNo}).Scan(&payOrder)
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
		return constant.ReturnStatusSuccess, nil
	}

	// 从支付平台查询订单状态
	switch payOrder.PaymentChannelId {
	case constant.PayChannelPnSafePay:
		var payResponse *pnsafepay.QueryOrderResponse
		payResponse, err = pnsafepay.QueryPayOrder(ctx, payOrder.OrderNo)
		if err != nil {
			global.MyLogger(ctx).Err(err).Msgf("QueryPayOrder failed")
			return
		}

		if strings.ToLower(payResponse.CheckStatus) != constant.ReturnStatusSuccess {
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
		resultStatus, orderRealityAmount = payResponse.ResultStatus, payResponse.OrderRealityAmount
	case constant.PayChannelUPay, constant.PayChannelBTCPay:
		var match bool
		match, err = upay.CheckBinanceOrder(ctx, time.Minute*30, payOrder.OrderAmount)
		if err != nil {
			global.MyLogger(ctx).Err(err).Msgf("CheckBinanceOrder failed")
			return
		}

		if !match {
			global.MyLogger(ctx).Info().Msgf("$$$$$$$$$$$$$$ orderNo: %s", orderNo)
			return constant.ReturnStatusWaiting, nil
		}
		resultStatus, orderRealityAmount = constant.ReturnStatusSuccess, payOrder.OrderAmount
	case constant.PayChannelBankCardPay:
		// 银行卡支付，后台无法校验，直接标记支付成功
		resultStatus, orderRealityAmount = constant.ReturnStatusSuccess, payOrder.OrderAmount
	case constant.PayChannelWebMoneyPay:
		var resp *webmoney.Operation
		resp, err = webmoney.QueryOrder(ctx, payOrder.OrderNo)
		if err != nil {
			global.MyLogger(ctx).Err(err).Msgf("QueryDeal failed")
			return
		}
		// -0 – simple (or protected, successfully completed),
		// -4 – protected (not completed),
		// -12 – protected (refunded).
		if resp == nil || resp.OperType == "4" {
			global.MyLogger(ctx).Info().Msgf("$$$$$$$$$$$$$$ orderNo: %s, waiting to pay (not completed)", orderNo)
			return constant.ReturnStatusWaiting, nil
		}
		if resp.OperType == "12" {
			global.MyLogger(ctx).Info().Msgf("$$$$$$$$$$$$$$ orderNo: %s, refunded", orderNo)
			return constant.ReturnStatusFail, nil
		}

		if resp.OperType == "0" {
			pass, err := CheckAmount(ctx, resp.Amount, payOrder.OrderAmount)
			if err != nil {
				return constant.ReturnStatusWaiting, err
			}

			if !pass {
				global.MyLogger(ctx).Info().Msgf("$$$$$$$$$$$$$$ OrderNo: %s, waiting to pay", orderNo)
				return constant.ReturnStatusWaiting, nil
			}
			resultStatus, orderRealityAmount = constant.ReturnStatusSuccess, resp.Amount
		} else {
			global.MyLogger(ctx).Info().Msgf("$$$$$$$$$$$$$$ orderNo: %s, waiting to pay", orderNo)
			return constant.ReturnStatusWaiting, nil
		}

		////0 – unpaid 1 – paid with protection 2 – fully paid or paid without protection 3 – rejected
		//if resp == nil || resp.State == 0 {
		//	global.MyLogger(ctx).Info().Msgf("$$$$$$$$$$$$$$ orderNo: %s, waiting to pay", orderNo)
		//	return constant.ReturnStatusWaiting, nil
		//}
		//if resp.State == 3 {
		//	global.MyLogger(ctx).Info().Msgf("$$$$$$$$$$$$$$ orderNo: %s, rejected", orderNo)
		//	return constant.ReturnStatusFail, nil
		//}
		//
		//if resp.State == 1 || resp.State == 2 {
		//	resultStatus, orderRealityAmount = constant.ReturnStatusSuccess, strconv.FormatFloat(resp.Amount, 'f', -1, 64)
		//} else {
		//	global.MyLogger(ctx).Info().Msgf("$$$$$$$$$$$$$$ orderNo: %s, waiting to pay", orderNo)
		//	return constant.ReturnStatusWaiting, nil
		//}

	case constant.PayChannelFreekassa_12, constant.PayChannelFreekassa_36,
		constant.PayChannelFreekassa_43, constant.PayChannelFreekassa_44, constant.PayChannelFreekassa_7:
		var pass bool
		if notifyData == nil {
			order, err := freekassa.QueryOrder(ctx, orderNo)
			if order == nil {
				global.MyLogger(ctx).Info().Msgf("$$$$$$$$$$$$$$ OrderNo: %s, waiting to pay", orderNo)
				return constant.ReturnStatusWaiting, nil
			}

			if order.Status == 1 {
				pass, err = checkAmountFloat64(ctx, order.Amount, payOrder.OrderAmount)
				if err != nil {
					return constant.ReturnStatusWaiting, err
				}
			}
			orderRealityAmount = payOrder.OrderAmount
		} else {
			notifyReq := notifyData.(*PayNotifyReq)
			pass, err = CheckAmount(ctx, notifyReq.Amount, payOrder.OrderAmount)
			if err != nil {
				return constant.ReturnStatusWaiting, err
			}
			orderRealityAmount = notifyReq.Amount
		}
		if !pass {
			global.MyLogger(ctx).Info().Msgf("$$$$$$$$$$$$$$ OrderNo: %s, waiting to pay", orderNo)
			return constant.ReturnStatusWaiting, nil
		}
		resultStatus = constant.ReturnStatusSuccess
	case constant.PayChannelApplePay:
		var (
			applePayReq    = notifyData.(*ApplePayNotifyReq)
			applePayStatus int
		)
		applePayStatus, err = applepay.AppleVerify(ctx, applePayReq.TransactionId, applePayReq.ReceiptData)
		if err != nil {
			global.MyLogger(ctx).Err(err).Msgf("AppleVerify failed")
			return constant.ReturnStatusWaiting, err
		}
		orderRealityAmount = payOrder.OrderAmount
		if applePayStatus != 0 {
			global.MyLogger(ctx).Info().Msgf("$$$$$$$$$$$$$$ OrderNo: %s, waiting to pay", orderNo)
			return constant.ReturnStatusWaiting, nil
		}
		paymentProof = applePayReq.TransactionId
		resultStatus = constant.ReturnStatusSuccess
	case constant.PayChannelRussPayBankCard, constant.PayChannelRussPaySBP, constant.PayChannelRussPaySBER:
		if payOrder.OrderData == "" {
			global.MyLogger(ctx).Info().Msgf("russpay OrderData is nil")
			return constant.ReturnStatusWaiting, nil
		}
		var resp *russpay.QueryOrderResponse
		resp, err = russpay.QueryOrder(ctx, russpay.QueryOrderReq{BillingNumber: payOrder.OrderData})
		if err != nil {
			global.MyLogger(ctx).Err(err).Msgf("russpay failed")
			return constant.ReturnStatusWaiting, err
		}

		if resp != nil &&
			strings.ToLower(resp.PaymentStatus) == strings.ToLower("SUCCESSFUL") &&
			resp.BillingNumber == payOrder.OrderData {
			resultStatus = constant.ReturnStatusSuccess
			orderRealityAmount = payOrder.OrderAmount
		} else {
			return constant.ReturnStatusWaiting, nil
		}
	case constant.PayChannelRussNewPayCard, constant.PayChannelRussNewPaySBP:
		var resp *russ_pay.QueryPaymentOrderRes
		resp, err = russ_pay.QueryOrder(ctx, russ_pay.QueryOrderReq{MerchantOrderID: orderNo})
		if err != nil {
			global.MyLogger(ctx).Err(err).Msgf("russ-new-pay failed")
			return constant.ReturnStatusWaiting, err
		}

		if resp != nil && strings.ToLower(resp.Data.Status) == "success" {
			var pass bool
			pass, err = CheckAmount(ctx, resp.Data.DepositAmount, payOrder.OrderData)
			if err != nil {
				return constant.ReturnStatusWaiting, err
			}
			if !pass {
				global.MyLogger(ctx).Info().Msgf("$$$$$$$$$$$$$$ OrderNo: %s, waiting to pay", orderNo)
				return constant.ReturnStatusWaiting, nil
			}

			resultStatus = constant.ReturnStatusSuccess
			orderRealityAmount = payOrder.OrderAmount
		} else {
			return constant.ReturnStatusWaiting, nil
		}
	}

	// 查询用户信息
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

	// 查询套餐信息
	err = dao.TGoods.Ctx(ctx).Where(do.TGoods{Id: payOrder.GoodsId}).Scan(&goodsEntity)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf(`goods not exist`)
		return
	}
	if goodsEntity == nil {
		err = fmt.Errorf(`goodsEntity not exist, userId: %d, GoodsId: %d`, payOrder.UserId, payOrder.GoodsId)
		global.MyLogger(ctx).Err(err).Msgf(`GoodsId not exist`)
		return
	}

	directUserEntity, directNewExpiredTime, directAddExpiredTime, _ = BuildDirectUserExpiredTime(ctx, userEntity.Id, goodsEntity)

	_ctx := ctx
	err = dao.TPayOrder.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		global.MyLogger(_ctx).Info().Msgf(`1`)
		updateDo := do.TPayOrder{
			ResultStatus:       resultStatus,
			OrderRealityAmount: orderRealityAmount,
			Version:            payOrder.Version + 1,
			UpdatedAt:          gtime.Now(),
		}
		if payOrder.PaymentChannelId == constant.PayChannelApplePay {
			updateDo.PaymentProof = paymentProof
		}

		// 订单支付成功时，需要执行相关操作
		if resultStatus == constant.ReturnStatusSuccess {
			// 修改订单状态
			updateDo.Status = constant.ParOrderStatusPaid
			global.MyLogger(_ctx).Info().Msgf(">>>> order status from(%s) to(%s)", payOrder.Status, constant.ParOrderStatusPaid)

			// 加时长
			var (
				newExpiredTime int64
				addExpiredTime = int64(goodsEntity.Period) * constant.DaySeconds
			)
			global.MyLogger(_ctx).Info().Msgf(`2`)
			if IsVIPExpired(userEntity) {
				newExpiredTime = time.Now().Unix() + addExpiredTime
				global.MyLogger(_ctx).Info().Msgf(`--------------11 newExpiredTime: %d`, newExpiredTime)
			} else {
				newExpiredTime = userEntity.ExpiredTime + addExpiredTime
				global.MyLogger(_ctx).Info().Msgf(`--------------12 newExpiredTime: %d`, newExpiredTime)
			}
			// 随机赠送
			giftDay := randomGiftDay(goodsEntity.Low, goodsEntity.High)
			global.MyLogger(_ctx).Info().Msgf(`3`)
			newExpiredTime += int64(giftDay * constant.DaySeconds)
			userUpdate := do.TUser{
				ExpiredTime: newExpiredTime,
				UpdatedAt:   gtime.Now(),
				Version:     userEntity.Version + 1,
				Kicked:      0,
			}
			// 加用户等级
			// 过期充值时，按套餐的等级来
			if IsVIPExpired(userEntity) {
				userUpdate.Level = goodsEntity.MType
			} else {
				// 没有过期的续费操作，按高等级来
				if goodsEntity.MType > userEntity.Level {
					userUpdate.Level = goodsEntity.MType
				}
			}

			global.MyLogger(_ctx).Info().Msgf(`4`)
			affected, err = dao.TUser.Ctx(ctx).Data(userUpdate).Where(do.TUser{
				Id:      userEntity.Id,
				Version: userEntity.Version,
			}).UpdateAndGetAffected()
			if err != nil {
				global.MyLogger(_ctx).Err(err).Msgf(`update t_user failed`)
				return err
			}
			global.MyLogger(_ctx).Info().Msgf(`5`)
			if affected != 1 {
				err = fmt.Errorf("update t_user affected(%d) != 1", affected)
				global.MyLogger(_ctx).Err(err).Msgf("update t_user failed")
				return err
			}

			global.MyLogger(_ctx).Info().Msgf(">>>> add(%d) user(%s) ExpiredTime from(%d) to(%d), giftDay(%d)",
				addExpiredTime, userEntity.Email, userEntity.ExpiredTime, newExpiredTime, giftDay)

			// 记录操作流水
			lastInsertId, err = dao.TUserVipAttrRecord.Ctx(ctx).Data(do.TUserVipAttrRecord{
				Email:           userEntity.Email,
				Source:          constant.UserVipAttrOpSourcePayOrder,
				OrderNo:         payOrder.OrderNo,
				ExpiredTimeFrom: userEntity.ExpiredTime,
				ExpiredTimeTo:   newExpiredTime,
				Desc:            fmt.Sprintf("ExpiredTime add[%d], giftDay(%d)", addExpiredTime, giftDay),
				CreatedAt:       gtime.Now(),
			}).InsertAndGetId()
			if err != nil {
				global.MyLogger(_ctx).Err(err).Msgf(`insert TUserVipAttrRecords failed`)
				return err
			}
			global.MyLogger(_ctx).Info().Msgf(">>>> insert TUserVipAttrRecords, lastInsertId: %d", lastInsertId)

			// direct process
			if directUserEntity != nil {
				global.MyLogger(_ctx).Info().Msgf(`6`)
				affected, err = dao.TUser.Ctx(ctx).Data(do.TUser{
					ExpiredTime: directNewExpiredTime,
					UpdatedAt:   gtime.Now(),
					Version:     directUserEntity.Version + 1,
					Kicked:      0,
				}).Where(do.TUser{
					Id:      directUserEntity.Id,
					Version: directUserEntity.Version,
				}).UpdateAndGetAffected()
				if err != nil {
					global.MyLogger(_ctx).Err(err).Msgf(`update direct t_user failed`)
					return err
				}
				if affected != 1 {
					err = fmt.Errorf("update direct t_user affected(%d) != 1", affected)
					global.MyLogger(_ctx).Err(err).Msgf("update direct t_user failed")
					return err
				}
				global.MyLogger(_ctx).Info().Msgf(`7`)

				global.MyLogger(_ctx).Info().Msgf("add(%d) direct user(%s) ExpiredTime from(%d) to(%d)",
					directAddExpiredTime, directUserEntity.Email, directUserEntity.ExpiredTime, directNewExpiredTime)

				// 记录操作流水
				lastInsertId, err = dao.TUserVipAttrRecord.Ctx(ctx).Data(do.TUserVipAttrRecord{
					Email:           directUserEntity.Email,
					Source:          constant.UserVipAttrOpSourceDirectGift,
					OrderNo:         orderNo,
					ExpiredTimeFrom: directUserEntity.ExpiredTime,
					ExpiredTimeTo:   directNewExpiredTime,
					Desc:            fmt.Sprintf("ExpiredTime add[%d]", directAddExpiredTime),
					CreatedAt:       gtime.Now(),
				}).InsertAndGetId()
				if err != nil {
					global.MyLogger(_ctx).Err(err).Msgf(`insert direct TUserVipAttrRecords failed`)
					return err
				}
				global.MyLogger(_ctx).Info().Msgf("insert direct TUserVipAttrRecords, lastInsertId: %d", lastInsertId)
			}
		}

		affected, err = dao.TPayOrder.Ctx(ctx).Data(updateDo).Where(do.TPayOrder{
			Id:      payOrder.Id,
			Version: payOrder.Version,
		}).UpdateAndGetAffected()
		if err != nil {
			global.MyLogger(_ctx).Err(err).Msgf("update TPayOrder failed")
			return err
		}

		if affected != 1 {
			err = fmt.Errorf("update TPayOrder affected(%d) != 1", affected)
			global.MyLogger(_ctx).Err(err).Msgf("update TPayOrder failed")
			return err
		}
		return nil
	})
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("sync order status failed")
		return
	}

	global.MyLogger(ctx).Info().Msgf("sync order status success, orderNo: %s", orderNo)
	return resultStatus, nil
}

func BuildDirectUserExpiredTime(ctx *gin.Context, userId int64, goodsEntity *entity.TGoods) (
	directUserEntity *entity.TUser, newExpiredTime, addExpiredTime int64, err error) {
	var (
		userTeam               *entity.TUserTeam
		giftDurationPercentage int64
	)
	// get user team
	err = dao.TUserTeam.Ctx(ctx).Where(do.TUserTeam{UserId: userId}).Scan(&userTeam)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf(`query user team failed`)
		return
	}
	if userTeam == nil {
		global.MyLogger(ctx).Info().Msgf("user team is nil")
		return
	}
	if userTeam.DirectId <= 0 {
		global.MyLogger(ctx).Info().Msgf("direct user is not exist")
		return
	}

	// get direct user
	err = dao.TUser.Ctx(ctx).Where(do.TUser{Id: userTeam.DirectId}).Scan(&directUserEntity)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("query direct user info failed")
		return
	}
	if directUserEntity == nil {
		err = fmt.Errorf(`direct user not exist, userId: %d, DirectId: %d`, userId, userTeam.DirectId)
		global.MyLogger(ctx).Err(err).Msgf(`direct user not exist`)
		return
	}

	// calc direct gift time
	giftDurationPercentage = int64(global.Config.PayConfig.GiftDurationPercentage)
	if giftDurationPercentage <= 0 && giftDurationPercentage > 100 {
		giftDurationPercentage = 20 // 默认20%
	}

	addExpiredTime = int64(goodsEntity.Period) * constant.DaySeconds * giftDurationPercentage / 100
	if IsVIPExpired(directUserEntity) {
		newExpiredTime = time.Now().Unix() + addExpiredTime
	} else {
		newExpiredTime = directUserEntity.ExpiredTime + addExpiredTime
	}
	return
}

func IsVIPExpired(user *entity.TUser) bool {
	if user.ExpiredTime < time.Now().Unix() {
		return true
	} else {
		return false
	}
}

func randomGiftDay(min, max int) int {
	rand.Seed(uint64(time.Now().UnixNano()))
	randomNumber := rand.Intn(max-min+1) + min
	if randomNumber < min {
		randomNumber = min
	}
	return randomNumber
}

func CheckAmount(ctx *gin.Context, amount, orderAmount string) (bool, error) {
	global.MyLogger(ctx).Info().Msgf("response amount: (%s), order amount: (%s)", amount, orderAmount)
	if amount == orderAmount {
		return true, nil
	}
	amountFloat, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("response amount ParseFloat failed, (%s)", amount)
		return false, err
	}

	orderAmountFloat, err := strconv.ParseFloat(orderAmount, 64)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("order amount ParseFloat failed, (%s)", orderAmount)
		return false, err
	}
	if amountFloat >= orderAmountFloat {
		return true, nil
	}

	err = fmt.Errorf("order amount is not eq, response amount: (%s), order amount: (%s)", amount, orderAmount)
	global.MyLogger(ctx).Err(err).Msgf("amount not eq")
	return false, err
}

func checkAmountFloat64(ctx *gin.Context, amount float64, orderAmount string) (bool, error) {
	global.MyLogger(ctx).Info().Msgf("response amount: (%f), order amount: (%s)", amount, orderAmount)
	orderAmountFloat, err := strconv.ParseFloat(orderAmount, 64)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("order amount ParseFloat failed, (%s)", orderAmount)
		return false, err
	}
	if amount >= orderAmountFloat {
		return true, nil
	}

	err = fmt.Errorf("order amount is not eq, response amount: (%s), order amount: (%s)", amount, orderAmount)
	global.MyLogger(ctx).Err(err).Msgf("amount not eq")
	return false, err
}
