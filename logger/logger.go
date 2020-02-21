package logger

import (
	"encoding/json"
	"log"
	"superdispatcher/config"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Logger wrap zap.Logger
type Logger struct {
	*zap.Logger
}

// NewLog creates customized Logger
func NewLog(config *config.Config) (*Logger, error) {
	var cfg zap.Config
	if err := json.Unmarshal([]byte(config.Constants.Log), &cfg); err != nil {
		log.Fatal(err)
	}
	var err error
	lg, err := cfg.Build()
	if err != nil {
		log.Fatal(err)
	}
	logger := &Logger{
		Logger: lg,
	}
	return logger, err
}

// Error logs fatal message
func (logger *Logger) Error(err error, msg string) {
	if err != nil {
		logger.Logger.Error(errors.Wrap(err, msg).Error())
	}
}

// FailOnError exists system on fail
func (logger *Logger) FailOnError(err error, msg string) {
	if err != nil {
		logger.Logger.Fatal(errors.Wrap(err, msg).Error())
	}
}
