package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"go-speed/global"
	"go-speed/initialize"
	"go-speed/service"
)

func main() {
	initialize.InitComponents()
	service.UpdateSysCache()
	engine := initialize.AdminRouters()
	if err := engine.Run(":" + global.Config.System.Addr); err != nil {
		global.Logger.Err(err).Msg("启动失败")
		return
	}
}
