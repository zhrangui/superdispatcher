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
	*zap.Logger
}

// New creates customized Logger
func New(loggerType int) (*Logger, error) {
	var (
		err error
	)
	logger := &Logger{}
	switch loggerType {
	case loggerTypeNop:
		logger.Logger = zap.NewNop()
	case loggerTypeDevelopment:
		logger.Logger, err = zap.NewDevelopment()
		if err != nil {
			err = errors.Wrap(err, "failed to initialize development logger")
			return nil, err
		}
	case loggerTypeProduction:
		logger.Logger, err = zap.NewProduction()
		if err != nil {
			err = errors.Wrap(err, "failed to initialize production logger")
			return nil, err
		}
	default:
		logger.Logger = zap.NewExample()
	}
	return logger, nil
}

// FailOnError exists system on fail
func (logger *Logger) FailOnError(err error, msg string) {
	if err != nil {
		logger.Sugar().Fatalf("%s: %s", msg, err)
	}
}
