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
	Dns            string `json:"dns" dc:"机器域名"`
	AndroidChannel string `json:"android_channel" dc:"安卓渠道"`
	WinChannel     string `json:"win_channel" dc:"windows电脑渠道"`
	MacChannel     string `json:"mac_channel" dc:"苹果电脑渠道"`
}

type PromotionDnsResponse struct {
	Total int               `json:"total" dc:"数据总条数"`
	List  []PromotionDnsRes `json:"list" dc:"数据明细"`
}
