package report

import (
	"go-speed/api/api/common/remote"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/response"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/os/gtime"
)

type ReportUserOpLogReq struct {
	Email        string `form:"email" json:"email"`
	UserId       uint64 `form:"user_id" json:"user_id"`
	DeviceId     string `form:"device_id" json:"device_id"`
	DeviceType   string `form:"device_type" json:"device_type"`
	PageName     string `form:"page_name" json:"page_name"`
	Content      string `form:"content" json:"content"`
	InterfaceUrl string `form:"interface_url" json:"interface_url"`
	ServerCode   string `form:"server_code" json:"server_code"`
	HttpCode     string `form:"http_code" json:"http_code"`
	TraceId      string `form:"trace_id" json:"trace_id"`
	CreateTime   string `form:"create_time" json:"create_time"`
	Version      string `form:"version" json:"version"`
	Result       string `form:"result" json:"result"`
}

func ReportUserOpLog(ctx *gin.Context) {
	var (
		err          error
		req          = new(ReportUserOpLogReq)
		email        string
		lastInsertId int64
	)

	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
		response.RespFail(ctx, i18n.RetMsgParamParseErr, nil)
		return
	}
	email = req.Email
	if req.UserId > 0 && email == "" {
		email, err = remote.GetUserEmailByUserId(ctx, req.UserId)
		if err != nil {
			global.MyLogger(ctx).Err(err).Msgf("GetUserEmailByUserId failed")
		}
	}

	lastInsertId, err = dao.TUserOpLog.Ctx(ctx).Data(do.TUserOpLog{
		Email:        email,
		UserId:       req.UserId,
		DeviceId:     req.DeviceId,
		DeviceType:   req.DeviceType,
		PageName:     req.PageName,
		Result:       req.Result,
		Content:      req.Content,
		InterfaceUrl: req.InterfaceUrl,
		ServerCode:   req.ServerCode,
		HttpCode:     req.HttpCode,
		TraceId:      req.TraceId,
		Version:      req.Version,
		CreateTime:   req.CreateTime,
		CreatedAt:    gtime.Now(),
	}).InsertAndGetId()
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("insert op log failed")
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}
	global.MyLogger(ctx).Debug().Msgf("lastInsertId: %d", lastInsertId)
	response.RespOk(ctx, i18n.RetMsgSuccess, nil)
	return
}
