package node

import (
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/entity"
	"go-speed/model/response"

	"github.com/gin-gonic/gin"
)

type NodeListReq struct {
	CountryName string `form:"country_name" dc:"国家名称"`
	IP          string `form:"ip" dc:"机器IP, eg: 45.150.236.6"`
	Server      string `form:"server" dc:"域名, eg: ru.workones.xyz"`
	Status      uint   `form:"status" dc:"状态。1-已上架；2-已下架"`
}

type NodeListRes struct {
	Items []Node `json:"items" dc:"机器列表"`
}

type Node struct {
	Id          uint64 `json:"id" dc:"机器节点ID"`
	CountryName string `json:"country_name" dc:"国家名称"`
	IP          string `json:"ip" dc:"机器IP, eg: 45.150.236.6"`
	Server      string `json:"server" dc:"域名, eg: ru.workones.xyz"`
	Path        string `json:"path" dc:"ws路径"`
	Port        uint   `json:"port" dc:"管控端口号，eg: 443"`
	MinPort     uint   `json:"min_port" dc:"监听的端口号，起始端口号, eg: 13001"`
	MaxPort     uint   `json:"max_port" dc:"监听的端口号，结束端口号, eg: 13005"`
	Weight      uint   `json:"weight" dc:"推荐权重,权重越大的节点优先连接"`
	Comment     string `json:"comment" dc:"备注信息"`
	Status      uint   `json:"status" dc:"状态。1-已上架；2-已下架"`
	CreatedAt   string `json:"created_at" dc:"创建时间"`
	UpdatedAt   string `json:"updated_at" dc:"更新时间"`
}

// NodeList 查询国家列表
func NodeList(ctx *gin.Context) {
	var (
		err   error
		items []entity.TNode
	)
	err = dao.TNode.Ctx(ctx).Scan(&items)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("get node failed")
		response.ResFail(ctx, err.Error())
		return
	}
	nodes := make([]Node, 0)
	for _, item := range items {
		nodes = append(nodes, Node{
			Id:          uint64(item.Id),
			CountryName: item.CountryEn,
			IP:          item.Ip,
			Server:      item.Ip,
			Path:        item.Ip,
			Port:        uint(item.Port),
			MinPort:     uint(item.MinPort),
			MaxPort:     uint(item.MaxPort),
			Weight:      item.Weight,
			Comment:     item.Comment,
			Status:      uint(item.Status),
			CreatedAt:   item.CreatedAt.String(),
			UpdatedAt:   item.UpdatedAt.String(),
		})
	}
	response.RespOk(ctx, i18n.RetMsgSuccess, NodeListRes{
		Items: nodes,
	})
}
