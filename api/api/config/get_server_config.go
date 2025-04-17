package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-speed/api/api/common"
	"go-speed/config"
	"go-speed/constant"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/response"
	"go-speed/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetServerConfigReq struct {
	UserId      uint64 `form:"user_id" binding:"required" json:"user_id"`
	CountryName string `form:"country_name" json:"country_name"`
}

type GetServerConfigRes struct {
}

func GetServerConfig(ctx *gin.Context) {
	var (
		err          error
		req          = new(GetServerConfigReq)
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

	v2rayServers := make([]Server, 0)
	for _, item := range nodeEntities {
		nodeId := item.Id
		nodePorts := []int{item.Port}
		for x := item.MinPort; x <= item.MaxPort; x++ {
			nodePorts = append(nodePorts, x)
		}
		global.MyLogger(ctx).Info().Msgf(">>>>> nodeId:%d, nodePorts: %+v", nodeId, nodePorts)
		var dnsList []entity.TNodeDns
		err = dao.TNodeDns.Ctx(ctx).Where(do.TNodeDns{
			NodeId: nodeId,
			Level:  userEntity.Level + 1,
		}).Scan(&dnsList)
		if err != nil {
			global.MyLogger(ctx).Err(err).Msgf(">>>>> FindNodeDnsByNodeId failed: %+v", err.Error())
			continue
		}
		for _, dns := range dnsList {
			for _, nodePort := range nodePorts {
				v2rayServers = append(v2rayServers, Server{Password: userEntity.V2RayUuid, Port: nodePort, Address: dns.Dns})
			}
		}
	}
	global.MyLogger(ctx).Info().Msgf(">>>>> v2rayServers: %+v", v2rayServers)

	v, err := json.Marshal(GenV2rayConfig(ctx, v2rayServers, winCountry.Name, false))
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("GenV2rayConfig failed")
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}
	ctx.String(http.StatusOK, fmt.Sprintf(string(v)))
}

func chooseCountryForUser(ctx *gin.Context, userId uint64, countryName string) (
	winServingCountry entity.TServingCountry, nodeEntities []entity.TNode, err error) {
	var (
		userEntity       *entity.TUser
		countryEntities  []entity.TServingCountry
		chooseCountry    string
		preferredCountry string
		recommendCountry string
		weightCountry    string
		winCountry       string
		weight           int = -1
		isADVersion          = len(global.GetClientVersion(ctx)) > 0
	)

	userEntity, err = common.CheckUserByUserId(ctx, userId)
	if err != nil {
		return
	}

	where := do.TServingCountry{Status: 1}
	if isADVersion && global.Config.ADConfig.KickExpiredUser == config.KickExpiredUserYes {
		global.MyLogger(ctx).Info().Msgf("######## %s clientVersion (%s) #######", userEntity.Email, global.GetClientVersion(ctx))
		if service.IsVIPExpired(userEntity) {
			err = errors.New(i18n.RetMsgAccountExpiredV2)
			global.MyLogger(ctx).Info().Msgf("######## %s expired #######", userEntity.Email)
			response.RespFail(ctx, i18n.RetMsgAccountExpiredV2, nil)
			return
		}
	} else {
		global.MyLogger(ctx).Info().Msgf("######## %s clientVersion is empty #######", userEntity.Email)

		// 过期用户只能选择免费节点
		if service.IsVIPExpired(userEntity) {
			global.MyLogger(ctx).Info().Msgf("######## %s expired, choose free site #######", userEntity.Email)
			where.IsFree = constant.IsFreeSiteYes
		}
	}

	err = dao.TServingCountry.Ctx(ctx).
		Where(where).
		Order(dao.TServingCountry.Columns().Weight, "Desc").
		Scan(&countryEntities)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("get serving country failed")
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}

	var match bool
	for _, country := range countryEntities {
		if country.Name == countryName || country.Name == userEntity.PreferredCountry {
			match = true
			break
		}
	}
	if service.IsVIPExpired(userEntity) && !match {
		err = errors.New(i18n.RetMsgAccountExpired)
		global.MyLogger(ctx).Warn().Msgf(">>>>>>>>> 过期 user: %s, ExpiredTime: %d, 用户选择的不是免费站点",
			userEntity.Uname, userEntity.ExpiredTime)
		response.RespFail(ctx, i18n.RetMsgAccountExpired, nil)
		return
	}

	for _, country := range countryEntities {
		if isADVersion && country.Level > userEntity.Level {
			global.MyLogger(ctx).Info().Msgf("skip country(%s), countryLevel: %d, userLevel: %d", country.Name, country.Level, userEntity.Level)
			continue
		}

		// 优先选择用户指定的国家
		if country.Name == countryName {
			chooseCountry = country.Name
			break
		}

		// 用户选择的默认节点
		if country.Name == userEntity.PreferredCountry {
			preferredCountry = country.Name
		}

		// 其次选择推荐的国家
		if country.IsRecommend == 1 {
			recommendCountry = country.Name
		}

		// 按权重选择
		if country.Weight > weight {
			weight = country.Weight
			weightCountry = country.Name
		}
	}
	global.MyLogger(ctx).Info().Msgf("chooseCountry: %s, preferredCountry: %s, "+
		"recommendCountry: %s, weightCountry: %s", chooseCountry, preferredCountry, recommendCountry, weightCountry)
	if chooseCountry != "" {
		winCountry = chooseCountry
	} else if preferredCountry != "" {
		winCountry = preferredCountry
	} else if recommendCountry != "" {
		winCountry = recommendCountry
	} else {
		winCountry = weightCountry
	}
	global.MyLogger(ctx).Info().Msgf("winCountry: %s", winCountry)
	// 查询节点
	err = dao.TNode.Ctx(ctx).Where(do.TNode{
		CountryEn: winCountry,
		Status:    1,
	}).Scan(&nodeEntities)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("数据库链接出错")
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}

	// 如果没有找到任何节点，则返回失败
	if len(nodeEntities) == 0 {
		global.MyLogger(ctx).Warn().Msg("未找到任何可用节点")
		response.RespFail(ctx, i18n.RetMsgNoAvailableNode, nil)
		return
	}

	// 遍历国家列表，找到最终选中的国家对象并赋值返回
	for i, c := range countryEntities {
		if c.Name == winCountry {
			winServingCountry = countryEntities[i]
			break
		}
	}
	return winServingCountry, nodeEntities, nil
}
