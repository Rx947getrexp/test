package service

import (
	"context"
	"fmt"
	"go-speed/constant"
	"go-speed/global"
	"go-speed/model"
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

func QueryOnlineUserDay(ctx context.Context, date, channelId int, email, orderType string, page, size int) (int64, []*model.TUserOnlineDay, error) {
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
	if channelId > 0 {
		sess = sess.Where(" channel_id = ?", channelId)
		sessCount = sessCount.Where(" channel_id = ?", channelId)
	}
	offset := 0
	if page > 1 {
		offset = (page - 1) * size
	}
	count, err := sessCount.Table(model.TUserOnlineDay{}).Count()
	if err != nil {
		return 0, nil, err
	}

	err = sess.Limit(size, offset).OrderBy(fmt.Sprintf("date %s, channel_id %s", order, order)).Find(&list)
	return count, list, err
}
