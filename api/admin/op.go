package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"go-speed/global"
	"go-speed/model"
	"go-speed/model/request"
	"go-speed/model/response"
	"go-speed/service"
	"time"
)

func MemberList(c *gin.Context) {

}

func MemberDevList(c *gin.Context) {

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
		MType:     param.MType,
		Title:     param.Title,
		TitleEn:   "",
		Price:     decimal.NewFromFloat(param.Price).Truncate(2).String(),
		Period:    param.Period,
		DevLimit:  param.DevLimit,
		FlowLimit: param.FlowLimit,
		Status:    1,
		Author:    user.Uname,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Comment:   "",
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
	rows, err := global.Db.Cols(cols...).Where("id = ?", param.Id).Update(bean)
	if err != nil || rows != 1 {
		global.Logger.Err(err).Msg("操作失败！")
		response.ResFail(c, "操作失败！")
		return
	}
	response.ResOk(c, "成功")
}

func LinkDetail(c *gin.Context) {

}

func EditLink(c *gin.Context) {

}

func SiteList(c *gin.Context) {

}

func AddSite(c *gin.Context) {

}

func EditSite(c *gin.Context) {

}

func DelSite(c *gin.Context) {

}

func GiveTeam(c *gin.Context) {

}

func GiveActivity(c *gin.Context) {

}

func GiveCombo(c *gin.Context) {

}

func GiveSummary(c *gin.Context) {

}

func PlantDaySummary(c *gin.Context) {

}

func PlantMonthSummary(c *gin.Context) {

}
