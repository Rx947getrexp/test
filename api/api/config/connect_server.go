package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-speed/api/api/common"
	"go-speed/constant"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/entity"
	"go-speed/model/request"
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

	winCountry, nodeEntities, err = chooseCountryForUser(ctx, req.UserId, req.CountryName)
	if err != nil {
		return
	}

	// 账号过期
	if len(nodeEntities) == 0 && userEntity.ExpiredTime < time.Now().Unix() {
		global.MyLogger(ctx).Error().Msgf(">>>>>>>>> 过期 user: %s, ExpiredTime: %d", userEntity.Uname, userEntity.ExpiredTime)
		response.RespFail(ctx, i18n.RetMsgAccountExpired, nil)
		return
	}

	nodeAddSubRequest := &request.NodeAddSubRequest{}
	if userEntity.ExpiredTime >= time.Now().Unix() || winCountry.IsFree == constant.IsFreeSiteYes {
		nodeAddSubRequest.Tag = "1"
	} else {
		nodeAddSubRequest.Tag = "2"
		global.MyLogger(ctx).Error().Msgf(">>>>>>>>> 过期 user: %s, ExpiredTime: %d", userEntity.Uname, userEntity.ExpiredTime)
	}

	//if user.V2rayUuid == "c541b521-17dd-11ee-bc4e-0c9d92c013fb" || user.V2rayUuid == "bf268a88-318f-d58f-0e9f-66d6f066be31" {
	//	//fmt.Printf("connect ok %s", req.Uuid)
	//	//response.ResOk(c, i18n.RetMsgSuccess)
	//	//return
	//	req.Tag = "1"
	//}
	//nodeAddSubRequest.Uuid = userEntity.V2rayUuid
	nodeAddSubRequest.Uuid = userEntity.V2RayUuid
	nodeAddSubRequest.Email = userEntity.Email
	nodeAddSubRequest.Level = fmt.Sprintf("%d", userEntity.Level)

	global.MyLogger(ctx).Info().Msgf(">>>>>>>>> nodeEntities: %+v", nodeEntities)
	for _, item := range nodeEntities {
		url := fmt.Sprintf("https://%s/site-api/node/add_sub", item.Server)
		if strings.Contains(item.Server, "http") {
			url = fmt.Sprintf("%s/node/add_sub", item.Server)
		}
		global.MyLogger(ctx).Info().Msgf(">>>>>>>>> url: %s", url)
		timestamp := fmt.Sprint(time.Now().Unix())
		headerParam := make(map[string]string)
		res := new(response.Response)
		headerParam["timestamp"] = timestamp
		headerParam["accessToken"] = util.MD5(fmt.Sprint(timestamp, constant.AccessTokenSalt))
		err = util.HttpClientPostV2(url, headerParam, nodeAddSubRequest, res)
		if res != nil {
			global.MyLogger(ctx).Info().Msgf(">>>>>>>>> nodeAddSubRequest: %+v, res: %+v", *nodeAddSubRequest, *res)
		} else {
			global.MyLogger(ctx).Info().Msgf(">>>>>>>>> nodeAddSubRequest: %+v, res: is nil", *nodeAddSubRequest)
		}
		if err != nil {
			global.MyLogger(ctx).Err(err).Msgf("email: %s, add_sub 发送失败", userEntity.Email)
			continue
		}
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
