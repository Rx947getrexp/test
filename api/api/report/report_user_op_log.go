package report

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/os/gtime"
	"go-speed/api/api/common"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/response"
)

type ReportUserOpLogReq struct {
	UserId     uint64 `form:"user_id" json:"user_id"`
	DeviceId   string `form:"device_id" json:"device_id"`
	DeviceType string `form:"device_type" json:"device_type"`
	PageName   string `form:"page_name" json:"page_name"`
	Content    string `form:"content" json:"content"`
	CreateTime string `form:"create_time" json:"create_time"`
	Result     string `form:"result" json:"result"`
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
		var userEntity *entity.TUser
		userEntity, err = common.CheckUserByUserId(ctx, req.UserId)
		if err != nil {
			return
		}
		email = userEntity.Email
	}
	if req.DeviceId != "" {
		_, err = common.CheckDevId(ctx, req.DeviceId)
		if err != nil {
			return
		}
	}

	lastInsertId, err = dao.TUserOpLog.Ctx(ctx).Data(do.TUserOpLog{
		Email:      email,
		DeviceId:   req.DeviceId,
		DeviceType: req.DeviceType,
		PageName:   req.PageName,
		Result:     req.Result,
		Content:    req.Content,
		CreateTime: req.CreateTime,
		CreatedAt:  gtime.Now(),
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
