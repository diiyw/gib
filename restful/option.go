package restful

type Option func() (code int, message string, data interface{})

func OK() ResponseJson {
	return Response(200, "ok", nil)
}

func Error(v interface{}) ResponseJson {
	if err, ok := v.(error); ok {
		return Response(300, err.Error(), nil)
	}
	return Response(300, v, nil)
}

func Success(v interface{}) ResponseJson {
	return Response(200, "success", v)
}
