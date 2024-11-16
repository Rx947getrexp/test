package ad

import (
	"github.com/gin-gonic/gin"
	"go-speed/api/common/builder"
	"go-speed/api/types"
	"go-speed/dao"
	"go-speed/i18n"
	"go-speed/model/entity"
	"go-speed/model/response"
)

type ADListReq struct{}

type ADListRes struct {
	Items []ADItem `json:"items" dc:"广告列表"`
}

type ADItem struct {
	Advertiser    string                   `json:"advertiser" dc:"广告主，客户名称"`
	Name          string                   `json:"name" dc:"广告名称-要求唯一"`
	Type          string                   `json:"type" dc:"广告类型. enum: text,image,video"`
	Url           string                   `json:"url" dc:"广告内容地址"`
	Logo          string                   `json:"logo" dc:"logo地址"`
	SlotLocations []types.SlotLocationItem `json:"slot_locations" dc:"广告位的位置以及在广告位中的排序"`
	TargetUrls    []types.TargetUrlItem    `json:"target_url" dc:"跳转地址，包括：pc,ios,android"`
	Devices       []string                 `json:"devices" dc:"投放设备"`
	Labels        []string                 `json:"labels" dc:"广告标签"`
	ExposureTime  int                      `json:"exposure_time" dc:"单次曝光时间，单位秒"`
	GiftDuration  int                      `json:"gift_duration" dc:"观看本条广告赠送时长，单位秒"`
	UserLevels    []int                    `json:"user_levels" dc:"用户等级"`
	StartTime     string                   `json:"start_time" dc:"广告生效时间"`
	EndTime       string                   `json:"end_time" dc:"广告失效时间"`
	Status        int                      `json:"status" dc:"状态:1-上架；2-下架"`
	CreatedAt     string                   `json:"created_at" dc:"创建时间"`
	UpdatedAt     string                   `json:"updated_at" dc:"更新时间"`
}

// ADList 添加广告位
func ADList(ctx *gin.Context) {
	var (
		err error
		ads []entity.TAd
	)

	err = dao.TAd.Ctx(ctx).Scan(&ads)
	if err != nil {
		response.ResFail(ctx, err.Error())
		return
	}
	items := make([]ADItem, 0)
	for _, item := range ads {
		items = append(items, ADItem{
			Advertiser:    item.Advertiser,
			Name:          item.Name,
			Type:          item.Type,
			Url:           item.Url,
			Logo:          item.Logo,
			SlotLocations: builder.BuildSlotLocations(item.SlotLocations),
			TargetUrls:    builder.BuildTargetUrls(item.TargetUrls),
			Devices:       builder.BuildStringArray(item.Devices),
			Labels:        builder.BuildStringArray(item.Labels),
			ExposureTime:  item.ExposureTime,
			GiftDuration:  item.GiftDuration,
			UserLevels:    builder.BuildIntArray(item.UserLevels),
			StartTime:     item.StartTime.String(),
			EndTime:       item.EndTime.String(),
			Status:        item.Status,
			CreatedAt:     item.CreatedAt.String(),
			UpdatedAt:     item.UpdatedAt.String(),
		})
	}
	response.RespOk(ctx, i18n.RetMsgSuccess, ADListRes{Items: items})
}
