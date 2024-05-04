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
	TimeFormat = "2006-01-02 15:04:05"
)
const (
	UserDevNormalStatus = 1
	UserDevBanStatus    = 2

	NetworkAutoMode   = 1
	NetworkManualMode = 2

	UserConfigStatusNormal  = 1
	UserConfigStatusDeleted = 2
)

const (
	UserStatusNormal    = 0  // 用户状态正常
	UserStatusFrozen    = 1  // 冻结
	UserStatusCancelled = 10 // 注销
)

const (
	NodeRecommendFlag = 1 // 推荐节点
)

const (
	SecretKey = "3f5202f0-4ed3-4456-80dd-13638c975bda" // 签名验证
	Healthy   = "1"
	Unhealthy = "2"
)
