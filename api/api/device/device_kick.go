package device

import (
	"github.com/gin-gonic/gin"
	"go-speed/api/api/common"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/response"
)

type KickDeviceReq struct {
	ClientId string `form:"client_id" json:"client_id" dc:"设备号"`
}

type KickDeviceRes struct {
}

func KickDevice(ctx *gin.Context) {
	var (
		err    error
		req    = new(KickDeviceReq)
		user   *entity.TUser
		device *entity.TUserDevice
	)

	// 绑定请求参数
	if err = ctx.ShouldBind(req); err != nil {
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

	err = dao.TUserDevice.Ctx(ctx).Where(do.TUserDevice{
		UserId:   user.Id,
		ClientId: req.ClientId,
		Kicked:   0,
	}).Scan(&device)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("query user devices failed")
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}

	if device != nil {
		_, err = dao.TUserDevice.Ctx(ctx).
			Where(do.TUserDevice{Id: device.Id, UserId: user.Id, ClientId: req.ClientId}).
			Update(do.TUserDevice{Kicked: 1})
		if err != nil {
			global.MyLogger(ctx).Err(err).Msgf("update user devices kicked failed")
			response.RespFail(ctx, i18n.RetMsgDBErr, nil)
			return
		}
		global.MyLogger(ctx).Info().Msgf("client_id: %s, kicked success", req.ClientId)
	}

	response.RespOk(ctx, i18n.RetMsgSuccess, KickDeviceRes{})
}
