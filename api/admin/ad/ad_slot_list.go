package ad

import (
	"github.com/gin-gonic/gin"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model/entity"
	"go-speed/model/response"
)

type ADSlotListReq struct{}

type ADSlotListRes struct {
	Items []ADSlotItem `json:"items" dc:"广告位列表"`
}

type ADSlotItem struct {
	Location  string `json:"location" dc:"广告位的位置，相当于ID"`
	Name      string `json:"name" dc:"广告位名称"`
	Desc      string `json:"desc" dc:"广告位描述"`
	Status    int    `json:"status" dc:"状态:1-上架；2-下架"`
	CreatedAt string `json:"created_at" dc:"创建时间"`
	UpdatedAt string `json:"updated_at" dc:"更新时间"`
}

// ADSlotList 查询国家列表
func ADSlotList(ctx *gin.Context) {
	var (
		err   error
		items []entity.TAdSlot
	)
	err = dao.TAdSlot.Ctx(ctx).Scan(&items)
	if err != nil {
		global.MyLogger(ctx).Err(err).Msgf("get ad slot failed")
		response.ResFail(ctx, err.Error())
		return
	}

	slots := make([]ADSlotItem, 0)
	for _, item := range items {
		slots = append(slots, ADSlotItem{
			Location:  item.Location,
			Name:      item.Name,
			Desc:      item.Desc,
			Status:    item.Status,
			CreatedAt: item.CreatedAt.String(),
			UpdatedAt: item.UpdatedAt.String(),
		})
	}
	response.RespOk(ctx, i18n.RetMsgSuccess, ADSlotListRes{Items: slots})
}
