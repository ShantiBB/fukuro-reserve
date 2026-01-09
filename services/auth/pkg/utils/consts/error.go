package consts

import "errors"

// Auth
var (
	UserNotFound        = errors.New("user not found")
	InvalidEmail        = errors.New("invalid email format")
	UniqueUserField     = errors.New("username or email already exists")
	ErrInvalidRole      = errors.New("invalid role status")
	PasswordHashing     = errors.New("error hashing password")
	InvalidPassword     = errors.New("minimum length 8 characters")
	InvalidCredentials  = errors.New("invalid credentials")
	InvalidRefreshToken = errors.New("invalid token")
	Unauthorized        = errors.New("unauthorized")
)

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
	InvalidSlug       = errors.New("invalid slug")
	InvalidQueryParam = errors.New("invalid query parameter")
	InternalServer    = errors.New("internal server error")
	Forbidden         = errors.New("forbidden")
	InvalidJSON       = errors.New("invalid JSON body")
)
