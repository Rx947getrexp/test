package response

type GetUserListResponse struct {
	Items []ClientItem `json:"items"`
}

type ClientItem struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUserTrafficResponse struct {
	Items []UserTrafficItem `json:"items"`
}

type UserTrafficItem struct {
	Email    string `json:"email"`
	UpLink   uint64 `json:"up_link"`
	DownLink uint64 `json:"down_link"`
}
