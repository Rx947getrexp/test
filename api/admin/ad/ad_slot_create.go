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

const (
	ADSlotStatusOnline  = 1
	ADSlotStatusOffline = 2
)

type ADSlotCreateReq struct {
	Location string `form:"location" binding:"required" json:"location" dc:"广告位的位置，相当于ID"`
	Name     string `form:"name" binding:"required" json:"name" dc:"广告位名称"`
	Desc     string `form:"desc" binding:"required" json:"desc" dc:"广告位描述"`
	Status   int    `form:"status" json:"status" dc:"状态:1-上架；2-下架（默认值）"`
}

type ADSlotCreateRes struct{}

// ADSlotCreate 添加广告位
func ADSlotCreate(ctx *gin.Context) {
	var (
		err          error
		req          = new(ADSlotCreateReq)
		adSlot       *entity.TAdSlot
		lastInsertId int64
	)
	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
		response.ResFail(ctx, err.Error())
		return
	}
	global.MyLogger(ctx).Info().Msgf("req: %+v", *req)

	if req.Status == 0 {
		req.Status = ADSlotStatusOffline
	}
	if req.Status != ADSlotStatusOnline && req.Status != ADSlotStatusOffline {
		response.ResFail(ctx, "参数 Status 设置值无效")
		return
	}

	err = dao.TAdSlot.Ctx(ctx).Where(do.TAdSlot{Location: req.Location}).Scan(&adSlot)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("query ad slot failed")
		response.ResFail(ctx, err.Error())
		return
	}
	if adSlot != nil {
		response.ResFail(ctx, fmt.Sprintf("%s 已经存在，不能创建相同的广告位", req.Location))
		return
	}

	lastInsertId, err = dao.TAdSlot.Ctx(ctx).Data(do.TAdSlot{
		Location:  req.Location,
		Name:      req.Name,
		Desc:      req.Desc,
		Status:    req.Status,
		CreatedAt: gtime.Now(),
		UpdatedAt: gtime.Now(),
	}).InsertAndGetId()
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("add ad slot failed")
		response.ResFail(ctx, err.Error())
		return
	}
	global.MyLogger(ctx).Debug().Msgf("lastInsertId: %d", lastInsertId)
	response.ResOk(ctx, "成功")
}
