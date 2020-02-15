package errors

type Option func(err *Error)

func String(s string) Option {
	return func(err *Error) {
		err.s = s
	}
}

func Code(i int) Option {
	return func(err *Error) {
		err.c = i
	}
}

func Wrap(err error) Option {
	return func(e *Error) {
		e.s = err.Error()
	}
}

func WrapString(s string, err error) Option {
	return func(e *Error) {
		e.s = s + err.Error()
	}
}
