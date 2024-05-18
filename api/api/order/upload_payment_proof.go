package order

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gogf/gf/v2/os/gtime"
	"go-speed/api/api/common"
	"go-speed/constant"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/request"
	"go-speed/model/response"
	"go-speed/service"
	"mime/multipart"
)

type UploadPaymentProofReq struct {
	OrderNo  string                `form:"order_no" json:"order_no" binding:"required" dc:"订单号"`
	File     *multipart.FileHeader `form:"file" json:"file" binding:"required"`
	FileType string                `form:"file_type,default=default" json:"file_type" `
}

type UploadPaymentProofRes struct {
	//Url string `json:"url" dc:"订单状态" dc:"success:成功，fail:支付失败,waiting：等待支付中"`
}

func UploadPaymentProof(ctx *gin.Context) {
	var (
		err      error
		req      = new(UploadPaymentProofReq)
		user     *entity.TUser
		payOrder *entity.TPayOrder
		fileUrl  string
		affected int64
	)
	if err = ctx.ShouldBindWith(req, binding.FormMultipart); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
		response.RespFail(ctx, i18n.RetMsgParamParseErr, nil)
		return
	}
	global.MyLogger(ctx).Info().Msgf("req: %+v", *req)

	// validate user
	user, err = common.ValidateClaims(ctx)
	if err != nil {
		return
	}

	// validate order
	payOrder, err = ValidateOrder(ctx, user.Email, req.OrderNo)
	if err != nil {
		return
	}

	if payOrder.PaymentChannelId != constant.PayChannelBankCardPay {
		err = fmt.Errorf("PaymentChannelId %s 无效", payOrder.PaymentChannelId)
		global.MyLogger(ctx).Err(err).Msgf("银行卡支付才需要上传凭证")
		response.RespFail(ctx, i18n.RetMsgParamInvalid, nil)
		return
	}

	// upload limit
	if payOrder.PaymentProof != "" {
		err = fmt.Errorf(i18n.RetMsgProofUploadLimit)
		global.MyLogger(ctx).Err(err).Msgf("PaymentProof: %s", payOrder.PaymentProof)
		response.RespFail(ctx, i18n.RetMsgProofUploadLimit, nil)
		return
	}

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
		Files:    req.File,
		FileType: req.FileType,
	}, fileMap, constant.UploadFilePathOrder)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("Upload failed")
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}
	global.MyLogger(ctx).Info().Msgf("fileUrl: %s", fileUrl)
	//data := new(response.DataResponse)
	//data.Data.Url = resUrl
	// save fileUrl
	affected, err = dao.TPayOrder.Ctx(ctx).Data(do.TPayOrder{
		PaymentProof: fileUrl,
		UpdatedAt:    gtime.Now(),
		Version:      payOrder.Version + 1,
	}).Where(do.TPayOrder{
		OrderNo: payOrder.OrderNo,
		Version: payOrder.Version,
	}).UpdateAndGetAffected()
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("modify order status failed")
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}
	global.MyLogger(ctx).Debug().Msgf("affected: %d", affected)
	response.ResOk(ctx, i18n.RetMsgSuccess)
}
