package config

import (
	"log"

	"github.com/spf13/viper"
)

func ViperInit() {
	once.Do(func() {
		viper.SetConfigName("settings-dev")
		viper.SetConfigType("toml")
		viper.AddConfigPath(".")
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalf("read config failed: %v", err)
		}
	})
}
