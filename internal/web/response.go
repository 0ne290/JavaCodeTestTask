package response

type Response struct {
	Status string
	Data   any
}

func Success(data any) *Response {
	return &Response{"Success.", data}
}

func Fail(data any) *Response {
	return &Response{"Fail.", data}
}
