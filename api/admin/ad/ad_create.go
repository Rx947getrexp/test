package ad

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/response"
	"go-speed/util"
	"strings"
)

type ADCreateReq struct {
	Advertiser    string             `form:"advertiser" json:"advertiser" dc:"广告主，客户名称"`
	Name          string             `form:"name" binding:"required" json:"name" dc:"广告名称-要求唯一"`
	Type          string             `form:"type" json:"type" dc:"广告类型. enum: text,image,video"`
	Url           string             `form:"url" json:"url" dc:"广告内容地址"`
	Logo          string             `form:"logo" json:"logo" dc:"logo地址"`
	SlotLocations []SlotLocationItem `form:"slot_locations" json:"slot_locations" dc:"广告位的位置以及在广告位中的排序"`
	TargetUrls    []TargetUrlItem    `form:"target_url" json:"target_url" dc:"跳转地址，包括：pc,ios,android"`
	Devices       []string           `form:"devices" json:"devices" dc:"投放设备"`
	Labels        []string           `form:"labels" json:"labels" dc:"广告标签"`
	ExposureTime  int                `form:"exposure_time" json:"exposure_time" dc:"单次曝光时间，单位秒"`
	GiftDuration  uint               `form:"gift_duration" json:"gift_duration" dc:"观看广告后赠送时间，单位秒"`
	UserLevels    []int              `form:"user_levels" json:"user_levels" dc:"用户等级"`
	StartTime     string             `form:"start_time" json:"start_time" dc:"广告生效时间"`
	EndTime       string             `form:"end_time" json:"end_time" dc:"广告失效时间"`
	Status        int                `form:"status" json:"status" dc:"状态:1-上架；2-下架"`
}

type SlotLocationItem struct {
	Location string `form:"location" json:"location" dc:"广告位的位置"`
	Sort     int    `form:"sort" json:"sort" dc:"在广告位置中的排序"`
}

type TargetUrlItem struct {
	Channel string `form:"channel" json:"channel" dc:"渠道，enum: pc, android, ios"`
	Url     string `form:"url" json:"url" dc:"跳转地址"`
}

func (req *ADCreateReq) GetLocationSlots() (out []string) {
	for _, i := range req.SlotLocations {
		s := strings.TrimSpace(i.Location)
		if s != "" && !util.IsInArrayIgnoreCase(s, out) {
			out = append(out, s)
		}
	}
	return
}

type ADCreateRes struct{}

// ADCreate 添加广告位
func ADCreate(ctx *gin.Context) {
	var (
		err          error
		req          = new(ADCreateReq)
		adSlots      []entity.TAdSlot
		ad           *entity.TAd
		lastInsertId int64
	)
	if err = ctx.ShouldBind(req); err != nil {
		response.ResFail(ctx, err.Error())
		return
	}
	global.MyLogger(ctx).Info().Msgf("req: %+v", *req)

	if req.Status == 0 {
		req.Status = ADSlotStatusOffline
	}
	if req.Status != ADSlotStatusOnline && req.Status != ADSlotStatusOffline {
		response.ResFail(ctx, "参数 Status 设置值无效")
		return
	}
	if err = validateStringTime(req.StartTime); err != nil {
		response.ResFail(ctx, "参数 StartTime 设置值无效")
		return
	}
	if err = validateStringTime(req.EndTime); err != nil {
		response.ResFail(ctx, "参数 EndTime 设置值无效")
		return
	}
	if req.GiftDuration < 0 || req.GiftDuration > 24*60*60 {
		response.ResFail(ctx, "参数 GiftDuration 设置值无效")
		return
	}

	err = dao.TAdSlot.Ctx(ctx).WhereIn(dao.TAdSlot.Columns().Location, req.GetLocationSlots()).Scan(&adSlots)
	if err != nil {
		response.ResFail(ctx, err.Error())
		return
	}
	if len(adSlots) != len(req.GetLocationSlots()) {
		response.ResFail(ctx, "广告位无效")
		return
	}

	err = dao.TAd.Ctx(ctx).Where(do.TAd{Name: req.Name}).Scan(&ad)
	if err != nil {
		response.ResFail(ctx, err.Error())
		return
	}
	if ad != nil {
		response.ResFail(ctx, fmt.Sprintf("广告名称 \"%s\" 已经存在", req.Name))
		return
	}

	lastInsertId, err = dao.TAd.Ctx(ctx).Data(do.TAd{
		Advertiser:    req.Advertiser,
		Name:          req.Name,
		Type:          req.Type,
		Url:           req.Url,
		Logo:          req.Logo,
		SlotLocations: req.SlotLocations,
		Devices:       req.Devices,
		TargetUrls:    req.TargetUrls,
		Labels:        req.Labels,
		ExposureTime:  req.ExposureTime,
		GiftDuration:  req.GiftDuration,
		UserLevels:    req.UserLevels,
		StartTime:     gtime.NewFromStr(req.StartTime),
		EndTime:       gtime.NewFromStr(req.EndTime),
		Status:        req.Status,
		CreatedAt:     gtime.Now(),
		UpdatedAt:     gtime.Now(),
	}).InsertAndGetId()
	if err != nil {
		response.ResFail(ctx, err.Error())
		return
	}
	global.MyLogger(ctx).Debug().Msgf("lastInsertId: %d", lastInsertId)
	response.ResOk(ctx, "成功")
}

func validateStringTime(in string) error {
	in = strings.TrimSpace(in)
	if len(in) == 0 {
		return nil
	}
	if gtime.NewFromStr(in) == nil {
		return gerror.Newf("%s invalid", in)
	}
	return nil
}
