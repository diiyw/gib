package gerr

var defaultError = new(Error)

func New(options ...Option) error {

	for _, option := range options {
		option(defaultError)
	}

	return defaultError
}

// Error is a trivial implementation of error.
type Error struct {
	C int    `json:"code"`
	M string `json:"message"`
}

func (e *Error) Error() string {
	return e.M
}