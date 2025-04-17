package config

type System struct {
	Env               string `mapstructure:"env"`                  // 环境值
	Addr              string `mapstructure:"addr"`                 // 端口值
	ShowSql           bool   `mapstructure:"show-sql"`             // 打印sql
	Sign              string `mapstructure:"sign"`                 // 签名环境
	Encrypt           bool   `mapstructure:"encrypt"`              // 加密配置信息
	HttpProxy         string `mapstructure:"http-proxy"`           // http代理
	Web3Plugin        string `mapstructure:"web3-plugin"`          // http代理
	V2rayConfigPath   string `mapstructure:"v2ray_config_path"`    // v2ray 配置文件 地址
	V2rayApiAddress   string `mapstructure:"v2ray_api_address"`    // v2ray api 地址
	APIServerAddr     string `mapstructure:"api_server_address"`   // api 地址
	APIServerIPs      string `mapstructure:"api_server_ips"`       // api 地址
	APIServerDNSList  string `mapstructure:"api_server_dns_list"`  // api dns 地址
	NodeServerDNSList string `mapstructure:"node_server_dns_list"` // node dns 地址
	UserNodeEnable    int    `mapstructure:"user_node_enable"`
	DomainDnsList     string `mapstructure:"domain_dns_list"`
}
