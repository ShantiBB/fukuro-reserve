package validation

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"

	"hotel/internal/repository/models"
	"hotel/pkg/utils/consts"
)

func CustomValidationError(err validator.FieldError) string {
	value := err.Value()
	param := err.Param()

	switch err.Tag() {
	case "required":
		return consts.FieldRequired.Error()
	case "min":
		return fmt.Sprintf(consts.FieldMin.Error(), param)
	case "max":
		return fmt.Sprintf(consts.FieldMax.Error(), param)
	case "gt":
		return fmt.Sprintf(consts.FieldGt.Error(), param, value)
	case "gte":
		return fmt.Sprintf(consts.FieldGte.Error(), param, value)
	case "lt":
		return fmt.Sprintf(consts.FieldLt.Error(), param, value)
	case "lte":
		return fmt.Sprintf(consts.FieldLte.Error(), param, value)
	case "email":
		return consts.FieldEmail.Error()
	case "uuid":
		return consts.FieldUUID.Error()
	case "datetime":
		return fmt.Sprintf(consts.FieldDatetime.Error(), param)
	case "room_status":
		return fmt.Sprintf("must be one of: %s", strings.Join(func() []string {
			vals := make([]string, len(models.RoomStatusValues))
			for i, v := range models.RoomStatusValues {
				vals[i] = string(v)
			}
			return vals
		}(), ", "))
	default:
		return fmt.Sprintf(consts.FieldInvalid.Error(), value)
	}
}
