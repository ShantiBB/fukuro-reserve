package consts

import "errors"

// Hotel
var (
	HotelNotFound    = errors.New("hotel not found")
	UniqueHotelField = errors.New("hotel title already exists")
)

// Room
var (
	RoomNotFound    = errors.New("room not found")
	UniqueRoomField = errors.New("room number already exists")
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
	FieldEmail           = "field must be a valid email"
	FieldUUID            = "field must be a valid UUID"
	FieldDatetime        = "field must be in the format %s"
	FieldEnum            = "field must be one of: %s"
)

var (
	InvalidID         = errors.New("invalid ID")
	InvalidQueryParam = errors.New("invalid query parameter")
	InternalServer    = errors.New("internal server error")
	InvalidJSON       = errors.New("invalid JSON body")
)
