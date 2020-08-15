package gog

import (
	"errors"
	"os"
)

type Option func(logger Logger) error

func Dir(dir string) Option {
	return func(logger Logger) error {
		l, ok := logger.(*FileLogger)
		if !ok {
			return errors.New("not *FileLogger")
		}
		l.Dir = "logs/" + dir
		return os.MkdirAll(l.Dir, 0777)
	}
}

func RawMessage() Option {
	return func(logger Logger) error {
		l, ok := logger.(*FileLogger)
		if !ok {
			return errors.New("not *FileLogger")
		}
		l.withRaw = true
		return nil
	}
}
