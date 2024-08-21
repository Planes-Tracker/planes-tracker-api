package dto

import "github.com/labstack/echo/v4"

func NewCollectionReponse(items interface{}, totalItems int64) *Response {
	return &Response{
		Ok: true,
		Data: &echo.Map{
			"items":      items,
			"totalItems": totalItems,
		},
	}
}
