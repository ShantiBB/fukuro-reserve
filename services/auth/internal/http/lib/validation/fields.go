package validation

import (
	"errors"

	"github.com/go-playground/validator/v10"

	"auth/internal/http/dto/response"
	"fukuro-reserve/pkg/utils/consts"
)

func getErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return consts.FieldRequired.Error()
	case "email":
		return consts.InvalidEmail.Error()
	case "min":
		return consts.InvalidPassword.Error()
	default:
		return consts.InternalServer.Error()
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
