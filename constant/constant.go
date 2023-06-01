package constant

const (
	FilePath       = "public"
	UploadFilePath = "upload"
	ImgFileType    = "img"
	OtherFileType  = "other"

	MaxPageSize = 1000
)

const (
	NodeVersion     = "v1.0.0"
	AccessTokenSalt = "2023@win"
	HeartbeatTime   = 5
	ReportDataTime  = 30

	ForgetSubject = "Speed找回密码"
	ForgetBody    = "<br>hello!</br>您本次的验证码是:<font color='red'>%s</font>"
)

const (
	UserDevNormalStatus = 1
	UserDevBanStatus    = 2

	NetworkAutoMode   = 1
	NetworkManualMode = 2
)
