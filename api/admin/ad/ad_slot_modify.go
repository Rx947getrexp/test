package ad

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/os/gtime"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/response"
)

type ADSlotModifyReq struct {
	Location string `form:"location" binding:"required" json:"location" dc:"广告位的位置，相当于ID"`
	Name     string `form:"name" binding:"required" json:"name" dc:"广告位名称"`
	Desc     string `form:"desc" binding:"required" json:"desc" dc:"广告位描述"`
	Status   int    `form:"status" json:"status" dc:"状态:1-上架；2-下架（默认值）"`
}

type ADSlotModifyRes struct{}

// ADSlotModify 修改广告位
func ADSlotModify(ctx *gin.Context) {
	var (
		err    error
		req    = new(ADSlotModifyReq)
		adSlot *entity.TAdSlot
	)
	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
		response.ResFail(ctx, err.Error())
		return
	}
	global.MyLogger(ctx).Info().Msgf("req: %+v", *req)

	if req.Status != 0 && req.Status != ADSlotStatusOnline && req.Status != ADSlotStatusOffline {
		response.ResFail(ctx, "参数 Status 设置值无效")
		return
	}

	err = dao.TAdSlot.Ctx(ctx).Where(do.TAdSlot{Location: req.Location}).Scan(&adSlot)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("query ad slot failed")
		response.ResFail(ctx, err.Error())
		return
	}
	if adSlot == nil {
		response.ResFail(ctx, fmt.Sprintf("%s 广告位不存在", req.Location))
		return
	}

	adSlotUpdate := do.TAdSlot{UpdatedAt: gtime.Now()}
	if req.Status > 0 {
		adSlotUpdate.Status = req.Status
	}
	if req.Desc != "" {
		adSlotUpdate.Desc = req.Desc
	}
	if req.Name != "" {
		adSlotUpdate.Name = req.Name
	}

	affected, err := dao.TAdSlot.Ctx(ctx).Data(adSlotUpdate).Where(do.TAdSlot{Location: req.Location}).UpdateAndGetAffected()
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("modify ad slot failed")
		response.ResFail(ctx, err.Error())
		return
	}
	global.MyLogger(ctx).Debug().Msgf("affected: %d", affected)
	response.ResOk(ctx, "成功")
}
