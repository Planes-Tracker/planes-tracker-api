package dto

type Response struct {
	Ok   bool        `json:"ok"`
	Data interface{} `json:"data"`
}

func NewResponse(data interface{}) *Response {
	return &Response{
		Ok:   true,
		Data: data,
	}
}
