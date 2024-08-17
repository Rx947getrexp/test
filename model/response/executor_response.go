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

type GetV2raySysStatsResponse struct {
	NumGoroutine uint32 `json:"num_goroutine"`
	NumGC        uint32 `json:"num_gc"`
	Alloc        uint64 `json:"alloc"`
	TotalAlloc   uint64 `json:"total_alloc"`
	Sys          uint64 `json:"sys"`
	Mallocs      uint64 `json:"mallocs"`
	Frees        uint64 `json:"frees"`
	LiveObjects  uint64 `json:"live_objects"`
	PauseTotalNs uint64 `json:"pause_total_ns"`
	Uptime       uint32 `json:"uptime"`
}
