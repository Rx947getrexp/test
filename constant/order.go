package constant

const (
	ParOrderStatusInit       = "init"        // 未支付
	ParOrderStatusUnpaid     = "unpaid"      // 未支付
	ParOrderStatusPaid       = "paid"        // 已支付
	ParOrderStatusPaidFailed = "paid-failed" // 支付失败
)

const (
	UserVipAttrOpSourcePayOrder = "pay-order"
	UserVipAttrOpSourceAdminSet = "admin-set"
)

const (
	ReturnStatusSuccess = "success"
)
