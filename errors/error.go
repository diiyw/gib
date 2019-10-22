package errors

import "github.com/diiyw/goutils/logger"

func HandleError(err error) {
	if err != nil {
		logger.Interface.Error(err)
	}
}