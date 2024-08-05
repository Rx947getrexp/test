package official_docs

import (
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/response"

	"github.com/gin-gonic/gin"
)

type OfficialDocsDeleteReq struct {
	Id uint64 `form:"id" json:"id" binding:"required" dc:"ID"`
}

type OfficialDocsDeleteRes struct {
}

func OfficialDocsDelete(ctx *gin.Context) {
	var (
		err error
		req = new(OfficialDocsDeleteReq)
	)
	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("bind-params-failed")
		response.ResFail(ctx, err.Error())
		return
	}
	global.MyLogger(ctx).Info().Msgf("req: %+v", *req)

	_, err = dao.TDoc.Ctx(ctx).Delete(do.TDoc{Id: req.Id})
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("delete doc failed")
		response.RespFail(ctx, err.Error(), nil)
		return
	}
	response.ResOk(ctx, i18n.RetMsgSuccess)
}
