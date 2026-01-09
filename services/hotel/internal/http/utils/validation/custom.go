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
		return consts.FieldRequired
	case "min":
		return fmt.Sprintf(consts.FieldMin, param)
	case "max":
		return fmt.Sprintf(consts.FieldMax, param)
	case "gt", "decimal_gt":
		return fmt.Sprintf(consts.FieldGt, param, value)
	case "gte":
		return fmt.Sprintf(consts.FieldGte, param, value)
	case "lt", "decimal_lt":
		return fmt.Sprintf(consts.FieldLt, param, value)
	case "lte":
		return fmt.Sprintf(consts.FieldLte, param, value)
	case "email":
		return consts.FieldEmail
	case "uuid":
		return consts.FieldUUID
	case "datetime":
		return fmt.Sprintf(consts.FieldDatetime, param)
	case "room_status":
		vals := make([]string, len(models.RoomStatusValues))
		for i, v := range models.RoomStatusValues {
			vals[i] = string(v)
		}
		return fmt.Sprintf(consts.FieldEnum, strings.Join(vals, ", "))
	case "room_type":
		vals := make([]string, len(models.RoomTypeValues))
		for i, v := range models.RoomTypeValues {
			vals[i] = string(v)
		}
		return fmt.Sprintf(consts.FieldEnum, strings.Join(vals, ", "))
	default:
		return fmt.Sprintf(consts.FieldInvalid, value)
	}
}
