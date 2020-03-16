package util

import (
	"github.com/spf13/viper"
)

func InitViper() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()
}

func InitViper1() {
	viper.SetConfigName("config")
	viper.AddConfigPath("../")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()
}

func InitViper2() {
	viper.SetConfigName("config")
	viper.AddConfigPath("../../")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()
}