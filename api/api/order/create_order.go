package order

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"go-speed/constant"
	"go-speed/dao"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/util/pay/pnsafepay"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go-speed/api/api/common"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/response"
)

const (
	ProductNoVIPMonth        = "vip-month" // vip月卡
	OrderAmountVIPMonth      = "500"       // vip月卡
	CurrencyRUB              = "RUB"       // 俄罗斯卢布
	RussianOnlineBankingCode = "29001"     // 俄罗斯网银
)

const (
	PayName  = "hsfly"
	PayEmail = "2233@gmail.com"
	PayPhone = "18818811881"
)

type CreateOrderReq struct {
	UserId      uint64 `form:"user_id" binding:"required" json:"user_id" dc:"用户ID"`
	ProductNo   string `form:"product_no" binding:"required" json:"product_no" dc:"产品编码"`
	Currency    string `form:"currency" binding:"required" json:"currency" dc:"货币类型"`
	OrderAmount int    `form:"order_amount" binding:"required" json:"order_amount" dc:"订单金额"`
}

type CreateOrderRes struct {
	OrderNo  string `json:"order_no" dc:"订单号"`
	OrderUrl string `json:"order_url" dc:"支付链接"`
	Status   string `json:"status" dc:"订单创建状态" eg:"success,fail"`
}

// CreateOrder 创建订单
func CreateOrder(ctx *gin.Context) {
	var (
		err            error
		req            = new(CreateOrderReq)
		userEntity     *entity.TUser
		affected       int64
		lastInsertId   int64
		orderNo        = generateOrderID()
		payRequest     *pnsafepay.PayRequest
		payResponse    *pnsafepay.PayResponse
		payOrderUpdate do.TPayOrder
		res            CreateOrderRes
	)

	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
		response.RespFail(ctx, i18n.RetMsgParamParseErr, nil)
		return
	}
	global.MyLogger(ctx).Debug().Msgf("request: %+v", *req)

	if req.ProductNo != ProductNoVIPMonth {
		global.MyLogger(ctx).Err(err).Msgf("params 'ProductNo'(%s) invalid", req.ProductNo)
		response.RespFail(ctx, i18n.RetMsgParamInvalid, nil)
		return
	}

	if req.Currency != CurrencyRUB {
		global.MyLogger(ctx).Err(err).Msgf("params 'Currency'(%s) invalid", req.Currency)
		response.RespFail(ctx, i18n.RetMsgParamInvalid, nil)
		return
	}

	if strconv.Itoa(req.OrderAmount) != OrderAmountVIPMonth {
		global.MyLogger(ctx).Err(err).Msgf("params 'OrderAmount'(%d) invalid", req.OrderAmount)
		response.RespFail(ctx, i18n.RetMsgParamInvalid, nil)
		return
	}

	userEntity, err = common.CheckUserByUserId(ctx, req.UserId)
	if err != nil {
		return
	}

	// 创建订单
	lastInsertId, err = dao.TPayOrder.Ctx(ctx).Data(do.TPayOrder{
		UserId:      req.UserId,
		Email:       userEntity.Email,
		OrderNo:     orderNo,
		OrderAmount: strconv.Itoa(req.OrderAmount),
		Currency:    req.Currency,
		PayTypeCode: RussianOnlineBankingCode,
		Status:      constant.ParOrderStatusInit,
		CreatedAt:   gtime.Now(),
		UpdatedAt:   gtime.Now(),
		Version:     1,
	}).InsertAndGetId()
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("insert pay order failed")
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}
	global.MyLogger(ctx).Debug().Msgf("lastInsertId: %d, email: %s, orderNo: %s", lastInsertId, userEntity.Email, orderNo)

	// 发起支付
	payRequest = &pnsafepay.PayRequest{
		MerNo:       global.Config.Pay.MerNo,
		OrderNo:     orderNo,
		OrderAmount: OrderAmountVIPMonth,
		PayName:     PayName,
		PayEmail:    PayEmail,
		PayPhone:    PayPhone,
		Currency:    req.Currency,
		PayTypeCode: RussianOnlineBankingCode,
	}
	payResponse, err = pnsafepay.CreatePayOrder(ctx, payRequest)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("CreatePayOrder failed")
		response.RespFail(ctx, i18n.RetMsgCreatePayOrderFailed, nil)
		return
	}

	// 修改订单支付信息
	if payResponse != nil {
		res.OrderUrl = payResponse.OrderData
		res.Status = payResponse.Status
		payOrderUpdate.ReturnStatus = payResponse.Status
		payOrderUpdate.StatusMes = payResponse.StatusMes
		payOrderUpdate.OrderData = payResponse.OrderData
	}
	payOrderUpdate.UpdatedAt = gtime.Now()
	payOrderUpdate.Version = 2
	affected, err = dao.TPayOrder.Ctx(ctx).Data(payOrderUpdate).Where(do.TPayOrder{
		UserId:  req.UserId,
		OrderNo: orderNo,
		Version: 1,
	}).UpdateAndGetAffected()
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("update pay order failed")
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}
	global.MyLogger(ctx).Debug().Msgf("affected: %d, email: %s, orderNo: %s", affected, userEntity.Email, orderNo)

	// 返回支付单信息
	res.OrderNo = orderNo
	response.RespOk(ctx, i18n.RetMsgSuccess, res)
	return
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
	randomNumber := rand.Intn(1000000) // 生成一个0到1000000之间的随机数

	return fmt.Sprintf("%04d%02d%02d%02d%02d%02d%06d", year, month, day, hour, minute, second, randomNumber)
}
