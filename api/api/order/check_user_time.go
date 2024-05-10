package order

import (
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/response"
	"time"

	"github.com/gin-gonic/gin"
)

type CheckUserTimeReq struct {
	UserId uint64 `form:"user_id" binding:"required" json:"user_id" dc:"用户ID"`
}

func CheckUserTime(ctx *gin.Context) {
	var (
		err error
		req = new(CheckUserTimeReq)
	)
	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
		response.RespFail(ctx, i18n.RetMsgParamParseErr, nil)
		return
	}
	global.MyLogger(ctx).Info().Msgf("request: %+v", *req)

	// 查询用户信息
	user := new(do.TUser)
	_, err = dao.TUser.Ctx(ctx).Where("user_id = ?", req.UserId).One(&user)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("数据库查询出错, userId: %d", req.UserId)
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}
	// 获取当前时间戳
	now := time.Now().Unix()
	expiredTime := user.ExpiredTime.(int64)
	// 计算距离到期的时间间隔
	expireDuration := expiredTime - now
	// 如果距离到期的时间不足三天，则发送到期提醒
	if expireDuration <= 3*24*3600 {
		global.MyLogger(ctx).Info().Msgf("用户 %d 的会员即将到期，请及时续费", req.UserId)
		response.RespOk(ctx, i18n.RetMsgMemberExpirationReminder, nil)
		return
	}
	response.RespOk(ctx, i18n.RetMsgSuccess, nil)
}
