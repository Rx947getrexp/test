package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/glog"

	"go-speed/global"
	"go-speed/i18n"
	"go-speed/initialize"
)

func main() {
	glog.SetLevel(glog.LEVEL_ALL)

	initialize.InitComponents()
	i18n.Init()
	engine := initialize.ApiRouters()
	if err := engine.Run(":" + global.Config.System.Addr); err != nil {
		global.Logger.Err(err).Msg("启动失败")
		return
	}
}
