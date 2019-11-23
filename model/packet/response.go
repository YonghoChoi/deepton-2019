package packet

type Resp struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func New(code, message string, data interface{}) Resp {
	return Resp{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
