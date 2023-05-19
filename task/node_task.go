package task

import (
	"go-speed/constant"
	"go-speed/global"
	"go-speed/service"
	"time"
)

func NodeHeartbeatTask() {
	global.Recovery()
	global.Logger.Info().Msg("NodeHeartbeatTask start...")
	for {
		service.Heartbeat()
		time.Sleep(time.Second * constant.HeartbeatTime)
	}
}

func NodeReportDataTask() {
	global.Recovery()
	global.Logger.Info().Msg("NodeReportDataTask start...")
	for {
		service.ReportData()
		time.Sleep(time.Second * constant.ReportDataTime)
	}
}
