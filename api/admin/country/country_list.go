package country

import (
	"github.com/gin-gonic/gin"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/entity"
	"go-speed/model/response"
)

type CountryListReq struct {
}

type CountryListRes struct {
	Items []Country `json:"items" dc:"国家列表"`
}

type Country struct {
	Name   string `json:"name" dc:"国家名称"`
	NameCN string `json:"name_cn" dc:"国家名称(中文)"`
}

// CountryList 查询国家名称列表
func CountryList(ctx *gin.Context) {
	var (
		err   error
		items []entity.TCountry
	)
	global.MyLogger(ctx).Info().Msgf("here -----")
	err = dao.TCountry.Ctx(ctx).
		WhereNot(dao.TCountry.Columns().Name, "").Scan(&items)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("get country failed")
		response.ResFail(ctx, err.Error())
		return
	}
	global.MyLogger(ctx).Info().Msgf("items -----: %+v", items)
	countries := make([]Country, 0)
	for _, item := range items {
		countries = append(countries, Country{Name: item.Name, NameCN: item.NameCn})
	}
	global.MyLogger(ctx).Info().Msgf("countries: %+v", countries)
	response.RespOk(ctx, i18n.RetMsgSuccess, CountryListRes{
		Items: countries,
	})
}
