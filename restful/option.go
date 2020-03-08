package restful

import "github.com/diiyw/gib/errors"

type Option func() (code int, message string, data interface{})

func OK() ResponseJson {
	return Response(200, "ok", nil)
}

func Error(v interface{}) ResponseJson {
	return Response(300, v, nil)
}

func Success(v interface{}) ResponseJson {
	return Response(200, "success", v)
}

func WithError(err errors.Error) ResponseJson {
	return Response(err.Code(), err.Error(), nil)
}
