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
	Tag   string `form:"tag" binding:"required" json:"tag"`
}
