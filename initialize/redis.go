package initialize

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go-speed/global"
)

func initRedis() *redis.Client {
	if len(global.Config.Redis.Addr) == 0 {
		global.Logger.Warn().Msg("redis未配置")
		return nil
	}
	password := global.Config.Redis.Password
	rdb := redis.NewClient(&redis.Options{
		Addr:     global.Config.Redis.Addr,
		Password: password,
		DB:       global.Config.Redis.Db,
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		global.Logger.Err(err).Msgf("redis连接失败，err：%v", err)
		return nil
	}
	return rdb
}
