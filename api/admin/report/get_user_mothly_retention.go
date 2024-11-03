package report

import (
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/request"
	"go-speed/model/response"
	"go-speed/service"

	"github.com/gin-gonic/gin"
)

func GetUserMonthlyRetention(c *gin.Context) {
	param := new(request.GetUserMonthlyRetentionRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, err.Error())
		return
	}
	total, list, err := service.QueryDeviceMonthlyRetention(c, param.Date, param.Device, param.OrderType, param.Page, param.Size)
	if err != nil {
		global.Logger.Err(err).Msg("查询出错！")
		response.ResFail(c, err.Error())
		return
	}
	items := make([]response.TUserReportMonthly, 0)

	for _, item := range list {
		items = append(items, response.TUserReportMonthly{
			Id:            item.Id,
			StatMonth:     int(item.StatMonth),
			Os:            item.Os,
			UserCount:     int(item.UserCount),
			NewUsers:      int(item.NewUsers),
			RetainedUsers: int(item.RetainedUsers),
			CreatedAt:     item.CreatedAt,
		})
	}
	resp := response.GetTUserReportMonthlyResponse{Total: total, Items: items}
	response.RespOk(c, i18n.RetMsgSuccess, resp)
}
