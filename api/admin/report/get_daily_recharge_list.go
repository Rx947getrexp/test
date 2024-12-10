package report

import (
	"go-speed/constant"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/entity"
	"go-speed/model/response"
	"sort"
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

// 每日充值汇总信息
type DailyRechargeSummary struct {
	Date        uint    `json:"date"`
	New         uint    `json:"new"`
	TotalAmount float64 `json:"total_amount"`
}

func GetDailyRechargeList(ctx *gin.Context) {
	var (
		err      error
		req      = new(DailyRechargeRequest)
		entities []entity.TUserRechargeReportDay
	)

	// 绑定请求参数
	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
		response.ResFail(ctx, err.Error())
		return
	}

	// 如果未指定起始日期和结束日期，则使用默认日期
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

	// 调整日期顺序，确保 StartDate 小于 EndDate
	if req.StartDate > req.EndDate {
		req.StartDate, req.EndDate = req.EndDate, req.StartDate
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

	// 获取所有商品的价格信息并创建映射
	priceMap, err := getAllGoodsPrices(ctx)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("获取商品价格失败")
		response.ResFail(ctx, err.Error())
		return
	}

	// 查询数据
	model := dao.TUserRechargeReportDay.Ctx(ctx).WhereBetween("date", req.StartDate, req.EndDate)

	err = model.Order(req.OrderBy, req.OrderType).Offset(offset).Limit(size).Scan(&entities)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("获取用户月度充值记录失败")
		response.ResFail(ctx, err.Error())
		return
	}

	// 按日期分组并计算每个日期的总计
	dateGroupedData := make(map[uint]DailyRechargeList)

	for _, item := range entities {
		priceRub, exists := priceMap[item.GoodsId]
		if !exists {
			global.MyLogger(ctx).Warn().Msgf("商品ID %d 的价格未找到", item.GoodsId)
			continue // 如果找不到价格，则跳过该条目
		}

		// 计算总金额
		totalAmount := float64(item.New) * priceRub

		// 尝试获取已存在的记录，如果不存在则创建新记录
		dailyStat, ok := dateGroupedData[item.Date]
		if !ok {
			dailyStat = DailyRechargeList{
				Date: item.Date,
			}
		}

		// 累加 New 和 TotalAmount
		dailyStat.New += item.New
		dailyStat.TotalAmount += totalAmount

		// 更新分组数据
		dateGroupedData[item.Date] = dailyStat
	}

	// 构造按日期分组的响应数据
	groupedItems := make([]DailyRechargeSummary, 0)

	// 按日期顺序遍历分组
	for _, item := range dateGroupedData {
		groupedItems = append(groupedItems, DailyRechargeSummary{
			Date:        item.Date,
			New:         item.New,
			TotalAmount: item.TotalAmount,
		})
	}

	// 对最终的结果按照日期排序（假设是升序）
	sort.Slice(groupedItems, func(i, j int) bool {
		return groupedItems[i].Date < groupedItems[j].Date
	})

	// 使用groupedItems的长度作为total值
	total := len(groupedItems)

	// 返回成功响应
	response.RespOk(ctx, i18n.RetMsgSuccess, map[string]interface{}{
		"total": total,
		"items": groupedItems,
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
