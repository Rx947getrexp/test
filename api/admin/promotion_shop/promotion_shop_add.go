package promotion_shop

import (
	"fmt"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/response"
	"go-speed/service"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/os/gtime"
)

type PromotionShopAddRequest struct {
	TitleCn string `form:"title_cn" binding:"required" json:"title_cn" dc:"商店标题(中文)"`
	TitleEn string `form:"title_en" binding:"required" json:"title_en" dc:"商店标题(英文)"`
	TitleRu string `form:"title_ru" binding:"required" json:"title_ru" dc:"商店标题(俄语)"`
	Type    string `form:"type" json:"type" dc:"商店类型，苹果：ios，安卓：android，华为：huawei..."`
	Url     string `form:"url" binding:"required" json:"url" dc:"商店地址"`
	Cover   string `form:"cover" binding:"required" json:"cover" dc:"商店图标"`
	Comment string `form:"comment" json:"comment" dc:"备注信息"`
}

func PromotionShopAdd(c *gin.Context) {
	// 定义局部变量
	var (
		err    error
		req    = new(PromotionShopAddRequest)
		entity *do.TAppStore
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
		global.Logger.Err(err).Msg("用户不合法！")
		response.RespFail(c, "用户不合法！", nil)
		return
	}

	err = dao.TAppStore.Ctx(c).Where(do.TAppStore{Url: req.Url}).Scan(&entity)
	if err != nil {
		global.MyLogger(c).Err(err).Msgf("查询数据失败，error: %v", err)
		response.RespFail(c, "查询数据失败", nil)
		return
	}
	if entity != nil {
		msg := fmt.Sprintf("商店url:%s 已存在，现在对应的商店中文名是: %s", req.Url, entity.TitleCn)
		global.MyLogger(c).Warn().Msgf(msg)
		response.RespFail(c, msg, nil)
		return
	}

	// 获取当前时间
	now := gtime.Now()
	// 插入新的商店记录，并获取插入数据的ID
	lastInsertId, err := dao.TAppStore.Ctx(c).Data(do.TAppStore{
		TitleCn:   req.TitleCn,
		TitleEn:   req.TitleEn,
		TitleRu:   req.TitleRu,
		Type:      req.Type,
		Url:       req.Url,
		Cover:     req.Cover,
		Status:    1,
		Author:    adminUser.Uname,
		Comment:   req.Comment,
		CreatedAt: now,
		UpdatedAt: now,
	}).InsertAndGetId()

	if err != nil || lastInsertId == 0 {
		global.MyLogger(c).Err(err).Msgf("添加数据失败，error: %v", err)
		response.RespFail(c, "添加数据失败", nil)
		return
	}

	// 返回成功响应
	response.RespOk(c, i18n.RetMsgSuccess, nil)
}
