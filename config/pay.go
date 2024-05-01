package config

type Pay struct {
	CallBackUrl string `mapstructure:"call_back_url"`
	OrderAmount string `mapstructure:"order_amount"`
	MerNo       string `mapstructure:"mer_no"`
}
