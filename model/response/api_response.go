package response

type LoginClientParam struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"` //token
	//UserInfo  *ClientUserInfo `json:"user_info"`
}
