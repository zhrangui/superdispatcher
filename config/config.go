package config

import (
	"github.com/spf13/viper"

	"superdispatcher/logger"
)

type Config struct {
	Constants *Constants
	Logger    *logger.Logger
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

	config.Logger, err = logger.New(config.Constants.My.LoggerType)
	if err != nil {
		return nil, err
	}
	return config, nil
}
