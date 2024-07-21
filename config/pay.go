package config

type PNSafePay struct {
	MerNo       string `mapstructure:"mer_no"`
	ProductNo   string `mapstructure:"product_no"`
	OrderAmount string `mapstructure:"order_amount"`
	PayName     string `mapstructure:"pay_name"`
	PayEmail    string `mapstructure:"pay_email"`
	PayPhone    string `mapstructure:"pay_phone"`
	CallBackUrl string `mapstructure:"call_back_url"`
}

type PayConfig struct {
	MaxFreeTrialDays          int    `mapstructure:"max_free_trial_days"`
	GiftDurationPercentage    int    `mapstructure:"gift_duration_percentage"`
	OrderUnpaidLimitNum       int    `mapstructure:"order_unpaid_limit_num"`
	OrderClosedLimitNum       int    `mapstructure:"order_closed_limit_num"`
	OrderFailedLimitNum       int    `mapstructure:"order_failed_limit_num"`
	DisablePaymentCountryList string `mapstructure:"disable_payment_country_list"`
	DisablePaymentEmailList   string `mapstructure:"disable_payment_email_list"`
}

type WebMoneyConfig struct {
	WmId     string `mapstructure:"wmid"`
	Purse    string `mapstructure:"purse"`
	RandCode string `mapstructure:"rand_code"`
}

// '168.119.157.136', '168.119.60.227', '178.154.197.79', '51.250.54.238'
type FreekassaConfig struct {
	NotifyClientIp string `mapstructure:"notify_client_ip"`
}
