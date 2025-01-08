package ad

import (
	"github.com/gin-gonic/gin"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/response"
)

type ADDeleteReq struct {
	Name string `form:"name" binding:"required" json:"name" dc:"广告名称-要求唯一"`
}

type ADDeleteRes struct{}

// ADDelete 删除广告
func ADDelete(ctx *gin.Context) {
	var (
		err error
		req = new(ADDeleteReq)
		ad  *entity.TAd
	)
	if err = ctx.ShouldBind(req); err != nil {
		response.ResFail(ctx, err.Error())
		return
	}
	global.MyLogger(ctx).Info().Msgf("req: %+v", *req)

	err = dao.TAd.Ctx(ctx).Where(do.TAd{Name: req.Name}).Scan(&ad)
	if err != nil {
		response.ResFail(ctx, err.Error())
		return
	}
	if ad == nil {
		response.ResOk(ctx, "成功")
		return
	}
	_, err = dao.TAd.Ctx(ctx).Delete(do.TAd{Id: ad.Id, Name: req.Name})
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("delete TAd failed")
		response.RespFail(ctx, err.Error(), nil)
		return
	}

	response.ResOk(ctx, "成功")
}
