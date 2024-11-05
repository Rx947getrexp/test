package report

import (
	"go-speed/constant"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/response"

	"github.com/gin-gonic/gin"
)

type GetUserMonthlyRetentionRequest struct {
	StatMonth int    `form:"stat_month" json:"stat_month" dc:"报表日期，eg:20240101"`
	Os        string `form:"os" json:"os" dc:"渠道设备"`
	OrderBy   string `form:"order_by" json:"order_by" dc:"排序字段，eg: id|created_time"`
	OrderType string `form:"order_type" json:"order_type" dc:"排序类型，eg: asc|desc"`
	Page      int    `form:"page" json:"page" dc:"分页查询page, 从1开始"`
	Size      int    `form:"size" json:"size" dc:"分页查询size, 最大1000"`
}

type UserReportMonthly struct {
	Id            uint64 `json:"id"             dc:"自增主键ID"`
	StatMonth     int    `json:"stat_month"     dc:"统计月份"`
	Os            string `json:"os"             dc:"设备类型"`
	UserCount     int    `json:"user_count"     dc:"用户总数"`
	NewUsers      int    `json:"new_users"      dc:"新增用户量"`
	RetainedUsers int    `json:"retained_users" dc:"次月留存"`
	CreatedAt     string `json:"created_at" dc:"记录创建时间"`
}
type GetUserReportMonthlyResponse struct {
	Total int                 `json:"total" dc:"数据总条数"`
	Items []UserReportMonthly `json:"items" dc:"数据明细"`
}

// 查询月度留存用户数据
func GetUserMonthlyRetention(ctx *gin.Context) {
	var (
		err      error
		req      = new(GetUserMonthlyRetentionRequest)
		doWhere  do.TUserReportMonthly
		entities []entity.TUserReportMonthly
		total    int
	)
	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
		response.ResFail(ctx, i18n.RetMsgParamParseErr)
		return
	}
	if req.StatMonth > 0 {
		doWhere.StatMonth = req.StatMonth
	}
	if req.Os != "" {
		doWhere.Os = req.Os
	}
	size := req.Size
	if size < 1 || size > constant.MaxPageSize {
		size = 20
	}
	offset := 0
	if req.Page > 1 {
		offset = (req.Page - 1) * size
	}

	model := dao.TUserReportMonthly.Ctx(ctx).Where(doWhere)

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
	items := make([]UserReportMonthly, 0)
	for _, item := range entities {
		items = append(items, UserReportMonthly{
			Id:            item.Id,
			StatMonth:     int(item.StatMonth),
			Os:            item.Os,
			UserCount:     int(item.UserCount),
			NewUsers:      int(item.NewUsers),
			RetainedUsers: int(item.RetainedUsers),
			CreatedAt:     item.CreatedAt.String(),
		})
	}
	response.RespOk(ctx, i18n.RetMsgSuccess, GetUserReportMonthlyResponse{
		Total: total,
		Items: items,
	})
}
