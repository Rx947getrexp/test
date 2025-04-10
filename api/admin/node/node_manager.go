package node

import (
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/response"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/os/gtime"
)

type NodeManagerReq struct {
	Server string `form:"server" binding:"required" json:"server" dc:"公网域名"`
	Status int    `form:"status" binding:"required" json:"status" dc:"状态。1-已上架；2-已下架"`
}

func NodeManager(c *gin.Context) {
	ip := c.ClientIP()
	if ip != "185.22.154.21" {
		global.Logger.Warn().Msgf("非法请求IP:%v", ip)
		response.RespFail(c, i18n.RetMsgOperateFailed, nil)
		return
	}
	req := new(NodeManagerReq)
	if err := c.ShouldBind(req); err != nil {
		global.Logger.Err(err).Msgf("参数校验失败，err:%v", err.Error())
		response.RespFail(c, i18n.RetMsgParamInvalid, nil)
		return
	}
	_, err := dao.TNode.Ctx(c).Where("server", req.Server).Data(do.TNode{
		Status:    req.Status,
		UpdatedAt: gtime.Now(),
	}).Update()

	if err != nil {
		global.Logger.Err(err).Msgf("更新节点状态失败，err:%v", err.Error())
		response.RespFail(c, "更新节点状态失败", nil)
		return
	}

	response.RespOk(c, i18n.RetMsgSuccess, nil)

}
