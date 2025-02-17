package report

import (
	"fmt"
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

type MonthlyChannelrechargeRenewalsRequest struct {
	Month     int    `form:"month" json:"date" dc:"数据月份, 202301"`
	Channel   string `form:"channel" json:"channel" dc:"注册渠道id名称"`
	OrderBy   string `form:"order_by" json:"order_by" dc:"排序字段，eg: id|created_time"`
	OrderType string `form:"order_type" json:"order_type" dc:"排序类型，eg: asc|desc"`
	Page      int    `form:"page" json:"page" dc:"分页查询page, 从1开始"`
	Size      int    `form:"size" json:"size" dc:"分页查询size, 最大1000"`
}

type MonthlyChannelrechargeRenewals struct {
	Month          int    `description:"统计数据月份，整数类型，格式为 YYYYMM，例如202501"`
	Channel        string `description:"渠道id"`
	RechargeUsers  int    `description:"付费用户数量"`
	RechargeAmount string `description:"付费用户充值总金额"`
	Retained       int    `description:"充值用户次月留存数量"`
	RenewalsUsers  int    `description:"次月续费人数"`
	RenewalsAmount string `description:"次月续费充值总金额"`
	CreatedAt      string `description:"记录创建时间，默认值为当前时间"`
}

type MonthlyChannelrechargeRenewalsResponse struct {
	Total int                              `json:"total" dc:"数据总条数"`
	Items []MonthlyChannelrechargeRenewals `json:"items" dc:"数据明细"`
}

func GetMonthlyChannelrechargeRenewalsList(ctx *gin.Context) {
	var (
		err      error
		req      = new(MonthlyChannelrechargeRenewalsRequest)
		entities []entity.TChannelRechargeRenewalsMonthly
		total    int
	) // 绑定请求参数
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

	if req.Month == 0 {
		req.Month = getLastMonthInt()
	}

	// 设置排序字段和排序方式
	if req.OrderBy == "" {
		req.OrderBy = "month" // 默认按数据月份排序
	}
	if req.OrderType == "" {
		req.OrderType = "desc"
	}
	// 查询数据
	model := dao.TChannelRechargeRenewalsMonthly.Ctx(ctx)
	if req.Channel != "" {
		model = model.Where("channel", req.Channel)
	}

	model = model.Where("month", req.Month)

	total, err = model.Count()

	err = model.Order(req.OrderBy, req.OrderType).Offset(offset).Limit(size).Scan(&entities)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("获取报表数据失败")
		response.ResFail(ctx, "获取报表数据失败")
		return
	}

	items := make([]MonthlyChannelrechargeRenewals, 0)
	for _, entity := range entities {
		items = append(items, MonthlyChannelrechargeRenewals{
			Month:          entity.Month,
			Channel:        entity.Channel,
			RechargeUsers:  entity.RechargeUsers,
			RechargeAmount: entity.RechargeAmount,
			Retained:       entity.Retained,
			RenewalsUsers:  entity.RenewalsUsers,
			RenewalsAmount: entity.RenewalsAmount,
			CreatedAt:      entity.CreatedAt.String(),
		})
	}

	response.RespOk(ctx, i18n.RetMsgSuccess, MonthlyChannelrechargeRenewalsResponse{
		Total: total,
		Items: items,
	})
}

func getLastMonthInt() int {
	// 获取当前时间
	now := time.Now()

	// 获取上个月的时间
	lastMonth := now.AddDate(0, -1, 0)

	// 格式化时间为 "YYYYMM" 格式
	formattedDate := lastMonth.Format("200601")

	// 将格式化后的字符串转换为整数
	result, err := strconv.Atoi(formattedDate)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return -1 // 返回一个错误码或处理错误
	}

	return result
}
