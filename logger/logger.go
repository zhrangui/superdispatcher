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

// Logger wrap zap.Logger
type Logger struct {
	Zap *zap.Logger
}

func New(loggerType int) (*Logger, error) {
	var (
		err error
	)
	logger := new(Logger)
	switch loggerType {
	case loggerTypeNop:
		logger.Zap = zap.NewNop()
	case loggerTypeDevelopment:
		logger.Zap, err = zap.NewDevelopment()
		if err != nil {
			err = errors.Wrap(err, "failed to initialize development logger")
			return nil, err
		}
	case loggerTypeProduction:
		logger.Zap, err = zap.NewProduction()
		if err != nil {
			err = errors.Wrap(err, "failed to initialize production logger")
			return nil, err
		}
	default:
		logger.Zap = zap.NewExample()
	}
	return logger, nil
}
