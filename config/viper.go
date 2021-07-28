package config

import (
	"github.com/spf13/viper"
	"log"
)

func ViperInit() {
	viper.SetConfigName("settings-dev")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}
}
