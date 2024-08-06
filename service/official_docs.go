package service

import (
	"github.com/gin-gonic/gin"
	"go-speed/api/api/official_docs"
	"go-speed/constant"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/model/entity"
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
