package goods

import (
	"github.com/gin-gonic/gin"
	"go-speed/api/api/common"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/entity"
	"go-speed/model/response"
	"go-speed/service"
)

type GoodsListReq struct {
}

type GoodsListRes struct {
	Items []Goods `json:"items" dc:"套餐列表"`
}

type Goods struct {
	Id          int64   `json:"id" dc:"套餐ID，创建订单时传此参数"`
	MType       int     `json:"mt_type" dc:"会员类型：1-vip1；2-vip2"`
	Title       string  `json:"title" dc:"套餐标题"`
	TitleEn     string  `json:"title_en" dc:"套餐标题（英文）"`
	TitleRus    string  `json:"title_rus" dc:"套餐标题（俄文）"`
	Price       float64 `json:"price" dc:"单价(U)"`
	UsdPayPrice float64 `json:"usd_pay_price" dc:"usd_pay价格(U)"`
	Period      int     `json:"period" dc:"有效期（天）"`
	DevLimit    int     `json:"dev_limit" dc:"设备限制数"`
	FlowLimit   int64   `json:"flow_limit" dc:"流量限制数；单位：字节；0-不限制"`
	IsDiscount  int     `json:"is_discount" dc:"是否优惠:1-是；2-否"`
	Low         int     `json:"low" dc:"最低赠送(天)"`
	High        int     `json:"high" dc:"最高赠送(天)"`
}

func GoodsList(ctx *gin.Context) {
	var (
		err           error
		goodsEntities []entity.TGoods
	)

	claims := ctx.MustGet("claims").(*service.CustomClaims)
	_, err = common.CheckUserByUserId(ctx, uint64(claims.UserId))
	if err != nil {
		return
	}
	global.MyLogger(ctx).Info().Msgf("OrderNo: %s", claims.UserId)

	err = dao.TGoods.Ctx(ctx).Scan(&goodsEntities)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("query goods failed")
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}

	items := make([]Goods, 0)
	for _, item := range goodsEntities {
		items = append(items, Goods{
			Id:          item.Id,
			MType:       item.MType,
			Title:       item.Title,
			TitleEn:     item.TitleEn,
			TitleRus:    item.TitleRus,
			Price:       item.Price,
			UsdPayPrice: item.UsdPayPrice,
			Period:      item.Period,
			IsDiscount:  item.IsDiscount,
			DevLimit:    item.DevLimit,
			FlowLimit:   item.FlowLimit,
			Low:         item.Low,
			High:        item.High,
		})
	}
	response.RespOk(ctx, i18n.RetMsgSuccess, GoodsListRes{Items: items})
}
