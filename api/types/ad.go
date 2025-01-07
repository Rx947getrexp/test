package types

type SlotLocationItem struct {
	Location string `json:"location" dc:"广告位的位置"`
	Sort     int    `json:"sort" dc:"在广告位置中的排序"`
}

type TargetUrlItem struct {
	Channel string `json:"channel" dc:"渠道，enum: pc, android, ios"`
	Url     string `json:"url" dc:"跳转地址"`
}
