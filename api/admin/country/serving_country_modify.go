package country

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/os/gtime"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/response"
)

type ServingCountryModifyReq struct {
	Name        string `form:"name" binding:"required" json:"name" dc:"国家名称"`
	NameDisplay string `form:"name_display" json:"name_display" dc:"用于在用户侧展示的国家名称"`
	LogoLink    string `form:"logo_link" json:"logo_link" dc:"国家图片地址"`
	PingUrl     string `form:"ping_url" json:"ping_url" dc:"ping的地址，供前端使用"`
	IsRecommend uint   `form:"is_recommend" json:"is_recommend" dc:"是否为推荐的国家，0:否，1：是"`
	Weight      uint   `form:"weight" json:"weight" dc:"推荐权重,权重越大的国家展示在越靠前"`
	Status      uint   `form:"status" json:"status" dc:"状态:0:未上架，1-已上架；2-已下架"`
	Level       *int   `form:"level" json:"level" dc:"等级约束：0-所有用户都可以选择；1-青铜、铂金会员可选择；2-铂金会员可选择"`
	IsFree      *int   `form:"is_free" json:"is_free" dc:"是否为免费站点，0: 不免费,1: 免费"`
}

type ServingCountryModifyRes struct {
}

// ServingCountryModify 修改国家信息
func ServingCountryModify(ctx *gin.Context) {
	var (
		err           error
		req           = new(ServingCountryModifyReq)
		countryEntity *entity.TServingCountry
		affected      int64
	)
	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
		response.ResFail(ctx, i18n.RetMsgParamParseErr)
		return
	}
	global.MyLogger(ctx).Info().Msgf("req: %+v", *req)
	err = dao.TServingCountry.Ctx(ctx).Where(do.TServingCountry{Name: req.Name}).Scan(&countryEntity)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("query serving country failed")
		response.ResFail(ctx, err.Error())
		return
	}
	if countryEntity == nil {
		global.MyLogger(ctx).Err(err).Msgf("country name invalid")
		response.ResFail(ctx, i18n.RetMsgParamInvalid)
		return
	}
	if req.Level != nil && (*req.Level < 0 || *req.Level > 2) {
		global.MyLogger(ctx).Err(err).Msgf(`param "Level" invalid`)
		response.ResFail(ctx, fmt.Sprintf(`param "Level"(%d) invalid`, *req.Level))
		return
	}

	updateData := do.TServingCountry{UpdatedAt: gtime.Now()}
	if req.NameDisplay != "" {
		updateData.Display = req.NameDisplay
	}
	if req.LogoLink != "" {
		updateData.LogoLink = req.LogoLink
	}
	if req.PingUrl != "" {
		updateData.PingUrl = req.PingUrl
	}
	if req.IsRecommend != 0 {
		updateData.IsRecommend = req.IsRecommend
	}
	if req.Weight != 0 {
		updateData.Weight = req.Weight
	}
	if req.Status != 0 {
		updateData.Status = req.Status
	}
	if req.Level != nil {
		updateData.Level = *req.Level
	}
	if req.IsFree != nil {
		updateData.IsFree = *req.IsFree
	}
	affected, err = dao.TServingCountry.Ctx(ctx).Data(updateData).Where(do.TServingCountry{
		Name: req.Name,
	}).UpdateAndGetAffected()
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("modify serving country failed")
		response.ResFail(ctx, err.Error())
		return
	}
	global.MyLogger(ctx).Debug().Msgf("affected: %d", affected)
	response.ResOk(ctx, "成功")
}
