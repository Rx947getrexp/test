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

// EditPromotion 编辑推广渠道信息
func EditPromotion(ctx *gin.Context) {
	// 从请求中解析参数
	param := new(request.EditPromotionRequest)
	if err := ctx.ShouldBind(param); err != nil {
		global.MyLogger(ctx).Err(err).Msg("参数解析失败")
		response.RespFail(ctx, i18n.RetMsgParamParseErr, nil)
		return
	}
	// 获取 ID
	id := param.Id
	if id <= 0 {
		global.MyLogger(ctx).Error().Msgf("无效的参数, id: %d", id)
		response.RespFail(ctx, "无效的参数", nil)
		return
	}
	// 查询数据库，找到要更新的记录
	PromotionChannel := model.TPromotionChannels{Id: param.Id}
	has, err := global.Db.Get(&PromotionChannel)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msg("查询推广渠道失败")
		response.RespFail(ctx, "查询失败", nil)
		return
	}
	if !has {
		response.RespFail(ctx, "渠道不存在", nil)
		return
	}
	if param.PromoterName != "" {
		PromotionChannel.PromoterName = param.PromoterName
	}
	if param.PromotionDomain != "" {
		PromotionChannel.PromotionDomain = param.PromotionDomain
	}
	if param.Channel != "" {
		PromotionChannel.Channel = param.Channel
	}
	// 更新 CreatedAt 时间
	PromotionChannel.CreatedAt = time.Now()

	// 将更新后的记录存回数据库
	affected, err := global.Db.ID(PromotionChannel.Id).Update(&PromotionChannel)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msg("更新推广渠道失败")
		response.RespFail(ctx, "更新失败", nil)
		return
	}

	if affected == 0 {
		response.RespFail(ctx, "更新失败", nil)
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
	response.RespOk(ctx, "更新成功", gin.H{
		"Id":              PromotionChannel.Id,
		"PromoterName":    PromotionChannel.PromoterName,
		"PromotionDomain": PromotionChannel.PromotionDomain,
		"Channel":         PromotionChannel.Channel,
		"CreatedAt":       PromotionChannel.CreatedAt,
		"registerCount":   registerCount,
		"rechargeCount":   rechargeCount,
		"totalAmount":     totalAmount,
	})
}
