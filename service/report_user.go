package service

import (
	"context"
	"fmt"
	"go-speed/constant"
	"go-speed/global"
	"go-speed/model"
	"strconv"
	"strings"
)

func QueryUserReportDay(ctx context.Context, date, channelId int, orderType string, page, size int) (int64, []*model.TUserReportDay, error) {
	order := "desc"
	if strings.ToLower(orderType) == "asc" {
		order = "asc"
	}
	if size > constant.MaxPageSize {
		size = constant.MaxPageSize
	}
	if size == 0 {
		size = 20
	}
	var err error
	var list []*model.TUserReportDay
	sessCount := global.Db2.Context(ctx)
	sess := global.Db2.Context(ctx)
	if date > 0 {
		sess = sess.Where(" date = ?", date)
		sessCount = sessCount.Where(" date = ?", date)
	}
	if channelId > 0 {
		sess = sess.Where(" channel_id = ?", channelId)
		sessCount = sessCount.Where(" channel_id = ?", channelId)
	}
	offset := 0
	if page > 1 {
		offset = (page - 1) * size
	}
	count, err := sessCount.Table(model.TUserReportDay{}).Count()
	if err != nil {
		return 0, nil, err
	}

	err = sess.Limit(size, offset).OrderBy(fmt.Sprintf("date %s, channel_id %s", order, order)).Find(&list)
	return count, list, err
}
func QueryUserChannelDay(ctx context.Context, date int, Channel string, orderType string, page, size int) (int64, []*model.TUserChannelDay, error) {
	order := "desc"
	if strings.ToLower(orderType) == "asc" {
		order = "asc"
	}
	if size > constant.MaxPageSize {
		size = constant.MaxPageSize
	}
	if size == 0 {
		size = 20
	}
	var err error
	var list []*model.TUserChannelDay
	sessCount := global.Db2.Context(ctx)
	sess := global.Db2.Context(ctx)
	if date > 0 {
		sess = sess.Where(" date = ?", date)
		sessCount = sessCount.Where(" date = ?", date)
	}
	if Channel != "" {
		sess = sess.Where(" channel = ?", Channel)
		sessCount = sessCount.Where(" channel = ?", Channel)
	}
	offset := 0
	if page > 1 {
		offset = (page - 1) * size
	}
	count, err := sessCount.Table(model.TUserChannelDay{}).Count()
	if err != nil {
		return 0, nil, err
	}
	err = sess.Limit(size, offset).OrderBy(fmt.Sprintf("date %s, channel %s", order, order)).Find(&list)
	return count, list, err
}
func QueryUserPromotionChannelDay(ctx context.Context, startDate, endDate, date int, Channel string, orderType string, page, size int) (int64, []*model.TUserChannelDay, error) {
	order := "desc"
	if strings.ToLower(orderType) == "asc" {
		order = "asc"
	}
	if size > constant.MaxPageSize {
		size = constant.MaxPageSize
	}
	if size == 0 {
		size = 20
	}
	var err error
	var list []*model.TUserChannelDay
	sessCount := global.Db2.Context(ctx)
	sess := global.Db2.Context(ctx)
	if date > 0 {
		sess = sess.Where(" date = ?", date)
		sessCount = sessCount.Where(" date = ?", date)
	}
	if startDate > 0 && endDate > 0 {
		sess = sess.Where(" date BETWEEN ? AND ?", startDate, endDate)
		sessCount = sessCount.Where(" date BETWEEN ? AND ?", startDate, endDate)
	}
	if Channel == "" {
		sess = sess.Where("channel BETWEEN ? AND ?", constant.MinChannel, constant.MaxChannel)
		sessCount = sessCount.Where("channel BETWEEN ? AND ?", constant.MinChannel, constant.MaxChannel)
	} else {
		channelInt, err := strconv.Atoi(Channel)
		if err != nil || channelInt <= constant.MinChannel || channelInt >= constant.MaxChannel {
			return 0, list, nil
		}
		sess = sess.Where("channel = ?", Channel)
		sessCount = sessCount.Where("channel = ?", Channel)
	}
	offset := 0
	if page > 1 {
		offset = (page - 1) * size
	}
	count, err := sessCount.Table(model.TUserChannelDay{}).Count()
	if err != nil {
		return 0, nil, err
	}
	err = sess.Limit(size, offset).OrderBy(fmt.Sprintf("date %s, channel %s", order, order)).Find(&list)
	return count, list, err
}
func QueryOnlineUserDay(ctx context.Context, date int, channelId string, email, orderType string, page, size int) (int64, []*model.TUserOnlineDay, error) {
	order := "desc"
	if strings.ToLower(orderType) == "asc" {
		order = "asc"
	}
	if size > constant.MaxPageSize {
		size = constant.MaxPageSize
	}
	if size == 0 {
		size = 20
	}
	var err error
	var list []*model.TUserOnlineDay
	sessCount := global.Db2.Context(ctx)
	sess := global.Db2.Context(ctx)
	if date > 0 {
		sess = sess.Where(" date = ?", date)
		sessCount = sessCount.Where(" date = ?", date)
	}
	if email != "" {
		sess = sess.Where(" email = ?", email)
		sessCount = sessCount.Where(" email = ?", email)
	}
	if channelId != "" {
		sess = sess.Where(" channel = ?", channelId)
		sessCount = sessCount.Where(" channel = ?", channelId)
	}
	offset := 0
	if page > 1 {
		offset = (page - 1) * size
	}
	count, err := sessCount.Table(model.TUserOnlineDay{}).Count()
	if err != nil {
		return 0, nil, err
	}

	//err = sess.Limit(size, offset).OrderBy(fmt.Sprintf("date %s, channel %s", order, order)).Find(&list)
	err = sess.Limit(size, offset).OrderBy(fmt.Sprintf("id %s", order)).Find(&list)
	return count, list, err
}
func QueryNodeDay(ctx context.Context, date int, Ip string, orderType string, page, size int) (int64, []*model.TUserNodeDay, error) {
	order := "desc"
	if strings.ToLower(orderType) == "asc" {
		order = "asc"
	}
	if size > constant.MaxPageSize {
		size = constant.MaxPageSize
	}
	if size == 0 {
		size = 20
	}
	var err error
	var list []*model.TUserNodeDay
	sessCount := global.Db2.Context(ctx)
	sess := global.Db2.Context(ctx)
	if date > 0 {
		sess = sess.Where(" date = ?", date)
		sessCount = sessCount.Where(" date = ?", date)
	}
	if Ip != "" {
		sess = sess.Where(" ip = ?", Ip)
		sessCount = sessCount.Where(" ip = ?", Ip)
	}
	offset := 0
	if page > 1 {
		offset = (page - 1) * size
	}
	count, err := sessCount.Table(model.TUserOnlineDay{}).Count()
	if err != nil {
		return 0, nil, err
	}

	err = sess.Limit(size, offset).OrderBy(fmt.Sprintf("date %s, ip %s", order, order)).Find(&list)
	return count, list, err
}
func QueryNodeOnlineUserDay(ctx context.Context, date int, channelId string, email, orderType string, page, size int) (int64, []*model.TUserNodeOnlineDay, error) {
	order := "desc"
	if strings.ToLower(orderType) == "asc" {
		order = "asc"
	}
	if size > constant.MaxPageSize {
		size = constant.MaxPageSize
	}
	if size == 0 {
		size = 20
	}
	var err error
	var list []*model.TUserNodeOnlineDay
	sessCount := global.Db2.Context(ctx)
	sess := global.Db2.Context(ctx)
	if date > 0 {
		sess = sess.Where(" date = ?", date)
		sessCount = sessCount.Where(" date = ?", date)
	}
	if email != "" {
		sess = sess.Where(" email = ?", email)
		sessCount = sessCount.Where(" email = ?", email)
	}
	if channelId != "" {
		sess = sess.Where(" channel = ?", channelId)
		sessCount = sessCount.Where(" channel = ?", channelId)
	}
	offset := 0
	if page > 1 {
		offset = (page - 1) * size
	}
	count, err := sessCount.Table(model.TUserNodeOnlineDay{}).Count()
	if err != nil {
		return 0, nil, err
	}

	//err = sess.Limit(size, offset).OrderBy(fmt.Sprintf("date %s, channel %s", order, order)).Find(&list)
	err = sess.Limit(size, offset).OrderBy(fmt.Sprintf("id %s", order)).Find(&list)
	return count, list, err
}
func QueryUserRechargeReportDay(ctx context.Context, date, GoodsId int, orderType string, page, size int) (int64, []*model.TUserRechargeReportDay, error) {
	order := "desc"
	if strings.ToLower(orderType) == "asc" {
		order = "asc"
	}
	if size > constant.MaxPageSize {
		size = constant.MaxPageSize
	}
	if size == 0 {
		size = 20
	}
	var err error
	var list []*model.TUserRechargeReportDay
	sessCount := global.Db2.Context(ctx)
	sess := global.Db2.Context(ctx)
	if date > 0 {
		sess = sess.Where(" date = ?", date)
		sessCount = sessCount.Where(" date = ?", date)
	}
	if GoodsId > 0 {
		sess = sess.Where(" goods_id = ?", GoodsId)
		sessCount = sessCount.Where(" goods_id = ?", GoodsId)
	}
	offset := 0
	if page > 1 {
		offset = (page - 1) * size
	}
	count, err := sessCount.Table(model.TUserRechargeReportDay{}).Count()
	if err != nil {
		return 0, nil, err
	}

	err = sess.Limit(size, offset).OrderBy(fmt.Sprintf("date %s", order)).Find(&list)
	return count, list, err
}
func QueryUserRechargeTimesReportDay(ctx context.Context, date, GoodsId int, orderType string, page, size int) (int64, []*model.TUserRechargeTimesReportDay, error) {
	order := "desc"
	if strings.ToLower(orderType) == "asc" {
		order = "asc"
	}
	if size > constant.MaxPageSize {
		size = constant.MaxPageSize
	}
	if size == 0 {
		size = 20
	}
	var err error
	var list []*model.TUserRechargeTimesReportDay
	sessCount := global.Db2.Context(ctx)
	sess := global.Db2.Context(ctx)
	if date > 0 {
		sess = sess.Where(" date = ?", date)
		sessCount = sessCount.Where(" date = ?", date)
	}
	if GoodsId > 0 {
		sess = sess.Where(" goods_id = ?", GoodsId)
		sessCount = sessCount.Where(" goods_id = ?", GoodsId)
	}
	offset := 0
	if page > 1 {
		offset = (page - 1) * size
	}
	count, err := sessCount.Table(model.TUserRechargeTimesReportDay{}).Count()
	if err != nil {
		return 0, nil, err
	}

	err = sess.Limit(size, offset).OrderBy(fmt.Sprintf("date %s", order)).Find(&list)
	return count, list, err
}
func QueryChannelUserRechargeTimesReportDay(ctx context.Context, startDate, endDate, date int, Channel string, orderType string, page, size int) (int64, []*model.TUserChannelRechargeDay, error) {
	order := "desc"
	if strings.ToLower(orderType) == "asc" {
		order = "asc"
	}
	if size > constant.MaxPageSize {
		size = constant.MaxPageSize
	}
	if size == 0 {
		size = 20
	}
	var err error
	var list []*model.TUserChannelRechargeDay
	sessCount := global.Db2.Context(ctx)
	sess := global.Db2.Context(ctx)
	if date > 0 {
		sess = sess.Where(" date = ?", date)
		sessCount = sessCount.Where(" date = ?", date)
	}
	if startDate > 0 && endDate > 0 {
		sess = sess.Where(" date BETWEEN ? AND ?", startDate, endDate)
		sessCount = sessCount.Where(" date BETWEEN ? AND ?", startDate, endDate)
	}
	if Channel != "" {
		sess = sess.Where(" channel = ?", Channel)
		sessCount = sessCount.Where(" channel = ?", Channel)
	}
	offset := 0
	if page > 1 {
		offset = (page - 1) * size
	}
	count, err := sessCount.Table(model.TUserChannelRechargeDay{}).Count()
	if err != nil {
		return 0, nil, err
	}
	err = sess.Limit(size, offset).OrderBy(fmt.Sprintf("date %s, channel %s", order, order)).Find(&list)
	return count, list, err
}
