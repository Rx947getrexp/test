package report

import (
	"go-speed/constant"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/entity"
	"go-speed/model/response"
	"log"
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
	Date    uint `description:"数据日期, 20230101"`
	GoodsId uint `description:"商品套餐id"`
	New     uint `description:"新增用户充值数量"`
}

type DailyRechargeResponse struct {
	Total int                 `json:"total" dc:"数据总条数"`
	Items []DailyRechargeList `json:"items" dc:"数据明细"`
}

func GetDailyRechargeList(ctx *gin.Context) {
	log.Println("GetDailyRechargeList")
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
		size = 20
	}
	offset := 0
	if req.Page > 1 {
		offset = (req.Page - 1) * size
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
		items = append(items, DailyRechargeList{
			Date:    item.Date,
			GoodsId: item.GoodsId,
			New:     item.New,
		})
	}

	response.RespOk(ctx, i18n.RetMsgSuccess, DailyRechargeResponse{
		Total: total,
		Items: items,
	})
}
