package restful

type ResponseJson struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(t Type) *ResponseJson {
	c, m, v := t()
	return &ResponseJson{
		Code:    c,
		Message: m,
		Data:    v,
	}
}
