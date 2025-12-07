package validation

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type ErrorSchema struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type ValidateError struct {
	Errors map[string]string `json:"errors"`
}

func ErrorResp(err error) *ErrorSchema {
	return &ErrorSchema{Type: "error", Message: err.Error()}
}

func formatValidationErrors(
	validationErrors validator.ValidationErrors,
	customErr func(validator.FieldError) string,
) map[string]string {
	errorMessages := make(map[string]string)

	for _, err := range validationErrors {
		errorMessages[err.Field()] = customErr(err)
	}

	return errorMessages
}

func CheckErrors(v any, customErr func(validator.FieldError) string) *ValidateError {
	validate := validator.New()
	if err := validate.Struct(v); err != nil {
		var validateErr validator.ValidationErrors
		errors.As(err, &validateErr)

		errorsResp := ValidateError{
			Errors: formatValidationErrors(validateErr, customErr),
		}
		return &errorsResp
	}

	return nil
}
