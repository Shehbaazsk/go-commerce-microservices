package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Debug   bool
	DBURL   string
	AppHost string
	AppPort string
}

func LoadConfig() Config {
	viper.AutomaticEnv()

	return Config{
		DBURL:   viper.GetString("DB_URL"),
		AppPort: viper.GetString("APP_PORT"),
		AppHost: viper.GetString("APP_HOST"),
		Debug:   viper.GetBool("DEBUG"),
	}
}
