package response

type Response struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}

func Success(data any) *Response {
	return &Response{"Success.", data}
}

func Fail(data any) *Response {
	return &Response{"Fail.", data}
}
