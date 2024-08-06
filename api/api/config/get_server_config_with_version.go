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
	"go-speed/util/geo"
	"net/http"
	"sync"
)

var (
	UserGetConfigCounter map[string]uint64
	mu                   sync.Mutex // 使用Mutex来保证线程安全
)

func init() {
	UserGetConfigCounter = make(map[string]uint64)
}

// GetCounter 安全地获取当前计数器的值
func GetCounter(country string) uint64 {
	mu.Lock()
	defer mu.Unlock()
	if count, ok := UserGetConfigCounter[country]; ok {
		return count
	} else {
		return 0
	}
}

func IncrementCounter(country string) uint64 {
	mu.Lock()
	defer mu.Unlock()
	if count, ok := UserGetConfigCounter[country]; ok {
		if count > 100000000000 {
			count = 0
		}
		UserGetConfigCounter[country] = count + 1
	} else {
		UserGetConfigCounter[country] = 1
	}
	return UserGetConfigCounter[country]
}

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
	if userEntity.Email == "test7@qq.com" {
		GetConfigOld(ctx)
		return
	}

	winCountry, nodeEntities, err = chooseCountryForUser(ctx, req.UserId, req.CountryName)
	if err != nil {
		return
	}
	nodesLen := len(nodeEntities)
	global.MyLogger(ctx).Info().Msgf(">>> winCountry: %s, len(nodeEntities): %d", winCountry.Name, nodesLen)

	dnss := make([]string, 0)
	v2rayServers := make([]Server, 0)

	index, counter := 0, GetCounter(winCountry.Name)
	if nodesLen > 1 {
		index = int(counter % uint64(nodesLen))
	}

	global.MyLogger(ctx).Info().Msgf("[choose-node-for-user-1] userId: %d, index: %d, counter: %d, nodesLen: %d, country: %s", userEntity.Id, index, counter, nodesLen, winCountry.Name)
	for i, item := range nodeEntities {
		if i != index {
			continue
		}
		global.MyLogger(ctx).Info().Msgf("[choose-node-for-user-2] (%s) (%s) (%s) (%s)", userEntity.Email, item.Ip, item.CountryEn, winCountry.Name)

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
			if !geo.IsInArrayIgnoreCase(dns.Dns, dnss) {
				dnss = append(dnss, dns.Dns)
			}
		}
	}
	global.MyLogger(ctx).Info().Msgf(">>>>> dns: %+v", dnss)
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
