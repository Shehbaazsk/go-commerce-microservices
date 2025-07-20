package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Debug          bool
	AppHost        string
	AppPort        string
	JWTSecret      string
	UserServiceURL string
	AllowedOrigins []string
}

func LoadConfig() Config {
	viper.AutomaticEnv()

	return Config{
		Debug:          viper.GetBool("DEBUG"),
		AppHost:        viper.GetString("APP_HOST"),
		AppPort:        viper.GetString("APP_PORT"),
		JWTSecret:      viper.GetString("JWT_SECRET"),
		UserServiceURL: viper.GetString("USER_SERVICE_URL"),
		AllowedOrigins: strings.Split(viper.GetString("ALLOWED_ORIGINS"), ","),
	}
}
