package config

import "github.com/spf13/viper"

func Init() error {

	viper.AddConfigPath("./config")
	viper.AddConfigPath("config")
	return viper.ReadInConfig()
}
