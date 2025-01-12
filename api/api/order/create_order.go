package order

import (
	"context"
	"fmt"
	"go-speed/api/api/common"
	"go-speed/util/pay/freekassa"
	"go-speed/util/pay/russpay"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"

	"go-speed/constant"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/response"
	"go-speed/util/pay/pnsafepay"
)

const (
	CurrencyRUB              = "RUB"   // 俄罗斯卢布
	CurrencyU                = "USD"   // U支付
	CurrencyBTC              = "BTC"   //bitcoin支付
	CurrencyWMZ              = "WMZ"   // webmoney
	CurrencyUAH              = "UAH"   // UAH
	RussianOnlineBankingCode = "29001" // 俄罗斯网银
)

type CreateOrderReq struct {
	ChannelId   string `form:"channel_id" json:"channel_id" binding:"required" dc:"支付渠道ID"`
	GoodsId     int64  `form:"goods_id" json:"goods_id" binding:"required" dc:"套餐ID"`
	DeviceType  string `form:"device_type" json:"device_type" dc:"客户端设备系统os"`
	RedirectURL string `form:"redirect_url" json:"redirect_url" dc:"支付页面完成以后，跳转到的结果页地址"`
}

type CreateOrderRes struct {
	Status      string  `json:"status" dc:"订单创建状态" eg:"success,fail"`
	OrderNo     string  `json:"order_no" dc:"订单号"`
	Currency    string  `json:"currency" dc:"交易币种, eg: U：usd支付，RUB：卢布"`
	OrderAmount float64 `json:"order_amount" dc:"订单金额，支付渠道为U支付时，订单金额要重新计算"`
	OrderUrl    string  `json:"order_url" dc:"支付平台链接. (u支付和银行卡支付此字段无效)"`
	IsGifted    bool    `json:"is_gifted" dc:"本次是否因为支付渠道关闭而赠送了时长. (u支付和银行卡支付此字段无效)"`
	GiftedDays  int     `json:"gifted_days" dc:"本次是否因为支付渠道关闭而赠送的天数 (u支付和银行卡支付此字段无效)" eg:"success,fail"`
	Purse       string  `json:"purse" dc:"purse"`
	Commission  float64 `json:"commission" dc:"手续费"`
}

// CreateOrder 创建订单
func CreateOrder(ctx *gin.Context) {
	var (
		err            error
		req            = new(CreateOrderReq)
		res            CreateOrderRes
		userEntity     *entity.TUser
		payRequest     *pnsafepay.PayRequest
		payResponse    *pnsafepay.PayResponse
		payOrderUpdate do.TPayOrder
		goodsEntity    *entity.TGoods
		paymentEntity  *entity.TPaymentChannel
		isNeedGift     bool
		affected       int64
		lastInsertId   int64
	)
	// gen order id
	res.OrderNo = generateOrderID()

	// 参数解析
	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
		response.RespFail(ctx, i18n.RetMsgParamParseErr, nil)
		return
	}
	global.MyLogger(ctx).Info().Msgf("request: %+v", *req)

	// validate user
	userEntity, err = common.ValidateClaims(ctx)
	if err != nil {
		return
	}

	// validate goods
	goodsEntity, err = ValidateGoods(ctx, req.GoodsId)
	if err != nil {
		return
	}

	// validate pay channel
	paymentEntity, err = ValidatePayChannel(ctx, req.ChannelId)
	if err != nil {
		return
	}

	// validate unpaid order
	err = ValidateOrderLimit(ctx, userEntity, paymentEntity)
	if err != nil {
		return
	}

	// create order
	//amount, amountString, currency := genPayAmount(goodsEntity.Price, goodsEntity.UsdPayPrice, goodsEntity.WebmoneyPayPrice, req.ChannelId)
	amount, commission, amountString, currency, err := genPayAmount(goodsEntity, paymentEntity)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("genPayAmount failed")
		response.RespFail(ctx, i18n.RetMsgParamParseErr, nil)
		return
	}
	global.MyLogger(ctx).Info().Msgf(
		"channel: %s, amount: %f, commission: %f, amountString: %s, currency: %s",
		req.ChannelId, amount, commission, amountString, currency,
	)

	lastInsertId, err = dao.TPayOrder.Ctx(ctx).Data(do.TPayOrder{
		UserId:           userEntity.Id,
		Email:            userEntity.Email,
		OrderNo:          res.OrderNo,
		PaymentChannelId: req.ChannelId,
		GoodsId:          req.GoodsId,
		OrderAmount:      amountString,
		Currency:         currency,
		PayTypeCode:      paymentEntity.PayTypeCode,
		Status:           constant.ParOrderStatusInit,
		CreatedAt:        gtime.Now(),
		UpdatedAt:        gtime.Now(),
		Version:          constant.VersionInit,
		Commission:       commission,
		DeviceType:       req.DeviceType,
	}).InsertAndGetId()
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("insert pay order failed")
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}
	global.MyLogger(ctx).Info().Msgf("lastInsertId: %d, email: %s, orderNo: %s", lastInsertId, userEntity.Email, res.OrderNo)

	// 发起支付
	switch req.ChannelId {
	case constant.PayChannelPnSafePay: // 用户选择的是 pnsafepay 支付渠道
		payRequest = &pnsafepay.PayRequest{
			MerNo:       paymentEntity.MerNo,
			OrderNo:     res.OrderNo,
			OrderAmount: amountString,
			PayName:     global.Config.PNSafePay.PayName,
			PayEmail:    global.Config.PNSafePay.PayEmail,
			PayPhone:    global.Config.PNSafePay.PayPhone,
			Currency:    CurrencyRUB,
			PayTypeCode: paymentEntity.PayTypeCode,
		}
		payResponse, err = pnsafepay.CreatePayOrder(ctx, payRequest)
		if err != nil && (payResponse == nil) {
			global.MyLogger(ctx).Err(err).Msgf("CreatePayOrder failed")
			response.RespFail(ctx, i18n.RetMsgCreatePayOrderFailed, nil)
			return
		}

		// 获取支付渠道返回的结果
		isNeedGift = payResponse.InterIsNeedGift         // 是否为渠道关闭需要赠送时长
		res.OrderUrl = payResponse.OrderData             // 返回支付链接
		res.Status = payResponse.Status                  // 返回支付渠道支付单创建的状态
		payOrderUpdate.ReturnStatus = payResponse.Status // 记录支付渠道状态
		payOrderUpdate.StatusMes = payResponse.StatusMes // 记录支付渠道状态描述信息
		payOrderUpdate.OrderData = payResponse.OrderData // 记录支付链接
		if isNeedGift {
			payOrderUpdate.PaymentChannelErr = constant.PaymentChannelErrYes
		}

	case constant.PayChannelUPay:
		// 用户选择的是 usd pay 支付渠道
		// U支付只需要返回支付连接给前端
		res.Status = constant.ReturnStatusSuccess
		payOrderUpdate.ReturnStatus = constant.ReturnStatusSuccess
		payOrderUpdate.StatusMes = "U-Pay directly returns the payment code"
	case constant.PayChannelBTCPay:
		// 用户选择的是 btc pay 支付渠道
		// BTC支付只需要返回支付连接给前端
		res.Status = constant.ReturnStatusSuccess
		payOrderUpdate.ReturnStatus = constant.ReturnStatusSuccess
		payOrderUpdate.StatusMes = "BTC-Pay directly returns the payment code"
	case constant.PayChannelBankCardPay:
		// 用户选择银行卡支付
		// 银行卡支付只需要返回银行卡给前端
		res.Status = constant.ReturnStatusSuccess
		payOrderUpdate.ReturnStatus = constant.ReturnStatusSuccess
		payOrderUpdate.StatusMes = "BankCard-Pay directly returns the bank card number"

	case constant.PayChannelWebMoneyPay:
		res.Status = constant.ReturnStatusSuccess
		payOrderUpdate.ReturnStatus = constant.ReturnStatusSuccess
		payOrderUpdate.StatusMes = "Webmoney"

	case constant.PayChannelFreekassa_12, constant.PayChannelFreekassa_36,
		constant.PayChannelFreekassa_43, constant.PayChannelFreekassa_44,
		constant.PayChannelFreekassa_7:
		resp, err := freekassa.CreateOrder(ctx, freekassa.CreateOrderReq{
			PaymentId: res.OrderNo,
			I:         paymentEntity.FreekassaCode,
			Email:     userEntity.Email,
			Ip:        ctx.ClientIP(),
			Amount:    amount,
			Currency:  currency,
		})
		if err != nil && (resp == nil) {
			global.MyLogger(ctx).Err(err).Msgf("freekassa CreatePayOrder failed")
			response.RespFail(ctx, i18n.RetMsgCreatePayOrderFailed, nil)
			return
		}
		res.Status = resp.Type
		res.OrderUrl = resp.Location
		payOrderUpdate.ReturnStatus = resp.Type
		payOrderUpdate.StatusMes = resp.Location
		payOrderUpdate.OrderData = resp.OrderId

	case constant.PayChannelApplePay:
		res.Status = constant.ReturnStatusSuccess
		payOrderUpdate.ReturnStatus = constant.ReturnStatusSuccess
		payOrderUpdate.StatusMes = "apple-pay"

	case constant.PayChannelRussPayBankCard, constant.PayChannelRussPaySBP, constant.PayChannelRussPaySBER:
		resp, err := russpay.CreateOrder(ctx, russpay.CreateOrderReq{
			ChannelId:   req.ChannelId,
			Amount:      amountString,
			OrderNumber: res.OrderNo,
			CompanyPage: req.RedirectURL,
			DeviceType:  req.DeviceType,
		})
		if err != nil && resp == nil {
			global.MyLogger(ctx).Err(err).Msgf("russpay CreateOrder failed")
			response.RespFail(ctx, i18n.RetMsgCreatePayOrderFailed, nil)
			return
		}
		res.Status = constant.ReturnStatusSuccess
		res.OrderUrl = resp.PayUrl
		payOrderUpdate.ReturnStatus = constant.ReturnStatusSuccess
		payOrderUpdate.StatusMes = resp.PayUrl
		payOrderUpdate.OrderData = resp.BillingNumber

	default:
		err = fmt.Errorf("ChannelId %s 无效", req.ChannelId)
		global.MyLogger(ctx).Err(err).Msgf("ChannelId not exist")
		response.RespFail(ctx, i18n.RetMsgParamInvalid, nil)
		return
	}

	// 支付结果记录到DB
	payOrderUpdate.UpdatedAt = gtime.Now()
	payOrderUpdate.Version = constant.VersionInit + 1
	affected, err = dao.TPayOrder.Ctx(ctx).Data(payOrderUpdate).Where(do.TPayOrder{
		UserId:  userEntity.Id,
		OrderNo: res.OrderNo,
		Version: constant.VersionInit,
	}).UpdateAndGetAffected()
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("update pay order failed")
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}
	global.MyLogger(ctx).Info().Msgf("affected: %d, email: %s, orderNo: %s", affected, userEntity.Email, res.OrderNo)

	// 赠送逻辑
	if isNeedGift && paymentEntity.FreeTrialDays > 0 {
		gifted, _err := GiftDurationForPaymentChannelClosed(ctx, res.OrderNo, userEntity, paymentEntity)
		if _err != nil {
			global.MyLogger(ctx).Err(_err).Msgf("赠送失败")
			response.RespFail(ctx, i18n.RetMsgCreatePayOrderFailed, nil)
			return
		}
		if gifted {
			// 赠送成功，返回描述信息
			res.IsGifted = true
			res.GiftedDays = paymentEntity.FreeTrialDays
		}
	}
	res.OrderAmount, res.Commission, res.Currency, res.Purse = amount, commission, currency, paymentEntity.PayTypeCode
	global.MyLogger(ctx).Info().Msgf("res: %+v", res)
	response.RespOk(ctx, i18n.RetMsgSuccess, res)
	return
}

func ValidateOrderLimit(ctx *gin.Context, user *entity.TUser, channel *entity.TPaymentChannel) (err error) {
	if channel.ChannelId == constant.PayChannelWebMoneyPay {
		return nil
	}
	// query unpaid order
	var orders []entity.TPayOrder
	err = dao.TPayOrder.Ctx(ctx).
		Where(do.TPayOrder{UserId: user.Id, PaymentChannelId: channel.ChannelId}).
		WhereGTE(dao.TPayOrder.Columns().CreatedAt, getNDurationAgoTime(time.Minute*time.Duration(channel.TimeoutDuration))).
		WhereNotIn(dao.TPayOrder.Columns().Status, []string{constant.ParOrderStatusPaid}).
		Order(dao.TPayOrder.Columns().Id, constant.OrderTypeDesc).Scan(&orders)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("query order failed")
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}

	var unpaidNum, closedNum, failedNum int
	for _, order := range orders {
		switch order.Status {
		case constant.ParOrderStatusInit, constant.ParOrderStatusUnpaid:
			unpaidNum++

		case constant.ParOrderStatusClosed, constant.ParOrderStatusTimeout:
			closedNum++

		case constant.ParOrderStatusPaidFailed:
			failedNum++
		}
	}

	if unpaidNum > global.Config.PayConfig.OrderUnpaidLimitNum {
		err = fmt.Errorf(i18n.RetMsgOrderUnpaidLimit)
		global.MyLogger(ctx).Err(err).Msgf("unpaidNum: %d", unpaidNum)
		response.RespFail(ctx, i18n.RetMsgOrderUnpaidLimit, nil)
		return
	}

	if closedNum > global.Config.PayConfig.OrderClosedLimitNum {
		err = fmt.Errorf(i18n.RetMsgOrderClosedLimit)
		global.MyLogger(ctx).Err(err).Msgf("closedNum: %d", closedNum)
		response.RespFail(ctx, i18n.RetMsgOrderClosedLimit, nil)
		return
	}

	if failedNum > global.Config.PayConfig.OrderFailedLimitNum {
		err = fmt.Errorf(i18n.RetMsgOrderFailedLimit)
		global.MyLogger(ctx).Err(err).Msgf("failedNum: %d", failedNum)
		response.RespFail(ctx, i18n.RetMsgOrderFailedLimit, nil)
		return
	}
	return nil
}

// generateOrderID 生成一个订单号
func generateOrderID() string {
	now := time.Now()
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(999) // 生成一个0到1000000之间的随机数
	return fmt.Sprintf("%02d%02d%02d%02d%02d%02d%03d", year-2014, month, day, hour, minute, second, randomNumber)
}

// 系统赠送时长
func GiftDurationForPaymentChannelClosed(ctx *gin.Context, orderNo string,
	userEntity *entity.TUser, paymentEntity *entity.TPaymentChannel) (gifted bool, err error) {
	// 需要全部失败
	// 查找全部的支付渠道
	var (
		paymentChannels []entity.TPaymentChannel
		payOrders       []entity.TPayOrder
	)
	err = dao.TPaymentChannel.Ctx(ctx).
		Where(do.TPaymentChannel{IsActive: constant.PaymentChannelIsActiveYes}).
		Scan(&paymentChannels)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf(`query payment channels failed`)
		return
	}

	err = dao.TPayOrder.Ctx(ctx).
		Where(do.TPayOrder{UserId: userEntity.Id}).
		WhereGTE(dao.TPayOrder.Columns().CreatedAt, getNDaysAgoTime(constant.PayChannelErrTimeWindow)).
		Order(dao.TPayOrder.Columns().Id, constant.OrderTypeDesc).
		Scan(&payOrders)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf(`query pay order failed`)
		return
	}

	for _, payChannel := range paymentChannels {
		// U支付、BTC支付、银行卡支付不参与赠送逻辑
		if payChannel.ChannelId == constant.PayChannelUPay || payChannel.ChannelId == constant.PayChannelBTCPay || payChannel.ChannelId == constant.PayChannelBankCardPay {
			continue
		}

		var errFlag bool
		for _, payOrder := range payOrders {
			if payOrder.PaymentChannelId != payChannel.ChannelId {
				continue
			}

			if payOrder.PaymentChannelErr == constant.PaymentChannelErrNo {
				global.MyLogger(ctx).Info().Msgf("最近订单(%s)，没有支付通道异常的标记，不满足赠送的条件, email: %s, channel: %s",
					payOrder.OrderNo, userEntity.Email, payOrder.PaymentChannelId)
				return
			}

			// 最近一条记录标记支付渠道异常
			errFlag = true
			break
		}
		if !errFlag {
			global.MyLogger(ctx).Info().Msgf("最近通道(%s), 没有支付通道异常的标记，不满足赠送的条件, email: %s",
				payChannel.ChannelId, userEntity.Email)
			return
		}
	}
	// 通过上面的检查，支付通道全部失败

	// 如果过期直接赠送
	// 如果未过期就检查最近是否已经赠送过了
	if !isVIPExpired(userEntity) {
		// 检查在可以赠送的时间窗口内，是否已经赠送过了
		var items []entity.TUserVipAttrRecord
		err = dao.TUserVipAttrRecord.Ctx(ctx).
			Where(do.TUserVipAttrRecord{Email: userEntity.Email, Source: constant.UserVipAttrOpSourcePaymentChannelClosedGift}).
			WhereGTE(dao.TUserVipAttrRecord.Columns().CreatedAt, getNDaysAgoTime(paymentEntity.FreeTrialDays)).
			Scan(&items)
		if len(items) > 0 {
			global.MyLogger(ctx).Info().Msgf("最近已经赠送过了, email: %s, items: %+v", userEntity.Email, items)
			return
		}
	}

	// 检查通过，开始赠送
	var (
		newExpiredTime int64
		addExpiredTime = int64(paymentEntity.FreeTrialDays) * constant.DaySeconds
	)
	if isVIPExpired(userEntity) {
		newExpiredTime = time.Now().Unix() + addExpiredTime
	} else {
		newExpiredTime = userEntity.ExpiredTime + addExpiredTime
	}

	_ctx := ctx
	err = dao.TPayOrder.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		userUpdate := do.TUser{
			ExpiredTime: newExpiredTime,
			UpdatedAt:   gtime.Now(),
			Version:     userEntity.Version + 1,
			Kicked:      0,
		}
		var (
			affected     int64
			lastInsertId int64
		)
		affected, err = dao.TUser.Ctx(ctx).Data(userUpdate).Where(do.TUser{
			Id:      userEntity.Id,
			Version: userEntity.Version,
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

		global.MyLogger(_ctx).Info().Msgf("add(%d) user(%s) ExpiredTime from(%d) to(%d)",
			addExpiredTime, userEntity.Email, userEntity.ExpiredTime, newExpiredTime)

		// 记录操作流水
		lastInsertId, err = dao.TUserVipAttrRecord.Ctx(ctx).Data(do.TUserVipAttrRecord{
			Email:           userEntity.Email,
			Source:          constant.UserVipAttrOpSourcePaymentChannelClosedGift,
			OrderNo:         orderNo,
			ExpiredTimeFrom: userEntity.ExpiredTime,
			ExpiredTimeTo:   newExpiredTime,
			Desc:            fmt.Sprintf("ExpiredTime add giftDay(%d)", paymentEntity.FreeTrialDays),
			CreatedAt:       gtime.Now(),
		}).InsertAndGetId()
		if err != nil {
			global.MyLogger(_ctx).Err(err).Msgf(`insert TUserVipAttrRecords failed`)
			return err
		}
		global.MyLogger(_ctx).Info().Msgf(">>>> insert TUserVipAttrRecords, lastInsertId: %d", lastInsertId)
		return nil
	})
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("sync order status failed")
		return
	}
	return true, nil
}

func isVIPExpired(user *entity.TUser) bool {
	if user.ExpiredTime < time.Now().Unix() {
		return true
	} else {
		return false
	}
}

func getNDaysAgoTime(n int) time.Time {
	return time.Now().AddDate(0, 0, -1*n)
}

func getNDurationAgoTime(n time.Duration) time.Time {
	return time.Now().Add(-1 * n)
}

func genUPayAmountDecimalPartValue() int {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(9999) // 生成一个0到9999之间的随机数
	return randomNumber % 10000
}
func genBTCAmountDecimalPartValue(exchangeRateBtc float64) float64 {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Float64() * exchangeRateBtc // 随机取0~最大值的随机数
	return randomNumber
}

func genUPayAmount2DecimalPartValue() int {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(99) // 生成一个0到9999之间的随机数
	return randomNumber % 100
}

//
//func genPayAmount(price, priceUSD, webMoneyPrice float64, channelId string) (amount float64, amountString string, currency string) {
//	switch channelId {
//	case constant.PayChannelUPay:
//		if priceUSD > 1.0 {
//			priceUSD = priceUSD - 1
//		}
//		amountInt := int(priceUSD*10000) + genUPayAmountDecimalPartValue()
//		amountString = fmt.Sprintf("%d.%d", amountInt/10000, amountInt%10000)
//		amount, _ = strconv.ParseFloat(amountString, 64)
//		return amount, amountString, CurrencyU
//	case constant.PayChannelBankCardPay:
//		amountInt := int((price-1)*100) + genUPayAmount2DecimalPartValue()
//		amountString = fmt.Sprintf("%d.%d", amountInt/100, amountInt%100)
//		amount, _ = strconv.ParseFloat(amountString, 64)
//		return amount, amountString, CurrencyRUB
//	case constant.PayChannelWebMoneyPay:
//		return webMoneyPrice, fmt.Sprintf("%.2f", webMoneyPrice), CurrencyWMZ
//	default:
//		return price, fmt.Sprintf("%.2f", price), CurrencyRUB
//	}
//}

/*
update t_payment_channel currency_type = 'RUB' where channel_id in ('bankcard', 'pnsafepay', 'freekassa-12', 'freekassa-36', 'freekassa-43', 'freekassa-44');
update t_payment_channel currency_type = 'WMZ' where channel_id in ('webmoney');
update t_payment_channel currency_type = 'USD' where channel_id in ('usd');
update t_payment_channel currency_type = 'UAH' where channel_id in ('freekassa-7');
*/
func genPayAmount(goodsEntity *entity.TGoods, paymentEntity *entity.TPaymentChannel) (
	amount, commission float64, amountString string, currency string, err error) {
	switch paymentEntity.ChannelId {
	case constant.PayChannelUPay:
		priceUSD := goodsEntity.PriceUsd
		if priceUSD > 1.0 {
			priceUSD = priceUSD - 1
		}
		amountInt := int(priceUSD*10000) + genUPayAmountDecimalPartValue()
		amountString = fmt.Sprintf("%d.%d", amountInt/10000, amountInt%10000)
		amount, _ = strconv.ParseFloat(amountString, 64)
		return amount, paymentEntity.Commission, amountString, CurrencyU, nil
	case constant.PayChannelBTCPay:
		const exchangeRateBtc = 0.00001688
		priceBTC := goodsEntity.PriceBtc
		if priceBTC > exchangeRateBtc {
			priceBTC = priceBTC - exchangeRateBtc
		}
		finalAmount := priceBTC + genBTCAmountDecimalPartValue(exchangeRateBtc)
		amountInt := int(finalAmount * 100000000)
		amountString = fmt.Sprintf("%d.%08d", amountInt/100000000, amountInt%100000000)
		amount, _ = strconv.ParseFloat(amountString, 64)
		return amount, paymentEntity.Commission, amountString, CurrencyBTC, nil
	case constant.PayChannelBankCardPay:
		priceRUB := goodsEntity.PriceRub
		if priceRUB > 1.0 {
			priceRUB = priceRUB - 1
		}
		amountInt := int((priceRUB-1)*100) + genUPayAmount2DecimalPartValue()
		amountString = fmt.Sprintf("%d.%d", amountInt/100, amountInt%100)
		amount, _ = strconv.ParseFloat(amountString, 64)
		return amount, paymentEntity.Commission, amountString, CurrencyRUB, nil

	case constant.PayChannelPnSafePay,
		constant.PayChannelFreekassa_12, constant.PayChannelFreekassa_36,
		constant.PayChannelFreekassa_43, constant.PayChannelFreekassa_44,
		constant.PayChannelApplePay,
		constant.PayChannelRussPayBankCard, constant.PayChannelRussPaySBP, constant.PayChannelRussPaySBER:
		return goodsEntity.PriceRub, paymentEntity.Commission, fmt.Sprintf("%.2f", goodsEntity.PriceRub), CurrencyRUB, nil

	case constant.PayChannelWebMoneyPay:
		return goodsEntity.PriceWmz, paymentEntity.Commission, fmt.Sprintf("%.2f", goodsEntity.PriceWmz), CurrencyWMZ, nil

	case constant.PayChannelFreekassa_7:
		return goodsEntity.PriceUah, paymentEntity.Commission, fmt.Sprintf("%.2f", goodsEntity.PriceUah), CurrencyUAH, nil

	default:
		return 0.0, 0.0, "", "", fmt.Errorf("invalid `PayChannelID`")
	}
}

func ValidateGoods(ctx *gin.Context, goodsId int64) (gs *entity.TGoods, err error) {
	err = dao.TGoods.Ctx(ctx).Where(do.TGoods{Id: goodsId}).Scan(&gs)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("query goods failed")
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}
	if gs == nil {
		err = fmt.Errorf("goodsId %d 无效", goodsId)
		global.MyLogger(ctx).Err(err).Msgf("goods not exist")
		response.RespFail(ctx, i18n.RetMsgParamInvalid, nil)
		return
	}
	return
}

func ValidatePayChannel(ctx *gin.Context, channelId string) (pc *entity.TPaymentChannel, err error) {
	err = dao.TPaymentChannel.Ctx(ctx).Where(
		do.TPaymentChannel{
			ChannelId: channelId,
			IsActive:  constant.PaymentChannelIsActiveYes,
		}).Scan(&pc)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("query TPaymentChannels failed")
		response.ResFail(ctx, err.Error())
		return
	}
	if pc == nil {
		err = fmt.Errorf("channelId %s 无效", channelId)
		global.MyLogger(ctx).Err(err).Msgf("channelId not exist")
		response.RespFail(ctx, i18n.RetMsgParamInvalid, nil)
		return
	}
	if pc.ChannelId == constant.PayChannelWebMoneyPay {
		pc.PayTypeCode = global.Config.WebMoneyConfig.Purse
	}
	return
}
