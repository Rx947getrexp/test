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

type DeviceListReq struct {
}

type DeviceListRes struct {
	Items []UserDevice `json:"items" dc:"设备列表"`
}

type UserDevice struct {
	ClientId  string `json:"client_id" dc:"用户设备号"`
	OS        string `json:"os" dc:"操作系统"`
	CreatedAt string `json:"created_at" dc:"设备号记录时间"`
	UpdatedAt string `json:"updated_at" dc:"最后更新时间"`
}

func DeviceList(ctx *gin.Context) {
	var (
		err            error
		deviceEntities []entity.TUserDevice
		user           *entity.TUser
	)
	user, err = common.ValidateClaims(ctx)
	if err != nil {
		return
	}
	global.MyLogger(ctx).Info().Msgf("user: %s", user.Email)

	err = dao.TUserDevice.Ctx(ctx).Where(do.TUserDevice{
		UserId: user.Id,
		Kicked: 0,
	}).Scan(&deviceEntities)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("query user devices failed")
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}

	items := make([]UserDevice, 0)
	for _, item := range deviceEntities {
		items = append(items, UserDevice{
			ClientId:  item.ClientId,
			OS:        item.Os,
			CreatedAt: item.CreatedAt.String(),
			UpdatedAt: item.UpdatedAt.String(),
		})
	}
	response.RespOk(ctx, i18n.RetMsgSuccess, DeviceListRes{Items: items})
}
