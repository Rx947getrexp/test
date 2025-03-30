package report

import (
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/response"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/os/gtime"
)

type ReportUserADLogReq struct {
	UserId     uint64 `form:"user_id" json:"user_id" dc:"用户ID"`
	ADLocation string `form:"ad_location" json:"ad_location" dc:"广告位"`
	ADName     string `form:"ad_name" json:"ad_name" dc:"广告名称"`
	DeviceType string `form:"device_type" json:"device_type" dc:"设备类型"`
	APPVersion string `form:"app_version" json:"app_version" dc:"前端版本"`
	Type       string `form:"type" json:"type" dc:"前端自定义"`
	Content    string `form:"content" json:"content" dc:"前端自定义"`
	Result     string `form:"result" json:"result" dc:"前端自定义"`
	ReportTime string `form:"report_time" json:"report_time" dc:"上报时间"`
}

func ReportUserADLog(ctx *gin.Context) {
	var (
		err          error
		req          = new(ReportUserADLogReq)
		lastInsertId int64
	)

	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
		response.RespFail(ctx, i18n.RetMsgParamParseErr, nil)
		return
	}

	lastInsertId, err = dao.TUserAdLog.Ctx(ctx).Data(do.TUserAdLog{
		UserId:     req.UserId,
		AdLocation: req.ADLocation,
		AdName:     req.ADName,
		DeviceType: req.DeviceType,
		AppVersion: req.APPVersion,
		ClientId:   global.GetClientId(ctx),
		Type:       req.Type,
		Content:    req.Content,
		Result:     req.Result,
		ReportTime: req.ReportTime,
		CreatedAt:  gtime.Now(),
		AppName:    global.GetAppVersion(ctx),
	}).InsertAndGetId()
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("insert ad log failed")
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}
	global.MyLogger(ctx).Debug().Msgf("lastInsertId: %d", lastInsertId)
	response.RespOk(ctx, i18n.RetMsgSuccess, nil)
	return
}
