package response

type LoginClientParam struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"` //token
	//UserInfo  *ClientUserInfo `json:"user_info"`
}

type UserInfoResponse struct {
	Id          int64  `json:"id"`
	Uname       string `json:"uname"`
	Uuid        string `json:"uuid"`
	MemberType  int    `json:"member_type"`
	ExpiredTime int64  `json:"expired_time"`
	SurplusFlow int64  `json:"surplus_flow"`
}

type TeamListResponse struct {
	Uname       string `json:"uname"`
	MemberType  int    `json:"member_type"`
	CreatedTime string `json:"created_time"`
}

type TeamInfoResponse struct {
	Fans       int64       `json:"fans"`
	AwardHour  string      `json:"award_hour"`
	AwardMoney string      `json:"award_money"`
	AwardList  []AwardInfo `json:"award_list"`
}

type AwardInfo struct {
	Uname   string `json:"uname"`
	Title   string `json:"title"`
	GiftSec int    `json:"gift_sec"`
	TimeStr string `json:"time_str"`
}

type ComboListResponse struct {
	MType     int         `json:"m_type"`
	ComboList []ComboInfo `json:"combo_list"`
}

type ComboInfo struct {
}

type ListNodeForReport struct {
	Items []NodeItem `json:"items"`
}

type NodeItem struct {
	Ip string `json:"ip"`
}

type PromotionDnsRes struct {
	AndroidChannel string `json:"android_channel" dc:"安卓渠道"`
	WinChannel     string `json:"win_channel" dc:"windows电脑渠道"`
	MacChannel     string `json:"mac_channel" dc:"苹果电脑渠道"`
	CreatedAt      string `json:"created_at" dc:"创建时间"`
	UpdatedAt      string `json:"updated_at" dc:"更新时间"`
}

type PromotionDnsResponse struct {
	List []PromotionDnsRes `json:"list" dc:"dns列表"`
}

type PromotionShopRes struct {
	TitleCn string `form:"title_cn" json:"title" dc:"商店标题(中文)"`
	TitleEn string `form:"title_en" json:"title_en" dc:"商店标题(英文)"`
	TitleRu string `form:"title_ru" json:"title_ru" dc:"商店标题(俄语)"`
	Type    string `form:"type" json:"type" dc:"商店类型，苹果：ios，安卓：android"`
	Url     string `form:"url" json:"url" dc:"商店地址"`
	Cover   string `form:"cover" json:"cover" dc:"商店图标"`
}

type PromotionShopResponse struct {
	List []PromotionShopRes `json:"list" dc:"数据明细"`
}
