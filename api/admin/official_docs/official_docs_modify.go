package official_docs

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/response"

	"github.com/gin-gonic/gin"
)

type OfficialDocsModifyReq struct {
	Id      uint64  `form:"id" json:"id" binding:"required" dc:"ID"`
	Type    *string `form:"type" json:"type" dc:"文档类型"`
	Name    *string `form:"name" json:"name" dc:"文档名称"`
	Desc    *string `form:"desc" json:"desc" dc:"文档描述"`
	Content *string `form:"content" json:"content" dc:"文档内容"`
}

type OfficialDocsModifyRes struct {
}

func OfficialDocsModify(ctx *gin.Context) {
	var (
		err      error
		req      = new(OfficialDocsModifyReq)
		affected int64
		doc      *entity.TDoc
	)
	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("bind-params-failed")
		response.ResFail(ctx, err.Error())
		return
	}
	global.MyLogger(ctx).Info().Msgf("req: %+v", *req)

	err = dao.TDoc.Ctx(ctx).Where(do.TDoc{Id: req.Id}).Scan(&doc)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("query doc failed")
		response.RespFail(ctx, err.Error(), nil)
		return
	}
	if doc == nil {
		err = fmt.Errorf("doc not exist")
		global.MyLogger(ctx).Err(err).Msgf("Id: %d", req.Id)
		response.ResFail(ctx, err.Error())
		return
	}

	updateData := do.TDoc{UpdatedAt: gtime.Now()}
	if req.Type != nil {
		updateData.Type = *req.Type
	}
	if req.Name != nil {
		updateData.Name = *req.Name
	}
	if req.Desc != nil {
		updateData.Desc = *req.Desc
	}
	if req.Content != nil {
		updateData.Content = *req.Content
	}

	global.MyLogger(ctx).Info().Msgf("param `updateData` is: %#v", updateData)
	affected, err = dao.TDoc.Ctx(ctx).Data(updateData).Where(do.TDoc{Id: req.Id}).UpdateAndGetAffected()
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("modify doc failed")
		response.ResFail(ctx, err.Error())
		return
	}
	global.MyLogger(ctx).Info().Msgf("affected: %d", affected)
	response.ResOk(ctx, i18n.RetMsgSuccess)
}
