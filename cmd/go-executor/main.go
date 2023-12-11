package main

import (
	"go-speed/api/executor"
	"go-speed/global"
	"go-speed/initialize"
	"google.golang.org/grpc"
	"time"
)

func main() {
	initialize.InitComponentsV2()
	engine := initialize.ExecutorRouters()
	if err := engine.Run(":" + global.Config.System.Addr); err != nil {
		global.Logger.Err(err).Msg("启动失败")
		return
	}
}

func CountUserTraffic() {
	userList := []string{
		"yyy1@qq.com",
		"a1@qq.com",
		"zzz@qq.com",
		"a101@qq.com",
		"well@gmail.com",
		"303468504456@gmail.com",
	}
	// "127.0.0.1:10088"
	conn, err := grpc.Dial(global.Config.System.V2rayApiAddress, grpc.WithInsecure())
	if err != nil {
		global.Logger.Err(err).Msg("grpc.Dial failed, err: " + err.Error())
		panic(err)
	}
	defer conn.Close()
	var index uint64
	for {
		index = index + 1

		global.Logger.Info().Msgf("-------------------------------------index: %d", index)
		err = executor.GetSysTraffic(conn)
		if err != nil {
			global.Logger.Err(err).Msg(" get sys traffic failed")
		}
		global.Logger.Info().Msg("###############################################################")
		for _, user := range userList {
			uplink, downlink, err := executor.GetUserTrafficByEmail(conn, user)
			if err != nil {
				global.Logger.Err(err).Msg(user + " get traffic failed")
			} else {
				global.Logger.Info().Msgf("------------- %s uplink: %d, downlink: %d", user, uplink, downlink)
			}
		}
		time.Sleep(time.Second * 10)
	}
}
