package logger

import (
	"go.uber.org/zap"
)

type Logger struct {
	*zap.SugaredLogger
	driver *zap.Logger
}

var StdOut = new(Logger)

func init() {
	if StdOut.driver != nil {
		return
	}
	StdOut = NewLogger()
}

func NewLogger() *Logger {
	logger := new(Logger)
	logger.driver, _ = zap.NewProduction()
	logger.SugaredLogger = logger.driver.Sugar()
	return logger
}
