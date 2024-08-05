package constant

const (
	ParOrderStatusInit               = "init"                 // 初始状态，还未提交到支付平台
	ParOrderStatusUnpaid             = "unpaid"               // 未支付（等待支付）
	ParOrderStatusPaid               = "paid"                 // 已支付
	ParOrderStatusPaidFailed         = "paid-failed"          // 支付失败
	ParOrderStatusClosed             = "closed"               // 用户取消关闭
	ParOrderStatusTimeout            = "timeout"              // 超时关闭
	ParOrderStatusAdminConfirmClosed = "admin-confirm-closed" // 管理员审核无效，关闭订单
	ParOrderStatusAdminConfirmPassed = "admin-confirm-passed" // 管理员审核通过
)

const (
	PaymentChannelErrNo  = 0
	PaymentChannelErrYes = 1
)

const (
	UserVipAttrOpSourcePayOrder                 = "pay-order"
	UserVipAttrOpSourceDirectGift               = "direct-gift" // 推荐人赠送
	UserVipAttrOpSourceAdminSet                 = "admin-set"
	UserVipAttrOpSourceAdminRevertOrder         = "admin-revert-order"
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

const (
	UploadFilePathOrder        = "order"
	UploadFilePathPayment      = "payment"
	UploadFilePathOfficialDocs = "official_docs"
)

const (
	OrderRevertFlag = 1
)
