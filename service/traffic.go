package service

import (
	"encoding/json"
	"fmt"
	"go-speed/global"
	"go-speed/model"
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

func GetUserTrafficByEmail(email, ip string, date int) (*model.TUserTraffic, error) {
	userTraffic := new(model.TUserTraffic)
	ok, err := global.Db.Where("email = ? and date = ? and ip = ?", email, date, ip).Get(userTraffic)
	if err != nil {
		global.Logger.Err(err).Msgf("get user traffic failed. email: %s, ip: %s, date: %s", email, ip, date)
		return nil, err
	}
	if !ok {
		return nil, nil
	}
	return userTraffic, nil
}

// CreateUserTraffic 插入新的用量统计记录
func CreateUserTraffic(email, ip string, date int, uplink, downlink uint64) error {
	traffic := &model.TUserTraffic{
		Email:     email,
		Ip:        ip,
		Date:      date,
		Uplink:    uplink,
		Downlink:  downlink,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	rows, err := global.Db.Insert(traffic)
	if err != nil {
		global.Logger.Err(err).Msgf("insert TUserTraffic failed, traffic: %+v", *traffic)
		return err
	}
	if rows != 1 {
		return fmt.Errorf("insert TUserTraffic failed，rows:%d is not eq 1, traffic: %+v", rows, *traffic)
	}
	return nil
}

// CreateUserTrafficLog 插入新的用量流水
func CreateUserTrafficLog(email, ip, dataTime string, uplink, downlink uint64) error {
	trafficLog := &model.TUserTrafficLog{
		Email:     email,
		Ip:        ip,
		DateTime:  dataTime,
		Uplink:    uplink,
		Downlink:  downlink,
		CreatedAt: time.Now(),
	}
	rows, err := global.Db.Insert(trafficLog)
	if err != nil {
		global.Logger.Err(err).Msgf(">>>>>>>>>>> insert TUserTrafficLog failed, trafficLog: %+v", *trafficLog)
		return err
	}
	if rows != 1 {
		return fmt.Errorf(">>>>>>>>>>> insert TUserTrafficLog failed，rows:%d is not eq 1, trafficLog: %+v", rows, *trafficLog)
	}
	return nil
}

// UpdateUserTraffic 更新统计记录
func UpdateUserTraffic(item *model.TUserTraffic, uplink, downlink uint64) error {
	traffic := &model.TUserTraffic{
		Uplink:    item.Uplink + uplink,
		Downlink:  item.Downlink + downlink,
		UpdatedAt: time.Now(),
	}
	rows, err := global.Db.Where("id = ?", item.Id).Update(traffic)
	if err != nil {
		global.Logger.Err(err).Msgf("Update User traffic failed, Email: %s, Traffic[%d, %d]",
			item.Email, uplink, downlink)
		return err
	}
	if rows != 1 {
		global.Logger.Error().Msgf("Update User traffic failed, rows: %d, Email: %s, Traffic[%d, %d]",
			rows, item.Email, uplink, downlink)
		return fmt.Errorf("update TUserTraffic failed，rows:%d is not eq 1, traffic: %+v", rows, *traffic)
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
