package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBURL   string
	AppPort string
}

func LoadConfig() Config {
	viper.AutomaticEnv()

	return Config{
		DBURL:   viper.GetString("DB_URL"),
		AppPort: viper.GetString("APP_PORT"),
	}
}
