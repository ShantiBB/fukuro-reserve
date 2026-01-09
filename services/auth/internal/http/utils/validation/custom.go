package validation

import (
	"github.com/go-playground/validator/v10"

	"auth/pkg/utils/consts"
)

func CustomValidationError(err validator.FieldError) string {
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
