package helper

import (
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{Message: message, Code: code, Status: status}
	respon := Response{Meta: meta, Data: data}

	return respon
}

func ErrorResponse(err error) []string {
	var errorss []string
	switch err.(type) {
	case validator.FieldError:
		for _, e := range err.(validator.ValidationErrors) {
			errorss = append(errorss, e.Error())
		}
		break
	case validator.ValidationErrors:
		for _, e := range err.(validator.ValidationErrors) {
			errorss = append(errorss, e.Error())
		}
		break
	}
	return errorss
}
