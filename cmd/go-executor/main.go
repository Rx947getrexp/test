package main

import (
	"go-speed/global"
	"go-speed/initialize"
	"go-speed/task"
)

func main() {
	initialize.InitComponentsV2()
	go task.NodeHeartbeatTask()
	go task.NodeReportDataTask()
	engine := initialize.ExecutorRouters()
	if err := engine.Run(":" + global.Config.System.Addr); err != nil {
		global.Logger.Err(err).Msg("启动失败")
		return
	}
}
