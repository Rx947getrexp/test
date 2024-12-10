package official_docs

import (
	"go-speed/i18n"
	"go-speed/model/response"
	"go-speed/service"

	"github.com/gin-gonic/gin"
)

type OfficialDocsListReq struct {
}

type OfficialDocsListRes struct {
	Items []DocItem `json:"items" dc:"文档列表"`
}

type DocItem struct {
	Id        uint64 `json:"id" dc:"用户uid"`
	Type      string `json:"type" dc:"文档类型"`
	Name      string `json:"name" dc:"文档名称"`
	Desc      string `json:"desc" dc:"文档描述"`
	Content   string `json:"content" dc:"文档内容"`
	CreatedAt string `json:"created_at" dc:"创建时间"`
	UpdatedAt string `json:"updated_at" dc:"更新时间"`
}

func OfficialDocsList(ctx *gin.Context) {
	resp, err := service.OfficialDocsList(ctx)
	if err != nil {
		response.RespFail(ctx, err.Error(), nil)
		return
	}
	response.RespOk(ctx, i18n.RetMsgSuccess, *resp)
}

// OfficialDocsById 通过文档id获取文档
func OfficialDocById(ctx *gin.Context) {
	// 调用 service.OfficialDoc 并解构返回值
	doc, err := service.OfficialDoc(ctx)
	if err != nil {
		// 如果有错误，这里已经由 OfficialDoc 函数处理过了，无需重复处理
		return
	}

	// 正常情况下返回成功响应
	response.RespOk(ctx, i18n.RetMsgSuccess, doc)
}
