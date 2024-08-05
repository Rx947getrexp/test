package official_docs

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-speed/constant"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/request"
	"go-speed/model/response"
	"go-speed/service"
	"mime/multipart"
)

type UploadOfficialDocsImageReq struct {
	Files    *multipart.FileHeader `form:"files" json:"files" binding:"required"`
	FileType string                `form:"file_type,default=default" json:"file_type" `
}

type UploadOfficialDocsImageRes struct {
	Url string `json:"url" dc:"地址"`
}

func UploadOfficialDocsImage(ctx *gin.Context) {
	var (
		err     error
		req     = new(UploadOfficialDocsImageReq)
		fileUrl string
	)
	if err = ctx.ShouldBindWith(req, binding.FormMultipart); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
		response.RespFail(ctx, i18n.RetMsgParamParseErr, nil)
		return
	}
	global.MyLogger(ctx).Info().Msgf("req: %+v", *req)

	// upload
	if req.FileType != constant.ImgFileType &&
		req.FileType != constant.OtherFileType {
		err = fmt.Errorf("FileType %s 无效", req.FileType)
		global.MyLogger(ctx).Err(err).Msgf("FileType invalid")
		response.RespFail(ctx, i18n.RetMsgParamInvalid, nil)
		return
	}

	fileMap := make(map[string]bool)
	if req.FileType == constant.ImgFileType {
		fileMap[".png"] = true
		fileMap[".jpg"] = true
		fileMap[".jpeg"] = true
	}
	global.MyLogger(ctx).Info().Msgf("fileMap: %+v", fileMap)

	fileUrl, err = service.Upload(&request.UploadFile{
		Files:    req.Files,
		FileType: req.FileType,
	}, fileMap, constant.UploadFilePathOfficialDocs)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("Upload failed")
		response.ResFail(ctx, err.Error())
		return
	}
	global.MyLogger(ctx).Info().Msgf("fileUrl: %s", fileUrl)
	response.RespOk(ctx, i18n.RetMsgSuccess, UploadOfficialDocsImageRes{Url: fileUrl})
}
