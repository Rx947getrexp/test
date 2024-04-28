package config

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-speed/api/api/common"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/response"
	"net/http"
)

type GetServerConfigWithoutRulesReq struct {
	UserId      uint64 `form:"user_id" binding:"required" json:"user_id"`
	CountryName string `form:"country_name" json:"country_name"`
}

type GetServerConfigWithoutRulesRes struct {
}

func GetServerConfigWithoutRules(ctx *gin.Context) {
	var (
		err          error
		req          = new(GetServerConfigWithoutRulesReq)
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

	v, err := json.Marshal(GenV2rayConfig(ctx, v2rayServers, winCountry.Name, true))
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("GenV2rayConfig failed")
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}
	global.MyLogger(ctx).Info().Msgf(">>>>> V2rayConfig: %s", string(v))
	ctx.String(http.StatusOK, fmt.Sprintf(string(v)))
}
