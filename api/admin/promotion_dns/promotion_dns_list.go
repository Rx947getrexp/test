package promotion_dns

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

type PromotionDnsListRequest struct {
	Id             int64   `form:"id" json:"id" dc:"机器id"`
	Dns            *string `form:"dns" json:"dns" dc:"机器域名"`
	Ip             *string `form:"ip" json:"ip" dc:"ip地址"`
	Status         *int    `form:"status" json:"status" dc:"状态，1：上架，2：下架"`
	MacChannel     *string `form:"mac_channel" json:"mac_channel" dc:"苹果电脑渠道"`
	WinChannel     *string `form:"win_channel" json:"win_channel" dc:"windows电脑渠道"`
	AndroidChannel *string `form:"android_channel" json:"android_channel" dc:"安卓渠道"`
	Promoter       *string `form:"promoter" json:"promoter" dc:"推广人员"`
	Page           int     `form:"page" json:"page" dc:"分页查询page, 从1开始"`
	Size           int     `form:"size" json:"size" dc:"分页查询size, 最大100"`
	OrderBy        string  `form:"order_by" json:"order_by" dc:"排序字段，eg: id"`
	OrderType      string  `form:"order_type" json:"order_type" dc:"排序类型，eg: asc|desc"`
}

type PromotionDnsListRes struct {
	Id             int64  `json:"id" dc:"机器id"`
	Dns            string `form:"dns" json:"dns" dc:"机器域名"`
	Ip             string `form:"ip" json:"ip" dc:"ip地址"`
	Status         int    `form:"status" json:"status" dc:"状态，1：上架，2：下架"`
	MacChannel     string `form:"mac_channel" json:"mac_channel" dc:"苹果电脑渠道"`
	WinChannel     string `form:"win_channel" json:"win_channel" dc:"windows电脑渠道"`
	AndroidChannel string `form:"android_channel" json:"android_channel" dc:"安卓渠道"`
	Promoter       string `form:"promoter" json:"promoter" dc:"推广人员"`
	Comment        string `form:"comment" json:"comment" dc:"备注"`
	Author         string `form:"author" json:"author" dc:"作者"`
	CreatedAt      string `form:"created_at" json:"created_at" dc:"创建时间"`
	UpdatedAt      string `form:"updated_at" json:"updated_at" dc:"更新时间"`
}
type PromotionDnsListResponse struct {
	Total int                   `json:"total" dc:"数据总条数"`
	List  []PromotionDnsListRes `json:"list" dc:"数据明细"`
}

func PromotionDnsList(c *gin.Context) {
	// 定义局部变量
	var (
		err      error
		req      = new(PromotionDnsListRequest)
		entities []entity.TPromotionDns
	)

	if err = c.ShouldBind(req); err != nil {
		global.Logger.Err(err).Msg(err.Error())
		response.RespFail(c, "绑定参数失败", nil)
		return
	}

	// 分页参数处理
	if req.Page < 1 {
		req.Page = 1
	}
	size := req.Size
	if size > 100 || size < 1 {
		size = constant.MinPageSize
	}

	// 设置排序字段和排序方式
	if req.OrderBy == "" {
		req.OrderBy = "id" // 默认按id排序
	}
	if req.OrderType == "" {
		req.OrderType = "desc"
	}

	// 初始化模型查询
	model := dao.TPromotionDns.Ctx(c)

	where := do.TPromotionDns{}

	if req.Dns != nil {
		where.Dns = req.Dns
	}

	if req.Ip != nil {
		where.Ip = req.Ip
	}

	if req.Status != nil {
		where.Status = req.Status
	}

	if req.MacChannel != nil {
		where.MacChannel = req.MacChannel
	}

	if req.WinChannel != nil {
		where.WinChannel = req.WinChannel
	}

	if req.AndroidChannel != nil {
		where.AndroidChannel = req.AndroidChannel
	}

	if req.Promoter != nil {
		where.Promoter = req.Promoter
	}

	model = model.Where(where)

	// 查询总记录数
	total, err := model.Count()
	if err != nil {
		global.MyLogger(c).Err(err).Msgf(err.Error())
		response.RespFail(c, "获取数据总数失败", nil)
		return
	}

	// 执行查询，获取分页数据
	err = model.Order(req.OrderBy, req.OrderType).Page(req.Page, size).Scan(&entities)
	if err != nil {
		global.Logger.Err(err).Msg(err.Error())
		response.RespFail(c, "查询数据失败", nil)
		return
	}

	items := make([]PromotionDnsListRes, 0)
	for _, entity := range entities {
		items = append(items, PromotionDnsListRes{
			Id:             entity.Id,
			Dns:            entity.Dns,
			Ip:             entity.Ip,
			Status:         entity.Status,
			MacChannel:     entity.MacChannel,
			WinChannel:     entity.WinChannel,
			AndroidChannel: entity.AndroidChannel,
			Promoter:       entity.Promoter,
			Comment:        entity.Comment,
			Author:         entity.Author,
			UpdatedAt:      entity.UpdatedAt.String(),
			CreatedAt:      entity.CreatedAt.String(),
		})
	}

	// 返回成功响应
	response.RespOk(c, i18n.RetMsgSuccess, PromotionDnsListResponse{
		Total: total,
		List:  items,
	})
}
