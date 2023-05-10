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

type UserListAdminRequest struct {
	Id           int64  `form:"id" json:"id"`
	UName        string `form:"uname" json:"uname"`
	NickName     string `form:"nick_name" json:"nick_name"`
	UserType     int    `form:"user_type" json:"user_type"`
	KindType     int    `form:"kind_type" json:"kind_type"`
	Classify     int    `form:"classify" json:"classify"`
	FrozenStatus int    `form:"frozen_status" json:"frozen_status"`
	Page         int    `form:"page" binding:"required" json:"page"`
	Size         int    `form:"size" binding:"required" json:"size"`
}

type AgentUserListAdminRequest struct {
	Id       int64  `form:"id" json:"id"`
	UName    string `form:"uname" json:"uname"`
	NickName string `form:"nick_name" json:"nick_name"`
	Phone    string `form:"phone" json:"phone"`
	KindType int    `form:"kind_type" json:"kind_type"`
	Classify int    `form:"classify" json:"classify"`
	Page     int    `form:"page" binding:"required" json:"page"`
	Size     int    `form:"size" binding:"required" json:"size"`
}

type UserBankListAdminRequest struct {
	Id       int64  `form:"id" json:"id"`
	UserId   int64  `form:"user_id" json:"user_id"`
	UName    string `form:"uname" json:"uname"`
	Status   int    `form:"status" json:"status"`
	BankType int    `form:"bank_type" json:"bank_type"`
	Page     int    `form:"page" binding:"required" json:"page"`
	Size     int    `form:"size" binding:"required" json:"size"`
}

type UserInviteCodeListAdminRequest struct {
	UserId     int64  `form:"user_id" json:"user_id"`
	UName      string `form:"uname" json:"uname"`
	InviteCode string `form:"invite_code" json:"invite_code"`
	Page       int    `form:"page" binding:"required" json:"page"`
	Size       int    `form:"size" binding:"required" json:"size"`
}

type SetUserRiskConfigAdminRequest struct {
}

type AgentListAdminRequest struct {
	UserId      int64  `form:"user_id" json:"user_id"`
	UName       string `form:"uname" json:"uname"`
	DirectId    int64  `form:"direct_id" json:"direct_id"`
	DirectUName string `form:"direct_uname" json:"direct_uname"`
	QueueId     int64  `form:"queue_id" json:"queue_id"`
	QueueUName  string `form:"queue_uname" json:"queue_uname"`
	Page        int    `form:"page" binding:"required" json:"page"`
	Size        int    `form:"size" binding:"required" json:"size"`
}

type AgentDayProfitListAdminRequest struct {
	UserId    int64  `form:"user_id" json:"user_id"`
	UName     string `form:"uname" json:"uname"`
	StartTime string `form:"start_time" json:"start_time"` //开始时间
	EndTime   string `form:"end_time" json:"end_time"`     //结束时间
	Page      int    `form:"page" binding:"required" json:"page"`
	Size      int    `form:"size" binding:"required" json:"size"`
}

type AgentProfitListAdminRequest struct {
	UserId      int64  `form:"user_id" json:"user_id"`
	UName       string `form:"uname" json:"uname"`
	DirectId    int64  `form:"direct_id" json:"direct_id"`
	DirectUName string `form:"direct_uname" json:"direct_uname"`
	OrderId     int64  `form:"order_id" json:"order_id"`
	StartTime   string `form:"start_time" json:"start_time"` //开始时间
	EndTime     string `form:"end_time" json:"end_time"`     //结束时间
	Page        int    `form:"page" binding:"required" json:"page"`
	Size        int    `form:"size" binding:"required" json:"size"`
}

type MerchantListAdminRequest struct {
	MerchantUid int64  `form:"merchant_uid" json:"merchant_uid"` //商户号
	UName       string `form:"uname" json:"uname"`
	Nickname    string `form:"nickname" json:"nickname"`
	Page        int    `form:"page" binding:"required" json:"page"`
	Size        int    `form:"size" binding:"required" json:"size"`
}

type MerchantApiListAdminRequest struct {
	MerchantUid int64  `form:"merchant_uid" json:"merchant_uid"` //商户号
	Nickname    string `form:"nickname" json:"nickname"`
	Page        int    `form:"page" binding:"required" json:"page"`
	Size        int    `form:"size" binding:"required" json:"size"`
}

type MerchantBankListAdminRequest struct {
	MerchantUid int64  `form:"merchant_uid" json:"merchant_uid"` //商户号
	Nickname    string `form:"nickname" json:"nickname"`
	Page        int    `form:"page" binding:"required" json:"page"`
	Size        int    `form:"size" binding:"required" json:"size"`
}

type CollectListAdminRequest struct {
	Id              int64  `form:"id" json:"id"`                               //平台id
	MerchantOrderId string `form:"merchant_order_id" json:"merchant_order_id"` //商户订单id
	MerchantUid     int64  `form:"merchant_uid" json:"merchant_uid"`           //商户号
	RecUid          int64  `form:"rec_uid" json:"rec_uid"`                     //收单人
	RecUname        string `form:"rec_uname" json:"rec_uname"`                 //收单人用户名
	ProductId       int    `form:"product_id" json:"product_id"`               //支付产品编号
	IsNormal        int    `form:"is_normal" json:"is_normal"`                 //1-正常；2-异常
	Status          int    `form:"status" json:"status"`                       //状态
	StartTime       string `form:"start_time" json:"start_time"`               //开始时间
	EndTime         string `form:"end_time" json:"end_time"`                   //结束时间
	Page            int    `form:"page" binding:"required" json:"page"`
	Size            int    `form:"size" binding:"required" json:"size"`
}

type CollectAbnormalListAdminRequest struct {
	Id              int64  `form:"id" json:"id"`                               //流水号id
	OpId            int64  `form:"op_id" json:"op_id"`                         //关联的平台业务订单id
	MerchantOrderId string `form:"merchant_order_id" json:"merchant_order_id"` //商户订单id
	MerchantUid     int64  `form:"merchant_uid" json:"merchant_uid"`           //商户号
	MerchantUname   string `form:"merchant_uname" json:"merchant_uname"`       //商户昵称
	RecUid          int64  `form:"rec_uid" json:"rec_uid"`                     //收单人
	RecUname        string `form:"rec_uname" json:"rec_uname"`                 //收单人用户名
	StartTime       string `form:"start_time" json:"start_time"`               //开始时间
	EndTime         string `form:"end_time" json:"end_time"`                   //结束时间
	Page            int    `form:"page" binding:"required" json:"page"`
	Size            int    `form:"size" binding:"required" json:"size"`
}

type CollectAbnormalEditAdminRequest struct {
	Id        int64  `form:"id" binding:"required" json:"id"`                 //平台订单id
	Status    int    `form:"status" binding:"required" json:"status"`         //变更的订单状态
	Auth2Code string `form:"auth2_code" binding:"required" json:"auth2_code"` //谷歌验证码
}

type AddUserAdminRequest struct {
	Uname    string `form:"uname" binding:"required" json:"uname"`          //用户名
	Pass     string `form:"pass"  binding:"required" json:"pass"`           //密码
	KindType int    `form:"kind_type" binding:"required"  json:"kind_type"` //1-跑币；2-话费；3-电商；4-三方
	Classify int    `form:"classify" binding:"required" json:"classify"`    //1-C端刷手；2-xx四方
	AreaCode string `form:"area_code" binding:"required" json:"area_code"`
	Phone    string `form:"phone" binding:"required" json:"phone"`
}

type EditUserAdminRequest struct {
	Id      int64  `form:"id" binding:"required" json:"id"` //用户id
	Status  string `form:"status" json:"status"`
	IsAuth  int    `form:"is_auth" json:"is_auth"`
	IsRisk  string `form:"is_risk" json:"is_risk"`
	Comment string `form:"comment" json:"comment"`
}

type EditMerchantApiAdminRequest struct {
	Id          int64 `form:"id" binding:"required" json:"id"` //商户id
	Status      int   `form:"status" json:"status"`
	CollectOpen int   `form:"collect_open" json:"collect_open"`
	PaymentOpen int   `form:"payment_open" json:"payment_open"`
}

type EditUserBankAdminRequest struct {
	Id      int64  `form:"id" binding:"required" json:"id"` //bank-id
	Status  int    `form:"status" json:"status"`
	IsRisk  string `form:"is_risk" json:"is_risk"`
	Comment string `form:"comment" json:"comment"`
}

type SetDictAdminRequest struct {
	Key   string `form:"key" binding:"required" json:"key"`
	Value string `form:"value" binding:"required" json:"value"`
	IsDel string `form:"is_del" json:"is_del"`
}

type DepositRechargeListAdminRequest struct {
	Id           int64  `form:"id" json:"id"`
	Uname        string `form:"uname" json:"uname"`
	Userid       int64  `form:"user_id" json:"user_id"`
	RechargeType int    `form:"recharge_type" json:"recharge_type"`
	Status       int    `form:"status" json:"status"`
	Nickname     string `form:"nickname" json:"nickname"`
	BankType     int    `form:"bank_type" json:"bank_type"`
	StartTime    string `form:"start_time" json:"start_time"` //开始时间
	EndTime      string `form:"end_time" json:"end_time"`     //结束时间
	Page         int    `form:"page" binding:"required" json:"page"`
	Size         int    `form:"size" binding:"required" json:"size"`
}

type DepositRechargeEditAdminRequest struct {
	Id        int64  `form:"id" binding:"required" json:"id"`
	Status    int    `form:"status" binding:"required" json:"status"`
	Comment   string `form:"comment" json:"comment"`
	Auth2Code string `form:"auth2_code" json:"auth2_code"` //谷歌验证码
}

type WithdrawListAdminRequest struct {
	Id           int64  `form:"id" json:"id"`
	Userid       int64  `form:"user_id" json:"user_id"`
	WithdrawType int    `form:"withdraw_type" json:"withdraw_type"`
	Status       int    `form:"status" json:"status"`
	Uname        string `form:"uname" json:"uname"`
	BankType     int    `form:"bank_type" json:"bank_type"`
	Nickname     string `form:"nickname" json:"nickname"`
	StartTime    string `form:"start_time" json:"start_time"` //开始时间
	EndTime      string `form:"end_time" json:"end_time"`     //结束时间
	Page         int    `form:"page" binding:"required" json:"page"`
	Size         int    `form:"size" binding:"required" json:"size"`
}

type WithdrawEditAdminRequest struct {
	Id        int64  `form:"id" binding:"required" json:"id"`
	Status    int    `form:"status" binding:"required"  json:"status"`
	Comment   string `form:"comment" json:"comment"`
	Auth2Code string `form:"auth2_code" json:"auth2_code"` //谷歌验证码
}

type PaymentListAdminRequest struct {
	Id               int64  `form:"id" json:"id"`
	MerchantOrderId  string `form:"merchant_order_id" json:"merchant_order_id"` //商户订单id
	MerchantUid      int64  `form:"merchant_uid" json:"merchant_uid"`           //商户号
	MerchantUname    string `form:"merchant_uname" json:"merchant_uname"`       //商户用户名
	MerchantNickname string `form:"merchant_nickname" json:"merchant_nickname"` //商户用户名
	BankType         int    `form:"bank_type" json:"bank_type"`
	ProductId        int    `form:"product_id" json:"product_id"`
	Status           int    `form:"status" json:"status"`
	StartTime        string `form:"start_time" json:"start_time"` //开始时间
	EndTime          string `form:"end_time" json:"end_time"`     //结束时间
	Page             int    `form:"page" binding:"required" json:"page"`
	Size             int    `form:"size" binding:"required" json:"size"`
}

type PaymentEditAdminRequest struct {
	Id        int64  `form:"id" binding:"required" json:"id"`
	Status    int    `form:"status" binding:"required" json:"status"`
	Comment   string `form:"comment" json:"comment"`
	Auth2Code string `form:"auth2_code" json:"auth2_code"` //谷歌验证码
}

type PublicBankListAdminRequest struct {
	Id       int64  `form:"id" json:"id"`
	CardNo   string `form:"card_no" json:"card_no"`
	BankType int    `form:"bank_type" json:"bank_type"`
	Page     int    `form:"page" binding:"required" json:"page"`
	Size     int    `form:"size" binding:"required" json:"size"`
}

type PublicBankAddAdminRequest struct {
	BankType    int    `form:"bank_type" binding:"required" json:"bank_type"`
	CardType    int    `form:"card_type" binding:"required" json:"card_type"`
	CardNo      string `form:"card_no" binding:"required"json:"card_no"`
	RealName    string `form:"real_name" json:"real_name"`
	CardBank    string `form:"card_bank" json:"card_bank"`
	SingleLimit string `form:"single_limit" json:"single_limit"`
	DayLimit    string `form:"day_limit" json:"day_limit"`
	Auth2Code   string `form:"auth2_code" binding:"required" json:"auth2_code"` //谷歌验证码
}

type PublicBankEditAdminRequest struct {
	Id          int64  `form:"id" json:"id"`
	BankType    int    `form:"bank_type" json:"bank_type"`
	CardType    int    `form:"card_type" json:"card_type"`
	CardNo      string `form:"card_no"   json:"card_no"`
	RealName    string `form:"real_name" json:"real_name"`
	CardBank    string `form:"card_bank" json:"card_bank"`
	SingleLimit string `form:"single_limit" json:"single_limit"`
	DayLimit    string `form:"day_limit" json:"day_limit"`
	Status      string `form:"status"    json:"status"`
	OpenStatus  string `form:"open_status"    json:"open_status"`
	IsRisk      string `form:"is_risk"   json:"is_risk"`
	Auth2Code   string `form:"auth2_code" json:"auth2_code"` //谷歌验证码
}

type PublicBankOrderListAdminRequest struct {
	Id        int64  `form:"id" json:"id"`
	CardNo    string `form:"card_no" json:"card_no"`
	StartTime string `form:"start_time" json:"start_time"` //开始时间
	EndTime   string `form:"end_time" json:"end_time"`     //结束时间
	Page      int    `form:"page" binding:"required" json:"page"`
	Size      int    `form:"size" binding:"required" json:"size"`
}

type PublicBankSummaryListAdminRequest struct {
	Id        int64  `form:"id" json:"id"`
	TimeType  int    `form:"time_type" json:"time_type"`   //1-按日期；2-按月份
	StartTime string `form:"start_time" json:"start_time"` //开始时间
	EndTime   string `form:"end_time" json:"end_time"`     //结束时间
	Page      int    `form:"page" binding:"required" json:"page"`
	Size      int    `form:"size" binding:"required" json:"size"`
}

type ChannelListAdminRequest struct {
	Id     int64  `form:"id" json:"id"`
	Name   string `form:"name" json:"name"`
	Status int    `form:"status" json:"status"`
	IsRisk int    `form:"is_risk" json:"is_risk"`
	Page   int    `form:"page" binding:"required" json:"page"`
	Size   int    `form:"size" binding:"required" json:"size"`
}

type ChannelAddAdminRequest struct {
	Category  int    `form:"category" binding:"required" json:"category"`
	Name      string `form:"name" binding:"required" json:"name"`
	Code      string `form:"code" binding:"required" json:"code"`
	CostRate  string `form:"cost_rate"  json:"cost_rate"`
	Fee       string `form:"fee"  json:"fee"`
	Auth2Code string `form:"auth2_code" binding:"required" json:"auth2_code"` //谷歌验证码
}

type ChannelEditAdminRequest struct {
	Id        int64  `form:"id" binding:"required" json:"id"`
	Category  int    `form:"category" json:"category"`
	Name      string `form:"name" json:"name"`
	Code      string `form:"code" json:"code"`
	CostRate  string `form:"cost_rate" json:"cost_rate"`
	Fee       string `form:"fee"  json:"fee"`
	Status    string `form:"status"    json:"status"`
	IsRisk    string `form:"is_risk"   json:"is_risk"`
	Auth2Code string `form:"auth2_code" json:"auth2_code"` //谷歌验证码
}

type ProductListAdminRequest struct {
	Id          int64  `form:"id" json:"id"`
	Name        string `form:"name" json:"name"`
	ChannelId   string `form:"channel_id" json:"channel_id"`
	ChannelName string `form:"channel_name" json:"channel_name"`
	Status      int    `form:"status" json:"status"`
	IsRisk      int    `form:"is_risk" json:"is_risk"`
	MerchantUid int64  `form:"merchant_uid" json:"merchant_uid"`
	Page        int    `form:"page" binding:"required" json:"page"`
	Size        int    `form:"size" binding:"required" json:"size"`
}

type ProductAddAdminRequest struct {
	ChannelId int64  `form:"channel_id" binding:"required"  json:"channel_id"`
	Category  int    `form:"category" binding:"required" json:"category"`
	Name      string `form:"name" binding:"required"  json:"name"`
	Code      string `form:"code" binding:"required"  json:"code"`
	BankType  int    `form:"bank_type" binding:"required"  json:"bank_type"`
	Fee       string `form:"fee"     json:"fee"`
	FeeRate   string `form:"fee_rate"    json:"fee_rate"`
	Auth2Code string `form:"auth2_code" binding:"required" json:"auth2_code"` //谷歌验证码
}

type ProductEditAdminRequest struct {
	Id        int64  `form:"id" binding:"required" json:"id"`
	ChannelId int64  `form:"channel_id" json:"channel_id"`
	Name      string `form:"name" json:"name"`
	Code      string `form:"code" json:"code"`
	BankType  int    `form:"bank_type" json:"bank_type"`
	Fee       string `form:"fee"    json:"fee"`
	FeeRate   string `form:"fee_rate"    json:"fee_rate"`
	Status    string `form:"status"    json:"status"`
	IsRisk    string `form:"is_risk"   json:"is_risk"`
	Auth2Code string `form:"auth2_code" json:"auth2_code"` //谷歌验证码
}

type MerchantRateListAdminRequest struct {
	Id        int64  `form:"id" json:"id"`
	UserId    int64  `form:"user_id" json:"user_id"`
	Nickname  string `form:"nickname" json:"nickname"`
	ChannelId string `form:"channel_id"  json:"channel_id"`
	ProductId string `form:"product_id"  json:"product_id"`
	Page      int    `form:"page" binding:"required" json:"page"`
	Size      int    `form:"size" binding:"required" json:"size"`
}

type FeeRateCheckListAdminRequest struct {
	Id          int64  `form:"id" json:"id"`
	UserId      int64  `form:"user_id" json:"user_id"`
	Nickname    string `form:"nickname" json:"nickname"`
	AgentUid    int64  `form:"agent_uid" json:"agent_uid"`
	AgentUname  string `form:"agent_uname" json:"agent_uname"`
	ProductId   int64  `form:"product_id" json:"product_id"`
	ProductName string `form:"agent_uname" json:"agent_uname"`
	ChannelId   int64  `form:"channel_id" json:"channel_id"`
	ChannelName string `form:"channel_name" json:"channel_name"`
	Page        int    `form:"page" binding:"required" json:"page"`
	Size        int    `form:"size" binding:"required" json:"size"`
}

type CollectPoolListAdminRequest struct {
	Id         int64  `form:"id" json:"id"`
	AgentUid   int64  `form:"agent_uid" json:"agent_uid"`
	AgentUname string `form:"agent_uname" json:"agent_uname"`
	IsRisk     string `form:"is_risk" json:"is_risk"`
	Classify   int    `form:"classify" json:"classify"`
	BankType   int    `form:"bank_type" json:"bank_type"`
	CardNo     string `form:"card_no" json:"card_no"`
	Page       int    `form:"page" binding:"required" json:"page"`
	Size       int    `form:"size" binding:"required" json:"size"`
}

type FinanceRepairListAdminRequest struct {
	Id        int64  `form:"id" json:"id"`
	UserId    int64  `form:"user_id" json:"user_id"`
	Uname     string `form:"uname" json:"uname"`
	Nickname  string `form:"nickname" json:"nickname"`
	UserType  int    `form:"user_type"  json:"user_type"`
	WorkType  int    `form:"work_type" json:"work_type"`
	OpType    int    `form:"op_type" json:"op_type"`
	StartTime string `form:"start_time" json:"start_time"` //开始时间
	EndTime   string `form:"end_time" json:"end_time"`     //结束时间
	Page      int    `form:"page" binding:"required" json:"page"`
	Size      int    `form:"size" binding:"required" json:"size"`
}

type FinanceRepairAddAdminRequest struct {
	UserId    int64  `form:"user_id" binding:"required" json:"user_id"`
	UserType  int    `form:"user_type"  json:"user_type"`
	WorkType  int    `form:"work_type" json:"work_type"` //1-代表代收账户app_amount;2-代表代付账户
	OpType    int    `form:"op_type" binding:"required" json:"op_type"`
	OpQty     string `form:"op_qty" binding:"required" json:"op_qty"`
	Title     string `form:"title" binding:"required" json:"title"`
	Comment   string `form:"comment" binding:"required" json:"comment"`
	Auth2Code string `form:"auth2_code" binding:"required" json:"auth2_code"` //谷歌验证码
}

type FinanceLogsListAdminRequest struct {
	Id        int64  `form:"id" json:"id"`
	UserId    int64  `form:"user_id" json:"user_id"`
	Uname     string `form:"uname" json:"uname"`
	Nickname  string `form:"nickname" json:"nickname"`
	StartTime string `form:"start_time" json:"start_time"` //开始时间
	EndTime   string `form:"end_time" json:"end_time"`     //结束时间
	Page      int    `form:"page" binding:"required" json:"page"`
	Size      int    `form:"size" binding:"required" json:"size"`
}

type RootRateEditRequest struct {
	//Id		 int64 `form:"id" binding:"required" json:"id"`                 //ID编号
	Rate1     float64 `form:"rate1" json:"rate1"`                              //费率1
	Rate2     float64 `form:"rate2" json:"rate2"`                              //费率2
	Rate3     float64 `form:"rate3" json:"rate3"`                              //费率3
	Rate4     float64 `form:"rate4" json:"rate4"`                              //费率4
	Rate5     float64 `form:"rate5" json:"rate5"`                              //费率5
	Rate6     float64 `form:"rate6" json:"rate6"`                              //费率6
	Rate7     float64 `form:"rate7" json:"rate7"`                              //费率7
	Rate8     float64 `form:"rate8" json:"rate8"`                              //费率8
	Auth2Code string  `form:"auth2_code" binding:"required" json:"auth2_code"` //谷歌验证码
}

type SetCashInOutRiskRequest struct {
	RiskStatus bool   `form:"risk_status" json:"risk_status"`                  //风控状态：true-打开风控；false-关闭风控
	Auth2Code  string `form:"auth2_code" binding:"required" json:"auth2_code"` //谷歌验证码
}
