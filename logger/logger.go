package logger

import (
	"go.uber.org/zap"
)

type Logger struct {
	*zap.SugaredLogger
	driver *zap.Logger
}

var Interface = new(Logger)

func init() {
	if Interface.driver != nil {
		return
	}
	Interface = NewLogger()
}

func NewLogger() *Logger {
	logger := new(Logger)
	logger.driver, _ = zap.NewProduction()
	logger.SugaredLogger = logger.driver.Sugar()
	return logger
}
