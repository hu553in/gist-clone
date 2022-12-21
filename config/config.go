package config

import (
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

func Init(env string) {
	config = viper.New()

	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("./config")

	err := config.ReadInConfig()
	if err != nil {
		log.Fatal("Failed to parse configuration file", err)
	}
}

func GetConfig() *viper.Viper {
	return config
}
