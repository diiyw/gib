package logger

import (
	"go.uber.org/zap"
)

type Logger struct {
	*zap.SugaredLogger
	driver *zap.Logger
}

var DefaultInterface = new(Logger)

func init() {
	if DefaultInterface.driver != nil {
		return
	}
	DefaultInterface = NewLogger()
}

func NewLogger() *Logger {
	logger := new(Logger)
	logger.driver, _ = zap.NewProduction()
	logger.SugaredLogger = logger.driver.Sugar()
	return logger
}
