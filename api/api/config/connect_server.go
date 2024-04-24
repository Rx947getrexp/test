package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-speed/api/api/common"
	"go-speed/constant"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/entity"
	"go-speed/model/response"
	"go-speed/util"
	"strings"
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

	// 账号过期
	if userEntity.ExpiredTime < time.Now().Unix() {
		global.MyLogger(ctx).Error().Msgf(">>>>>>>>> 过期 user: %s, ExpiredTime: %d", userEntity.Uname, userEntity.ExpiredTime)
		response.RespFail(ctx, i18n.RetMsgAccountExpired, nil)
		return
	}

	winCountry, nodeEntities, err = chooseCountryForUser(ctx, req.UserId, req.CountryName)
	if err != nil {
		return
	}

	for _, item := range nodeEntities {
		url := fmt.Sprintf("https://%s/site-api/node/add_sub", item.Server)
		if strings.Contains(item.Server, "http") {
			url = fmt.Sprintf("%s/node/add_sub", item.Server)
		}
		timestamp := fmt.Sprint(time.Now().Unix())
		headerParam := make(map[string]string)
		res := new(response.Response)
		headerParam["timestamp"] = timestamp
		headerParam["accessToken"] = util.MD5(fmt.Sprint(timestamp, constant.AccessTokenSalt))
		err = util.HttpClientPostV2(url, headerParam, req, res)
		if err != nil {
			global.MyLogger(ctx).Err(err).Msgf("email: %s, add_sub 发送失败", userEntity.Email)
			continue
		}
	}

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
