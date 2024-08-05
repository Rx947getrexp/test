package official_docs

import (
	"github.com/gogf/gf/v2/os/gtime"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/response"

	"github.com/gin-gonic/gin"
)

type OfficialDocsCreateReq struct {
	Type    string `form:"type" json:"type" binding:"required" dc:"文档类型"`
	Name    string `form:"name" json:"name" binding:"required" dc:"文档名称"`
	Desc    string `form:"desc" json:"desc" dc:"文档描述"`
	Content string `form:"content" json:"content" binding:"required" dc:"文档内容"`
}

type OfficialDocsCreateRes struct {
}

func OfficialDocsCreate(ctx *gin.Context) {
	var (
		err          error
		req          = new(OfficialDocsCreateReq)
		lastInsertId int64
	)
	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("bind-params-failed")
		response.ResFail(ctx, err.Error())
		return
	}
	global.MyLogger(ctx).Info().Msgf("req: %+v", *req)

	lastInsertId, err = dao.TDoc.Ctx(ctx).Data(do.TDoc{
		Type:      req.Type,
		Name:      req.Name,
		Desc:      req.Desc,
		Content:   req.Content,
		CreatedAt: gtime.Now(),
		UpdatedAt: gtime.Now(),
	}).InsertAndGetId()
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("insert doc failed")
		response.RespFail(ctx, err.Error(), nil)
		return
	}
	global.MyLogger(ctx).Info().Msgf("lastInsertId: %d", lastInsertId)
	response.ResOk(ctx, i18n.RetMsgSuccess)
}
