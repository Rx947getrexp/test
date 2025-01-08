package ad

import (
	"github.com/gin-gonic/gin"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/response"
)

type ADSlotDeleteReq struct {
	Location string `form:"location" binding:"required" json:"location" dc:"广告位的位置，相当于ID"`
}

type ADSlotDeleteRes struct{}

// ADSlotDelete 删除广告位
func ADSlotDelete(ctx *gin.Context) {
	var (
		err    error
		req    = new(ADSlotDeleteReq)
		adSlot *entity.TAdSlot
	)
	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
		response.ResFail(ctx, err.Error())
		return
	}
	global.MyLogger(ctx).Info().Msgf("req: %+v", *req)

	err = dao.TAdSlot.Ctx(ctx).Where(do.TAdSlot{Location: req.Location}).Scan(&adSlot)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("query ad slot failed")
		response.ResFail(ctx, err.Error())
		return
	}
	if adSlot == nil {
		response.ResOk(ctx, "成功")
		return
	}

	_, err = dao.TAdSlot.Ctx(ctx).Delete(do.TAdSlot{Id: adSlot.Id, Location: req.Location})
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("delete TAdSlot failed")
		response.RespFail(ctx, err.Error(), nil)
		return
	}

	response.ResOk(ctx, "成功")
}
