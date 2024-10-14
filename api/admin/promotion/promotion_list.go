package promotion

import (
	"go-speed/constant"
	"go-speed/global"
	"go-speed/model"
	"go-speed/model/response"

	"github.com/gin-gonic/gin"
)

type GetPromotionListRequest struct {
	Id        int    `form:"id" json:"id" dc:"唯一标识符"`
	OrderType string `form:"order_type" json:"order_type" dc:"排序类型，eg: asc|desc"`
	Page      int    `form:"page" json:"page" dc:"分页查询page, 从1开始"`
	Size      int    `form:"size" json:"size" dc:"分页查询size, 最大1000"`
}

// PromotionList 获取推广渠道列表
func PromotionList(ctx *gin.Context) {
	param := new(GetPromotionListRequest)
	if err := ctx.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(ctx, "参数错误")
		return
	}
	var promotionChannels []model.TPromotionChannels
	query := global.Db.Where("1=1")
	if param.OrderType != "" {
		query = query.OrderBy(param.OrderType)
	}
	if param.Page > 0 && param.Size > 0 {
		if param.Size > 1000 {
			param.Size = 1000
		}
		offset := (param.Page - 1) * param.Size
		err := query.Limit(param.Size, offset).Find(&promotionChannels)
		if err != nil {
			global.MyLogger(ctx).Err(err).Msg("查询推广渠道列表失败")
			response.RespFail(ctx, "查询失败", nil)
			return
		}
	} else {
		err := query.Find(&promotionChannels)
		if err != nil {
			global.MyLogger(ctx).Err(err).Msg("查询推广渠道列表失败")
			response.RespFail(ctx, "查询失败", nil)
			return
		}
	}
	// 构建返回数据
	var data []map[string]interface{}
	for _, params := range promotionChannels {
		// 统计注册人数
		var registerCount int64
		registerQuery := "SELECT COUNT(*) FROM t_user WHERE channel = ?"
		err := global.Db.DB().QueryRow(registerQuery, params.Channel).Scan(&registerCount)
		if err != nil {
			global.MyLogger(ctx).Err(err).Msg("查询注册人数失败")
			response.RespFail(ctx, "查询失败", nil)
			return
		}

		// 统计充值人数
		var rechargeCount int64
		rechargeQuery := "SELECT COUNT(*) FROM t_pay_order INNER JOIN t_user ON t_pay_order.email = t_user.email WHERE t_user.channel = ? AND t_pay_order.status = ?"
		err = global.Db.DB().QueryRow(rechargeQuery, params.Channel, "paid").Scan(&rechargeCount)
		if err != nil {
			global.MyLogger(ctx).Err(err).Msg("查询充值人数失败")
			response.RespFail(ctx, "查询失败", nil)
			return
		}

		// 统计充值金额
		var totalAmount float64
		amountQuery := "SELECT currency, order_reality_amount FROM t_pay_order INNER JOIN t_user ON t_pay_order.email = t_user.email WHERE t_user.channel = ? AND t_pay_order.status = ?"
		rows, err := global.Db.DB().Query(amountQuery, params.Channel, "paid")
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

		item := map[string]interface{}{
			"Id":              params.Id,
			"PromoterName":    params.PromoterName,
			"PromotionDomain": params.PromotionDomain,
			"Channel":         params.Channel,
			"CreatedAt":       params.CreatedAt,
			"registerCount":   registerCount,
			"rechargeCount":   rechargeCount,
			"totalAmount":     totalAmount,
		}
		data = append(data, item)
	}

	resp := map[string]interface{}{
		"total": len(data), // 数据总条数
		"items": data,      // 数据明细
	}
	// 返回推广渠道列表
	response.RespOk(ctx, "查询成功", resp)
}
