package promotion_shop

import (
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/response"

	"github.com/gin-gonic/gin"
)

type PromotionShopDeleteRequest struct {
	Id int64 `form:"id" binding:"required" json:"id" dc:"自增id"`
}

func PromotionShopDelete(c *gin.Context) {
	// 定义局部变量
	var (
		err    error
		req    = new(PromotionShopDeleteRequest)
		entity *do.TAppStore
	)

	if err = c.ShouldBind(req); err != nil {
		global.Logger.Err(err).Msg(err.Error())
		response.RespFail(c, "绑定参数失败", nil)
		return
	}

	err = dao.TAppStore.Ctx(c).Where(do.TAppStore{Id: req.Id}).Scan(&entity)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("查询数据失败，error: %v", err)
		response.RespFail(c, "查询数据失败", nil)
		return
	}
	if entity == nil {
		global.MyLogger(c).Error().Msgf("数据不存在，id: %d", req.Id)
		response.RespFail(c, "数据不存在", nil)
		return
	}

	_, err = dao.TAppStore.Ctx(c).Where(do.TAppStore{Id: req.Id}).Delete()
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("delete failed，error: %v", err)
		response.RespFail(c, "删除数据失败", nil)
		return
	}

	// 返回成功
	response.RespOk(c, i18n.RetMsgSuccess, nil)
}
