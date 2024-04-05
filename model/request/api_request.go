package request

type RegRequest struct {
	Account     string `form:"account" binding:"required" json:"account"`
	Passwd      string `form:"passwd" binding:"required" json:"passwd"`
	EnterPasswd string `form:"enter_passwd" binding:"required" json:"enter_passwd"`
	InviteCode  string `form:"invite_code" json:"invite_code"`
}

type LoginRequest struct {
	Account string `form:"account" binding:"required" json:"account"`
	Passwd  string `form:"passwd" binding:"required" json:"passwd"`
}

type ForgetRequest struct {
	Account     string `form:"account" binding:"required" json:"account"`
	VerifyCode  string `form:"verify_code" binding:"required" json:"verify_code"`
	Passwd      string `form:"passwd" binding:"required" json:"passwd"`
	EnterPasswd string `form:"enter_passwd" binding:"required" json:"enter_passwd"`
}

type SendEmailRequest struct {
	Email string `form:"email" binding:"required" json:"email"`
}

type ChangePasswdRequest struct {
	OldPasswd   string `form:"old_passwd" binding:"required" json:"old_passwd"`
	Passwd      string `form:"passwd" binding:"required" json:"passwd"`
	EnterPasswd string `form:"enter_passwd" binding:"required" json:"enter_passwd"`
}

type TeamListRequest struct {
	Page int `form:"page" binding:"required" json:"page"`
	Size int `form:"size" binding:"required" json:"size"`
}

type NoticeListRequest struct {
	Page int `form:"page" binding:"required" json:"page"`
	Size int `form:"size" binding:"required" json:"size"`
}

type NoticeDetailRequest struct {
	Id int64 `form:"id" binding:"required" json:"id"`
}

type UploadLogRequest struct {
	Content string `form:"content" binding:"required" json:"content"`
}

type DevListRequest struct {
	Page int `form:"page" binding:"required" json:"page"`
	Size int `form:"size" binding:"required" json:"size"`
}

type BanDevRequest struct {
	DevId  string `form:"dev_id" binding:"required" json:"dev_id"`
	NodeId int64  `form:"node_id" json:"node_id"`
}

type ChangeNetworkRequest struct {
	WorkMode int   `form:"work_mode" binding:"required" json:"work_mode"` //1-智能；2-手选IP
	NodeLine int64 `form:"node_line" json:"node_line"`                    //工作线路
}
type ConnectDevRequest struct {
	NodeId int64 `form:"node_id" binding:"required" json:"node_id"`
}
type SwitchButtonStatusRequest struct {
	Status int `form:"status" binding:"required" json:"status"` //1-开启；2-关闭
}

type CreateOrderRequest struct {
	GoodsId int64 `form:"goods_id" binding:"required" json:"goods_id"`
}

type OrderListRequest struct {
	Page int `form:"page" binding:"required" json:"page"`
	Size int `form:"size" binding:"required" json:"size"`
}

type SaveUserConfigRequest struct {
	NodeId int64 `form:"node_id" binding:"required" json:"node_id"`
}

type GetUserConfigResponse struct {
	UserId    int64  `json:"user_id"`
	NodeId    int64  `json:"node_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ReportNodePingResultRequest struct {
	UserId     int64            `form:"user_id" binding:"required" json:"UserId"`
	ReportTime string           `form:"report_time" json:"ReportTime"`
	Items      []PingResultItem `form:"items" json:"Items"`
}

type PingResultItem struct {
	Ip   string `form:"ip" json:"Ip"`
	Code string `form:"code" json:"Code" dc:"ping结果"`
	Cost string `form:"cost" json:"Cost" dc:"ping耗时"`
}
