package restful

import "github.com/diiyw/gib/errors"

type Type func() (code int, message string, data interface{})

func OK() Type {
	return func() (code int, message string, data interface{}) {
		return 200, "ok", nil
	}
}

func Error(s string) Type {
	return func() (code int, message string, data interface{}) {
		return 300, s, nil
	}
}

func Success(v interface{}) Type {
	return func() (code int, message string, data interface{}) {
		return 200, "error", v
	}
}

func WithError(err errors.Error) Type {
	return func() (code int, message string, data interface{}) {
		return err.Code(), err.Error(), nil
	}
}
