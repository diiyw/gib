package restful

import "github.com/diiyw/gib/gerr"

type Option func() (code int, message string, data interface{})

func OK() ResponseJson {
	return Response(200, "ok", nil)
}

func Error(v interface{}) ResponseJson {
	if err, ok := v.(gerr.Error); ok {
		return Response(err.C, err.M, nil)
	}
	if err, ok := v.(error); ok {
		return Response(300, err.Error(), nil)
	}
	return Response(300, v, nil)
}

func Success(v interface{}) ResponseJson {
	return Response(200, "success", v)
}
