package promotion_shop

import (
	"go-speed/constant"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/response"

	"github.com/gin-gonic/gin"
)

type PromotionShopListRequest struct {
	Id        *int64  `form:"id" json:"id" dc:"自增id"`
	Type      *string `form:"type" json:"type" dc:"商店类型，苹果：ios，安卓：android"`
	Url       *string `form:"url" json:"url" dc:"商店地址"`
	Cover     *string `form:"cover" json:"cover" dc:"商店图标"`
	Status    *int    `form:"status" json:"status" dc:"状态，1：正常，2：已软删"`
	Page      int     `form:"page" json:"page" dc:"分页查询page, 从1开始"`
	Size      int     `form:"size" json:"size" dc:"分页查询size, 最大100"`
	OrderBy   string  `form:"order_by" json:"order_by" dc:"排序字段，eg: id"`
	OrderType string  `form:"order_type" json:"order_type" dc:"排序类型，eg: asc|desc"`
}

type PromotionShopListRes struct {
	Id        int64  `form:"id" json:"id" dc:"自增id"`
	TitleCn   string `form:"title_cn" json:"title_cn" dc:"商店标题(中文)"`
	TitleEn   string `form:"title_en" json:"title_en" dc:"商店标题(英文)"`
	TitleRu   string `form:"title_ru" json:"title_ru" dc:"商店标题(俄语)"`
	Type      string `form:"type" json:"type" dc:"商店类型，苹果：ios，安卓：android"`
	Url       string `form:"url" json:"url" dc:"商店地址"`
	Cover     string `form:"cover" json:"cover" dc:"商店图标"`
	Status    int    `form:"status" json:"status" dc:"状态，1：正常，2：已软删"`
	Comment   string `form:"comment" json:"comment" dc:"备注"`
	CreatedAt string `form:"created_at" json:"created_at" dc:"创建时间"`
	UpdatedAt string `form:"updated_at" json:"updated_at" dc:"更新时间"`
}
type PromotionShopListResponse struct {
	Total int                    `json:"total" dc:"数据总条数"`
	List  []PromotionShopListRes `json:"list" dc:"数据明细"`
}

func PromotionShopList(c *gin.Context) {
	// 定义局部变量
	var (
		err      error
		req      = new(PromotionShopListRequest)
		entities []entity.TAppStore
	)

	if err = c.ShouldBind(req); err != nil {
		global.Logger.Err(err).Msg(err.Error())
		response.RespFail(c, "绑定参数失败", nil)
		return
	}

	// 分页参数处理
	if req.Page < 1 {
		req.Page = 1
	}
	size := req.Size
	if size > 100 || size < 1 {
		size = constant.MinPageSize
	}

	// 设置排序字段和排序方式
	if req.OrderBy == "" {
		req.OrderBy = "id" // 默认按id排序
	}
	if req.OrderType == "" {
		req.OrderType = "desc"
	}

	// 初始化模型查询
	model := dao.TAppStore.Ctx(c)

	where := do.TAppStore{}
	if req.Id != nil {
		where.Id = req.Id
	}
	if req.Type != nil {
		where.Type = req.Type
	}
	if req.Url != nil {
		where.Url = req.Url
	}
	if req.Cover != nil {
		where.Cover = req.Cover
	}
	if req.Status != nil {
		where.Status = req.Status
	}

	model = model.Where(where)

	// 查询总记录数
	total, err := model.Count()
	if err != nil {
		global.MyLogger(c).Err(err).Msgf(err.Error())
		response.RespFail(c, "获取数据总数失败", nil)
		return
	}

	// 执行查询，获取分页数据
	err = model.Order(req.OrderBy, req.OrderType).Page(req.Page, size).Scan(&entities)
	if err != nil {
		global.Logger.Err(err).Msg(err.Error())
		response.RespFail(c, "查询数据失败", nil)
		return
	}

	items := make([]PromotionShopListRes, 0)
	for _, entity := range entities {
		items = append(items, PromotionShopListRes{
			Id:        entity.Id,
			TitleCn:   entity.TitleCn,
			TitleEn:   entity.TitleEn,
			TitleRu:   entity.TitleRu,
			Type:      entity.Type,
			Url:       entity.Url,
			Cover:     entity.Cover,
			Status:    entity.Status,
			Comment:   entity.Comment,
			UpdatedAt: entity.UpdatedAt.String(),
			CreatedAt: entity.CreatedAt.String(),
		})
	}

	// 返回成功响应
	response.RespOk(c, i18n.RetMsgSuccess, PromotionShopListResponse{
		Total: total,
		List:  items,
	})

}
