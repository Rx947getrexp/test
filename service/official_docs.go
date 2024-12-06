package service

import (
	"fmt"
	"go-speed/api/api/official_docs"
	"go-speed/constant"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/entity"
	"go-speed/model/response"

	"github.com/gin-gonic/gin"
)

func OfficialDocsList(ctx *gin.Context) (resp *official_docs.OfficialDocsListRes, err error) {
	var items []entity.TDoc

	err = dao.TDoc.Ctx(ctx).Order(dao.TDoc.Columns().Id, constant.OrderTypeDesc).Scan(&items)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("query doc failed")
		return
	}
	var docs []official_docs.DocItem
	for _, i := range items {
		docs = append(docs, official_docs.DocItem{
			Id:        i.Id,
			Type:      i.Type,
			Name:      i.Name,
			Desc:      i.Desc,
			Content:   i.Content,
			CreatedAt: i.CreatedAt.String(),
			UpdatedAt: i.UpdatedAt.String(),
		})
	}
	return &official_docs.OfficialDocsListRes{
		Items: docs,
	}, nil
}

func OfficialDoc(ctx *gin.Context) (docEntity *entity.TDoc, err error) {
	id := ctx.Query("Id")

	if id == "" {
		global.MyLogger(ctx).Err(err).Msg("Doc id is required")
		response.RespFail(ctx, i18n.RetMsgParamInputInvalid, nil)
		return nil, fmt.Errorf("doc id is required")
	}

	err = dao.TDoc.Ctx(ctx).Where("id", id).Scan(&docEntity)

	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("doc %s 查询db失败", id)
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return nil, fmt.Errorf("query doc %s failed", id)
	}

	if docEntity == nil {
		err = fmt.Errorf("doc id无效 %s", id)
		global.MyLogger(ctx).Warn().Msgf("doc id无效 %s", id)
		response.RespFail(ctx, i18n.RetMsgOperateFailed, nil)
		return nil, err
	}
	return docEntity, nil
}
