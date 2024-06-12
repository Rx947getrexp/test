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

type ServingCountryCreateReq struct {
	Name        string `form:"name" binding:"required" json:"name" dc:"国家名称"`
	NameDisplay string `form:"name_display" binding:"required" json:"name_display" dc:"用于在用户侧展示的国家名称"`
	LogoLink    string `form:"logo_link" binding:"required" json:"logo_link" dc:"国家图片地址"`
	PingUrl     string `form:"ping_url" binding:"required" json:"ping_url" dc:"ping的地址，供前端使用"`
	IsRecommend uint   `form:"is_recommend" json:"is_recommend" dc:"是否为推荐的国家，0:否，1：是"`
	Weight      uint   `form:"weight" json:"weight" dc:"推荐权重,权重越大的国家展示在越靠前"`
	Level       int    `form:"level" json:"level" dc:"等级约束：0-所有用户都可以选择；1-青铜、铂金会员可选择；2-铂金会员可选择"`
}

type ServingCountryCreateRes struct {
}

// ServingCountryCreate 添加新的国家
func ServingCountryCreate(ctx *gin.Context) {
	var (
		err           error
		req           = new(ServingCountryCreateReq)
		countryEntity *entity.TCountry
		lastInsertId  int64
	)
	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
		response.ResFail(ctx, i18n.RetMsgParamParseErr)
		return
	}
	global.MyLogger(ctx).Info().Msgf("req: %+v", *req)
	err = dao.TCountry.Ctx(ctx).Where(do.TCountry{Name: req.Name}).Scan(&countryEntity)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("query country failed")
		response.ResFail(ctx, err.Error())
		return
	}
	if countryEntity == nil {
		global.MyLogger(ctx).Err(err).Msgf("country name invalid")
		response.ResFail(ctx, i18n.RetMsgParamInvalid)
		return
	}
	if req.Level < 0 || req.Level > 2 {
		global.MyLogger(ctx).Err(err).Msgf(`param "Level" invalid`)
		response.ResFail(ctx, fmt.Sprintf(`param "Level"(%d) invalid`, req.Level))
		return
	}

	lastInsertId, err = dao.TServingCountry.Ctx(ctx).Data(do.TServingCountry{
		Name:        req.Name,
		Display:     req.NameDisplay,
		LogoLink:    req.LogoLink,
		PingUrl:     req.PingUrl,
		IsRecommend: req.IsRecommend,
		Weight:      req.Weight,
		Status:      1,
		CreatedAt:   gtime.Now(),
		UpdatedAt:   gtime.Now(),
	}).InsertAndGetId()
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("add serving country failed")
		response.ResFail(ctx, err.Error())
		return
	}
	global.MyLogger(ctx).Debug().Msgf("lastInsertId: %d", lastInsertId)
	response.ResOk(ctx, "成功")
}
