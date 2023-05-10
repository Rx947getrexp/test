package initialize

import (
	"fmt"
	"github.com/rs/zerolog"
	"go-speed/global"
	"go-speed/util"
)

func initLog() zerolog.Logger {
	writer := util.GetLogWriter(global.Config.Log)
	level, err := zerolog.ParseLevel(global.Config.Log.Level)
	if err != nil {
		fmt.Println("日志等级配置错误，默认使用Info级别", err)
		level = zerolog.InfoLevel
	}
	return zerolog.New(writer).Level(level).With().Caller().Timestamp().Logger()
}
