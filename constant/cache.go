package constant

import "time"

const (
	MenuFullTreeKey = "menuFullTree"
	RoleTreeMapKey  = "roleTreeMap"
	RoleUrlMapKey   = "roleUrlMap"
	FullTreeKey     = "fullTree"

	DictCacheMapKey = "dictCacheMap"
)

// 获取后台配置的推广人员与渠道映射关系 缓存配置
const (
	PromotionDnsMappingCacheKey   = "promotionDnsMappingCacheKey:"
	PromotionDnsMappingExpiration = 24 * time.Hour
)
