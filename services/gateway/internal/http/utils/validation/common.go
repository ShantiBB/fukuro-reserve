package validation

import (
	"errors"
	"net/mail"
	"regexp"
	"strings"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/consts"
	"github.com/google/uuid"
)

var (
	currencyRegexp = regexp.MustCompile(consts.CurrencyPattern)
	amountRegexp   = regexp.MustCompile(consts.AmountPattern)
)

// ValidateUUID validates that a string is a valid UUID format.
func ValidateUUID(value string) error {
	if _, err := uuid.Parse(value); err != nil {
		return errors.New(consts.ErrInvalidUUIDFormat)
	}
	return nil
}

// ValidateCurrency validates that a string matches the currency pattern (3 uppercase letters).
func ValidateCurrency(value string) error {
	if !currencyRegexp.MatchString(value) {
		return errors.New(consts.ErrInvalidCurrency)
	}
	return nil
}

// ValidateAmount validates that a string matches the amount pattern (numeric with up to 18 decimal places).
func ValidateAmount(value string) error {
	if !amountRegexp.MatchString(value) {
		return errors.New(consts.ErrInvalidAmount)
	}
	return nil
}

func validateEmail(email string) error {
	if strings.TrimSpace(email) == "" {
		return errors.New(consts.ErrFieldValueRequired)
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return errors.New(consts.ErrInvalidEmail)
	}
	return nil
}
