package report

import (
	"fmt"
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
	UserId       uint64 `form:"user_id" json:"user_id"`
	DeviceId     string `form:"device_id" json:"device_id"`
	DeviceType   string `form:"device_type" json:"device_type"`
	PageName     string `form:"page_name" json:"page_name"`
	Content      string `form:"content" json:"content"`
	InterfaceUrl string `form:"interfaceUrl" json:"interfaceUrl"`
	ServerCode   string `form:"serverCode" json:"serverCode"`
	HttpCode     string `form:"httpCode" json:"httpCode"`
	TraceId      string `form:"traceId" json:"traceId"`
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
	if req.UserId > 0 {
		email, err = remote.GetUserEmailByUserId(ctx, req.UserId)
	}
	if email == "" {
		email = fmt.Sprintf("%d", req.UserId)
	}

	lastInsertId, err = dao.TUserOpLog.Ctx(ctx).Data(do.TUserOpLog{
		Email:        email,
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
