package initialize

import (
	"github.com/rs/zerolog"
	"go-speed/global"
	"time"
)

// InitComponents 初始化基本组件
func InitComponents() {
	// 日志打印的时间格式
	zerolog.TimeFieldFormat = time.RFC3339
	global.Viper = initViper()
	global.Logger = InitLog()
	global.Db = initMysqlDb(global.Config.Db)
	if len(global.Config.Db2.Host) > 0 {
		// 初始化第二个数据源
		global.Db2 = initMysqlDb(global.Config.Db2)
	}
	global.Redis = initRedis()
	//global.Kafka = initKafkaClient()
	global.HttpClient = initClient()

}

func InitComponentsV2() {
	// 日志打印的时间格式
	zerolog.TimeFieldFormat = time.RFC3339
	global.Viper = initViper()
	global.Logger = InitLog()
}
