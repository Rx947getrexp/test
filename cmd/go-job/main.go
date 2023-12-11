package main

import (
	"go-speed/global"
	"go-speed/initialize"
	"go-speed/task"
)

func main() {
	initialize.InitComponents()
	go task.DeleteExpiredUser()
	go task.CollectUserTraffic()
	engine := initialize.JobRouters()
	if err := engine.Run(":" + global.Config.System.Addr); err != nil {
		global.Logger.Err(err).Msg("启动失败")
		return
	}
}
