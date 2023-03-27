package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DB string
}

func InitConfig() *Config {
	defaultEnv := ".env"

	viper.SetConfigFile(defaultEnv)
	if err := viper.ReadInConfig(); err != nil {
		panic("Error when reading config file")
	}

	config := Config{DB: viper.GetString("mysql")}

	return &config
}
