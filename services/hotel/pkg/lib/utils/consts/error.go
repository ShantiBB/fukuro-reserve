package consts

import "errors"

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
	FieldSlug            = "must contain only lowercase letters, numbers and hyphens (e.g., 'my-hotel-slug')"

	MsgHotelNotFound     = "hotel not found"
	MsgUniqueHotelField  = "hotel title already exists"
	MsgRoomNotFound      = "room not found"
	MsgUniqueRoomField   = "room number already exists"
	MsgInvalidHotelID    = "invalid hotel id"
	MsgInvalidRoomID     = "invalid room id"
	MsgInvalidPrice      = "invalid price"
	MsgInvalidQueryParam = "invalid query parameter"
	MsgInternalServer    = "internal server error"
	MsgInvalidJSON       = "invalid JSON body"
)

var (
	ErrHotelNotFound     = errors.New(MsgHotelNotFound)
	ErrUniqueHotelField  = errors.New(MsgUniqueHotelField)
	ErrRoomNotFound      = errors.New(MsgRoomNotFound)
	ErrUniqueRoomField   = errors.New(MsgUniqueRoomField)
	ErrInvalidHotelID    = errors.New(MsgInvalidHotelID)
	ErrInvalidRoomID     = errors.New(MsgInvalidRoomID)
	ErrInvalidPrice      = errors.New(MsgInvalidRoomID)
	ErrInvalidQueryParam = errors.New(MsgInvalidQueryParam)
	ErrInternalServer    = errors.New(MsgInternalServer)
	ErrInvalidJSON       = errors.New(MsgInvalidJSON)
)
