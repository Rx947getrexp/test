package constant

const (
	ParOrderStatusInit       = "init"        // 未支付
	ParOrderStatusUnpaid     = "unpaid"      // 未支付
	ParOrderStatusPaid       = "paid"        // 已支付
	ParOrderStatusPaidFailed = "paid-failed" // 支付失败
	ParOrderStatusClosed     = "closed"      // 用户取消关闭
	ParOrderStatusTimeout    = "timeout"     // 超时关闭
)

const (
	PaymentChannelErrNo  = 0
	PaymentChannelErrYes = 1
)

const (
	UserVipAttrOpSourcePayOrder                 = "pay-order"
	UserVipAttrOpSourceDirectGift               = "direct-gift" // 推荐人赠送
	UserVipAttrOpSourceAdminSet                 = "admin-set"
	UserVipAttrOpSourcePaymentChannelClosedGift = "payment-channel-closed-gift"
)

const (
	ReturnStatusSuccess = "success"
	ReturnStatusFail    = "fail"
	ReturnStatusWaiting = "waiting"
)

const (
	PayChannelErrTimeWindow = 3 // 3 天内支付通道故障窗口
)
