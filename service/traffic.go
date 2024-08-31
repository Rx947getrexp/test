package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"go-speed/dao"
	"go-speed/global"
	"go-speed/model"
	"go-speed/model/do"
	"go-speed/model/entity"
	"time"
)

func GetUserTrafficCurrentMonth(email string) ([]model.TUserTraffic, error) {
	userTraffic := make([]model.TUserTraffic, 0)
	now := time.Now()
	// 计算当前月份的第一天
	firstDay := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location()).Format("20060102")
	// 计算当前月份的最后一天
	lastDay := time.Date(now.Year(), now.Month()+1, 0, 0, 0, 0, 0, now.Location()).Format("20060102")

	err := global.Db.Where("email = ? and date >= ? and date <= ?", email, firstDay, lastDay).Find(&userTraffic)
	if err != nil {
		global.Logger.Err(err).Msgf("GetUserTrafficCurrentMonth failed. email: %s, firstDay: %s, lastDay: %s", email, firstDay, lastDay)
		return nil, err
	}
	return userTraffic, nil
}

func GetUserTrafficByEmail(ctx context.Context, email, ip string, date int) (userTraffic *entity.TV2RayUserTraffic, err error) {
	err = dao.TV2RayUserTraffic.Ctx(ctx).Where(do.TV2RayUserTraffic{
		Email: email,
		Date:  date,
		Ip:    ip,
	}).Scan(&userTraffic)
	if err != nil {
		global.Logger.Err(err).Msgf("get user traffic failed. email: %s, ip: %s, date: %s", email, ip, date)
		return nil, err
	}
	return userTraffic, nil
}

// CreateUserTraffic 插入新的用量统计记录
func CreateUserTraffic(ctx context.Context, email, ip string, date int, uplink, downlink uint64) error {
	lastId, err := dao.TV2RayUserTraffic.Ctx(ctx).Data(do.TV2RayUserTraffic{
		Email:     email,
		Date:      date,
		Ip:        ip,
		Uplink:    uplink,
		Downlink:  downlink,
		CreatedAt: gtime.Now(),
		UpdatedAt: gtime.Now(),
	}).InsertAndGetId()
	if err != nil {
		global.Logger.Err(err).Msgf("insert TV2RayUserTraffic failed, email: %+v, date: %s", email, date)
		return err
	}
	global.Logger.Info().Msgf("insert TV2RayUserTraffic success, email: %+v, date: %s, lastId: %d", email, date, lastId)
	return nil
}

// CreateUserTrafficLog 插入新的用量流水
func CreateUserTrafficLog(ctx context.Context, email, ip, dataTime string, uplink, downlink uint64) error {
	lastId, err := dao.TV2RayUserTrafficLog.Ctx(ctx).Data(do.TV2RayUserTrafficLog{
		Email:     email,
		Ip:        ip,
		DateTime:  dataTime,
		Uplink:    uplink,
		Downlink:  downlink,
		CreatedAt: gtime.Now(),
	}).InsertAndGetId()
	if err != nil {
		global.Logger.Err(err).Msgf("insert TV2RayUserTrafficLog failed, email: %+v, dataTime: %s", email, dataTime)
		return err
	}
	global.Logger.Info().Msgf("insert TV2RayUserTrafficLog success, email: %+v, dataTime: %s, lastId: %d", email, dataTime, lastId)
	return nil
}

// UpdateUserTraffic 更新统计记录
func UpdateUserTraffic(ctx context.Context, item *entity.TV2RayUserTraffic, uplink, downlink uint64) error {
	affected, err := dao.TV2RayUserTraffic.Ctx(ctx).
		Where(do.TV2RayUserTraffic{
			Id: item.Id,
		}).
		UpdateAndGetAffected(do.TV2RayUserTraffic{
			Uplink:    item.Uplink + uplink,
			Downlink:  item.Downlink + downlink,
			UpdatedAt: gtime.Now(),
		})
	if err != nil {
		global.Logger.Err(err).Msgf("Update TV2RayUserTraffic failed, Email: %s, Traffic[%d, %d]",
			item.Email, uplink, downlink)
		return err
	}
	if affected != 1 {
		err = fmt.Errorf("update TV2RayUserTraffic failed, affected: %d, Email: %s, Traffic[%d, %d]",
			affected, item.Email, uplink, downlink)
		global.Logger.Error().Msgf(err.Error())
		return err
	}
	return nil
}

func AppendTrafficLog(dataLog string, dataTime string, uplink, downlink uint64) (string, error) {
	var log []TrafficItem
	err := json.Unmarshal([]byte(dataLog), &log)
	if err != nil {
		return "", err
	}
	log = append(log, TrafficItem{DateTime: dataTime, Uplink: uplink, Downlink: downlink})
	marshal, _ := json.Marshal(log)
	return string(marshal), nil
}

func NewTrafficLog(dataTime string, uplink, downlink uint64) string {
	data := []TrafficItem{{DateTime: dataTime, Uplink: uplink, Downlink: downlink}}
	marshal, _ := json.Marshal(data)
	return string(marshal)
}

type TrafficLog []TrafficItem

type TrafficItem struct {
	DateTime string `json:"t"`
	Uplink   uint64 `json:"up"`
	Downlink uint64 `json:"down"`
}
