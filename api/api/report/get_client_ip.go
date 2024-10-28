package report

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/response"
)

var nodeIPs map[string]struct{}

func init() {
	nodeIPs = make(map[string]struct{})

	var nodes []entity.TNode
	err := dao.TNode.Ctx(context.Background()).Where(do.TNode{Status: 1}).Scan(&nodes)
	if err != nil {
		global.Logger.Err(err).Msgf("get TNode failed")
		return
	}
	for _, node := range nodes {
		nodeIPs[node.Ip] = struct{}{}
	}
	global.Logger.Info().Msgf("get TNode nodeIPs: %+v", nodeIPs)
}

type GetClientIPResponse struct {
	ClientIP        string `json:"client_ip"`
	CheckNodeResult int    `json:"check_node_result"`
}

func GetClientIP(c *gin.Context) {
	global.MyLogger(c).Info().Msgf("ClientIP: %+v", c.ClientIP())

	response.RespOk(c, i18n.RetMsgSuccess, GetClientIPResponse{ClientIP: c.ClientIP(), CheckNodeResult: GetNodes(c)})
}

func GetNodes(c *gin.Context) int {
	if nodeIPs == nil || len(nodeIPs) == 0 {
		return -1
	}

	if _, ok := nodeIPs[c.ClientIP()]; ok {
		return 200
	} else {
		return 404
	}
}
