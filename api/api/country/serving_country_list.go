package country

import (
	"github.com/gin-gonic/gin"
	"go-speed/api/api/common"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/response"
	"go-speed/service"
)

type ServingCountryListReq struct {
	UserId uint64 `form:"user_id" binding:"required" json:"user_id"`
}

type ServingCountryListRes struct {
	PreferredCountry string           `json:"preferred_country" dc:"用户倾向的国家名称"`
	Items            []ServingCountry `json:"items" dc:"在役的国家列表"`
}

type ServingCountry struct {
	Name        string `json:"name" dc:"国家名称，不可以修改，后端当ID用"`
	NameDisplay string `json:"name_display" dc:"用于在用户侧展示的国家名称"`
	LogoLink    string `json:"logo_link" dc:"国家图片地址"`
	Weight      int    `json:"weight" dc:"权重。权重越高越靠前"`
	Level       int    `json:"level" dc:"等级约束：0-所有用户都可以选择；1-青铜、铂金会员可选择；2-铂金会员可选择"`
	IsFree      int    `json:"is_free" dc:"是否为免费站点，0: 不免费,1: 免费"`
}

// ServingCountryList 查询国家列表
func ServingCountryList(ctx *gin.Context) {
	var (
		err        error
		req        = new(ServingCountryListReq)
		items      []entity.TServingCountry
		userEntity *entity.TUser
		userId     uint64
	)

	if err = ctx.ShouldBind(req); err != nil {
		claims := ctx.MustGet("claims").(*service.CustomClaims)
		user, err := service.GetUserByClaims(claims)
		if err != nil {
			global.MyLogger(ctx).Err(err).Msgf("用户token鉴权失败")
			response.RespFail(ctx, i18n.RetMsgAuthFailed, nil, response.CodeTokenExpired)
			return
		}
		userId = uint64(user.Id)
		//global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
		//response.RespFail(ctx, i18n.RetMsgParamParseErr, nil)
		//return
	} else {
		userId = req.UserId
	}

	userEntity, err = common.CheckUserByUserId(ctx, userId)
	if err != nil {
		return
	}

	err = dao.TServingCountry.Ctx(ctx).
		Where(do.TServingCountry{Status: 1}).
		Order(dao.TServingCountry.Columns().Weight, "Desc").
		Scan(&items)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("get serving country failed")
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}
	countries := make([]ServingCountry, 0)
	for _, item := range items {
		countries = append(countries, ServingCountry{
			Name:        item.Name,
			NameDisplay: item.Display,
			LogoLink:    item.LogoLink,
			Weight:      item.Weight,
			Level:       item.Level,
			IsFree:      item.IsFree,
		})
	}
	response.RespOk(ctx, i18n.RetMsgSuccess, ServingCountryListRes{
		PreferredCountry: userEntity.PreferredCountry,
		Items:            countries,
	})
}
