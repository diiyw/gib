package errors

import "github.com/diiyw/gib/logger"

func HandleError(err error) {
	if err != nil {
		logger.DefaultInterface.Error(err)
	}
}