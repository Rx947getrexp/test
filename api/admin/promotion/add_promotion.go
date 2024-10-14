package promotion

import (
	"go-speed/constant"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model"
	"go-speed/model/request"
	"go-speed/model/response"
	"time"

	"github.com/gin-gonic/gin"
)

// AddPromotion 添加渠道代理人域名
func AddPromotion(ctx *gin.Context) {
	param := new(request.AddPromotionRequest)
	if err := ctx.ShouldBind(param); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("参数解析失败")
		response.RespFail(ctx, i18n.RetMsgParamParseErr, nil)
		return
	}

	// 创建新的推广渠道记录
	NewPromotionChannel := model.TPromotionChannels{
		PromoterName:    param.PromoterName,
		PromotionDomain: param.PromotionDomain,
		Channel:         param.Channel,
		CreatedAt:       time.Now(),
	}

	// 插入数据到数据库
	_, err := global.Db.Insert(&NewPromotionChannel)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msg("插入推广渠道失败")
		response.RespFail(ctx, "添加失败", nil)
		return
	}
	// 统计注册人数
	var registerCount int64
	registerQuery := "SELECT COUNT(*) FROM t_user WHERE channel = ?"
	err = global.Db.DB().QueryRow(registerQuery, param.Channel).Scan(&registerCount)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msg("查询注册人数失败")
		response.RespFail(ctx, "查询失败", nil)
		return
	}
	// 统计充值人数
	var rechargeCount int64
	rechargeQuery := "SELECT COUNT(*) FROM t_pay_order INNER JOIN t_user ON t_pay_order.email = t_user.email WHERE t_user.channel = ? AND t_pay_order.status = ?"
	err = global.Db.DB().QueryRow(rechargeQuery, param.Channel, "paid").Scan(&rechargeCount)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msg("查询充值人数失败")
		response.RespFail(ctx, "查询失败", nil)
		return
	}
	// 统计充值金额
	var totalAmount float64
	amountQuery := "SELECT currency, order_reality_amount FROM t_pay_order INNER JOIN t_user ON t_pay_order.email = t_user.email WHERE t_user.channel = ? AND t_pay_order.status = ?"
	rows, err := global.Db.DB().Query(amountQuery, param.Channel, "paid")
	if err != nil {
		global.MyLogger(ctx).Err(err).Msg("查询充值金额数据失败")
		response.RespFail(ctx, "查询失败", nil)
		return
	}
	defer rows.Close()
	// 循环处理数据并累加
	for rows.Next() {
		var currency string
		var orderRealityAmount float64
		if err := rows.Scan(&currency, &orderRealityAmount); err != nil {
			global.MyLogger(ctx).Err(err).Msg("扫描数据失败")
			response.RespFail(ctx, "查询失败", nil)
			return
		}
		switch currency {
		case "USD":
			totalAmount += orderRealityAmount * constant.ExchangeRateUSD
		case "WMZ":
			totalAmount += orderRealityAmount * constant.ExchangeRateWMZ
		default:
			totalAmount += orderRealityAmount
		}
	}
	// 检查循环过程中是否出错
	if err := rows.Err(); err != nil {
		global.MyLogger(ctx).Err(err).Msg("迭代数据时发生错误")
		response.RespFail(ctx, "查询失败", nil)
		return
	}
	// 返回新增的渠道及统计信息
	response.RespOk(ctx, "添加成功", gin.H{
		"Id":              NewPromotionChannel.Id,
		"PromoterName":    NewPromotionChannel.PromoterName,
		"PromotionDomain": NewPromotionChannel.PromotionDomain,
		"Channel":         NewPromotionChannel.Channel,
		"CreatedAt":       NewPromotionChannel.CreatedAt,
		"registerCount":   registerCount,
		"rechargeCount":   rechargeCount,
		"totalAmount":     totalAmount,
	})
}
