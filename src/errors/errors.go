package errors

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	ErrorMessage string     `json:"errorMessage"`
	Errors       []ErrorMsg `json:"errors"`
}

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	}
	return "Unknown error"
}

func ParseValidationErrors(err error) ErrorResponse {
	var ve validator.ValidationErrors

	if errors.As(err, &ve) {
		out := make([]ErrorMsg, len(ve))
		for i, fe := range ve {
			out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
		}

		return ErrorResponse{ErrorMessage: "validation error", Errors: out}
	}

	return ErrorResponse{ErrorMessage: "unknown error"}
}

func FilterNonRequiredErrors(errResponse ErrorResponse) ErrorResponse {
	var errors []ErrorMsg

	for ind, e := range errResponse.Errors {
		if !strings.Contains(e.Message, "required") {
			errors = append(errors, errResponse.Errors[ind])
		}
	}

	errResponse.Errors = errors

	return errResponse
}
