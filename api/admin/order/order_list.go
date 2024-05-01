package order

import (
	"github.com/gin-gonic/gin"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/response"
)

type PayOrderListReq struct {
	Email   string `form:"email" json:"email"`
	OrderNo string `form:"order_no" json:"order_no"`
	Page    int    `form:"page" binding:"required" json:"page"`
	Size    int    `form:"size" binding:"required" json:"size"`
}

type PayOrderListRes struct {
	Total int        `json:"total" dc:"数据总条数"`
	Items []PayOrder `json:"items" dc:"支付订单列表"`
}

type PayOrder struct {
	UserId             uint64 `json:"user_id"              description:"用户uid"`
	Email              string `json:"email"                description:"用户邮箱"`
	OrderNo            string `json:"order_no"             description:"订单号"`
	OrderAmount        string `json:"order_amount"         description:"交易金额"`
	Currency           string `json:"currency"             description:"交易币种"`
	PayTypeCode        string `json:"pay_type_code"        description:"支付类型编码"`
	Status             string `json:"status"               description:"状态"`
	ReturnStatus       string `json:"return_status"        description:"支付平台返回的结果"`
	StatusMes          string `json:"status_mes"           description:"状态描述"`
	OrderData          string `json:"order_data"           description:"创建订单时支付平台返回的信息"`
	ResultStatus       string `json:"result_status"        description:"查询结果，实际订单状态"`
	OrderRealityAmount string `json:"order_reality_amount" description:"实际交易金额"`
	CreatedAt          string `json:"created_at"           description:"创建时间"`
	UpdatedAt          string `json:"updated_at"           description:"更新时间"`
	Version            int    `json:"version"              description:"数据版本号"`
}

// PayOrderList 支付订单表
func PayOrderList(ctx *gin.Context) {
	var (
		err       error
		req       = new(PayOrderListReq)
		totalNum  int
		payOrders []entity.TPayOrder
	)
	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
		response.ResFail(ctx, i18n.RetMsgParamParseErr)
		return
	}

	totalNum, err = dao.TPayOrder.Ctx(ctx).Where(do.TPayOrder{Email: req.Email, OrderNo: req.OrderNo}).Count()
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("query payOrder failed")
		response.ResFail(ctx, err.Error())
		return
	}

	err = dao.TPayOrder.Ctx(ctx).Where(do.TPayOrder{Email: req.Email, OrderNo: req.OrderNo}).Scan(&payOrders)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("query payOrder failed")
		response.ResFail(ctx, err.Error())
		return
	}
	items := make([]PayOrder, 0)
	for _, order := range payOrders {
		items = append(items, PayOrder{
			UserId:             order.UserId,
			Email:              order.Email,
			OrderNo:            order.OrderNo,
			OrderAmount:        order.OrderAmount,
			Currency:           order.Currency,
			PayTypeCode:        order.PayTypeCode,
			Status:             order.Status,
			ReturnStatus:       order.ReturnStatus,
			StatusMes:          order.StatusMes,
			OrderData:          order.OrderData,
			ResultStatus:       order.ResultStatus,
			OrderRealityAmount: order.OrderRealityAmount,
			CreatedAt:          order.CreatedAt.String(),
			UpdatedAt:          order.UpdatedAt.String(),
			Version:            order.Version,
		})
	}
	response.RespOk(ctx, i18n.RetMsgSuccess, PayOrderListRes{
		Total: totalNum,
		Items: items,
	})
}
