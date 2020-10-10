package gerr

import "fmt"

var defaultError = new(Error)

func New(options ...Option) error {

	for _, option := range options {
		option(defaultError)
	}

	return defaultError
}

func NewError(code int, message string) *Error {
	return &Error{code, message}
}

// Error is a trivial implementation of error.
type Error struct {
	C int    `json:"code"`
	M string `json:"message"`
}

func (e *Error) Error() string {
	return e.M
}

func (e *Error) With(a ...interface{}) *Error {
	return &Error{
		C: e.C,
		M: fmt.Sprintf(e.M, a...),
	}
}
