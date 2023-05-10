package initialize

import (
	"github.com/spf13/viper"
	"go-speed/global"
)

func initViper() *viper.Viper {
	path := "./config.yaml"
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(&global.Config); err != nil {
		panic(err)
	}
	return v
}
