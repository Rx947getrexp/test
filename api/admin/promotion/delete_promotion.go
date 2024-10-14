package promotion

import (
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model"
	"go-speed/model/request"
	"go-speed/model/response"

	"github.com/gin-gonic/gin"
)

// DeletePromotion 删除推广渠道
func DeletePromotion(ctx *gin.Context) {
	param := new(request.DeletePromotionIdRequest)
	if err := ctx.ShouldBind(param); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("参数解析失败")
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

	// 检查记录是否存在
	var channel model.TPromotionChannels
	has, err := global.Db.Where("id = ?", id).Get(&channel)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("查询推广渠道失败, id: %d", id)
		response.RespFail(ctx, "查询失败", nil)
		return
	}
	if !has {
		global.MyLogger(ctx).Error().Msgf("推广渠道不存在, id: %d", id)
		response.RespFail(ctx, "推广渠道不存在", nil)
		return
	}

	// 执行删除操作
	_, err = global.Db.Exec("DELETE FROM t_promotion_channels WHERE id = ?", id)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("删除推广渠道失败, id: %d", id)
		response.RespFail(ctx, "删除失败", nil)
		return
	}
	response.RespOk(ctx, "删除成功", nil)
}
