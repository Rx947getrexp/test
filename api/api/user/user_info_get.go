package user

import (
	"github.com/gin-gonic/gin"
	"go-speed/api/api/common"
	"go-speed/api/types/api"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/entity"
	"go-speed/model/response"
)

func DescribeUserInfo(ctx *gin.Context) {
	var (
		err  error
		req  = new(api.DescribeUserInfoReq)
		user *entity.TUser
	)
	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
		response.ResFail(ctx, err.Error())
		return
	}

	if req.UserId > 0 {
		user, err = common.CheckUserByUserId(ctx, req.UserId)
		if err != nil {
			return
		}
	} else if req.Email != "" {
		user, err = common.GetUserByEmail(ctx, req.Email)
		if err != nil {
			return
		}
	} else {
		response.ResFail(ctx, "无效参数")
		return
	}
	resp := api.DescribeUserInfoRes{
		Id:               user.Id,
		Uname:            user.Uname,
		Email:            user.Email,
		Phone:            user.Phone,
		Level:            user.Level,
		ExpiredTime:      user.ExpiredTime,
		V2RayUuid:        user.V2RayUuid,
		V2RayTag:         user.V2RayTag,
		Channel:          user.Channel,
		ChannelId:        user.ChannelId,
		Status:           user.Status,
		CreatedAt:        user.CreatedAt.String(),
		UpdatedAt:        user.UpdatedAt.String(),
		Comment:          user.Comment,
		ClientId:         user.ClientId,
		LastLoginIp:      user.LastLoginIp,
		LastLoginCountry: user.LastLoginCountry,
		PreferredCountry: user.PreferredCountry,
		Version:          user.Version,
	}
	response.RespOk(ctx, i18n.RetMsgSuccess, resp)
}
