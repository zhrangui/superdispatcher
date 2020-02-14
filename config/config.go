package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Constants *Constants
}

/*New create configuration*/
func New(configName string, configPath string) (*Config, error) {

	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	viper.AutomaticEnv()
	var config *Config = &Config{}
	config.Constants = &Constants{}

	err = viper.Unmarshal(config.Constants)
	if err != nil {
		return nil, err
	}

	return config, nil
}
