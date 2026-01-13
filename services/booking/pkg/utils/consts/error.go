package consts

import "errors"

// Bookings
var (
	BookingNotFound      = errors.New("booking not found")
	BookingRoomNotFound  = errors.New("booking room not found")
	RoomLockNotFound     = errors.New("room lock room not found")
	RoomLockAlreadyExist = errors.New("room lock already exists")
	ErrInvalidDates      = errors.New("invalid booking dates")
	ErrPriceChanged      = errors.New("expected total amount does not match calculated total")
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
	InvalidHotelID               = errors.New("invalid hotel ID")
	InvalidRoomID                = errors.New("invalid room ID")
	InvalidPricePerNightID       = errors.New("invalid price per night. example: 123.45")
	InvalidExpectedTotalAmountID = errors.New("invalid expected total amount. example: 123.45")
	InvalidQueryParam            = errors.New("invalid query parameter")
	InternalServer               = errors.New("internal server error")
	InvalidJSON                  = errors.New("invalid JSON body")
)
