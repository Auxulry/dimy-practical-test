// Package configs is describe all configuration for this source
package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	JWTSecret string `mapstructure:"JWT_SECRET"`
	DBDNS     string `mapstructure:"DB_DNS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	viper.WatchConfig()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}