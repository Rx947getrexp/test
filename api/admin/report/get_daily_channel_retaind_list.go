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

type DailyChannelRetaindRequest struct {
	StartDate int    `form:"start_date" json:"start_date" dc:"数据日期, 20230101"`
	EndDate   int    `form:"end_date" json:"end_date" dc:"数据日期, 20230101"`
	Date      int    `form:"date" json:"date" dc:"数据日期, 20230101"`
	Channel   string `form:"channel" json:"ad_name" dc:"广告名称"`
	OrderBy   string `form:"order_by" json:"order_by" dc:"排序字段，eg: id|created_time"`
	OrderType string `form:"order_type" json:"order_type" dc:"排序类型，eg: asc|desc"`
	Page      int    `form:"page" json:"page" dc:"分页查询page, 从1开始"`
	Size      int    `form:"size" json:"size" dc:"分页查询size, 最大1000"`
}

type DailyChannelRetaind struct {
	Date          int    `description:"数据日期, 20230101"`
	Channel       string `description:"注册渠道名称"`
	NewUsers      int    `description:"新增用户数量"`
	Day2Retained  int    `description:"新增用户次日留存数量"`
	Day7Retained  int    `description:"新增用户7日留存数量"`
	Day15Retained int    `description:"新增用户15日留存数量"`
	Day30Retained int    `description:"新增用户30日留存数量"`
	CreatedAt     string `description:"数据统计时间"`
}

type DailyChannelRetaindResponse struct {
	Total int                   `json:"total" dc:"数据总条数"`
	Items []DailyChannelRetaind `json:"items" dc:"数据明细"`
}

func GetDailyChannelRetaindList(ctx *gin.Context) {
	var (
		err      error
		req      = new(DailyChannelRetaindRequest)
		entities []entity.TChannelRetaindDaily
		total    int
	)
	// 绑定请求参数
	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
		response.ResFail(ctx, err.Error())
		return
	}
	if req.Date == 0 {
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
	// 查询数据
	model := dao.TChannelRetaindDaily.Ctx(ctx)

	if req.Date == 0 {
		model = model.WhereBetween("date", req.StartDate, req.EndDate)
	} else {
		model = model.Where("date", req.Date)
	}

	if req.Channel != "" {
		model = model.Where("channel", req.Channel)
	}

	total, err = model.Count()

	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("获取报表数据出错")
		response.ResFail(ctx, "获取报表数据出错")
		return
	}
	err = model.Order(req.OrderBy, req.OrderType).Offset(offset).Limit(size).Scan(&entities)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("获取报表数据失败")
		response.ResFail(ctx, "获取报表数据失败")
		return
	}

	items := make([]DailyChannelRetaind, 0)
	for _, entity := range entities {
		items = append(items, DailyChannelRetaind{
			Date:          entity.Date,
			Channel:       entity.Channel,
			NewUsers:      entity.NewUsers,
			Day2Retained:  entity.Day2Retained,
			Day7Retained:  entity.Day7Retained,
			Day15Retained: entity.Day15Retained,
			Day30Retained: entity.Day30Retained,
			CreatedAt:     entity.CreatedAt.String(),
		})
	}

	response.RespOk(ctx, i18n.RetMsgSuccess, DailyChannelRetaindResponse{
		Total: total,
		Items: items,
	})
}
