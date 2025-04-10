package promotion_shop

import (
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/response"
	"go-speed/service"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/os/gtime"
)

type PromotionShopEditRequest struct {
	Id      int64  `form:"id" binding:"required" json:"id" dc:"机器id"`
	TitleCn string `form:"title_cn" json:"title_cn" dc:"商店标题(中文)"`
	TitleEn string `form:"title_en" json:"title_en" dc:"商店标题(英文)"`
	TitleRu string `form:"title_ru" json:"title_ru" dc:"商店标题(俄语)"`
	Type    string `form:"type" json:"type" dc:"商店类型，苹果：ios，安卓：android"`
	Url     string `form:"url" json:"url" dc:"商店地址"`
	Cover   string `form:"cover" json:"cover" dc:"商店图标"`
	Status  int64  `form:"status" json:"status" dc:"状态，1：正常，2：已软删"`
	Comment string `form:"comment" json:"comment" dc:"备注信息"`
}

func PromotionShopEdit(c *gin.Context) {
	// 定义局部变量
	var (
		err error
		req = new(PromotionShopEditRequest)
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

	// 更新域名信息
	updateData := make(map[string]interface{}) // 存储要更新的字段
	time := gtime.Now()
	updateData["updated_at"] = time        // 更新时间必须更新
	updateData["author"] = adminUser.Uname // 更新此次操作人

	if req.TitleCn != "" {
		updateData["title_cn"] = req.TitleCn
	}
	if req.TitleEn != "" {
		updateData["title_en"] = req.TitleEn
	}
	if req.TitleRu != "" {
		updateData["title_ru"] = req.TitleRu
	}
	if req.Type != "" {
		updateData["type"] = req.Type
	}
	if req.Url != "" {
		updateData["url"] = req.Url
	}
	if req.Cover != "" {
		updateData["cover"] = req.Cover
	}
	if req.Status != 0 {
		updateData["status"] = req.Status
	}
	if req.Comment != "" {
		updateData["comment"] = req.Comment
	}

	// 更新数据
	_, err = dao.TAppStore.Ctx(c).Where("id", req.Id).Data(updateData).Update()
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("更新数据失败，error: %v", err)
		response.RespFail(c, "更新数据失败", nil)
		return
	}

	service.ResetPromotionShopCache()

	// 返回成功
	response.RespOk(c, i18n.RetMsgSuccess, nil)
}
