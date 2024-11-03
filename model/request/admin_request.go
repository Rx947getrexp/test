package request

// 登陆后台
type LoginAdminRequest struct {
	UserName string `form:"user_name" binding:"required" json:"user_name"`
	Pass     string `form:"pass" binding:"required" json:"pass"`
}

type EditPasswdRequest struct {
	OldPass   string `form:"old_pass" binding:"required" json:"old_pass"`
	NewPass   string `form:"new_pass" binding:"required" json:"new_pass"`
	EnterPass string `form:"enter_pass" binding:"required" json:"enter_pass"`
}

type AddResourceRequest struct {
	Pid     int    `form:"pid" json:"pid"`
	Name    string `form:"name" binding:"required" json:"name"`
	Url     string `form:"url" binding:"required" json:"url"`
	ResType int    `form:"res_type" binding:"required" json:"res_type"`
	Icon    string `form:"icon" json:"icon"`
}

type EditResourceRequest struct {
	Id      int    `form:"id" binding:"required" json:"id"`
	Pid     int    `form:"pid" json:"pid"`
	Name    string `form:"name" binding:"required" json:"name"`
	Url     string `form:"url" binding:"required" json:"url"`
	ResType int    `form:"res_type" binding:"required" json:"res_type"`
	Icon    string `form:"icon" json:"icon"`
}

type DelResourceRequest struct {
	Id int `form:"id" binding:"required" json:"id"`
}

type AddRoleRequest struct {
	Name   string `form:"name" binding:"required" json:"name"`
	Remark string `form:"remark" json:"remark"`
	IsUsed int    `form:"is_used" binding:"required" json:"is_used"`
	ResIds string `form:"res_ids" json:"res_ids"`
}

type EditRoleRequest struct {
	Id     int    `form:"id" binding:"required" json:"id"`
	Name   string `form:"name" binding:"required" json:"name"`
	Remark string `form:"remark" json:"remark"`
	IsUsed int    `form:"is_used" binding:"required" json:"is_used"`
	ResIds string `form:"res_ids" json:"res_ids"`
}

type AddUserRequest struct {
	Username string `form:"username" binding:"required" json:"username"`
	Password string `form:"password" binding:"required" json:"password"`
	Nickname string `form:"nickname" binding:"required" json:"nickname"`
	RoleIds  string `form:"role_ids" json:"role_ids"`
}

type EditUserRoleRequest struct {
	UserId   int    `form:"user_id" binding:"required" json:"user_id"`
	Password string `form:"password" json:"password"`
	Nickname string `form:"nickname" json:"nickname"`
	Status   int    `form:"status" json:"status"`
	RoleIds  string `form:"role_ids" json:"role_ids"`
}

type UserRoleRequest struct {
	Id int `form:"id" binding:"required" json:"id"`
}

type AccountListAdminRequest struct {
	Account  string `form:"account" json:"account"`
	NickName string `form:"nick_name" json:"nick_name"`
	Page     int    `form:"page" binding:"required" json:"page"`
	Size     int    `form:"size" binding:"required" json:"size"`
}

type RoleListAdminRequest struct {
	RoleName string `form:"role_name" json:"role_name"`
	Page     int    `form:"page" binding:"required" json:"page"`
	Size     int    `form:"size" binding:"required" json:"size"`
}

type AccountEditAdminRequest struct {
	AccountId int64  `form:"account_id" binding:"required" json:"account_id"`
	Password  string `form:"password" json:"password"`
	NickName  string `form:"nick_name" json:"nick_name"`
	RoleId    int    `form:"role_id" json:"role_id"`
	IsDel     string `form:"is_del" json:"is_del"`
	Status    string `form:"status" json:"status"`
	IsReset   string `form:"is_reset" json:"is_reset"`
	Auth2Code string `form:"auth2_code" json:"auth2_code"`
	Channel   string `form:"channel" json:"channel"`
}

type AccountAddAdminRequest struct {
	Account   string `form:"account" binding:"required" json:"account"`
	Password  string `form:"password" binding:"required" json:"password"`
	NickName  string `form:"nick_name" binding:"required" json:"nick_name"`
	RoleId    int    `form:"role_id" binding:"required" json:"role_id"`
	Auth2Code string `form:"auth2_code" json:"auth2_code"`
	Channel   string `form:"channel" json:"channel"`
}

type RoleEditAdminRequest struct {
	Id     int64  `form:"id" binding:"required" json:"id"`
	Name   string `form:"name"  json:"name"`
	Remark string `form:"remark" json:"remark"`
	ResIds string `form:"res_ids" json:"res_ids"`
	IsDel  string `form:"is_del" json:"is_del"`
}

type RoleAddAdminRequest struct {
	Name   string `form:"name" binding:"required" json:"name"`
	Remark string `form:"remark" json:"remark"`
	ResIds string `form:"res_ids" binding:"required" json:"res_ids"`
}

type ChangeAuth2KeyAdminRequest struct {
	//SmsCode   string `form:"sms_code" binding:"required" json:"sms_code"`     //sms验证码
	Auth2Code string `form:"auth2_code" binding:"required" json:"auth2_code"` //谷歌验证码
	Auth2Key  string `form:"auth2_key" binding:"required" json:"auth2_key"`   //谷歌两步验证器私钥
}

type HeartbeatAdminRequest struct {
	//NodeId      int64   `form:"node_id" binding:"required" json:"node_id"`
	//NodeIp 	    string	`form:"node_ip" binding:"required" json:"node_ip"`	//内网IP
	//NodeName    string	`form:"node_name" binding:"required" json:"node_name"`
	NodeVersion string `form:"node_version" binding:"required" json:"node_version"`
	//Disk			int64	`form:"disk" binding:"required" json:"disk"`
	//Memory			int64	`form:"memory" binding:"required" json:"memory"`
	//Cpu				int64	`form:"cpu" binding:"required" json:"cpu"`
	//Net         	int64	`form:"net" binding:"required" json:"net"`
}

type ReportDataAdminRequest struct {
	//NodeId      	int64   `form:"node_id" binding:"required" json:"node_id"`
	//NodeIp 	    	string	`form:"node_ip" binding:"required" json:"node_ip"`	//内网IP
	DiskUsed    int64 `form:"disk_used" binding:"required" json:"disk_used"`
	MemoryUsed  int64 `form:"memory_used" binding:"required" json:"memory_used"`
	CpuUsed     int64 `form:"cpu_used" binding:"required" json:"cpu_used"`
	NetFlowUsed int64 `form:"net_flow_used" binding:"required" json:"net_flow_used"`
}

type AddAdAdminRequest struct {
	Name    string `form:"name" binding:"required" json:"name"`
	Logo    string `form:"logo" binding:"required" json:"logo"`
	Link    string `form:"link" binding:"required" json:"link"`
	Tag     string `form:"tag" binding:"required" json:"tag"`
	Content string `form:"content" binding:"required" json:"content"`
	AdType  int    `form:"ad_type" binding:"required" json:"ad_type"`
}

type EditAdAdminRequest struct {
	Id      int64  `form:"id" binding:"required" json:"id"`
	Name    string `form:"name" json:"name"`
	Logo    string `form:"logo" json:"logo"`
	Link    string `form:"link" json:"link"`
	Tag     string `form:"tag" json:"tag"`
	Content string `form:"content" json:"content"`
	AdType  int    `form:"ad_type" json:"ad_type"`
	Status  int    `form:"status" json:"status"`
}

type AdListAdminRequest struct {
	Name   string `form:"name" json:"name"`
	Tag    string `form:"tag" json:"tag"`
	AdType int    `form:"ad_type" json:"ad_type"`
	Page   int    `form:"page" binding:"required" json:"page"`
	Size   int    `form:"size" binding:"required" json:"size"`
}

type AddNoticeAdminRequest struct {
	Title   string `form:"title" binding:"required" json:"title"`
	Tag     string `form:"tag" binding:"required" json:"tag"`
	Content string `form:"content" binding:"required" json:"content"`
}

type EditNoticeAdminRequest struct {
	Id      int64  `form:"id" binding:"required" json:"id"`
	Title   string `form:"title" json:"title"`
	Tag     string `form:"tag" json:"tag"`
	Content string `form:"content" json:"content"`
	Status  int    `form:"status" json:"status"`
}

type NoticeListAdminRequest struct {
	Title string `form:"title" json:"title"`
	Tag   string `form:"tag" json:"tag"`
	Page  int    `form:"page" binding:"required" json:"page"`
	Size  int    `form:"size" binding:"required" json:"size"`
}

type AddGoodsAdminRequest struct {
	Title            string  `form:"title" binding:"required" json:"title"`
	MType            int     `form:"m_type" binding:"required" json:"m_type"`
	DevLimit         int     `form:"dev_limit" binding:"required" json:"dev_limit"`
	FlowLimit        int64   `form:"flow_limit" json:"flow_limit"`
	Period           int     `form:"period" binding:"required" json:"period"`
	Price            float64 `form:"price" binding:"required" json:"price"`
	PriceUnit        string  `form:"price_unit" binding:"required" json:"price_unit"`
	UsdPayPrice      float64 `form:"usd_pay_price" binding:"required" json:"usd_pay_price"`
	WebmoneyPayPrice float64 `form:"webmoney_pay_price" binding:"required" json:"webmoney_pay_price"`
	UsdPriceUnit     string  `form:"usd_price_unit" binding:"required" json:"usd_price_unit"`
	IsDiscount       int     `form:"is_discount" binding:"required" json:"is_discount"`
	Low              int     `form:"low" json:"low"`
	High             int     `form:"high" json:"high"`
	PriceRUB         float64 `form:"price_rub" binding:"required" json:"price_rub" dc:"卢布价格"`
	PriceWMZ         float64 `form:"price_wmz" binding:"required" json:"price_wmz" dc:"WMZ价格"`
	PriceUSD         float64 `form:"price_usd" binding:"required" json:"price_usd" dc:"USD价格"`
	PriceUAH         float64 `form:"price_uah" binding:"required" json:"price_uah" dc:"UAH价格"`
}

type EditGoodsAdminRequest struct {
	Id               int64   `form:"id" binding:"required" json:"id"`
	Title            string  `form:"title" json:"title"`
	MType            int     `form:"m_type" json:"m_type"`
	DevLimit         int     `form:"dev_limit" json:"dev_limit"`
	FlowLimit        int64   `form:"flow_limit" json:"flow_limit"`
	Period           int     `form:"period" json:"period"`
	Price            float64 `form:"price" json:"price"`
	UsdPayPrice      float64 `form:"usd_pay_price" json:"usd_pay_price"`
	WebmoneyPayPrice float64 `form:"webmoney_pay_price" json:"webmoney_pay_price"`
	PriceUnit        string  `form:"price_unit" json:"price_unit"`
	UsdPriceUnit     string  `form:"usd_price_unit" json:"usd_price_unit"`
	IsDiscount       int     `form:"is_discount" json:"is_discount"`
	Low              int     `form:"low" json:"low"`
	High             int     `form:"high" json:"high"`
	Status           int     `form:"status" json:"status"`
	PriceRUB         float64 `form:"price_rub" json:"price_rub"`
	PriceWMZ         float64 `form:"price_wmz" json:"price_wmz"`
	PriceUSD         float64 `form:"price_usd" json:"price_usd"`
	PriceUAH         float64 `form:"price_uah" json:"price_uah"`
}

type GoodsListAdminRequest struct {
	Title string `form:"title" json:"title"`
	Page  int    `form:"page" binding:"required" json:"page"`
	Size  int    `form:"size" binding:"required" json:"size"`
}

type AddNodeAdminRequest struct {
	Title   string `form:"title" binding:"required" json:"title"`
	Name    string `form:"name" binding:"required" json:"name"`
	Country string `form:"country" binding:"required" json:"country"`
	Ip      string `form:"ip" binding:"required" json:"ip"`
	Server  string `form:"server" binding:"required" json:"server"`
	Port    int    `form:"port" binding:"required" json:"port"`
	Cpu     int    `form:"cpu" json:"cpu"`
	Flow    int64  `form:"flow" json:"flow"`
	Disk    int64  `form:"disk" json:"disk"`
	Memory  int64  `form:"memory" json:"memory"`
}

type EditNodeAdminRequest struct {
	Id          int64  `form:"id" binding:"required" json:"id"`
	Title       string `form:"title" json:"title"`
	Name        string `form:"name" json:"name"`
	Country     string `form:"country" json:"country"`
	Ip          string `form:"ip" json:"ip"`
	Server      string `form:"server" json:"server"`
	Port        int    `form:"port" json:"port"`
	Cpu         int    `form:"cpu" json:"cpu"`
	Flow        int64  `form:"flow" json:"flow"`
	Disk        int64  `form:"disk" json:"disk"`
	Memory      int64  `form:"memory" json:"memory"`
	Status      int    `form:"status" json:"status"`
	IsRecommend int    `form:"is_recommend" json:"is_recommend"`
}

type NodeListAdminRequest struct {
	Id      int64  `form:"id" json:"id"`
	Title   string `form:"title" json:"title"`
	Name    string `form:"name" json:"name"`
	Country string `form:"country" json:"country"`
	Ip      string `form:"ip" json:"ip"`
	Page    int    `form:"page" binding:"required" json:"page"`
	Size    int    `form:"size" binding:"required" json:"size"`
}

type NodeUuidListAdminRequest struct {
	Id       int64  `form:"id" json:"id"`
	UName    string `form:"uname" json:"uname"`
	NodeName string `form:"node_name" json:"node_name"`
	NodeId   int64  `form:"node_id" json:"node_id"`
	Uuid     string `form:"uuid" json:"uuid"`
	Page     int    `form:"page" binding:"required" json:"page"`
	Size     int    `form:"size" binding:"required" json:"size"`
}

type OrderListAdminRequest struct {
	Id        int64  `form:"id" json:"id"`
	UserId    int64  `form:"user_id" json:"user_id"`
	Uname     string `form:"uname" json:"uname"`
	StartTime string `form:"start_time" json:"start_time"`
	EndTime   string `form:"end_time" json:"end_time"`
	Page      int    `form:"page" binding:"required" json:"page"`
	Size      int    `form:"size" binding:"required" json:"size"`
}

type SiteListAdminRequest struct {
	Site string `form:"site" json:"site"`
	Ip   string `form:"ip" json:"ip"`
	Page int    `form:"page" binding:"required" json:"page"`
	Size int    `form:"size" binding:"required" json:"size"`
}

type AddSiteAdminRequest struct {
	Site string `form:"site" binding:"required" json:"site"`
	Ip   string `form:"ip" binding:"required" json:"ip"`
}

type EditSiteAdminRequest struct {
	Id     int64  `form:"id" binding:"required" json:"id"`
	Site   string `form:"site" json:"site"`
	Ip     string `form:"ip" json:"ip"`
	Status int    `form:"status" json:"status"`
}

type DictDetailAdminRequest struct {
	FilterPac    string `form:"filter_pac" binding:"required" json:"filter_pac"`
	FilterRefuse string `form:"filter_refuse" binding:"required" json:"filter_refuse"`
}

type DictEditAdminRequest struct {
	FilterPac    string `form:"filter_pac" binding:"required" json:"filter_pac"`
	FilterRefuse string `form:"filter_refuse" binding:"required" json:"filter_refuse"`
}

type GiftListAdminRequest struct {
	UserId int64  `form:"user_id" json:"user_id"`
	Uname  string `form:"uname" json:"uname"`
	GType  int    `form:"g_type" json:"g_type"`
	Page   int    `form:"page" binding:"required" json:"page"`
	Size   int    `form:"size" binding:"required" json:"size"`
}

type ActivityListAdminRequest struct {
	UserId int64  `form:"user_id" json:"user_id"`
	Uname  string `form:"uname" json:"uname"`
	Page   int    `form:"page" binding:"required" json:"page"`
	Size   int    `form:"size" binding:"required" json:"size"`
}

type SpeedLogsAdminRequest struct {
	UserId    int64  `form:"user_id" json:"user_id"`
	Uname     string `form:"uname" json:"uname"`
	DevId     int64  `form:"dev_id" json:"dev_id"`
	StartTime string `form:"start_time" json:"start_time"`
	EndTime   string `form:"end_time" json:"end_time"`
	Page      int    `form:"page" binding:"required" json:"page"`
	Size      int    `form:"size" binding:"required" json:"size"`
}

type DevLogsAdminRequest struct {
	UserId    int64  `form:"user_id" json:"user_id"`
	Uname     string `form:"uname" json:"uname"`
	DevId     int64  `form:"dev_id" json:"dev_id"`
	StartTime string `form:"start_time" json:"start_time"`
	EndTime   string `form:"end_time" json:"end_time"`
	Page      int    `form:"page" binding:"required" json:"page"`
	Size      int    `form:"size" binding:"required" json:"size"`
}

type MemberListAdminRequest struct {
	UserId     int64  `form:"user_id" json:"user_id"`         //用户ID
	Uname      string `form:"uname" json:"uname"`             //用户名
	DirectId   int64  `form:"direct_id" json:"direct_id"`     //上级ID
	DirectName string `form:"direct_name" json:"direct_name"` //上级用户名
	TeamId     int64  `form:"team_id" json:"team_id"`         //团队长ID
	TeamName   string `form:"team_name" json:"team_name"`     //团队长用户名
	Page       int    `form:"page" binding:"required" json:"page"`
	Size       int    `form:"size" binding:"required" json:"size"`
}

type MemberDevListAdminRequest struct {
	UserId int64  `form:"user_id" json:"user_id"`
	Uname  string `form:"uname" json:"uname"`
	DevId  int64  `form:"dev_id" json:"dev_id"`
	Page   int    `form:"page" binding:"required" json:"page"`
	Size   int    `form:"size" binding:"required" json:"size"`
}

type EditMemberAdminRequest struct {
	Id     int64  `form:"id" binding:"required" json:"id"` //用户ID
	Status string `form:"status" json:"status"`            //状态
}

type EditMemberDevAdminRequest struct {
	Id     int64  `form:"id" binding:"required" json:"id"` //用户ID
	Status string `form:"status" json:"status"`            //状态
}

type ChannelListAdminRequest struct {
	Name string `form:"name" json:"name"`
	Code string `form:"code" json:"code"`
	Page int    `form:"page" binding:"required" json:"page"`
	Size int    `form:"size" binding:"required" json:"size"`
}

type AddChannelAdminRequest struct {
	Name string `form:"name" binding:"required" json:"name"`
	Code string `form:"code" binding:"required" json:"code"`
	Link string `form:"link" binding:"required" json:"link"`
}

type EditChannelAdminRequest struct {
	Id     int64  `form:"id" binding:"required" json:"id"` //用户ID
	Name   string `form:"name" json:"name"`
	Code   string `form:"code" json:"code"`
	Link   string `form:"link" json:"link"`
	Status int    `form:"status" json:"status"`
}

type AppVersionListAdminRequest struct {
	AppType int `form:"app_type" json:"app_type"`
	Page    int `form:"page" binding:"required" json:"page"`
	Size    int `form:"size" binding:"required" json:"size"`
}

type AddAppVersionAdminRequest struct {
	AppType int    `form:"app_type" binding:"required" json:"app_type"`
	Version string `form:"version" binding:"required" json:"version"`
	Link    string `form:"link" binding:"required" json:"link"`
}

type EditAppVersionAdminRequest struct {
	Id      int64  `form:"id" binding:"required" json:"id"` //编号ID
	Version string `form:"version" json:"version"`
	Link    string `form:"link" json:"link"`
	Status  int    `form:"status" json:"status"`
}

type IosAccountListAdminRequest struct {
	Name        string `form:"name" json:"name"`
	Account     string `form:"account" json:"account"`
	AccountType int    `form:"account_type" json:"account_type"`
	Page        int    `form:"page" binding:"required" json:"page"`
	Size        int    `form:"size" binding:"required" json:"size"`
}

type AddIosAccountAdminRequest struct {
	Name        string `form:"name" json:"name"`
	Account     string `form:"account" binding:"required" json:"account"`
	Pass        string `form:"pass" binding:"required" json:"pass"`
	AccountType int    `form:"account_type" binding:"required" json:"account_type"`
}

type EditIosAccountAdminRequest struct {
	Id          int64  `form:"id" binding:"required" json:"id"` //编号ID
	Name        string `form:"name" json:"name"`
	Account     string `form:"account" json:"account"`
	Pass        string `form:"pass" json:"pass"`
	AccountType int    `form:"account_type" json:"account_type"`
	Status      int    `form:"status" json:"status"`
}

type AppDnsListAdminRequest struct {
	Dns      string `form:"dns" json:"dns"`
	Ip       string `form:"ip" json:"ip"`
	SiteType int    `form:"site_type" json:"site_type"`
	Level    int    `form:"level" json:"level"`
	Page     int    `form:"page" binding:"required" json:"page"`
	Size     int    `form:"size" binding:"required" json:"size"`
}

type AddAppDnsAdminRequest struct {
	Dns      string `form:"dns" binding:"required" json:"dns"`
	Ip       string `form:"ip" binding:"required" json:"ip"`
	SiteType int    `form:"site_type" binding:"required" json:"site_type"`
	Level    int    `form:"level" binding:"required" json:"level"`
}

type EditAppDnsAdminRequest struct {
	Id       int64  `form:"id" binding:"required" json:"id"`
	Dns      string `form:"dns" json:"dns"`
	Ip       string `form:"ip" json:"ip"`
	SiteType int    `form:"site_type" json:"site_type"`
	Level    int    `form:"level" json:"level"`
	Status   int    `form:"status" json:"status"`
}

type NodeDnsListAdminRequest struct {
	Dns    string `form:"dns" json:"dns"`
	Ip     string `form:"ip" json:"ip"`
	Level  int    `form:"level" json:"level"`
	NodeId int    `form:"node_id" json:"node_id"`
	Page   int    `form:"page" binding:"required" json:"page"`
	Size   int    `form:"size" binding:"required" json:"size"`
}

type AddNodeDnsAdminRequest struct {
	Dns    string `form:"dns" binding:"required" json:"dns"`
	Ip     string `form:"ip" binding:"required" json:"ip"`
	NodeId int64  `form:"node_id" binding:"required" json:"node_id"`
	Level  int    `form:"level" binding:"required" json:"level"`
}

type EditNodeDnsAdminRequest struct {
	Id     int64  `form:"id" binding:"required" json:"id"`
	Dns    string `form:"dns" json:"dns"`
	Ip     string `form:"ip" json:"ip"`
	NodeId int64  `form:"node_id" json:"node_id"`
	Level  int    `form:"level" json:"level"`
	Status int    `form:"status" json:"status"`
}

type GetReportUserDayListRequest struct {
	Date      int    `form:"date" json:"date" dc:"报表日期，eg:20240101"`
	ChannelId int    `form:"channel_id" json:"channel_id" dc:"渠道ID"`
	OrderType string `form:"order_type" json:"order_type" dc:"排序类型，eg: asc|desc"`
	Page      int    `form:"page" json:"page" dc:"分页查询page, 从1开始"`
	Size      int    `form:"size" json:"size" dc:"分页查询size, 最大1000"`
}
type GetChannelUserDayListRequest struct {
	StartDate int    `form:"startDate" json:"startDate" dc:"报表日期，eg:20240101"`
	EndDate   int    `form:"endDate" json:"endDate" dc:"报表日期，eg:20240101"`
	Date      int    `form:"date" json:"date" dc:"报表日期，eg:20240101"`
	Channel   string `form:"channel" json:"channel" dc:"推广渠道ID"`
	OrderType string `form:"order_type" json:"order_type" dc:"排序类型，eg: asc|desc"`
	Page      int    `form:"page" json:"page" dc:"分页查询page, 从1开始"`
	Size      int    `form:"size" json:"size" dc:"分页查询size, 最大1000"`
}
type GetChannelUserRechargeRequest struct {
	StartDate int    `form:"startDate" json:"startDate" dc:"报表日期，eg:20240101"`
	EndDate   int    `form:"endDate" json:"endDate" dc:"报表日期，eg:20240101"`
	Date      int    `form:"date" json:"date" dc:"报表日期，eg:20240101"`
	Channel   string `form:"channel" json:"channel" dc:"推广渠道ID"`
	OrderType string `form:"order_type" json:"order_type" dc:"排序类型，eg: asc|desc"`
	Page      int    `form:"page" json:"page" dc:"分页查询page, 从1开始"`
	Size      int    `form:"size" json:"size" dc:"分页查询size, 最大1000"`
}
type GetOnlineUserDayListRequest struct {
	Date      int    `form:"date" json:"date" dc:"报表日期，eg:20240101"`
	Email     string `form:"email" json:"email" dc:"用户email"`
	ChannelId string `form:"channel_id" json:"channel_id" dc:"渠道ID"`
	OrderBy   string `form:"order_by" json:"order_by" dc:"排序字段，eg: id|online_duration|uplink|downlink"`
	OrderType string `form:"order_type" json:"order_type" dc:"排序类型，eg: asc|desc"`
	Page      int    `form:"page" json:"page" dc:"分页查询page, 从1开始"`
	Size      int    `form:"size" json:"size" dc:"分页查询size, 最大1000"`
}
type GetNodeOnlineUserDayListRequest struct {
	Date      int    `form:"date" json:"date" dc:"报表日期，eg:20240101"`
	Email     string `form:"email" json:"email" dc:"用户email"`
	ChannelId string `form:"channel_id" json:"channel_id" dc:"渠道ID"`
	OrderBy   string `form:"order_by" json:"order_by" dc:"排序字段，eg: id|online_duration|uplink|downlink"`
	OrderType string `form:"order_type" json:"order_type" dc:"排序类型，eg: asc|desc"`
	Page      int    `form:"page" json:"page" dc:"分页查询page, 从1开始"`
	Size      int    `form:"size" json:"size" dc:"分页查询size, 最大1000"`
}
type GetNodeDayListRequest struct {
	Date      int    `form:"date" json:"date" dc:"报表日期，eg:20240101"`
	Ip        string `form:"ip" json:"ip" dc:"节点ip"`
	OrderType string `form:"order_type" json:"order_type" dc:"排序类型，eg: asc|desc"`
	Page      int    `form:"page" json:"page" dc:"分页查询page, 从1开始"`
	Size      int    `form:"size" json:"size" dc:"分页查询size, 最大1000"`
}
type GetReportUserRechargeDayListRequest struct {
	Date      int    `form:"date" json:"date" dc:"报表日期，eg:20240101"`
	GoodsId   int    `form:"goods_id" json:"goods_id" dc:"商品套餐id"`
	OrderType string `form:"order_type" json:"order_type" dc:"排序类型，eg: asc|desc"`
	Page      int    `form:"page" json:"page" dc:"分页查询page, 从1开始"`
	Size      int    `form:"size" json:"size" dc:"分页查询size, 最大1000"`
}
type GetReportUserRechargeTimesDayListRequest struct {
	Date      int    `form:"date" json:"date" dc:"报表日期，eg:20240101"`
	GoodsId   int    `form:"goods_id" json:"goods_id" dc:"商品套餐id"`
	OrderType string `form:"order_type" json:"order_type" dc:"排序类型，eg: asc|desc"`
	Page      int    `form:"page" json:"page" dc:"分页查询page, 从1开始"`
	Size      int    `form:"size" json:"size" dc:"分页查询size, 最大1000"`
}
type GetChannelUserRechargeListRequest struct {
	StartDate int    `form:"startDate" json:"startDate" dc:"报表日期，eg:20240101"`
	EndDate   int    `form:"endDate" json:"endDate" dc:"报表日期，eg:20240101"`
	Date      int    `form:"date" json:"date" dc:"报表日期，eg:20240101"`
	Channel   string `form:"channel" json:"channel" dc:"推广渠道ID"`
	OrderType string `form:"order_type" json:"order_type" dc:"排序类型，eg: asc|desc"`
	Page      int    `form:"page" json:"page" dc:"分页查询page, 从1开始"`
	Size      int    `form:"size" json:"size" dc:"分页查询size, 最大1000"`
}
type GetRechargeClickByDeviceListRequest struct {
	Date      int    `form:"date" json:"date" dc:"报表日期，eg:20240101"`
	Device    string `form:"device" json:"device" dc:"渠道设备"`
	OrderType string `form:"order_type" json:"order_type" dc:"排序类型，eg: asc|desc"`
	Page      int    `form:"page" json:"page" dc:"分页查询page, 从1开始"`
	Size      int    `form:"size" json:"size" dc:"分页查询size, 最大1000"`
}
type GetReportDeviceDayListRequest struct {
	Date      int    `form:"date" json:"date" dc:"报表日期，eg:20240101"`
	Device    string `form:"device" json:"device" dc:"渠道设备"`
	OrderType string `form:"order_type" json:"order_type" dc:"排序类型，eg: asc|desc"`
	Page      int    `form:"page" json:"page" dc:"分页查询page, 从1开始"`
	Size      int    `form:"size" json:"size" dc:"分页查询size, 最大1000"`
}
type GetReportDeviceRetentionRequest struct {
	Date      int    `form:"date" json:"date" dc:"报表日期，eg:20240101"`
	Device    string `form:"device" json:"device" dc:"渠道设备"`
	OrderType string `form:"order_type" json:"order_type" dc:"排序类型，eg: asc|desc"`
	Page      int    `form:"page" json:"page" dc:"分页查询page, 从1开始"`
	Size      int    `form:"size" json:"size" dc:"分页查询size, 最大1000"`
}
type GetReportDeviceMonthlyRetentionRequest struct {
	Date      int    `form:"date" json:"date" dc:"报表日期，eg:20240101"`
	Device    string `form:"device" json:"device" dc:"渠道设备"`
	OrderType string `form:"order_type" json:"order_type" dc:"排序类型，eg: asc|desc"`
	Page      int    `form:"page" json:"page" dc:"分页查询page, 从1开始"`
	Size      int    `form:"size" json:"size" dc:"分页查询size, 最大1000"`
}
type GetChannelUserRechargeByMonthRequest struct {
	Date      int    `form:"date" json:"date" dc:"报表日期，eg:202408"`
	Channel   string `form:"channel" json:"channel" dc:"推广渠道ID"`
	OrderType string `form:"order_type" json:"order_type" dc:"排序类型，eg: asc|desc"`
	Page      int    `form:"page" json:"page" dc:"分页查询page, 从1开始"`
	Size      int    `form:"size" json:"size" dc:"分页查询size, 最大1000"`
}
type DeletePromotionIdRequest struct {
	Id int64 `form:"id" binding:"required" json:"id"`
}
type AddPromotionRequest struct {
	PromoterName    string `form:"promoter_name" json:"promoter_name" dc:"推广人姓名" binding:"required"`
	PromotionDomain string `form:"promotion_domain" json:"promotion_domain" dc:"推广域名" binding:"required"`
	Channel         string `form:"channel" json:"channel" dc:"推广域名对应渠道"`
}
type EditPromotionRequest struct {
	Id              int64  `form:"id" binding:"required" json:"id"`
	PromoterName    string `form:"promoter_name" json:"promoter_name" dc:"推广人姓名"`
	PromotionDomain string `form:"promotion_domain" json:"promotion_domain" dc:"推广域名"`
	Channel         string `form:"channel" json:"channel" dc:"推广域名对应渠道"`
}
