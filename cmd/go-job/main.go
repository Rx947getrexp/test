package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"go-speed/global"
	"go-speed/initialize"
	"go-speed/task"
)

func main() {
	initialize.InitComponents()
	go task.DeleteExpiredUser()
	//go task.CollectUserTraffic()
	go task.SyncPayOrderStatus()
	engine := initialize.JobRouters()
	if err := engine.Run(":" + global.Config.System.Addr); err != nil {
		global.Logger.Err(err).Msg("启动失败")
		return
	}
}
