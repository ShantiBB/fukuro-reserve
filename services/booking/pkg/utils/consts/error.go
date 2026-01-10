package consts

import "errors"

// Booking
var (
	BookingNotFound    = errors.New("booking not found")
	UniqueBookingField = errors.New("booking title already exists")
)

// Field
const (
	ValidationUnregister = "failed to register validation"
	FieldRequired        = "field is required"
	FieldInvalid         = "field is invalid: %v"
	FieldMin             = "field must be at least %s"
	FieldMax             = "field must be at most %s"
	FieldGt              = "field must be > %s, got %v"
	FieldGte             = "field must be ≥ %s, got %v"
	FieldLt              = "field must be < %s, got %v"
	FieldLte             = "field must be ≤ %s, got %v"
	FieldUUID            = "field must be a valid UUID"
	FieldDatetime        = "field must be in the format %s"
	FieldEnum            = "field must be one of: %s"
	FieldSlug            = "must contain only lowercase letters, numbers and hyphens (e.g., 'my-booking-slug')"
)

var (
	InvalidID         = errors.New("invalid ID")
	InvalidQueryParam = errors.New("invalid query parameter")
	InternalServer    = errors.New("internal server error")
	InvalidJSON       = errors.New("invalid JSON body")
)
