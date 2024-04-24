package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/os/gtime"
	"go-speed/api/api/common"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/response"
)

type SetPreferredCountryReq struct {
	UserId      uint64 `form:"user_id" binding:"required" json:"user_id"`
	CountryName string `form:"country_name" json:"country_name"`
}

type SetPreferredCountryRes struct {
}

// SetPreferredCountry 用户设置国家站点
func SetPreferredCountry(ctx *gin.Context) {
	var (
		err           error
		req           = new(SetPreferredCountryReq)
		userEntity    *entity.TUser
		countryEntity *entity.TServingCountry
		affected      int64
	)

	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
		response.RespFail(ctx, i18n.RetMsgParamParseErr, nil)
		return
	}

	userEntity, err = common.CheckUserByUserId(ctx, req.UserId)
	if err != nil {
		return
	}

	if req.CountryName != "" {
		err = dao.TServingCountry.Ctx(ctx).
			Where(do.TServingCountry{Name: req.CountryName, Status: 1}).
			Scan(&countryEntity)
		if err != nil {
			global.MyLogger(ctx).Err(err).Msgf("get serving country failed")
			response.RespFail(ctx, i18n.RetMsgDBErr, nil)
			return
		}
		if countryEntity == nil {
			global.MyLogger(ctx).Err(err).Msgf("get serving country failed")
			response.RespFail(ctx, i18n.RetMsgParamInvalid, nil)
			return
		}
	}

	if userEntity.PreferredCountry == req.CountryName {
		response.RespOk(ctx, i18n.RetMsgSuccess, nil)
		return
	}

	affected, err = dao.TUser.Ctx(ctx).Data(do.TUser{
		PreferredCountry: req.CountryName,
		UpdatedAt:        gtime.Now(),
	}).Where(do.TUser{Id: req.UserId}).UpdateAndGetAffected()
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("get serving country failed")
		response.RespFail(ctx, i18n.RetMsgParamInvalid, nil)
		return
	}
	global.MyLogger(ctx).Debug().Msgf("affected: %d", affected)
	response.RespOk(ctx, i18n.RetMsgSuccess, nil)
	return
}
