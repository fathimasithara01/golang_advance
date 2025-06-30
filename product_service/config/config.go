package config

import (
	"log"

	"github.com/spf13/viper"
)

func Load() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("config load error: %v", err)
	}
}
