package request

type NodeAddEmailRequest struct {
	Email   string `form:"email" binding:"required" json:"email"`
	Uuid    string `form:"uuid" binding:"required" json:"uuid"`
	Level   string `form:"level" binding:"required" json:"level"`
	AlterId string `form:"alterId" binding:"required" json:"alterId"`
}

type NodeRemoveEmailRequest struct {
	Email string `form:"email" binding:"required" json:"email"`
}

type NodeAddSubRequest struct {
	Email string `form:"email" binding:"required" json:"email"`
	Uuid  string `form:"uuid" binding:"required" json:"uuid"`
	Level string `form:"level" binding:"required" json:"level"`
	Tag   string `form:"tag" binding:"required" json:"tag"`
}

type GetUserTrafficRequest struct {
	Emails []string `form:"emails" json:"emails"`
	All    bool     `form:"all" json:"all"`
	Reset  bool     `form:"reset" json:"reset"`
}

type Request struct {
}
