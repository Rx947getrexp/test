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

type NodeModifyReq struct {
	Id      uint64 `binding:"required" form:"id" dc:"机器节点ID"`
	IP      string `form:"ip" dc:"机器IP, eg: 45.150.236.6"`
	Server  string `form:"server" dc:"域名, eg: ru.workones.xyz"`
	Port    uint   `form:"port" dc:"管控端口号，eg: 443"`
	MinPort uint   `form:"min_port" dc:"监听的端口号，起始端口号, eg: 13001"`
	MaxPort uint   `form:"max_port" dc:"监听的端口号，结束端口号, eg: 13005"`
	Weight  uint   `form:"weight" dc:"推荐权重,权重越大的节点优先连接"`
	Comment string `form:"comment" dc:"备注信息"`
	Status  uint   `form:"status" dc:"状态。1-已上架；2-已下架"`
}

type NodeModifyRes struct {
}

// NodeModify 添加代理机器节点
func NodeModify(ctx *gin.Context) {
	var (
		err        error
		req        = new(NodeModifyReq)
		nodeEntity *entity.TNode
		affected   int64
	)
	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
		response.ResFail(ctx, i18n.RetMsgParamParseErr)
		return
	}
	global.MyLogger(ctx).Info().Msgf("req: %+v", *req)
	err = dao.TNode.Ctx(ctx).Where(do.TNode{Id: req.Id}).Scan(&nodeEntity)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("query node failed")
		response.ResFail(ctx, err.Error())
		return
	}
	if nodeEntity == nil {
		global.MyLogger(ctx).Err(err).Msgf("param `Id` invalid")
		response.ResFail(ctx, i18n.RetMsgParamInvalid)
		return
	}

	updateData := do.TNode{UpdatedAt: gtime.Now()}
	if req.IP != "" {
		updateData.Ip = req.IP
	}
	if req.Server != "" {
		updateData.Server = req.Server
	}
	if req.Port != 0 {
		updateData.Port = req.Port
	}
	if req.MinPort != 0 {
		updateData.MinPort = req.MinPort
	}
	if req.MaxPort != 0 {
		updateData.MaxPort = req.MaxPort
	}
	if req.Weight != 0 {
		updateData.Weight = req.Weight
	}
	if req.Comment != "" {
		updateData.Comment = req.Comment
	}
	if req.Status != 0 {
		updateData.Status = req.Status
	}
	affected, err = dao.TNode.Ctx(ctx).Data(updateData).Where(do.TNode{
		Id: req.Id,
	}).UpdateAndGetAffected()
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("modify node failed")
		response.ResFail(ctx, err.Error())
		return
	}
	global.MyLogger(ctx).Debug().Msgf("affected: %d", affected)
	response.ResOk(ctx, "成功")
}
