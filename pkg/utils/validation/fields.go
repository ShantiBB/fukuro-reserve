package validation

import (
	"errors"
	"reflect"
	"strings"

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
	errorMessages := make(map[string]string, len(validationErrors))

	for _, fe := range validationErrors {
		ns := fe.Namespace()
		if i := strings.IndexByte(ns, '.'); i >= 0 {
			ns = ns[i+1:]
		}

		errorMessages[ns] = customErr(fe)
	}

	return errorMessages
}

func CheckErrors(v any, customErr func(validator.FieldError) string) *ValidateError {
	validate := validator.New()

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" || name == "" {
			return fld.Name
		}
		return name
	})

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
