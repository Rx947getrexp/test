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
	MaxFreeTrialDays       int `mapstructure:"max_free_trial_days"`
	GiftDurationPercentage int `mapstructure:"gift_duration_percentage"`
}
