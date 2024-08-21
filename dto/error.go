package dto

type Error struct {
	Ok      bool        `json:"ok"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewError(message string) *Error {
	return &Error{
		Ok:      false,
		Message: message,
		Data:    nil,
	}
}
