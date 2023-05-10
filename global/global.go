package global

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"go-speed/config"
	"net/http"
	"runtime/debug"
	"xorm.io/xorm"
)

var (
	Db         *xorm.Engine
	Db2        *xorm.Engine
	Config     config.Server
	Logger     zerolog.Logger
	Viper      *viper.Viper
	Redis      *redis.Client
	Kafka      sarama.Client
	HttpClient *http.Client
)

var Recovery = func() {
	if r := recover(); r != nil {
		// 同时打印到日志文件和标准输出中
		Logger.Error().Msgf("%v\n\n%v", r, string(debug.Stack()))
		fmt.Printf("%v\n\n%v\n", r, string(debug.Stack()))
	}
}
