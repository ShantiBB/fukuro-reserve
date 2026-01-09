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
var (
	FieldRequired = errors.New("field is required")
	FieldInvalid  = errors.New("field is invalid: %v")
	FieldMin      = errors.New("field must be at least %s")
	FieldMax      = errors.New("field must be at most %s")
	FieldGt       = errors.New("field must be > %s, got %v")
	FieldGte      = errors.New("field must be ≥ %s, got %v")
	FieldLt       = errors.New("field must be < %s, got %v")
	FieldLte      = errors.New("field must be ≤ %s, got %v")
	FieldEmail    = errors.New("field must be a valid email")
	FieldUUID     = errors.New("field must be a valid UUID")
	FieldDatetime = errors.New("field must be in the format %s")
)

var (
	InvalidID         = errors.New("invalid ID")
	InvalidQueryParam = errors.New("invalid query parameter")
	InternalServer    = errors.New("internal server error")
	InvalidJSON       = errors.New("invalid JSON body")
)
