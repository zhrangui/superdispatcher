package logger

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const (
	loggerTypeNop = iota
	loggerTypeDevelopment
	loggerTypeExample
	loggerTypeProduction
)

func New(loggerType int) (*zap.Logger, error) {
	var (
		logger *zap.Logger
		err    error
	)

	switch loggerType {
	case loggerTypeNop:
		logger = zap.NewNop()
	case loggerTypeDevelopment:
		logger, err = zap.NewDevelopment()
		if err != nil {
			err = errors.Wrap(err, "failed to initialize development logger")
			return nil, err
		}
	case loggerTypeProduction:
		logger, err = zap.NewProduction()
		if err != nil {
			err = errors.Wrap(err, "failed to initialize production logger")
			return nil, err
		}
	default:
		logger = zap.NewExample()
	}
	return logger, nil
}
