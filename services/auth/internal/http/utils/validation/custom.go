package validation

import (
	"github.com/go-playground/validator/v10"

	"auth/pkg/lib/utils/consts"
)

func CustomValidationError(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return consts.ErrFieldRequired.Error()
	case "email":
		return consts.ErrInvalidEmail.Error()
	case "min":
		return consts.ErrInvalidPassword.Error()
	default:
		return consts.ErrInternalServer.Error()
	}
}
