package config

type Kafka struct {
	Addr []string `mapstructure:"addr"`
}
