package validation

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func slugFormatValidator(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	match, _ := regexp.MatchString(`^[a-z0-9]+(-[a-z0-9]+)*$`, value)
	return match
}
