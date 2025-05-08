package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	ENVIRONMENT       string `mapstructure:"ENVIRONMENT"`
	DBHost            string `mapstructure:"DB_HOST"`
	DBPort            string `mapstructure:"DB_PORT"`
	DBName            string `mapstructure:"DB_DATABASE"`
	DBUsername        string `mapstructure:"DB_USERNAME"`
	DBPassword        string `mapstructure:"DB_PASSWORD"`
	HTTPServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`
	MidtransServerKey string `mapstructure:"MIDTRANS_SERVER_KEY"`
	SlackToken        string `mapstructure:"SLACK_TOKEN"`
	MandrillKey       string `mapstructure:"MANDRILL_KEY"`
}

func Load() Config {
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error: invalid configuration: %w", err))
	}

	var c Config
	err = viper.Unmarshal(&c)
	if err != nil {
		panic(fmt.Errorf("fatal error: unable decode config into struct: %v", err))
	}

	return c
}
