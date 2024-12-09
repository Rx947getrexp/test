package report

import (
	"go-speed/constant"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/entity"
	"go-speed/model/response"
	"time"

	"github.com/gin-gonic/gin"
)

type DailyRechargeRequest struct {
	StartDate int    `form:"start_date" json:"start_date" dc:"数据日期, 20230101"`
	EndDate   int    `form:"end_date" json:"end_date" dc:"数据日期, 20230101"`
	OrderBy   string `form:"order_by" json:"order_by" dc:"排序字段，eg: id|created_time"`
	OrderType string `form:"order_type" json:"order_type" dc:"排序类型，eg: asc|desc"`
	Page      int    `form:"page" json:"page" dc:"分页查询page, 从1开始"`
	Size      int    `form:"size" json:"size" dc:"分页查询size, 最大1000"`
}

type DailyRechargeList struct {
	Date        uint    `description:"数据日期, 20230101"`
	GoodsId     uint    `description:"商品套餐id"`
	New         uint    `description:"新增用户充值数量"`
	TotalAmount float64 `description:"总充值金额"`
}

type DailyRechargeResponse struct {
	Total int                 `json:"total" dc:"数据总条数"`
	Items []DailyRechargeList `json:"items" dc:"数据明细"`
}

func GetDailyRechargeList(ctx *gin.Context) {
	// log.Println("GetDailyRechargeList")
	var (
		err error
		req = new(DailyRechargeRequest)
		// doWhere  do.TUserReportDay
		entities []entity.TUserRechargeReportDay
		total    int
	)

	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
		response.ResFail(ctx, err.Error())
		return
	}

	if req.StartDate <= 0 || req.EndDate <= 0 {
		// 获取当前时间
		currentTime := time.Now()
		if req.StartDate <= 0 {
			// 计算前第20天的日期
			DaysAgo := currentTime.AddDate(0, 0, -20)
			req.StartDate = getFormatDateToInt(DaysAgo)
		}
		if req.EndDate <= 0 {
			req.EndDate = getFormatDateToInt(currentTime)
		}
	}

	if req.StartDate > req.EndDate {
		req.StartDate, req.EndDate = req.EndDate, req.StartDate
	}

	size := req.Size
	if size < 1 || size > constant.MaxPageSize {
		size = 8
	}
	offset := 0
	if req.Page > 1 {
		offset = (req.Page - 1) * size
	}

	if req.OrderBy == "" {
		req.OrderBy = "date" // 默认按数据日期排序
	}

	if req.OrderType == "" {
		req.OrderType = "desc"
	}

	// 获取所有商品的价格信息并创建映射
	priceMap, err := getAllGoodsPrices(ctx)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("get all goods prices failed")
		response.ResFail(ctx, err.Error())
		return
	}

	model := dao.TUserRechargeReportDay.Ctx(ctx).WhereBetween("date", req.StartDate, req.EndDate)
	total, err = model.Count()

	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("count user monthly retention failed")
		response.ResFail(ctx, err.Error())
		return
	}

	err = model.Order(req.OrderBy, req.OrderType).Offset(offset).Limit(size).Scan(&entities)

	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("get user monthly retention failed")
		response.ResFail(ctx, err.Error())
		return
	}

	items := make([]DailyRechargeList, 0)
	for _, item := range entities {
		priceRub, exists := priceMap[item.GoodsId]
		if !exists {
			global.MyLogger(ctx).Warn().Msgf("price for GoodsId %d not found", item.GoodsId)
			continue // 如果找不到价格，则跳过该条目
		}
		// 计算总金额
		totalAmount := float64(item.New) * priceRub
		items = append(items, DailyRechargeList{
			Date:        item.Date,
			GoodsId:     item.GoodsId,
			New:         item.New,
			TotalAmount: totalAmount,
		})
	}

	response.RespOk(ctx, i18n.RetMsgSuccess, DailyRechargeResponse{
		Total: total,
		Items: items,
	})
}

// 新增函数 getAllGoodsPrices 来一次性查询所有商品的价格信息
func getAllGoodsPrices(ctx *gin.Context) (map[uint]float64, error) {
	priceMap := make(map[uint]float64)

	var goodsPrices []struct {
		ID       uint    `gorm:"column:id"`
		PriceRub float64 `gorm:"column:price_rub"`
	}

	err := dao.TGoods.Ctx(ctx).Scan(&goodsPrices)
	if err != nil { // 只检查 err 是否为 nil
		return nil, err
	}

	for _, gp := range goodsPrices {
		priceMap[gp.ID] = gp.PriceRub
	}

	return priceMap, nil
}
