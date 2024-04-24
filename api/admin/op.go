package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/model"
	"go-speed/model/request"
	"go-speed/model/response"
	"go-speed/service"
	"strconv"
	"time"
)

func MemberList(c *gin.Context) {
	param := new(request.MemberListAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	if param.TeamName != "" {
		teamU := new(model.TUser)
		has, err := global.Db.Where("uname = ?", param.TeamName).Get(teamU)
		if err != nil || !has {
			global.Logger.Err(err).Msg("团队长不存在！")
			response.ResFail(c, "查询出错！")
			return
		}
		param.TeamId = teamU.Id
	}
	session := service.MemberAdminList(param, user)
	count, err := service.MemberAdminList(param, user).Count()
	if err != nil {
		global.Logger.Err(err).Msg("查询出错！")
		response.ResFail(c, "查询出错！")
		return
	}
	cols := "u1.id as id,u1.uname as uname,u1.created_at,u1.level,u1.expired_time as time1," +
		"u2.id as uid2,u2.uname as uname2,u2.level,u2.expired_time as time2,u1.v2ray_uuid"
	session.Cols(cols)
	session.OrderBy("t.id desc")
	dataList, _ := commonPageListV2(param.Page, param.Size, count, session)
	response.RespOk(c, "成功", dataList)
}

func MemberDevList(c *gin.Context) {
	param := new(request.MemberDevListAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	session := service.MemberDevAdminList(param, user)
	count, err := service.MemberDevAdminList(param, user).Count()
	if err != nil {
		global.Logger.Err(err).Msg("查询出错！")
		response.ResFail(c, "查询出错！")
		return
	}
	cols := "ud.*," +
		"u.uname," +
		"d.os,d.network"
	session.Cols(cols)
	session.OrderBy("ud.id desc")
	dataList, _ := commonPageListV2(param.Page, param.Size, count, session)
	response.RespOk(c, "成功", dataList)
}

func EditMember(c *gin.Context) {
	param := new(request.EditMemberAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	bean := new(model.TGoods)
	bean.UpdatedAt = time.Now()
	bean.Author = user.Uname
	cols := []string{"updated_at", "author"}
	if param.Status != "" {
		bean.Status, _ = strconv.Atoi(param.Status)
		cols = append(cols, "status")
	}
	rows, err := global.Db.Cols(cols...).Where("id = ?", param.Id).Update(bean)
	if err != nil || rows != 1 {
		global.Logger.Err(err).Msg("操作失败！")
		response.ResFail(c, "操作失败！")
		return
	}
	response.ResOk(c, "成功")
}

func EditMemberDev(c *gin.Context) {
	param := new(request.EditMemberDevAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	bean := new(model.TUserDev)
	bean.UpdatedAt = time.Now()
	bean.Comment = user.Uname
	cols := []string{"updated_at", "comment"}
	if param.Status != "" {
		bean.Status, _ = strconv.Atoi(param.Status)
		cols = append(cols, "status")
	}
	rows, err := global.Db.Cols(cols...).Where("id = ?", param.Id).Update(bean)
	if err != nil || rows != 1 {
		global.Logger.Err(err).Msg("操作失败！")
		response.ResFail(c, "操作失败！")
		return
	}
	response.ResOk(c, "成功")
}

func EditMemberExpiredTime(c *gin.Context) {
	param := new(request.EditMemberExpiredTimeRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}

	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	global.Logger.Err(err).Msgf("Admin uname: %s", user.Uname)

	if !isValidTimestamp(param.ExpiredTime) {
		global.Logger.Err(nil).Msgf("ExpiredTime: (%d) 参数无效！不能超过5年时间长度。", param.ExpiredTime)
		response.ResFail(c, "用户过期参数无效！不能超过5年时间长度。")
		return
	}

	bean := new(model.TUser)
	bean.UpdatedAt = time.Now()
	bean.ExpiredTime = param.ExpiredTime
	cols := []string{"updated_at", "expired_time"}
	rows, err := global.Db.Cols(cols...).Where("id = ?", param.Id).Update(bean)
	if err != nil || rows != 1 {
		global.Logger.Err(err).Msg("操作失败！")
		response.ResFail(c, "操作失败！")
		return
	}
	response.ResOk(c, "成功")
}

func isValidTimestamp(expiredTime int64) bool {
	// 将 ExpiredTime 转换为 time.Time 类型
	expiredTimeInTime := time.Unix(expiredTime, 0)

	// 获取当前时间
	now := time.Now()

	// 获取1年前的时间
	oneYearAgo := now.AddDate(-1, 0, 0)

	// 获取5年后的时间
	fiveYearsLater := now.AddDate(5, 0, 0)

	// 判断 ExpiredTime 是否在有效范围内
	if expiredTimeInTime.After(oneYearAgo) && expiredTimeInTime.Before(fiveYearsLater) {
		return true
	}
	return false
}

func GetReportUserDayList(c *gin.Context) {
	param := new(request.GetReportUserDayListRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}

	total, list, err := service.QueryUserReportDay(c, param.Date, param.ChannelId, param.OrderType, param.Page, param.Size)
	if err != nil {
		global.Logger.Err(err).Msg("查询出错！")
		response.ResFail(c, "查询出错！")
		return
	}
	items := make([]response.ReportUserDay, 0)
	for _, item := range list {
		items = append(items, response.ReportUserDay{
			Id:        item.Id,
			Date:      item.Date,
			ChannelId: item.ChannelId,
			Total:     item.Total,
			New:       item.New,
			Retained:  item.Retained,
			CreatedAt: item.CreatedAt.String(),
		})
	}
	resp := response.GetReportUserDayListResponse{Total: total, Items: items}
	response.RespOk(c, i18n.RetMsgSuccess, resp)
	return
}

func GetOnlineUserDayList(c *gin.Context) {
	param := new(request.GetOnlineUserDayListRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}

	total, list, err := service.QueryOnlineUserDay(c, param.Date, param.ChannelId, param.Email, param.OrderType, param.Page, param.Size)
	if err != nil {
		global.Logger.Err(err).Msg("查询出错！")
		response.ResFail(c, "查询出错！")
		return
	}
	items := make([]response.OnlineUserDay, 0)
	for _, item := range list {
		items = append(items, response.OnlineUserDay{
			Id:               item.Id,
			Date:             item.Date,
			Email:            item.Email,
			ChannelId:        item.ChannelId,
			OnlineDuration:   item.OnlineDuration,
			Uplink:           item.Uplink,
			Downlink:         item.Downlink,
			CreatedAt:        item.CreatedAt.String(),
			LastLoginCountry: item.LastLoginCountry,
		})
	}
	resp := response.GetOnlineUserDayListResponse{Total: total, Items: items}
	response.RespOk(c, i18n.RetMsgSuccess, resp)
	return
}

func ComboList(c *gin.Context) {
	param := new(request.GoodsListAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	session := service.GoodsAdminList(param, user)
	count, err := service.GoodsAdminList(param, user).Count()
	if err != nil {
		global.Logger.Err(err).Msg("查询出错！")
		response.ResFail(c, "查询出错！")
		return
	}
	cols := "g.*"
	session.Cols(cols)
	session.OrderBy("g.id desc")
	dataList, _ := commonPageListV2(param.Page, param.Size, count, session)
	response.RespOk(c, "成功", dataList)
}

func AddCombo(c *gin.Context) {
	param := new(request.AddGoodsAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	goods := &model.TGoods{
		MType:      param.MType,
		Title:      param.Title,
		TitleEn:    "",
		Price:      decimal.NewFromFloat(param.Price).Truncate(2).String(),
		Period:     param.Period,
		DevLimit:   param.DevLimit,
		FlowLimit:  param.FlowLimit,
		Status:     1,
		Author:     user.Uname,
		IsDiscount: param.IsDiscount,
		Low:        param.Low,
		High:       param.High,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Comment:    "",
	}
	rows, err := global.Db.Insert(goods)
	if err != nil || rows != 1 {
		global.Logger.Err(err).Msg("操作失败！")
		response.ResFail(c, "操作失败！")
		return
	}
	response.ResOk(c, "成功")
}

func EditCombo(c *gin.Context) {
	param := new(request.EditGoodsAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	bean := new(model.TGoods)
	bean.UpdatedAt = time.Now()
	bean.Author = user.Uname
	cols := []string{"updated_at", "author"}
	if param.MType > 0 {
		cols = append(cols, "m_type")
		bean.MType = param.MType
	}
	if param.Period > 0 {
		cols = append(cols, "period")
		bean.Period = param.Period
	}
	if param.DevLimit > 0 {
		cols = append(cols, "dev_limit")
		bean.DevLimit = param.DevLimit
	}
	if param.FlowLimit > 0 {
		cols = append(cols, "flow_limit")
		bean.FlowLimit = param.FlowLimit
	}
	if param.Price > 0 {
		cols = append(cols, "price")
		bean.Price = decimal.NewFromFloat(param.Price).Truncate(2).String()
	}
	if param.Title != "" {
		cols = append(cols, "title")
		bean.Title = param.Title
	}
	if param.Status > 0 {
		cols = append(cols, "status")
		bean.Status = param.Status
	}
	if param.IsDiscount > 0 {
		cols = append(cols, "is_discount")
		bean.IsDiscount = param.IsDiscount
	}
	if param.Low > 0 {
		cols = append(cols, "low")
		bean.Low = param.Low
	}
	if param.High > 0 {
		cols = append(cols, "high")
		bean.High = param.High
	}
	rows, err := global.Db.Cols(cols...).Where("id = ?", param.Id).Update(bean)
	if err != nil || rows != 1 {
		global.Logger.Err(err).Msg("操作失败！")
		response.ResFail(c, "操作失败！")
		return
	}
	response.ResOk(c, "成功")

}

func NoticeList(c *gin.Context) {
	param := new(request.NoticeListAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	session := service.NoticeAdminList(param, user)
	count, err := service.NoticeAdminList(param, user).Count()
	if err != nil {
		global.Logger.Err(err).Msg("查询出错！")
		response.ResFail(c, "查询出错！")
		return
	}
	cols := "n.*"
	session.Cols(cols)
	session.OrderBy("n.id desc")
	dataList, _ := commonPageListV2(param.Page, param.Size, count, session)
	response.RespOk(c, "成功", dataList)
}

func AddNotice(c *gin.Context) {
	param := new(request.AddNoticeAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	notice := &model.TNotice{
		Title:     param.Title,
		Tag:       param.Tag,
		Content:   param.Content,
		Author:    user.Uname,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status:    1,
		Comment:   "",
	}
	rows, err := global.Db.Insert(notice)
	if err != nil || rows != 1 {
		global.Logger.Err(err).Msg("操作失败！")
		response.ResFail(c, "操作失败！")
		return
	}
	response.ResOk(c, "成功")
}

func EditNotice(c *gin.Context) {
	param := new(request.EditNoticeAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	bean := new(model.TNotice)
	bean.UpdatedAt = time.Now()
	bean.Author = user.Uname
	cols := []string{"updated_at", "author"}
	if param.Title != "" {
		cols = append(cols, "title")
		bean.Title = param.Title
	}
	if param.Tag != "" {
		cols = append(cols, "tag")
		bean.Tag = param.Tag
	}
	if param.Content != "" {
		cols = append(cols, "content")
		bean.Content = param.Content
	}
	if param.Status > 0 {
		cols = append(cols, "status")
		bean.Status = param.Status
	}

	rows, err := global.Db.Cols(cols...).Where("id = ?", param.Id).Update(bean)
	if err != nil || rows != 1 {
		global.Logger.Err(err).Msg("操作失败！")
		response.ResFail(c, "操作失败！")
		return
	}
	response.ResOk(c, "成功")
}

func OrderList(c *gin.Context) {
	param := new(request.OrderListAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	session := service.OrderAdminList(param, user)
	count, err := service.OrderAdminList(param, user).Count()
	if err != nil {
		global.Logger.Err(err).Msg("查询出错！")
		response.ResFail(c, "查询出错！")
		return
	}
	cols := "o.*,u.uname"
	session.Cols(cols)
	session.OrderBy("o.id desc")
	dataList, _ := commonPageListV2(param.Page, param.Size, count, session)
	response.RespOk(c, "成功", dataList)
}

func OrderSummary(c *gin.Context) {

}

func AdList(c *gin.Context) {
	param := new(request.AdListAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	session := service.AdAdminList(param, user)
	count, err := service.AdAdminList(param, user).Count()
	if err != nil {
		global.Logger.Err(err).Msg("查询出错！")
		response.ResFail(c, "查询出错！")
		return
	}
	cols := "ad.*"
	session.Cols(cols)
	session.OrderBy("ad.id desc")
	dataList, _ := commonPageListV2(param.Page, param.Size, count, session)
	response.RespOk(c, "成功", dataList)
}

func AddAd(c *gin.Context) {
	param := new(request.AddAdAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	ad := &model.TAd{
		Status:    1,
		Sort:      0,
		Name:      param.Name,
		Logo:      param.Logo,
		Link:      param.Link,
		AdType:    param.AdType,
		Tag:       param.Tag,
		Content:   param.Content,
		Author:    user.Uname,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Comment:   "",
	}
	rows, err := global.Db.Insert(ad)
	if err != nil || rows != 1 {
		global.Logger.Err(err).Msg("操作失败！")
		response.ResFail(c, "操作失败！")
		return
	}
	response.ResOk(c, "成功")
}

func EditAd(c *gin.Context) {
	param := new(request.EditAdAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	ad := new(model.TAd)
	ad.UpdatedAt = time.Now()
	ad.Author = user.Uname
	cols := []string{"updated_at", "author"}
	if param.Status > 0 {
		cols = append(cols, "status")
		ad.Status = param.Status
	}
	if param.Name != "" {
		cols = append(cols, "name")
		ad.Name = param.Name
	}
	if param.Tag != "" {
		cols = append(cols, "tag")
		ad.Tag = param.Tag
	}
	if param.Link != "" {
		cols = append(cols, "link")
		ad.Link = param.Link
	}
	if param.Logo != "" {
		cols = append(cols, "logo")
		ad.Logo = param.Logo
	}
	if param.Content != "" {
		cols = append(cols, "content")
		ad.Content = param.Content
	}
	if param.AdType > 0 {
		cols = append(cols, "ad_type")
		ad.AdType = param.AdType
	}
	rows, err := global.Db.Cols(cols...).Where("id = ?", param.Id).Update(ad)
	if err != nil || rows != 1 {
		global.Logger.Err(err).Msg("操作失败！")
		response.ResFail(c, "操作失败！")
		return
	}
	response.ResOk(c, "成功")

}

func AdSummary(c *gin.Context) {

}

func NodeList(c *gin.Context) {
	param := new(request.NodeListAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	session := service.NodeAdminList(param, user)
	count, err := service.NodeAdminList(param, user).Count()
	if err != nil {
		global.Logger.Err(err).Msg("查询出错！")
		response.ResFail(c, "查询出错！")
		return
	}
	cols := "n.*"
	session.Cols(cols)
	session.OrderBy("n.id desc")
	dataList, _ := commonPageListV2(param.Page, param.Size, count, session)
	list := dataList.List.([]map[string]interface{})
	if len(list) > 0 {
		for _, item := range list {
			var dnsArray []map[string]interface{}
			nodeId := item["id"].(int64)
			var dnsList []*model.TNodeDns
			err = global.Db.Where("node_id = ? and status = 1", nodeId).Find(&dnsList)

			for _, dns := range dnsList {
				var dnsItem = make(map[string]interface{})
				dnsItem["id"] = dns.Id
				dnsItem["node_id"] = dns.NodeId
				dnsItem["dns"] = dns.Dns
				dnsItem["ip"] = dns.Ip
				dnsItem["level"] = dns.Level
				dnsArray = append(dnsArray, dnsItem)
			}

			item["dns_list"] = dnsArray
		}
	}
	response.RespOk(c, "成功", dataList)
}

func AddNode(c *gin.Context) {
	param := new(request.AddNodeAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	node := &model.TNode{
		Name:      param.Name,
		Title:     param.Title,
		TitleEn:   "",
		Country:   param.Country,
		CountryEn: "",
		Ip:        param.Ip,
		Server:    param.Server,
		Port:      param.Port,
		Cpu:       param.Cpu,
		Flow:      param.Flow,
		Disk:      param.Disk,
		Memory:    param.Memory,
		Status:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Author:    user.Uname,
		Comment:   "",
	}
	rows, err := global.Db.Insert(node)
	if err != nil || rows != 1 {
		global.Logger.Err(err).Msg("操作失败！")
		response.ResFail(c, "操作失败！")
		return
	}
	response.ResOk(c, "成功")
}

func EditNode(c *gin.Context) {
	param := new(request.EditNodeAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	bean := new(model.TNode)
	bean.UpdatedAt = time.Now()
	bean.Author = user.Uname
	cols := []string{"updated_at", "author"}
	if param.Cpu > 0 {
		cols = append(cols, "cpu")
		bean.Cpu = param.Cpu
	}
	if param.Port > 0 {
		cols = append(cols, "port")
		bean.Port = param.Port
	}
	if param.Disk > 0 {
		cols = append(cols, "disk")
		bean.Disk = param.Disk
	}
	if param.Memory > 0 {
		cols = append(cols, "memory")
		bean.Memory = param.Memory
	}
	if param.Flow > 0 {
		cols = append(cols, "flow")
		bean.Flow = param.Flow
	}
	if param.Title != "" {
		cols = append(cols, "title")
		bean.Title = param.Title
	}
	if param.Server != "" {
		cols = append(cols, "server")
		bean.Server = param.Server
	}
	if param.Ip != "" {
		cols = append(cols, "ip")
		bean.Ip = param.Ip
	}
	if param.Country != "" {
		cols = append(cols, "country")
		bean.Country = param.Country
	}
	if param.Name != "" {
		cols = append(cols, "name")
		bean.Name = param.Name
	}
	if param.Status > 0 {
		cols = append(cols, "status")
		bean.Status = param.Status
	}
	if param.IsRecommend > 0 {
		cols = append(cols, "is_recommend")
		bean.IsRecommend = param.IsRecommend
	}
	rows, err := global.Db.Cols(cols...).Where("id = ?", param.Id).Update(bean)
	if err != nil || rows != 1 {
		global.Logger.Err(err).Msg("操作失败！")
		response.ResFail(c, "操作失败！")
		return
	}
	response.ResOk(c, "成功")
}

func NodeUuidList(c *gin.Context) {
	param := new(request.NodeUuidListAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	session := service.NodeUuidAdminList(param, user)
	count, err := service.NodeUuidAdminList(param, user).Count()
	if err != nil {
		global.Logger.Err(err).Msg("查询出错！")
		response.ResFail(c, "查询出错！")
		return
	}
	cols := "nu.*,u.uname,n.name"
	session.Cols(cols)
	session.OrderBy("nu.id desc")
	dataList, _ := commonPageListV2(param.Page, param.Size, count, session)
	response.RespOk(c, "成功", dataList)
}

func AppInfo(c *gin.Context) {
	param := new(request.DictDetailAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	_, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}

	var list []*model.TDict
	err = global.Db.Where("key_id = ?", param.FilterPac).
		Or("key_id = ?", param.FilterRefuse).
		Find(&list)
	if err != nil {
		global.Logger.Err(err).Msg("key不存在！")
		response.ResFail(c, "失败！")
		return
	}
	var result = make(map[string]interface{})
	for _, item := range list {
		result[item.KeyId] = item.Value
	}
	response.RespOk(c, "成功", result)
}

func EditAppInfo(c *gin.Context) {
	param := new(request.DictEditAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	_, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	nowTime := time.Now()

	bean := new(model.TDict)
	bean.UpdatedAt = nowTime
	bean.Value = param.FilterPac
	rows, err := global.Db.Cols("updated_at", "value").Where("key_id = ?", "filter_pac").Update(bean)
	if err != nil || rows != 1 {
		global.Logger.Err(err).Msg("操作失败！")
		response.ResFail(c, "操作失败！")
		return
	}
	bean = new(model.TDict)
	bean.UpdatedAt = nowTime
	bean.Value = param.FilterRefuse
	rows, err = global.Db.Cols("updated_at", "value").Where("key_id = ?", "filter_refuse").Update(bean)
	if err != nil || rows != 1 {
		global.Logger.Err(err).Msg("操作失败！")
		response.ResFail(c, "操作失败！")
		return
	}
	response.ResOk(c, "成功")
}

func SiteList(c *gin.Context) {
	param := new(request.SiteListAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	session := service.SiteAdminList(param, user)
	count, err := service.SiteAdminList(param, user).Count()
	if err != nil {
		global.Logger.Err(err).Msg("查询出错！")
		response.ResFail(c, "查询出错！")
		return
	}
	cols := "s.*"
	session.Cols(cols)
	session.OrderBy("s.id desc")
	dataList, _ := commonPageListV2(param.Page, param.Size, count, session)
	response.RespOk(c, "成功", dataList)
}

func AddSite(c *gin.Context) {
	param := new(request.AddSiteAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	bean := &model.TSite{
		Site:      param.Site,
		Ip:        param.Ip,
		Status:    1,
		Author:    user.Uname,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Comment:   "",
	}
	rows, err := global.Db.Insert(bean)
	if err != nil || rows != 1 {
		global.Logger.Err(err).Msg("操作失败！")
		response.ResFail(c, "操作失败！")
		return
	}
	response.ResOk(c, "成功")
}

func EditSite(c *gin.Context) {
	param := new(request.EditSiteAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	bean := new(model.TSite)
	bean.UpdatedAt = time.Now()
	bean.Author = user.Uname
	cols := []string{"updated_at", "author"}
	if param.Site != "" {
		cols = append(cols, "site")
		bean.Site = param.Site
	}
	if param.Ip != "" {
		cols = append(cols, "ip")
		bean.Ip = param.Ip
	}
	if param.Status > 0 {
		cols = append(cols, "status")
		bean.Status = param.Status
	}
	rows, err := global.Db.Cols(cols...).Where("id = ?", param.Id).Update(bean)
	if err != nil || rows != 1 {
		global.Logger.Err(err).Msg("操作失败！")
		response.ResFail(c, "操作失败！")
		return
	}
	response.ResOk(c, "成功")
}

func GiftList(c *gin.Context) {
	param := new(request.GiftListAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	session := service.GiftAdminList(param, user)
	count, err := service.GiftAdminList(param, user).Count()
	if err != nil {
		global.Logger.Err(err).Msg("查询出错！")
		response.ResFail(c, "查询出错！")
		return
	}
	cols := "g.*,u.uname"
	session.Cols(cols)
	session.OrderBy("g.id desc")
	dataList, _ := commonPageListV2(param.Page, param.Size, count, session)
	response.RespOk(c, "成功", dataList)
}

func ActivityList(c *gin.Context) {
	param := new(request.ActivityListAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	session := service.ActivityAdminList(param, user)
	count, err := service.ActivityAdminList(param, user).Count()
	if err != nil {
		global.Logger.Err(err).Msg("查询出错！")
		response.ResFail(c, "查询出错！")
		return
	}
	cols := "a.*,u.uname"
	session.Cols(cols)
	session.OrderBy("a.id desc")
	dataList, _ := commonPageListV2(param.Page, param.Size, count, session)
	response.RespOk(c, "成功", dataList)
}

func SpeedLogs(c *gin.Context) {
	param := new(request.SpeedLogsAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	session := service.SpeedLogsAdminList(param, user)
	count, err := service.SpeedLogsAdminList(param, user).Count()
	if err != nil {
		global.Logger.Err(err).Msg("查询出错！")
		response.ResFail(c, "查询出错！")
		return
	}
	cols := "w.*,u.uname,dev.os,dev.network,node.name"
	session.Cols(cols)
	session.OrderBy("w.id desc")
	dataList, _ := commonPageListV2(param.Page, param.Size, count, session)
	response.RespOk(c, "成功", dataList)
}

func DevLogs(c *gin.Context) {
	param := new(request.DevLogsAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	session := service.DevLogsAdminList(param, user)
	count, err := service.DevLogsAdminList(param, user).Count()
	if err != nil {
		global.Logger.Err(err).Msg("查询出错！")
		response.ResFail(c, "查询出错！")
		return
	}
	cols := "up.*,u.uname,dev.os,dev.network"
	session.Cols(cols)
	session.OrderBy("up.id desc")
	dataList, _ := commonPageListV2(param.Page, param.Size, count, session)
	response.RespOk(c, "成功", dataList)
}

func ChannelList(c *gin.Context) {
	param := new(request.ChannelListAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	session := service.ChannelAdminList(param, user)
	count, err := service.ChannelAdminList(param, user).Count()
	if err != nil {
		global.Logger.Err(err).Msg("查询出错！")
		response.ResFail(c, "查询出错！")
		return
	}
	cols := "c.*"
	session.Cols(cols)
	session.OrderBy("c.id desc")
	dataList, _ := commonPageListV2(param.Page, param.Size, count, session)
	response.RespOk(c, "成功", dataList)
}

func AddChannel(c *gin.Context) {
	param := new(request.AddChannelAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	bean := &model.TChannel{
		Name:      param.Name,
		Code:      param.Code,
		Link:      param.Link,
		Status:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Author:    user.Uname,
		Comment:   "",
	}
	rows, err := global.Db.Insert(bean)
	if err != nil || rows != 1 {
		global.Logger.Err(err).Msg("操作失败！")
		response.ResFail(c, "操作失败！")
		return
	}
	response.ResOk(c, "成功")
}

func EditChannel(c *gin.Context) {
	param := new(request.EditChannelAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	bean := new(model.TSite)
	bean.UpdatedAt = time.Now()
	bean.Author = user.Uname
	cols := []string{"updated_at", "author"}
	if param.Name != "" {
		cols = append(cols, "name")
		bean.Site = param.Name
	}
	if param.Code != "" {
		cols = append(cols, "code")
		bean.Ip = param.Code
	}
	if param.Link != "" {
		cols = append(cols, "link")
		bean.Ip = param.Link
	}
	if param.Status > 0 {
		cols = append(cols, "status")
		bean.Status = param.Status
	}
	rows, err := global.Db.Cols(cols...).Where("id = ?", param.Id).Update(bean)
	if err != nil || rows != 1 {
		global.Logger.Err(err).Msg("操作失败！")
		response.ResFail(c, "操作失败！")
		return
	}
	response.ResOk(c, "成功")
}

func AppVersionList(c *gin.Context) {
	param := new(request.AppVersionListAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	session := service.AppVersionAdminList(param, user)
	count, err := service.AppVersionAdminList(param, user).Count()
	if err != nil {
		global.Logger.Err(err).Msg("查询出错！")
		response.ResFail(c, "查询出错！")
		return
	}
	cols := "v.*"
	session.Cols(cols)
	session.OrderBy("v.id desc")
	dataList, _ := commonPageListV2(param.Page, param.Size, count, session)
	response.RespOk(c, "成功", dataList)
}

func AddAppVersion(c *gin.Context) {
	param := new(request.AddAppVersionAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	global.MyLogger(c).Info().Msgf("param: %+v", *param)
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	bean := &model.TAppVersion{
		AppType:   param.AppType,
		Version:   param.Version,
		Link:      param.Link,
		Status:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Author:    user.Uname,
		Comment:   "",
	}
	rows, err := global.Db.Insert(bean)
	if err != nil || rows != 1 {
		global.Logger.Err(err).Msg("操作失败！")
		response.ResFail(c, "操作失败！")
		return
	}
	response.ResOk(c, "成功")
}

func EditAppVersion(c *gin.Context) {
	param := new(request.EditAppVersionAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	bean := new(model.TAppVersion)
	bean.UpdatedAt = time.Now()
	bean.Author = user.Uname
	cols := []string{"updated_at", "author"}
	if param.Version != "" {
		cols = append(cols, "version")
		bean.Version = param.Version
	}
	if param.Link != "" {
		cols = append(cols, "link")
		bean.Link = param.Link
	}
	if param.Status > 0 {
		cols = append(cols, "status")
		bean.Status = param.Status
	}
	rows, err := global.Db.Cols(cols...).Where("id = ?", param.Id).Update(bean)
	if err != nil || rows != 1 {
		global.Logger.Err(err).Msg("操作失败！")
		response.ResFail(c, "操作失败！")
		return
	}
	response.ResOk(c, "成功")
}

func AppDnsList(c *gin.Context) {
	param := new(request.AppDnsListAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	session := service.AppDnsAdminList(param, user)
	count, err := service.AppDnsAdminList(param, user).Count()
	if err != nil {
		global.Logger.Err(err).Msg("查询出错！")
		response.ResFail(c, "查询出错！")
		return
	}
	cols := "d.*"
	session.Cols(cols)
	session.OrderBy("d.id desc")
	dataList, _ := commonPageListV2(param.Page, param.Size, count, session)
	response.RespOk(c, "成功", dataList)
}

func AddAppDns(c *gin.Context) {
	param := new(request.AddAppDnsAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	bean := &model.TAppDns{
		SiteType:  param.SiteType,
		Dns:       param.Dns,
		Ip:        param.Ip,
		Level:     param.Level,
		Status:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Author:    user.Uname,
		Comment:   "",
	}
	rows, err := global.Db.Insert(bean)
	if err != nil || rows != 1 {
		global.Logger.Err(err).Msg("操作失败！")
		response.ResFail(c, "操作失败！")
		return
	}
	response.ResOk(c, "成功")
}

func EditAppDns(c *gin.Context) {
	param := new(request.EditAppDnsAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	bean := new(model.TAppDns)
	bean.UpdatedAt = time.Now()
	bean.Author = user.Uname
	cols := []string{"updated_at", "author"}
	if param.Dns != "" {
		cols = append(cols, "dns")
		bean.Dns = param.Dns
	}
	if param.Ip != "" {
		cols = append(cols, "ip")
		bean.Ip = param.Ip
	}
	if param.Status > 0 {
		cols = append(cols, "status")
		bean.Status = param.Status
	}
	if param.SiteType > 0 {
		cols = append(cols, "site_type")
		bean.SiteType = param.SiteType
	}
	if param.Level > 0 {
		cols = append(cols, "level")
		bean.Level = param.Level
	}
	rows, err := global.Db.Cols(cols...).Where("id = ?", param.Id).Update(bean)
	if err != nil || rows != 1 {
		global.Logger.Err(err).Msg("操作失败！")
		response.ResFail(c, "操作失败！")
		return
	}
	response.ResOk(c, "成功")
}

func NodeDnsList(c *gin.Context) {
	param := new(request.NodeDnsListAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	session := service.NodeDnsAdminList(param, user)
	count, err := service.NodeDnsAdminList(param, user).Count()
	if err != nil {
		global.Logger.Err(err).Msg("查询出错！")
		response.ResFail(c, "查询出错！")
		return
	}
	cols := "d.*"
	session.Cols(cols)
	session.OrderBy("d.id desc")
	dataList, _ := commonPageListV2(param.Page, param.Size, count, session)
	response.RespOk(c, "成功", dataList)
}

func AddNodeDns(c *gin.Context) {
	param := new(request.AddNodeDnsAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	bean := &model.TNodeDns{
		NodeId:    param.NodeId,
		Dns:       param.Dns,
		Ip:        param.Ip,
		Level:     param.Level,
		Status:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Author:    user.Uname,
		Comment:   "",
	}
	rows, err := global.Db.Insert(bean)
	if err != nil || rows != 1 {
		global.Logger.Err(err).Msg("操作失败！")
		response.ResFail(c, "操作失败！")
		return
	}
	response.ResOk(c, "成功")
}

func EditNodeDns(c *gin.Context) {
	param := new(request.EditNodeDnsAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	bean := new(model.TNodeDns)
	bean.UpdatedAt = time.Now()
	bean.Author = user.Uname
	cols := []string{"updated_at", "author"}
	if param.Dns != "" {
		cols = append(cols, "dns")
		bean.Dns = param.Dns
	}
	if param.Ip != "" {
		cols = append(cols, "ip")
		bean.Ip = param.Ip
	}
	if param.Status > 0 {
		cols = append(cols, "status")
		bean.Status = param.Status
	}
	if param.NodeId > 0 {
		cols = append(cols, "node_id")
		bean.NodeId = param.NodeId
	}
	if param.Level > 0 {
		cols = append(cols, "level")
		bean.Level = param.Level
	}
	rows, err := global.Db.Cols(cols...).Where("id = ?", param.Id).Update(bean)
	if err != nil || rows != 1 {
		global.Logger.Err(err).Msg("操作失败！")
		response.ResFail(c, "操作失败！")
		return
	}
	response.ResOk(c, "成功")
}

func IosAccountList(c *gin.Context) {
	param := new(request.IosAccountListAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	session := service.IosAccountAdminList(param, user)
	count, err := service.IosAccountAdminList(param, user).Count()
	if err != nil {
		global.Logger.Err(err).Msg("查询出错！")
		response.ResFail(c, "查询出错！")
		return
	}
	cols := "a.*"
	session.Cols(cols)
	session.OrderBy("a.id desc")
	dataList, _ := commonPageListV2(param.Page, param.Size, count, session)
	response.RespOk(c, "成功", dataList)
}

func AddIosAccount(c *gin.Context) {
	param := new(request.AddIosAccountAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	bean := &model.TIosAccount{
		Account:     param.Account,
		Pass:        param.Pass,
		Name:        param.Name,
		AccountType: param.AccountType,
		Status:      1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Author:      user.Uname,
		Comment:     "",
	}
	rows, err := global.Db.Insert(bean)
	if err != nil || rows != 1 {
		global.Logger.Err(err).Msg("操作失败！")
		response.ResFail(c, "操作失败！")
		return
	}
	response.ResOk(c, "成功")
}

func EditIosAccount(c *gin.Context) {
	param := new(request.EditIosAccountAdminRequest)
	if err := c.ShouldBind(param); err != nil {
		global.Logger.Err(err).Msg("绑定参数")
		response.ResFail(c, "参数错误")
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	user, err := service.GetAdminUserByClaims(claims)
	if err != nil {
		global.Logger.Err(err).Msg("不合法！")
		response.ResFail(c, "不合法！")
		return
	}
	bean := new(model.TIosAccount)
	bean.UpdatedAt = time.Now()
	bean.Author = user.Uname
	cols := []string{"updated_at", "author"}
	if param.Name != "" {
		cols = append(cols, "name")
		bean.Name = param.Name
	}
	if param.Account != "" {
		cols = append(cols, "account")
		bean.Account = param.Account
	}
	if param.Pass != "" {
		cols = append(cols, "pass")
		bean.Pass = param.Pass
	}
	if param.Status > 0 {
		cols = append(cols, "status")
		bean.Status = param.Status
	}
	if param.AccountType > 0 {
		cols = append(cols, "account_type")
		bean.AccountType = param.AccountType
	}
	rows, err := global.Db.Cols(cols...).Where("id = ?", param.Id).Update(bean)
	if err != nil || rows != 1 {
		global.Logger.Err(err).Msg("操作失败！")
		response.ResFail(c, "操作失败！")
		return
	}
	response.ResOk(c, "成功")
}

func GiveSummary(c *gin.Context) {

}

func PlantDaySummary(c *gin.Context) {

}

func PlantMonthSummary(c *gin.Context) {

}
