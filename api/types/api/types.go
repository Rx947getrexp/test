package api

type DescribeUserInfoReq struct {
	Email  string `form:"email" json:"email" dc:"邮箱"`
	UserId uint64 `form:"user_id" json:"user_id" dc:"UID"`
}

type DescribeUserInfoRes struct {
	Id               int64  `json:"id"                 description:"自增id"`
	Uname            string `json:"uname"              description:"用户名"`
	Email            string `json:"email"              description:"邮件"`
	Phone            string `json:"phone"              description:"电话"`
	Level            int    `json:"level"              description:"等级：0-vip0；1-vip1；2-vip2"`
	ExpiredTime      int64  `json:"expired_time"       description:"vip到期时间"`
	V2RayUuid        string `json:"v2ray_uuid"         description:"节点UUID"`
	V2RayTag         int    `json:"v2ray_tag"          description:"v2ray存在UUID标签:1-有；2-无"`
	Channel          string `json:"channel"            description:""`
	ChannelId        int    `json:"channel_id"         description:"渠道id"`
	Status           int    `json:"status"             description:"冻结状态：0-正常；1-冻结"`
	CreatedAt        string `json:"created_at"         description:"创建时间"`
	UpdatedAt        string `json:"updated_at"         description:"更新时间"`
	Comment          string `json:"comment"            description:"备注信息"`
	ClientId         string `json:"client_id"          description:""`
	LastLoginIp      string `json:"last_login_ip"      description:"最近一次登录的ip"`
	LastLoginCountry string `json:"last_login_country" description:"最近一次登录的国家"`
	PreferredCountry string `json:"preferred_country"  description:"用户选择的国家（国家名称）"`
	Version          int    `json:"version"            description:"数据版本号"`
}

type DescribeNodeListReq struct {
}

type DescribeNodeListRes struct {
	Items []NodeItem `json:"items"`
}

type NodeItem struct {
	Id          int64  `json:"id"           description:"自增id"`
	Name        string `json:"name"         description:"节点名称"`
	Title       string `json:"title"        description:"节点标题"`
	CountryEn   string `json:"country_en"   description:"国家（英文）"`
	Ip          string `json:"ip"           description:"内网IP"`
	Server      string `json:"server"       description:"公网域名"`
	Port        int    `json:"port"         description:"公网端口"`
	MinPort     int    `json:"min_port"     description:"最小端口"`
	MaxPort     int    `json:"max_port"     description:"最大端口"`
	Path        string `json:"path"         description:"ws路径"`
	IsRecommend int    `json:"is_recommend" description:"推荐节点1-是；2-否"`
	ChannelId   int    `json:"channel_id"   description:"市场渠道（默认0）-优选节点有效"`
	Status      int    `json:"status"       description:"状态:1-正常；2-已软删"`
	Weight      uint   `json:"weight"       description:"权重"`
}
