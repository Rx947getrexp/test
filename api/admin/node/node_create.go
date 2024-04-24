package node

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/os/gtime"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/response"
)

type NodeCreateReq struct {
	CountryName string `form:"country_name" binding:"required" json:"country_name" dc:"国家名称"`
	IP          string `form:"ip" binding:"required" json:"ip" dc:"机器IP, eg: 45.150.236.6"`
	Server      string `form:"server" binding:"required" json:"server" dc:"域名, eg: ru.workones.xyz"`
	Port        uint   `form:"port" binding:"required" json:"port" dc:"管控端口号，eg: 443"`
	MinPort     uint   `form:"min_port" binding:"required" json:"min_port" dc:"监听的端口号，起始端口号, eg: 13001"`
	MaxPort     uint   `form:"max_port" binding:"required" json:"max_port" dc:"监听的端口号，结束端口号, eg: 13005"`
	Weight      uint   `form:"weight" binding:"required" json:"weight" dc:"推荐权重,权重越大的节点优先连接"`
	Comment     string `form:"comment" json:"comment" dc:"备注信息"`
}

//NodeType     uint   `binding:"required" json:"node_type" dc:"节点类别:1-常规；2-高带宽"`
//Path string `binding:"required" json:"path" dc:"ws路径"`

type NodeCreateRes struct {
}

// NodeCreate 添加代理机器节点
func NodeCreate(ctx *gin.Context) {
	var (
		err           error
		req           = new(NodeCreateReq)
		countryEntity *entity.TServingCountry
		lastInsertId  int64
	)
	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
		response.ResFail(ctx, i18n.RetMsgParamParseErr)
		return
	}
	global.MyLogger(ctx).Info().Msgf("req: %+v", *req)
	err = dao.TServingCountry.Ctx(ctx).Where(do.TServingCountry{Name: req.CountryName}).Scan(&countryEntity)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("query serving country failed")
		response.ResFail(ctx, err.Error())
		return
	}
	if countryEntity == nil {
		global.MyLogger(ctx).Err(err).Msgf("country name invalid")
		response.ResFail(ctx, i18n.RetMsgParamInvalid)
		return
	}

	lastInsertId, err = dao.TNode.Ctx(ctx).Data(do.TNode{
		Name:        req.CountryName,
		Title:       countryEntity.Display,
		TitleEn:     countryEntity.Display,
		TitleRus:    countryEntity.Display,
		Country:     countryEntity.Display,
		CountryEn:   countryEntity.Name,
		CountryRus:  countryEntity.Display,
		Ip:          req.IP,
		Server:      req.Server,
		NodeType:    1,
		Port:        req.Port,
		MinPort:     req.MinPort,
		MaxPort:     req.MaxPort,
		Path:        "/work",
		Cpu:         0,
		Flow:        0,
		Disk:        0,
		Memory:      0,
		IsRecommend: countryEntity.IsRecommend,
		Status:      1,
		CreatedAt:   gtime.Now(),
		UpdatedAt:   gtime.Now(),
		Author:      "",
		Comment:     req.Comment,
		Weight:      req.Weight,
	}).InsertAndGetId()
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("add node failed")
		response.ResFail(ctx, err.Error())
		return
	}
	global.MyLogger(ctx).Debug().Msgf("lastInsertId: %d", lastInsertId)
	response.ResOk(ctx, "成功")
}
