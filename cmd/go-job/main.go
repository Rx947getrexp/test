package main

import (
	"fmt"
	"go-speed/global"
	"go-speed/initialize"
	"time"
)

func main() {
	initialize.InitComponents()
	go a()
	go b()
	engine := initialize.JobRouters()
	if err := engine.Run(":" + global.Config.System.Addr); err != nil {
		global.Logger.Err(err).Msg("启动失败")
		return
	}
}

func a() {
	fmt.Println("当前时间戳1：", time.Now().Unix())
}

func b() {
	fmt.Println("当前时间戳2：", time.Now().Unix())
}
