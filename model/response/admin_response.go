package response

// 登陆后台
type LoginAdminParam struct {
	UserId int64 `json:"user_id"`
	//IsReset   int               `json:"is_reset"`
	Token string `json:"token"` //token
	//PowerList []*PowerListParam `json:"power_list"`
}

type PowerDataParam struct {
	PowerList []*PowerListParam `json:"power_list"`
}

type PowerListParam struct {
	Path      string `json:"path"`
	Title     string `json:"title"`
	Check     string `json:"check"`
	Expansion string `json:"expansion"`
	Type      string `json:"type"`
	Id        string `json:"id"`
	ParentId  string `json:"parent_id"`
	Name      string `json:"name"`
	Sort      string `json:"sort"`
	Sign      string `json:"sign"`
}

type MenuTree struct {
	Id       int        `json:"id"`
	Pid      int        `json:"pid"`
	Name     string     `json:"name"`
	ResType  int        `json:"res_type"`
	Url      string     `json:"url"`
	Sort     int        `json:"sort"`
	Icon     string     `json:"icon"`
	Children []MenuTree `json:"children,omitempty"`
}

type Tree struct {
	Id       int    `json:"id"`
	Pid      int    `json:"pid"`
	Name     string `json:"name"`
	ResType  int    `json:"res_type"`
	Url      string `json:"url"`
	IsSelect bool   `json:"is_select"`
	Children []Tree `json:"children,omitempty"`
}

type UserRes struct {
	UserInfo UserInfo   `json:"user_info"`
	RoleInfo []RoleInfo `json:"role_info,omitempty"`
}

type RoleRes struct {
	RoleInfo RoleInfo `json:"role_info"`
	RoleTree []Tree
}

type UserRoleRes struct {
	UserInfo UserInfo   `json:"user_info"`
	RoleInfo []RoleInfo `json:"role_info"`
	RoleTree []Tree     `json:"role_tree"`
}

type RoleInfo struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Status     int    `json:"status"`
	Remark     string `json:"remark"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

type UserInfo struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	NickName   string `json:"nick_name"`
	CreateTime string `json:"create_time"`
	LastIp     string `json:"last_ip"`
	LastTime   string `json:"last_time"`
	Status     int    `json:"status"`
}

type GenerateAuth2KeyAdminResponse struct {
	Auth2Key string `json:"auth2_key"` //谷歌两步验证器私钥
}

type GetReportUserDayListResponse struct {
	Total int64           `json:"total" dc:"数据总条数"`
	Items []ReportUserDay `json:"items" dc:"数据明细"`
}

type ReportUserDay struct {
	Id        int64  `json:"id" dc:"自增主键ID"`
	Date      int    `json:"date" dc:"报表日期，eg:20240101"`
	ChannelId int    `json:"channel_id" dc:"渠道ID"`
	Total     int    `json:"total" dc:"截止到Date，通过渠道ID注册的用户总量"`
	New       int    `json:"new" dc:"Date日期，通过渠道ID注册的新增用户量"`
	Retained  int    `json:"retained" dc:"通过渠道ID注册的用户，Date日期内有使用APP的用户量（留存）"`
	CreatedAt string `json:"created_at" dc:"报表数据统计时间"`
}

type GetChannelUserDayListResponse struct {
	Total int64            `json:"total" dc:"数据总条数"`
	Items []ChannelUserDay `json:"items" dc:"数据明细"`
}

type ChannelUserDay struct {
	Id        int64  `json:"id" dc:"自增主键ID"`
	Date      int    `json:"date" dc:"报表日期，eg:20240101"`
	Channel   string `json:"channel" dc:"渠道ID"`
	Total     int    `json:"total" dc:"截止到Date，通过渠道ID注册的用户总量"`
	New       int    `json:"new" dc:"Date日期，通过渠道ID注册的新增用户量"`
	Retained  int    `json:"retained" dc:"通过渠道ID注册的用户，Date日期内有使用APP的用户量（留存）"`
	CreatedAt string `json:"created_at" dc:"报表数据统计时间"`
}
type GetNodeDayListResponse struct {
	Total int64         `json:"total" dc:"数据总条数"`
	Items []NodeUserDay `json:"items" dc:"数据明细"`
}
type NodeUserDay struct {
	Id        int64  `json:"id" dc:"自增主键ID"`
	Date      int    `json:"date" dc:"报表日期，eg:20240101"`
	Ip        string `json:"ip" dc:"节点ip"`
	Total     int    `json:"total" dc:"截止到Date，通过渠道ID注册的用户总量"`
	New       int    `json:"new" dc:"Date日期，通过渠道ID注册的新增用户量"`
	Retained  int    `json:"retained" dc:"通过渠道ID注册的用户，Date日期内有使用APP的用户量（留存）"`
	CreatedAt string `json:"created_at" dc:"报表数据统计时间"`
}
type GetOnlineUserDayListResponse struct {
	Total int64           `json:"total" dc:"数据总条数"`
	Items []OnlineUserDay `json:"items" dc:"数据明细"`
}

type OnlineUserDay struct {
	Id               int64  `json:"id" dc:"自增主键ID"`
	Date             int    `json:"date" dc:"报表日期，eg:20240101"`
	Email            string `json:"email" dc:"账号email"`
	Channel          string `json:"channel" dc:"渠道ID"`
	OnlineDuration   int    `json:"online_duration" dc:"用户在线时间戳长度，单位：秒"`
	Uplink           int64  `json:"uplink" dc:"上行流量，单位：字节"`
	Downlink         int64  `json:"downlink" dc:"下行流量，单位：字节"`
	CreatedAt        string `json:"created_at" dc:"报表数据统计时间"`
	LastLoginCountry string `json:"last_login_country" dc:"最后登陆的国家"`
}
type GetNodeOnlineUserDayListResponse struct {
	Total int64               `json:"total" dc:"数据总条数"`
	Items []NodeOnlineUserDay `json:"items" dc:"数据明细"`
}

type NodeOnlineUserDay struct {
	Id             int64  `json:"id" dc:"自增主键ID"`
	Date           int    `json:"date" dc:"报表日期，eg:20240101"`
	Email          string `json:"email" dc:"账号email"`
	Channel        string `json:"channel" dc:"渠道ID"`
	OnlineDuration int    `json:"online_duration" dc:"用户在线时间戳长度，单位：秒"`
	Uplink         int64  `json:"uplink" dc:"上行流量，单位：字节"`
	Downlink       int64  `json:"downlink" dc:"下行流量，单位：字节"`
	Node           string `json:"node" dc:"国家节点"`
	RegisterDate   string `json:"register_date" dc:"用户注册最早时间"`
	CreatedAt      string `json:"created_at" dc:"报表数据统计时间"`
}
