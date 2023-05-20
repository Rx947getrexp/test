package request

type RegRequest struct {
	Account     string `form:"account" binding:"required" json:"account"`
	Passwd      string `form:"passwd" binding:"required" json:"passwd"`
	EnterPasswd string `form:"enter_passwd" binding:"required" json:"enter_passwd"`
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
	Account     string `form:"account" binding:"required" json:"account"`
	OldPasswd   string `form:"old_passwd" binding:"required" json:"old_passwd"`
	Passwd      string `form:"passwd" binding:"required" json:"passwd"`
	EnterPasswd string `form:"enter_passwd" binding:"required" json:"enter_passwd"`
}
