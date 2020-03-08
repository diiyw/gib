package restful

type ResponseJson struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(c int, m, v interface{}) ResponseJson {
	return ResponseJson{
		Code:    c,
		Message: m,
		Data:    v,
	}
}
