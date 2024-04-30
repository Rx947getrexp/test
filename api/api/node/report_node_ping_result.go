package node

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/request"
	"go-speed/model/response"
)

func ReportNodePingResult(c *gin.Context) {
	param := new(request.ReportNodePingResultRequest)
	if err := c.ShouldBind(param); err != nil {
		global.MyLogger(c).Err(err).Msgf("绑定参数失败")
		response.RespFail(c, i18n.RetMsgParamParseErr, nil)
		return
	}
	global.MyLogger(c).Info().Msgf("param: %+v", *param)
	var (
		err  error
		user *entity.TUser
	)
	err = dao.TUser.Ctx(c).Where(do.TUser{Id: param.UserId}).Scan(&user)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("get user failed, param: %+v", *param)
		response.RespFail(c, i18n.RetMsgDBErr, nil)
		return
	}
	if user == nil {
		global.MyLogger(c).Error().Msgf("账号不存在！userId: %d， param: %+v", param.UserId, *param)
		response.RespFail(c, i18n.RetMsgAccountNotExist, nil)
		return
	}

	var inArray = g.Slice{}
	for _, item := range param.Items {
		inArray = append(inArray, do.TUserPing{
			Email:     user.Email,
			Host:      item.Ip,
			Code:      item.Code,
			Cost:      item.Cost,
			Time:      param.ReportTime,
			CreatedAt: gtime.Now(),
		})
	}
	err = dao.TUserPing.Transaction(c, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.TUserPing.Ctx(ctx).Data(inArray).Insert()
		return err
	})
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("insert ping result failed")
		response.RespFail(c, i18n.RetMsgDBErr, nil)
		return
	}
	response.RespOk(c, i18n.RetMsgSuccess, nil)
}
