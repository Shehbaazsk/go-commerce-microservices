package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	AppHost   string
	AppPort   string
	JWTSecret string
}

func LoadConfig() Config {
	viper.AutomaticEnv()

	return Config{
		AppHost:   viper.GetString("APP_HOST"),
		AppPort:   viper.GetString("APP_PORT"),
		JWTSecret: viper.GetString("JWT_SECRET"),
	}
}
