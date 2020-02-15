package errors

var defaultError = new(Error)

func Throw(options ...Option) error {

	for _, option := range options {
		option(defaultError)
	}

	return defaultError
}

// Error is a trivial implementation of error.
type Error struct {
	c int
	s string
}

func (e *Error) Error() string {
	return e.s
}

func (e *Error) Code() int {
	return e.c
}
