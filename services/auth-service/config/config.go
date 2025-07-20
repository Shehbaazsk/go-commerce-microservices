package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBURL            string
	AppPort          string
	JWTSecret        string
	JWTExpiryMinutes int
}

func LoadConfig() Config {
	viper.AutomaticEnv()

	return Config{
		DBURL:            viper.GetString("DB_URL"),
		AppPort:          viper.GetString("APP_PORT"),
		JWTSecret:        viper.GetString("JWT_SECRET"),
		JWTExpiryMinutes: viper.GetInt("JWT_EXPIRY_MINUTES"),
	}
}
