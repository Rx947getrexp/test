package report

import (
	"go-speed/constant"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/entity"
	"go-speed/model/response"

	"github.com/gin-gonic/gin"
)

type DailyPaymentAmoutRequest struct {
	StartDate int      `form:"start_date" json:"start_date" dc:"数据日期, 20230101"`
	EndDate   int      `form:"end_date" json:"end_date" dc:"数据日期, 20230101"`
	Channel   []string `form:"channel" json:"channel" dc:"支付渠道名称"` // 保持为切片类型
	OrderBy   string   `form:"order_by" json:"order_by" dc:"排序字段，eg: id|created_time"`
	OrderType string   `form:"order_type" json:"order_type" dc:"排序类型，eg: asc|desc"`
	Page      int      `form:"page" json:"page" dc:"分页查询page, 从1开始"`
	Size      int      `form:"size" json:"size" dc:"分页查询size, 最大1000"`
}

type DailyPaymentAmout struct {
	Date      int     `description:"数据日期, 20230101"`
	Channel   string  `description:"支付渠道名称"`
	Amount    float64 `description:"支付金额"`
	CreatedAt string  `description:"数据统计时间"`
}

type DailyPaymentAmoutResponse struct {
	Total int                 `json:"total" dc:"数据总条数"`
	Items []DailyPaymentAmout `json:"items" dc:"数据明细"`
}

func GetDailyPaymentAmoutList(ctx *gin.Context) {
	var (
		err      error
		req      = new(DailyPaymentAmoutRequest)
		entities []entity.TDailyPaymentTotalByChannel
		total    int
	)

	// 绑定请求参数
	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
		response.ResFail(ctx, err.Error())
		return
	}

	// 设置分页大小，默认 8，最大 1000
	size := req.Size
	if size < 1 || size > constant.MaxPageSize {
		size = constant.MaxPageSize
	}
	offset := 0
	if req.Page > 1 {
		offset = (req.Page - 1) * size
	}

	// 设置排序字段和排序方式
	if req.OrderBy == "" {
		req.OrderBy = "date" // 默认按数据日期排序
	}
	if req.OrderType == "" {
		req.OrderType = "desc"
	}

	// 查询数据
	model := dao.TDailyPaymentTotalByChannel.Ctx(ctx).WhereBetween("date", req.StartDate, req.EndDate)

	if len(req.Channel) > 0 {
		model = model.WhereIn("channel", req.Channel)
	}

	total, err = model.Count()

	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("获取支付报表数据出错")
		response.ResFail(ctx, "获取支付报表数据出错")
		return
	}
	err = model.Order(req.OrderBy, req.OrderType).Offset(offset).Limit(size).Scan(&entities)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("获取支付报表数据失败")
		response.ResFail(ctx, "获取支付报表数据失败")
		return
	}
	items := make([]DailyPaymentAmout, 0)
	for _, entity := range entities {
		items = append(items, DailyPaymentAmout{
			Date:      entity.Date,
			Channel:   entity.Channel,
			Amount:    entity.Amount,
			CreatedAt: entity.CreatedAt.String(),
		})
	}

	response.RespOk(ctx, i18n.RetMsgSuccess, DailyPaymentAmoutResponse{
		Total: total,
		Items: items,
	})
}
