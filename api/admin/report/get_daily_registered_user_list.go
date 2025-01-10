package report

import (
	"go-speed/constant"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/entity"
	"go-speed/model/response"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type DailyRegisteredUserRequest struct {
	StartDate int    `form:"start_date" json:"start_date" dc:"数据日期, 20230101"`
	EndDate   int    `form:"end_date" json:"end_date" dc:"数据日期, 20230101"`
	OrderBy   string `form:"order_by" json:"order_by" dc:"排序字段，eg: id|created_time"`
	OrderType string `form:"order_type" json:"order_type" dc:"排序类型，eg: asc|desc"`
	Page      int    `form:"page" json:"page" dc:"分页查询page, 从1开始"`
	Size      int    `form:"size" json:"size" dc:"分页查询size, 最大1000"`
}

type UserReportDay struct {
	Date          uint `description:"数据日期, 20230101"`
	New           uint `description:"新增用户"`
	Retained      uint `description:"日留存"`
	MonthRetained uint `description:"月留存"`
	// CreatedAt string `json:"created_at" dc:"记录创建时间"`
}

type UserReportDayResponse struct {
	Total int             `json:"total" dc:"数据总条数"`
	Items []UserReportDay `json:"items" dc:"数据明细"`
}

func getFormatDateToInt(t time.Time) int {
	formattedTime := t.Format("20060102")
	// log.Println("formattedTime1:", formattedTime)
	formattedDate, _ := strconv.Atoi(formattedTime)
	// log.Println("formattedTime2:", formattedDate)
	return formattedDate
}

func GetDailyRegisteredUser(ctx *gin.Context) {
	// 依赖表
	var (
		err error
		req = new(DailyRegisteredUserRequest)
		// doWhere  do.TUserReportDay
		entities []entity.TUserReportDay
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
		size = 20
	}
	offset := 0
	if req.Page > 1 {
		offset = (req.Page - 1) * size
	}

	model := dao.TUserReportDay.Ctx(ctx).WhereBetween("date", req.StartDate, req.EndDate)
	total, err = model.Count()

	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("count daily registered users failed")
		response.ResFail(ctx, err.Error())
		return
	}

	err = model.Order(req.OrderBy, req.OrderType).Offset(offset).Limit(size).Scan(&entities)

	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("get daily registered users failed")
		response.ResFail(ctx, err.Error())
		return
	}

	items := make([]UserReportDay, 0)
	for _, item := range entities {
		items = append(items, UserReportDay{
			Date:          item.Date,
			New:           item.New,
			Retained:      item.Retained,
			MonthRetained: item.MonthRetained,
			// CreatedAt: item.CreatedAt.Format(constant.TimeFormat),
		})
	}

	response.RespOk(ctx, i18n.RetMsgSuccess, UserReportDayResponse{
		Total: total,
		Items: items,
	})
}
