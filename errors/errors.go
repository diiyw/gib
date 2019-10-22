package errors

func New(code int, text string) error {
	return &Error{code, text}
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
