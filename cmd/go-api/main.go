package main

import (
	"go-speed/global"
	"go-speed/i18n"
	"go-speed/initialize"
)

func main() {
	initialize.InitComponents()
	i18n.Init()
	engine := initialize.ApiRouters()
	if err := engine.Run(":" + global.Config.System.Addr); err != nil {
		global.Logger.Err(err).Msg("启动失败")
		return
	}
}
