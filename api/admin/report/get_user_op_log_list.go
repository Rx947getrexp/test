package report

import (
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/response"

	"github.com/gin-gonic/gin"
)

type GetUserOpLogListReq struct {
	Email      string `form:"email" json:"email" dc:"用户email"`
	DeviceId   string `form:"dev_id" json:"dev_id" dc:"dev_id"`
	DeviceType string `form:"device_type" json:"device_type" dc:"DeviceType"`
	Result     string `form:"result" json:"result" dc:"result"`
	StartTime  string `form:"start_time" json:"start_time" dc:"数据创建的开始时间"`
	EndTime    string `form:"end_time" json:"end_time" dc:"数据创建的结束时间"`
	OrderBy    string `form:"order_by" json:"order_by" dc:"排序字段，eg: id|created_time"`
	OrderType  string `form:"order_type" json:"order_type" dc:"排序类型，eg: asc|desc"`
	Page       int    `form:"page" json:"page" dc:"分页查询page, 从1开始"`
	Size       int    `form:"size" json:"size" dc:"分页查询size, 最大1000"`
}

//PageName   string `form:"page_name" json:"page_name" dc:"PageName"`
//Result     string `form:"result" json:"result" dc:"Result"`

type GetUserOpLogListRes struct {
	Total int64       `json:"total" dc:"数据总条数"`
	Items []UserOpLog `json:"items" dc:"数据明细"`
}

type UserOpLog struct {
	Id           uint64 `json:"id"          dc:"自增id"`
	Email        string `json:"email"       dc:"用户账号"`
	DeviceId     string `json:"device_id"   dc:"设备ID"`
	DeviceType   string `json:"device_type" dc:"设备类型"`
	PageName     string `json:"page_name"   dc:"page_name"`
	Result       string `json:"result"      dc:"result"`
	Version      string `json:"version"     dc:"version"`
	InterfaceUrl string `json:"interfaceUrl"     dc:"interfaceUrl"`
	ServerCode   string `json:"serverCode"     dc:"serverCode"`
	HttpCode     string `json:"httpCode"     dc:"httpCode"`
	TraceId      string `json:"traceId"     dc:"traceId"`
	Content      string `json:"content"     dc:"content"`
	CreateTime   string `json:"create_time" dc:"提交时间"`
	CreatedAt    string `json:"created_at"  dc:"记录创建时间"`
}

// GetUserOpLogList 查询用户操作日志列表
func GetUserOpLogList(ctx *gin.Context) {
	var (
		err      error
		req      = new(GetUserOpLogListReq)
		doWhere  do.TUserOpLog
		entities []entity.TUserOpLog
		total    int
	)
	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
		response.ResFail(ctx, i18n.RetMsgParamParseErr)
		return
	}
	if req.Email != "" {
		doWhere.Email = req.Email
	}
	if req.DeviceId != "" {
		doWhere.DeviceId = req.DeviceId
	}
	if req.DeviceType != "" {
		doWhere.DeviceType = req.DeviceType
	}
	//if req.PageName != "" {
	//	doWhere.PageName = req.PageName
	//}
	if req.Result != "" {
		doWhere.Result = req.Result
	}
	size := req.Size
	if size < 1 || size > 8000 {
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
	model := dao.TUserOpLog.Ctx(ctx).Where(doWhere)
	if req.StartTime != "" {
		model = model.WhereGTE(dao.TUserOpLog.Columns().CreatedAt, req.StartTime)
	}
	if req.EndTime != "" {
		model = model.WhereLTE(dao.TUserOpLog.Columns().CreatedAt, req.EndTime)
	}
	total, err = model.Count()
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("count user op log failed")
		response.ResFail(ctx, err.Error())
		return
	}
	err = model.Order(req.OrderBy, req.OrderType).Offset(offset).Limit(size).Scan(&entities)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("count user op log failed")
		response.ResFail(ctx, err.Error())
		return
	}

	items := make([]UserOpLog, 0)
	for _, i := range entities {
		items = append(items, UserOpLog{
			Id:           i.Id,
			Email:        i.Email,
			DeviceId:     i.DeviceId,
			DeviceType:   i.DeviceType,
			PageName:     i.PageName,
			Result:       i.Result,
			Content:      i.Content,
			InterfaceUrl: i.InterfaceUrl,
			ServerCode:   i.ServerCode,
			HttpCode:     i.HttpCode,
			TraceId:      i.TraceId,
			Version:      i.Version,
			CreateTime:   i.CreateTime,
			CreatedAt:    i.CreatedAt.String(),
		})
	}
	response.RespOk(ctx, i18n.RetMsgSuccess, GetUserOpLogListRes{
		Total: int64(total),
		Items: items,
	})
}
