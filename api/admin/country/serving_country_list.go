package country

import (
	"github.com/gin-gonic/gin"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/entity"
	"go-speed/model/response"
)

type ServingCountryListReq struct {
}

type ServingCountryListRes struct {
	Items []ServingCountry `json:"items" dc:"在役的国家列表"`
}

type ServingCountry struct {
	Name        string `json:"name" dc:"国家名称，不可以修改，后端当ID用"`
	NameDisplay string `json:"name_display" dc:"用于在用户侧展示的国家名称"`
	LogoLink    string `json:"logo_link" dc:"国家图片地址"`
	PingUrl     string `json:"ping_url" dc:"ping的地址，供前端使用"`
	IsRecommend uint   `json:"is_recommend" dc:"是否为推荐的国家，0:否，1：是"`
	Weight      uint   `json:"weight" dc:"推荐权重,权重越大的国家展示在越靠前"`
	Status      uint   `json:"status" dc:"状态。1-已上架；2-已下架"`
	CreatedAt   string `json:"created_at" dc:"创建时间"`
	UpdatedAt   string `json:"updated_at" dc:"更新时间"`
}

// ServingCountryList 查询国家列表
func ServingCountryList(ctx *gin.Context) {
	var (
		err   error
		items []entity.TServingCountry
	)
	err = dao.TServingCountry.Ctx(ctx).Scan(&items)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("get serving country failed")
		response.ResFail(ctx, err.Error())
		return
	}
	countries := make([]ServingCountry, 0)
	for _, item := range items {
		countries = append(countries, ServingCountry{
			Name:        item.Name,
			NameDisplay: item.Display,
			LogoLink:    item.LogoLink,
			PingUrl:     item.PingUrl,
			IsRecommend: uint(item.IsRecommend),
			Weight:      uint(item.Weight),
			Status:      uint(item.Status),
			CreatedAt:   item.CreatedAt.String(),
			UpdatedAt:   item.UpdatedAt.String(),
		})
	}
	response.RespOk(ctx, i18n.RetMsgSuccess, ServingCountryListRes{
		Items: countries,
	})
}
