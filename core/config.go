package core

import (
	"fmt"
	"github.com/spf13/viper"
	"indulge/model/config"
)

func InitConfig() (*viper.Viper, *config.Config) {
	v := viper.New()
	v.SetConfigName("app")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	//TODO Config changing monitors
	//v.WatchConfig()
	//v.OnConfigChange
	conf := new(config.Config)
	if err = v.Unmarshal(&conf); err != nil {
		fmt.Println(err)
	}
	return v, conf
}
