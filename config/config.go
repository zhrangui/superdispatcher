package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config.staging")
	viper.AddConfigPath("config/")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
