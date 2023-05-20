package constant

const (
	SmsMsg        = "您的验证码为:%s"
	MaxCountMsg   = 5  //单号码一天最多次数 （超过冷却）
	IpMaxCountMsg = 20 //单IP一天最多次数 （超过冷却）
	WaitTimeMsg   = 30 //冷却时间30分钟

	TelMsgKey    = "code_%s"       //号码key（有效时间5min）
	TelDayMsgKey = "code_%s_%s"    //号码单日key（有效时间1天）
	IpDayMsgKey  = "ip_code_%s_%s" //IP单日key（有效时间1天）

	VerifyCountByHour = 10               //1小时同一手机号只能验证10次
	VerifySmsKey      = "verify_code_%s" //验证的key
)
