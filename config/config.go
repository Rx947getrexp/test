package config

type Server struct {
	System System `mapstructure:"system"`
	Db     Db     `mapstructure:"db"`
	Db2    Db     `mapstructure:"db2"`
	Log    Log    `mapstructure:"log"`
	Redis  Redis  `mapstructure:"redis"`
	Gas    Gas    `mapstructure:"gas"`
	Kafka  Kafka  `mapstructure:"kafka"`
	Pay    Pay    `mapstructure:"pay"`
}
