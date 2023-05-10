package config

type System struct {
	Env        string `mapstructure:"env"`         // 环境值
	Addr       string `mapstructure:"addr"`        // 端口值
	ShowSql    bool   `mapstructure:"show-sql"`    // 打印sql
	Sign       string `mapstructure:"sign"`        // 签名环境
	Encrypt    bool   `mapstructure:"encrypt"`     // 加密配置信息
	HttpProxy  string `mapstructure:"http-proxy"`  // http代理
	Web3Plugin string `mapstructure:"web3-plugin"` // http代理
}
