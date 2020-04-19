package errors

type Option func(err *Error)

func String(s string) Option {
	return func(err *Error) {
		err.M = s
	}
}

func Code(i int) Option {
	return func(err *Error) {
		err.C = i
	}
}

func Wrap(err error) Option {
	return func(e *Error) {
		e.M = err.Error()
	}
}

func WrapString(s string, err error) Option {
	return func(e *Error) {
		e.M = s + err.Error()
	}
}
