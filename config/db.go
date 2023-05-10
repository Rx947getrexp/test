package config

type Db struct {
	Host        string `mapstructure:"host"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	DbName      string `mapstructure:"db-name"`
	Port        string `mapstructure:"port"`
	MaxIdleConn int    `mapstructure:"max-idle-conns"`
	Log         Log    `mapstructure:"log"`
}
