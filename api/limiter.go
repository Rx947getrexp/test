package api

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go-speed/global"
	"time"
)

// Limiter 定义属性
type Limiter struct {
	// Redis client connection.
	rc *redis.Client
}

// NewLimiter 根据redisURL创建新的limiter并返回
func NewLimiter(redisURL string) (*Limiter, error) {
	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}

	rc := redis.NewClient(opts)
	if err := rc.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return &Limiter{rc: rc}, nil
}

func NewLimiterV2() *Limiter {
	return &Limiter{rc: global.Redis}
}

// Allow 通过redis的value判断第几次访问并返回是否允许访问
func (l *Limiter) Allow(key string, events int64, per time.Duration) bool {
	curr := l.rc.LLen(context.Background(), key).Val()
	if curr >= events {
		return false
	}

	if v := l.rc.Exists(context.Background(), key).Val(); v == 0 {
		pipe := l.rc.TxPipeline()
		pipe.RPush(context.Background(), key, key)
		//设置过期时间
		pipe.Expire(context.Background(), key, per)
		_, _ = pipe.Exec(context.Background())
	} else {
		l.rc.RPushX(context.Background(), key, key)
	}

	return true
}
