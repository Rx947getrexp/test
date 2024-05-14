package order

import (
	"github.com/gin-gonic/gin"
	"go-speed/constant"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/response"
)

type GetOrderListReq struct {
	ChannelId string `form:"channel_id" dc:"支付渠道"`
	Status    string `form:"status" dc:"订单状态"`
	Page      int    `form:"page" dc:"分页查询page, 从1开始"`
	Size      int    `form:"size" dc:"分页查询size, 最大1000"`
}

type GetOrderListRes struct {
	Total int64      `json:"total" dc:"数据总条数"`
	Items []PayOrder `json:"items" dc:"数据明细"`
}

type PayOrder struct {
	OrderNo            string `json:"order_no"             dc:"订单号"`
	PaymentChannelId   string `json:"payment_channel_id"   dc:"支付通道ID"`
	GoodsId            int    `json:"goods_id"             dc:"套餐ID"`
	OrderAmount        string `json:"order_amount"         dc:"交易金额"`
	Currency           string `json:"currency"             dc:"交易币种"`
	Status             string `json:"status"               dc:"状态:1-正常；2-已软删"`
	OrderRealityAmount string `json:"order_reality_amount" dc:"实际交易金额"`
	PaymentProof       string `json:"payment_proof"        dc:"支付凭证地址"`
	CreatedAt          string `json:"created_at"           dc:"创建时间"`
	UpdatedAt          string `json:"updated_at"           dc:"更新时间"`
}

func GetOrderList(ctx *gin.Context) {
	var (
		err       error
		req       = new(GetOrderListReq)
		payOrders []entity.TPayOrder
		doWhere   do.TPayOrder
		total     int
	)
	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
		response.ResFail(ctx, i18n.RetMsgParamParseErr)
		return
	}

	if req.ChannelId != "" {
		doWhere.PaymentChannelId = req.ChannelId
	}
	if req.Status != "" {
		doWhere.Status = req.Status
	}
	size := req.Size
	if size < 1 || size > 1000 {
		size = 20
	}
	offset := 0
	if req.Page > 1 {
		offset = (req.Page - 1) * size
	}
	//orderBy := "create_time"
	//if req.OrderBy != "" {
	//	orderBy = req.OrderBy
	//}
	model := dao.TPayOrder.Ctx(ctx).Where(doWhere)
	//if req.StartTime != "" {
	//	model = model.WhereGTE(dao.TUserOpLog.Columns().CreatedAt, req.StartTime)
	//}
	//if req.EndTime != "" {
	//	model = model.WhereLTE(dao.TUserOpLog.Columns().CreatedAt, req.EndTime)
	//}
	total, err = model.Count()
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("count pay order failed")
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}
	err = model.Order(dao.TPayOrder.Columns().CreatedAt, constant.OrderTypeDesc).Offset(offset).Limit(size).Scan(&payOrders)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("query pay order failed")
		response.RespFail(ctx, i18n.RetMsgDBErr, nil)
		return
	}

	items := make([]PayOrder, 0)
	for _, i := range payOrders {
		items = append(items, PayOrder{
			OrderNo:            i.OrderNo,
			PaymentChannelId:   i.PaymentChannelId,
			GoodsId:            i.GoodsId,
			OrderAmount:        i.OrderAmount,
			Currency:           i.Currency,
			Status:             i.Status,
			OrderRealityAmount: i.OrderRealityAmount,
			PaymentProof:       i.PaymentProof,
			CreatedAt:          i.CreatedAt.String(),
			UpdatedAt:          i.UpdatedAt.String(),
		})
	}
	response.RespOk(ctx, i18n.RetMsgSuccess, GetOrderListRes{
		Total: int64(total),
		Items: items,
	})
}
