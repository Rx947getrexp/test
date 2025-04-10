package promotion_shop

import (
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/response"
	"go-speed/service"

	"github.com/gin-gonic/gin"
)

type PromotionShopDeleteRequest struct {
	Id int64 `form:"id" binding:"required" json:"id" dc:"自增id"`
}

func PromotionDnsDelete(c *gin.Context) {
	// 定义局部变量
	var (
		err error
		req = new(PromotionShopDeleteRequest)
	)

	if err = c.ShouldBind(req); err != nil {
		global.Logger.Err(err).Msg(err.Error())
		response.RespFail(c, "绑定参数失败", nil)
		return
	}

	_, err = dao.TAppStore.Ctx(c).Where("id = ?", req.Id).Delete()
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("delete failed，error: %v", err)
		response.RespFail(c, "删除数据失败", nil)
		return
	}

	service.ResetPromotionShopCache()

	// 返回成功
	response.RespOk(c, i18n.RetMsgSuccess, nil)
}
