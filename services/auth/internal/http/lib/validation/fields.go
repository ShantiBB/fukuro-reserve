package validation

import (
	"errors"

	"github.com/go-playground/validator/v10"

	"auth/internal/http/lib/schemas/response"
	"fukuro-reserve/pkg/utils/errs"
)

func getErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return errs.FieldRequired.Error()
	case "email":
		return errs.InvalidEmail.Error()
	case "min":
		return errs.InvalidPassword.Error()
	default:
		return errs.InternalServer.Error()
	}
}

func formatValidationErrors(validationErrors validator.ValidationErrors) map[string]string {
	errorMessages := make(map[string]string)

	for _, err := range validationErrors {
		errorMessages[err.Field()] = getErrorMessage(err)
	}

	return errorMessages
}

func CheckErrors(v interface{}) *response.ValidateError {
	validate := validator.New()
	if err := validate.Struct(v); err != nil {
		var validateErr validator.ValidationErrors
		errors.As(err, &validateErr)

		errorsResp := response.ValidateError{
			Errors: formatValidationErrors(validateErr),
		}
		return &errorsResp
	}

	return nil
}
