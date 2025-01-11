package user

import (
	"github.com/gin-gonic/gin"
	"go-speed/api/types/api"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/response"
	"go-speed/service"
	"strings"
)

func DeleteCancelledUser(ctx *gin.Context) {
	var (
		err          error
		req          = new(api.DeleteCancelledUserReq)
		nodeEntities []entity.TNode
		//userInfo     *entity.TUser
	)
	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
		response.ResFail(ctx, err.Error())
		return
	}
	global.MyLogger(ctx).Debug().Msgf("req: %+v", *req)

	email := strings.TrimSpace(req.Email)
	if email == "" {
		global.MyLogger(ctx).Warn().Msgf("param Email is empty")
		response.ResFail(ctx, "email is empty")
		return
	}

	//err = dao.TUser.Ctx(ctx).Where(do.TUser{Email: email, V2RayUuid: req.UUID}).Scan(&userInfo)
	//if err != nil {
	//	global.MyLogger(ctx).Err(err).Msgf("query user info failed")
	//	response.ResFail(ctx, err.Error())
	//	return
	//}
	//if userInfo != nil && userInfo.Status != constant.UserStatusCancelled {
	//	global.MyLogger(ctx).Err(errors.New("user is not cancelled")).Msgf("%s, status: %d", email, userInfo.Status)
	//	response.ResFail(ctx, "user is not cancelled")
	//	return
	//}

	err = dao.TNode.Ctx(ctx).Where(do.TNode{Status: 1}).Scan(&nodeEntities)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("query db failed")
		response.ResFail(ctx, err.Error())
		return
	}

	for _, node := range nodeEntities {
		err = service.DeleteUserConfigForNode(ctx, req.Email, req.UUID, node.Ip)
		if err != nil {
			global.MyLogger(ctx).Err(err).Msgf("delete user config for node failed")
			response.ResFail(ctx, err.Error())
			return
		}
	}
	response.RespOk(ctx, i18n.RetMsgSuccess, nil)
}
