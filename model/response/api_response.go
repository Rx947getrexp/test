package response

type LoginClientParam struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"` //token
	//UserInfo  *ClientUserInfo `json:"user_info"`
}

type UserInfoResponse struct {
	Id          int64  `json:"id"`
	Uname       string `json:"uname"`
	MemberType  int    `json:"member_type"`
	ExpiredTime int64  `json:"expired_time"`
	SurplusFlow int64  `json:"surplus_flow"`
}

type TeamListResponse struct {
}

type TeamInfoResponse struct {
}

type NoticeListResponse struct {
}

type NoticeDetailResponse struct {
}
