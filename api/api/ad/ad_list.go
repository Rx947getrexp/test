package ad

import (
	"github.com/gin-gonic/gin"
	"go-speed/api/common/builder"
	"go-speed/api/types"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/do"
	"go-speed/model/entity"
	"go-speed/model/response"
	"go-speed/util"
	"math/rand"
	"time"
)

const (
	ADStatusOnline = 1

	QueryTypeWeightRand = "weight-rand"
)

type ADListReq struct {
	Locations []string `form:"locations" json:"locations" dc:"广告位的位置，相当于ID"`
	Devices   []string `form:"devices" json:"devices" dc:"投放设备"`
	QueryType string   `form:"query_type" json:"query_type" dc:"查询类型。weight-rand: 按权重随机"`
}

type ADListRes struct {
	Items []ADItem `json:"items" dc:"广告列表"`
}

type ADItem struct {
	Name          string                   `json:"name" dc:"广告名称-要求唯一"`
	Type          string                   `json:"type" dc:"广告类型. enum: text,image,video"`
	Url           string                   `json:"url" dc:"广告内容地址"`
	Logo          string                   `json:"logo" dc:"logo地址"`
	SlotLocations []types.SlotLocationItem `json:"slot_locations" dc:"广告位的位置以及在广告位中的排序"`
	TargetUrls    []types.TargetUrlItem    `json:"target_url" dc:"跳转地址，包括：pc,ios,android"`
	Devices       []string                 `json:"devices" dc:"投放设备"`
	Labels        []string                 `json:"labels" dc:"广告标签"`
	ExposureTime  int                      `json:"exposure_time" dc:"单次曝光时间，单位秒"`
	UserLevels    []int                    `json:"user_levels" dc:"用户等级"`
	StartTime     string                   `json:"start_time" dc:"广告生效时间"`
	EndTime       string                   `json:"end_time" dc:"广告失效时间"`
}

// ADList 添加广告位
func ADList(ctx *gin.Context) {
	var (
		err error
		req = new(ADListReq)
		ads []entity.TAd
	)
	if err = ctx.ShouldBind(req); err != nil {
		global.MyLogger(ctx).Err(err).Msgf("绑定参数失败")
		response.RespFail(ctx, i18n.RetMsgParamParseErr, nil)
		return
	}
	global.MyLogger(ctx).Info().Msgf("req: %+v", *req)

	if req.QueryType == QueryTypeWeightRand && len(req.Locations) != 1 {
		global.MyLogger(ctx).Err(err).Msgf("QueryType: %s, len(Locations) = %d not 1", req.QueryType, len(req.Locations))
		response.RespFail(ctx, i18n.RetMsgParamInvalid, nil)
		return
	}

	err = dao.TAd.Ctx(ctx).
		Where(do.TAd{Status: ADStatusOnline}).
		WhereLTE(dao.TAd.Columns().StartTime, time.Now().Format(time.DateTime)).
		WhereGTE(dao.TAd.Columns().EndTime, time.Now().Format(time.DateTime)).
		Scan(&ads)
	if err != nil {
		response.ResFail(ctx, err.Error())
		return
	}

	items := make([]ADItem, 0)
	for _, item := range ads {
		if !isPicked(ctx, req, item) {
			global.MyLogger(ctx).Debug().Msgf("ad %d/%s is not Picked", item.Id, item.Name)
			continue
		}
		items = append(items, ADItem{
			Name:          item.Name,
			Type:          item.Type,
			Url:           item.Url,
			Logo:          item.Logo,
			SlotLocations: builder.BuildSlotLocations(item.SlotLocations),
			TargetUrls:    builder.BuildTargetUrls(item.TargetUrls),
			Devices:       builder.BuildStringArray(item.Devices),
			Labels:        builder.BuildStringArray(item.Labels),
			ExposureTime:  item.ExposureTime,
			UserLevels:    builder.BuildIntArray(item.UserLevels),
			StartTime:     item.StartTime.String(),
			EndTime:       item.EndTime.String(),
		})
	}

	if req.QueryType == QueryTypeWeightRand {
		var (
			location = req.Locations[0]
			sorts    []int
		)
		for _, item := range items {
			for _, v := range item.SlotLocations {
				if location == v.Location {
					sorts = append(sorts, v.Sort)
				}
			}
		}
		index := weightedRand(ctx, sorts)
		items = []ADItem{items[index]}
	}

	global.MyLogger(ctx).Debug().Msgf("items: %+v", items)
	response.RespOk(ctx, i18n.RetMsgSuccess, ADListRes{Items: items})
}

func weightedRand(ctx *gin.Context, weights []int) int {
	global.MyLogger(ctx).Debug().Msgf("weights: %+v", weights)
	sum := 0
	for _, weight := range weights {
		sum += weight
	}

	rand.Seed(time.Now().UnixNano())
	randNum := 0
	if sum > 0 {
		randNum = rand.Intn(sum)
	}
	global.MyLogger(ctx).Debug().Msgf("randNum: %d", randNum)
	for i, weight := range weights {
		randNum -= weight
		if randNum < 0 {
			global.MyLogger(ctx).Debug().Msgf("i: %d", i)
			return i
		}
	}
	global.MyLogger(ctx).Debug().Msgf("i -> 0")
	return 0
}

func isPicked(ctx *gin.Context, req *ADListReq, item entity.TAd) bool {
	if len(req.Locations) > 0 {
		global.Logger.Debug().Msgf("input.locations: %+v", req.Locations)
		flag := false
		var locations []string
		for _, i := range builder.BuildSlotLocations(item.SlotLocations) {
			locations = append(locations, i.Location)
		}
		global.Logger.Debug().Msgf("item.locations: %+v", locations)
		for _, l := range req.Locations {
			if util.IsInArrayIgnoreCase(l, locations) {
				flag = true
				break
			}
		}
		if !flag {
			return false
		}
	}

	if len(req.Devices) > 0 {
		global.Logger.Debug().Msgf("input.Devices: %+v", req.Devices)
		flag := false
		devices := builder.BuildStringArray(item.Devices)
		global.Logger.Debug().Msgf("item.Devices: %+v", devices)

		for _, d := range req.Devices {
			if util.IsInArrayIgnoreCase(d, devices) {
				flag = true
				break
			}
		}
		if !flag {
			return false
		}
	}
	global.Logger.Debug().Msgf("picked: %s", item.Name)
	return true
}
