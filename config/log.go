package config

type Log struct {
	Adapter     string `mapstructure:"adapter"`
	Path        string `mapstructure:"path"`
	ReverseDays uint   `mapstructure:"reverse-days"`
	Level       string `mapstructure:"level"`
}
