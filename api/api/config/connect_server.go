package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-speed/api/api/common"
	"go-speed/api/api/internal"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/entity"
	"go-speed/model/response"
	"go-speed/service"
	"time"
)

type ConnectServerReq struct {
	UserId      uint64 `form:"user_id" binding:"required" json:"user_id"`
	CountryName string `form:"country_name" json:"country_name"`
}

type ConnectServerRes struct {
	CountryName        string `json:"country_name" dc:"连接的国家名"`
	CountryNameDisplay string `json:"country_name_display" dc:"连接的国家展示名称"`
}

func ConnectServer(ctx *gin.Context) {
	var (
		err          error
		req          = new(ConnectServerReq)
		userEntity   *entity.TUser
		winCountry   entity.TServingCountry
		nodeEntities []entity.TNode
	)
	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
		response.RespFail(ctx, i18n.RetMsgParamParseErr, nil)
		return
	}
	global.MyLogger(ctx).Info().Msgf(">>> req: %+v", *req)

	userEntity, err = common.CheckUserByUserId(ctx, req.UserId)
	if err != nil {
		return
	}
	err = internal.CheckClientIdNumLimitsWithErrRsp(ctx, userEntity)
	if err != nil {
		return
	}

	winCountry, nodeEntities, err = chooseCountryForUser(ctx, req.UserId, req.CountryName)
	if err != nil {
		return
	}

	// 账号过期
	if len(nodeEntities) == 0 && userEntity.ExpiredTime < time.Now().Unix() {
		global.MyLogger(ctx).Warn().Msgf(">>>>>>>>> 过期 user: %s, ExpiredTime: %d", userEntity.Uname, userEntity.ExpiredTime)
		response.RespFail(ctx, i18n.RetMsgAccountExpired, nil)
		return
	}

	global.MyLogger(ctx).Info().Msgf(">>>>>>>>> nodeEntities: %+v", nodeEntities)
	allFailed := true
	for _, item := range nodeEntities {
		err = service.AddUserConfigToNode(ctx, userEntity, &item)
		if err != nil {
			global.MyLogger(ctx).Err(err).Msgf("AddUserConfigToNode failed, %s node: %s", userEntity.Email, item.Ip)
			continue
		}
		allFailed = false
	}
	if allFailed {
		err = fmt.Errorf("allFailed is true")
		global.MyLogger(ctx).Err(err).Msgf("%s connect server failed", userEntity.Uname)
		response.RespFail(ctx, i18n.RetMsgGetV2rayConfigFailed, nil)
		return
	}

	count := IncrementCounter(winCountry.Name)
	global.MyLogger(ctx).Info().Msgf("[connect-node-for-user] user: %s, country: %s, count: %d", userEntity.Email, winCountry.Name, count)

	var result = make(map[string]interface{})
	result["node_name"] = winCountry.Name
	result["country_name"] = winCountry.Name
	result["country_name_display"] = winCountry.Display
	result["ping_url"] = winCountry.PingUrl
	result["logo_link"] = winCountry.LogoLink
	response.RespOk(ctx, i18n.RetMsgSuccess, result)
	global.MyLogger(ctx).Info().Msgf(">>>>>>>>> user: %s, result: %v", userEntity.Uname, result)
	return
}
