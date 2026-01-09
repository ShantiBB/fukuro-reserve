package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"

	"hotel/internal/repository/models"
)

func roomStatusValidator(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	for _, v := range models.RoomStatusValues {
		if string(v) == value {
			return true
		}
	}
	return false
}

func roomTypeValidator(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	for _, v := range models.RoomTypeValues {
		if string(v) == value {
			return true
		}
	}
	return false
}

func decimalGtValidator(fl validator.FieldLevel) bool {
	dec, ok := fl.Field().Interface().(decimal.Decimal)
	if !ok {
		return false
	}
	param := fl.Param()
	compareValue, err := decimal.NewFromString(param)
	if err != nil {
		return false
	}
	return dec.Cmp(compareValue) > 0
}

func decimalLtValidator(fl validator.FieldLevel) bool {
	dec, ok := fl.Field().Interface().(decimal.Decimal)
	if !ok {
		return false
	}
	param := fl.Param()
	compareValue, err := decimal.NewFromString(param)
	if err != nil {
		return false
	}
	return dec.Cmp(compareValue) < 0
}
