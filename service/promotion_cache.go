package service

import (
	"context"
	"fmt"
	"go-speed/constant"
	"go-speed/global"
)

func ResetPromotionDnsCache() {
	ctx := context.Background()
	keyProfix := fmt.Sprintf("%s*", constant.PromotionDnsMappingCacheKey)
	keys, _ := global.Redis.Keys(ctx, keyProfix).Result()
	for _, key := range keys {
		global.Redis.Del(ctx, key)
	}
}

func ResetPromotionShopCache() {
	ctx := context.Background()
	keyProfix := fmt.Sprintf("%s*", constant.PromotionStoreCacheKey)
	keys, _ := global.Redis.Keys(ctx, keyProfix).Result()
	for _, key := range keys {
		global.Redis.Del(ctx, key)
	}
}
