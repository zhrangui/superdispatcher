package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Config struct {
	Constants *Constants
	Logger    *zap.Logger
}

func New(configName string) (*Constants, error) {
	viper.SetConfigName(configName)
	viper.AddConfigPath("../resources")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	constants := new(Constants)
	err = viper.Unmarshal(constants)
	if err != nil {
		return nil, err
	}

	return constants, nil
}
