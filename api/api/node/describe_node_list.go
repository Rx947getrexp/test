package node

import (
	"github.com/gin-gonic/gin"
	"go-speed/api/types/api"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/response"
)

func DescribeNodeList(ctx *gin.Context) {
	var (
		err          error
		nodeEntities []entity.TNode
	)

	err = dao.TNode.Ctx(ctx).Where(do.TNode{Status: 1}).Scan(&nodeEntities)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("query db failed")
		response.ResFail(ctx, err.Error())
		return
	}

	var items []api.NodeItem
	for _, node := range nodeEntities {
		items = append(items, api.NodeItem{
			Id:          node.Id,
			Name:        node.Name,
			Title:       node.Title,
			CountryEn:   node.CountryEn,
			Ip:          node.Ip,
			Server:      node.Server,
			Port:        node.Port,
			MinPort:     node.MinPort,
			MaxPort:     node.MaxPort,
			Path:        node.Path,
			IsRecommend: node.IsRecommend,
			ChannelId:   node.ChannelId,
			Status:      node.Status,
			Weight:      node.Weight,
		})
	}
	response.RespOk(ctx, i18n.RetMsgSuccess, api.DescribeNodeListRes{Items: items})
}
