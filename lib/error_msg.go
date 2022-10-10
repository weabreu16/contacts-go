package lib

import (
	"errors"
	"reflect"

	"github.com/go-playground/validator/v10"
)

type ErrorMsg struct {
	Message string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {

	var message string = fe.Field() + " "

	switch fe.Tag() {
	case "required":
		message += "field is required"
	case "lte":
		message += "should be less than " + fe.Param()
	case "gte":
		message += "should be greater than " + fe.Param()
	case "min":
		if fe.Type().Kind() == reflect.String {
			message += "minimum length is " + fe.Param()
		} else {
			message += "should be greater than " + fe.Param()
		}
	case "max":
		if fe.Type().Kind() == reflect.String {
			message += "maximum length is " + fe.Param()
		} else {
			message += "should be less than " + fe.Param()
		}
	case "uuid":
		message += "should be a valid uuid"
	default:
		return "Unknown error"
	}

	return message
}

type ErrorMsgs []ErrorMsg

func GetErrorMsgs(err error) ErrorMsgs {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ErrorMsg, len(ve))
		for i, fe := range ve {
			out[i] = ErrorMsg{getErrorMsg(fe)}
		}
		return out
	}
	return ErrorMsgs{ErrorMsg{"Unknown error"}}
}
