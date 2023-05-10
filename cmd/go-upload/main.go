package main

import (
	"go-speed/global"
	"go-speed/initialize"
)

func main() {
	initialize.InitComponents()
	engine := initialize.RouterUpload()
	if err := engine.Run(":" + global.Config.System.Addr); err != nil {
		global.Logger.Err(err).Msg("启动失败")
	}
}
