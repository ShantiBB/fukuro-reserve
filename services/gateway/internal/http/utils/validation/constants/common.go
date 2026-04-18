package constants

const (
	CurrencyPattern = `^[A-Z]{3}$`
	AmountPattern   = `^[0-9]+(\.[0-9]{1,18})?$`
)

const (
	ErrWrapWithField      = "%s %w"
	ErrInvalidUUIDFormat  = "must be a valid UUID format"
	ErrInvalidCurrency    = "must be a valid 3-letter currency code (e.g., USD, EUR)"
	ErrInvalidAmount      = "must be a valid numeric amount (up to 18 decimal places)"
	ErrFieldValueRequired = "is required"
	ErrInvalidEmail       = "must be a valid email"
)

const (
	FieldEmail = "email"
)
