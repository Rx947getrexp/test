package goods

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-speed/api/api/common"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/entity"
	"go-speed/model/response"
)

type GoodsListReq struct {
}

type GoodsListRes struct {
	Items []Goods `json:"items" dc:"套餐列表"`
}

type Goods struct {
	Id            int64   `json:"id" dc:"套餐ID，创建订单时传此参数"`
	MType         int     `json:"mt_type" dc:"会员类型：1-vip1；2-vip2"`
	Title         string  `json:"title" dc:"套餐标题"`
	TitleEn       string  `json:"title_en" dc:"套餐标题（英文）"`
	TitleRus      string  `json:"title_rus" dc:"套餐标题（俄文）"`
	Price         float64 `json:"price" dc:"单价(U)"`
	PriceUnit     string  `json:"price_unit" dc:"价格单位"`
	UsdPayPrice   float64 `json:"usd_pay_price" dc:"usd_pay价格(U)"`
	UsdPriceUnit  string  `json:"usd_price_unit" dc:"USD支付的价格单位"`
	Period        int     `json:"period" dc:"有效期（天）"`
	DiscountTitle string  `json:"discount_title" dc:"赠送天数描述"`
	DevLimit      int     `json:"dev_limit" dc:"设备限制数"`
	FlowLimit     int64   `json:"flow_limit" dc:"流量限制数；单位：字节；0-不限制"`
	IsDiscount    int     `json:"is_discount" dc:"是否优惠:1-是；2-否"`
	Low           int     `json:"low" dc:"最低赠送(天)"`
	High          int     `json:"high" dc:"最高赠送(天)"`
}

func GoodsList(ctx *gin.Context) {
	var (
		err           error
		goodsEntities []entity.TGoods
		user          *entity.TUser
	)
	user, err = common.ValidateClaims(ctx)
	if err != nil {
		return
	}
	global.MyLogger(ctx).Info().Msgf("user: %s", user.Email)

	err = dao.TGoods.Ctx(ctx).Scan(&goodsEntities)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("query goods failed")
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}

	items := make([]Goods, 0)
	for _, item := range goodsEntities {
		items = append(items, Goods{
			Id:            item.Id,
			MType:         item.MType,
			Title:         item.Title,
			TitleEn:       item.TitleEn,
			TitleRus:      item.TitleRus,
			Price:         item.Price,
			PriceUnit:     item.PriceUnit,
			UsdPayPrice:   item.UsdPayPrice,
			UsdPriceUnit:  item.UsdPriceUnit,
			Period:        item.Period,
			IsDiscount:    item.IsDiscount,
			DevLimit:      item.DevLimit,
			FlowLimit:     item.FlowLimit,
			Low:           item.Low,
			High:          item.High,
			DiscountTitle: BuildDiscountTitle(ctx, item.Low, item.High),
		})
	}
	response.RespOk(ctx, i18n.RetMsgSuccess, GoodsListRes{Items: items})
}

func BuildDiscountTitle(ctx *gin.Context, l, h int) string {
	if h <= 0 {
		return ""
	}
	lang := global.GetLang(ctx)
	switch lang {
	case i18n.LangCN:
		return fmt.Sprintf("随机赠送%d-%d天", l, h)
	case i18n.LangRU:
		return fmt.Sprintf("Случайно дарить %d-%d дней", l, h)
	default:
		return fmt.Sprintf("Randomly give away %d-%d days", l, h)
	}
}
