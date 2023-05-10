package response

//登陆后台
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
