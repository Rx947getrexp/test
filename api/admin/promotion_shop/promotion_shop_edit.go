package promotion_shop

import (
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/response"
	"go-speed/service"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/os/gtime"
)

type PromotionShopEditRequest struct {
	Id      int64   `form:"id" binding:"required" json:"id" dc:"自增id"`
	TitleCn *string `form:"title_cn" json:"title_cn" dc:"商店标题(中文)"`
	TitleEn *string `form:"title_en" json:"title_en" dc:"商店标题(英文)"`
	TitleRu *string `form:"title_ru" json:"title_ru" dc:"商店标题(俄语)"`
	Type    *string `form:"type" json:"type" dc:"商店类型，苹果：ios，安卓：android"`
	Url     *string `form:"url" json:"url" dc:"商店地址"`
	Cover   *string `form:"cover" json:"cover" dc:"商店图标"`
	Status  *int64  `form:"status" json:"status" dc:"状态，1：正常，2：已软删"`
	Comment *string `form:"comment" json:"comment" dc:"备注信息"`
}

func PromotionShopEdit(c *gin.Context) {
	// 定义局部变量
	var (
		err    error
		req    = new(PromotionShopEditRequest)
		entity *entity.TAppStore
	)

	if err = c.ShouldBind(req); err != nil {
		global.Logger.Err(err).Msg(err.Error())
		response.RespFail(c, "绑定参数失败", nil)
		return
	}

	// 从上下文中获取用户信息
	claims := c.MustGet("claims").(*service.CustomClaims)
	// 根据用户信息获取管理员用户对象
	adminUser, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msgf("Illegal user, err: %v", err)
		response.RespFail(c, "用户不合法！", nil)
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

	// 更新域名信息
	time := gtime.Now()
	// 存储要更新的字段
	updateDo := do.TAppStore{
		UpdatedAt: time,            // 更新时间必须更新
		Author:    adminUser.Uname, // 更新此次操作人
	}

	if req.TitleCn != nil {
		updateDo.TitleCn = req.TitleCn
	}
	if req.TitleEn != nil {
		updateDo.TitleEn = req.TitleEn
	}
	if req.TitleRu != nil {
		updateDo.TitleRu = req.TitleRu
	}
	if req.Type != nil {
		updateDo.Type = req.Type
	}
	if req.Url != nil {
		updateDo.Url = req.Url
	}
	if req.Cover != nil {
		updateDo.Cover = req.Cover
	}
	if req.Status != nil {
		updateDo.Status = req.Status
	}
	if req.Comment != nil {
		updateDo.Comment = req.Comment
	}

	// 更新数据
	_, err = dao.TAppStore.Ctx(c).Where(do.TAppStore{Id: req.Id}).Data(updateDo).Update()
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("更新数据失败，error: %v", err)
		response.RespFail(c, "更新数据失败", nil)
		return
	}

	// 返回成功
	response.RespOk(c, i18n.RetMsgSuccess, nil)
}
