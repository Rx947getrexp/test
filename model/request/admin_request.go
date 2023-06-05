package request

//登陆后台
type LoginAdminRequest struct {
	UserName string `form:"user_name" binding:"required"`
	Pass     string `form:"pass" binding:"required"`
}

type EditPasswdRequest struct {
	OldPass   string `form:"old_pass" binding:"required"`
	NewPass   string `form:"new_pass" binding:"required"`
	EnterPass string `form:"enter_pass" binding:"required"`
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
	Name   string `form:"name" binding:"required"`
	Remark string `form:"remark"`
	IsUsed int    `form:"is_used" binding:"required"`
	ResIds string `form:"res_ids"`
}

type EditRoleRequest struct {
	Id     int    `form:"id" binding:"required"`
	Name   string `form:"name" binding:"required"`
	Remark string `form:"remark"`
	IsUsed int    `form:"is_used" binding:"required"`
	ResIds string `form:"res_ids"`
}

type AddUserRequest struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Nickname string `form:"nickname" binding:"required"`
	RoleIds  string `form:"role_ids"`
}

type EditUserRoleRequest struct {
	UserId   int    `form:"user_id" binding:"required"`
	Password string `form:"password"`
	Nickname string `form:"nickname"`
	Status   int    `form:"status"`
	RoleIds  string `form:"role_ids"`
}

type UserRoleRequest struct {
	Id int `form:"id" binding:"required"`
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
}

type AccountAddAdminRequest struct {
	Account   string `form:"account" binding:"required" json:"account"`
	Password  string `form:"password" binding:"required" json:"password"`
	NickName  string `form:"nick_name" binding:"required" json:"nick_name"`
	RoleId    int    `form:"role_id" binding:"required" json:"role_id"`
	Auth2Code string `form:"auth2_code" json:"auth2_code"`
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
	Name string `form:"name" json:"name"`
	Tag  string `form:"tag" json:"tag"`
	Page int    `form:"page" binding:"required" json:"page"`
	Size int    `form:"size" binding:"required" json:"size"`
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
	Title     string  `form:"title" binding:"required" json:"title"`
	MType     int     `form:"m_type" binding:"required" json:"m_type"`
	DevLimit  int     `form:"dev_limit" binding:"required" json:"dev_limit"`
	FlowLimit int64   `form:"flow_limit" binding:"required" json:"flow_limit"`
	Period    int     `form:"period" binding:"required" json:"period"`
	Price     float64 `form:"price" binding:"required" json:"price"`
}

type EditGoodsAdminRequest struct {
	Id        int64   `form:"id" binding:"required" json:"id"`
	Title     string  `form:"title" json:"title"`
	MType     int     `form:"m_type" json:"m_type"`
	DevLimit  int     `form:"dev_limit" json:"dev_limit"`
	FlowLimit int64   `form:"flow_limit" json:"flow_limit"`
	Period    int     `form:"period" json:"period"`
	Price     float64 `form:"price" json:"price"`
	Status    int     `form:"status" json:"status"`
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
	Cpu     int    `form:"cpu" binding:"required" json:"cpu"`
	Flow    int64  `form:"flow" binding:"required" json:"flow"`
	Disk    int64  `form:"disk" binding:"required" json:"disk"`
	Memory  int64  `form:"memory" binding:"required" json:"memory"`
}

type EditNodeAdminRequest struct {
	Id      int64  `form:"id" binding:"required" json:"id"`
	Title   string `form:"title" json:"title"`
	Name    string `form:"name" json:"name"`
	Country string `form:"country" json:"country"`
	Ip      string `form:"ip" json:"ip"`
	Server  string `form:"server" json:"server"`
	Port    int    `form:"port" json:"port"`
	Cpu     int    `form:"cpu" json:"cpu"`
	Flow    int64  `form:"flow" json:"flow"`
	Disk    int64  `form:"disk" json:"disk"`
	Memory  int64  `form:"memory" json:"memory"`
	Status  int    `form:"status" json:"status"`
}

type NodeListAdminRequest struct {
	Title   string `form:"title" json:"title"`
	Name    string `form:"name" json:"name"`
	Country string `form:"country" json:"country"`
	Page    int    `form:"page" binding:"required" json:"page"`
	Size    int    `form:"size" binding:"required" json:"size"`
}
