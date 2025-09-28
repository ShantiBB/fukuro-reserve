package validation

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"

	"auth_service/api/http/schemas"
)

func formatValidationErrors(validationErrors validator.ValidationErrors) map[string]string {
	errorMessages := make(map[string]string)

	for _, err := range validationErrors {
		errorMessages[err.Field()] = getErrorMessage(err)
	}

	return errorMessages
}

func getErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "Поле обязательно для заполнения"
	case "email":
		return "Неверный формат email"
	case "min":
		return fmt.Sprintf("Минимальная длина %s символов", err.Param())
	default:
		return "Некорректное значение"
	}
}

func CheckErrors(v interface{}) *schemas.ValidateErrorResponse {
	validate := validator.New()
	if err := validate.Struct(v); err != nil {
		var validateErr validator.ValidationErrors
		errors.As(err, &validateErr)

		errorsResp := schemas.ValidateErrorResponse{
			Errors: formatValidationErrors(validateErr),
		}
		return &errorsResp
	}

	return nil
}
