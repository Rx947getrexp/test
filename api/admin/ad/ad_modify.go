package ad

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/os/gtime"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/response"
	"go-speed/util"
	"strings"
)

type ADModifyReq struct {
	Name          string              `form:"name" binding:"required" json:"name" dc:"广告名称-要求唯一"`
	Advertiser    *string             `form:"advertiser" json:"advertiser" dc:"广告主，客户名称"`
	Type          *string             `form:"type" json:"type" dc:"广告类型. enum: text,image,video"`
	Url           *string             `form:"url" json:"url" dc:"广告内容地址"`
	Logo          *string             `form:"logo" json:"logo" dc:"logo地址"`
	SlotLocations *[]SlotLocationItem `form:"slot_locations" json:"slot_locations" dc:"广告位的位置以及在广告位中的排序"`
	TargetUrls    *[]TargetUrlItem    `form:"target_url" json:"target_url" dc:"跳转地址，包括：pc,ios,android"`
	Devices       *[]string           `form:"devices" json:"devices" dc:"投放设备"`
	Labels        *[]string           `form:"labels" json:"labels" dc:"广告标签"`
	ExposureTime  *uint               `form:"exposure_time" json:"exposure_time" dc:"单次曝光时间，单位秒"`
	GiftDuration  *uint               `form:"gift_duration" json:"gift_duration" dc:"观看广告后赠送时间，单位秒"`
	UserLevels    *[]int              `form:"user_levels" json:"user_levels" dc:"用户等级"`
	StartTime     *gtime.Time         `form:"start_time" json:"start_time" dc:"广告生效时间"`
	EndTime       *gtime.Time         `form:"end_time" json:"end_time" dc:"广告失效时间"`
	Status        *int                `form:"status" json:"status" dc:"状态:1-上架；2-下架"`
}

func (req *ADModifyReq) GetLocationSlots() (out []string) {
	if req.SlotLocations == nil {
		return
	}
	for _, i := range *req.SlotLocations {
		s := strings.TrimSpace(i.Location)
		if s != "" && !util.IsInArrayIgnoreCase(s, out) {
			out = append(out, s)
		}
	}
	return
}

type ADModifyRes struct{}

// ADModify 添加广告位
func ADModify(ctx *gin.Context) {
	var (
		err      error
		req      = new(ADModifyReq)
		adSlots  []entity.TAdSlot
		ad       *entity.TAd
		affected int64
	)
	if err = ctx.ShouldBind(req); err != nil {
		response.ResFail(ctx, err.Error())
		return
	}
	global.MyLogger(ctx).Info().Msgf("req: %+v", *req)

	if req.Status != nil && *req.Status != ADSlotStatusOnline && *req.Status != ADSlotStatusOffline {
		response.ResFail(ctx, "参数 Status 设置值无效")
		return
	}
	if req.GiftDuration != nil && (*req.GiftDuration < 0 || *req.GiftDuration > 24*60*60) {
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
	if ad == nil {
		response.ResFail(ctx, fmt.Sprintf("广告名称 \"%s\" 不存在", req.Name))
		return
	}

	updateDo := do.TAd{UpdatedAt: gtime.Now()}
	if req.Advertiser != nil {
		updateDo.Advertiser = *req.Advertiser
	}
	if req.Type != nil {
		updateDo.Type = *req.Type
	}
	if req.Url != nil {
		updateDo.Url = *req.Url
	}
	if req.Logo != nil {
		updateDo.Logo = *req.Logo
	}
	if req.SlotLocations != nil {
		updateDo.SlotLocations = *req.SlotLocations
	}
	if req.Devices != nil {
		updateDo.Devices = *req.Devices
	}
	if req.TargetUrls != nil {
		updateDo.TargetUrls = *req.TargetUrls
	}

	if req.Labels != nil {
		updateDo.Labels = *req.Labels
	}
	if req.ExposureTime != nil {
		updateDo.ExposureTime = *req.ExposureTime
	}
	if req.GiftDuration != nil {
		updateDo.GiftDuration = *req.GiftDuration
	}
	if req.UserLevels != nil {
		updateDo.UserLevels = *req.UserLevels
	}

	if req.StartTime != nil {
		updateDo.StartTime = req.StartTime
	}

	if req.EndTime != nil {
		updateDo.EndTime = req.EndTime
	}

	if req.Status != nil {
		updateDo.Status = *req.Status
	}

	affected, err = dao.TAd.Ctx(ctx).Where(do.TAd{Name: req.Name}).Data(updateDo).UpdateAndGetAffected()
	if err != nil {
		response.ResFail(ctx, err.Error())
		return
	}

	global.MyLogger(ctx).Debug().Msgf("affected: %d", affected)
	response.ResOk(ctx, "成功")
}
