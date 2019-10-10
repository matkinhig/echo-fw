package config

import (
	"github.com/spf13/viper"
)

type Schema struct {
	Mongo struct {
		Host     string `mapstructure:"Host"`
		Username string `mapstructure:"User"`
		Password string
		Port     string
	} `mapstructure:"MongoDB"`

	Encryption struct {
		JWTSecret string
	}
}

var Config Schema

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.ReadInConfig()
	viper.Unmarshal(&Config)

}
