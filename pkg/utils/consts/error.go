package consts

import "errors"

// auth
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

// hotel
var (
	HotelNotFound    = errors.New("hotel not found")
	UniqueHotelField = errors.New("hotel name already exists")
)

// room
var (
	RoomNotFound    = errors.New("room not found")
	UniqueRoomField = errors.New("room number already exists")
)

var (
	InvalidID         = errors.New("invalid ID")
	InvalidQueryParam = errors.New("invalid query parameter")
	FieldRequired     = errors.New("field is required")
	InternalServer    = errors.New("internal server error")
	Forbidden         = errors.New("forbidden")
	InvalidJSON       = errors.New("invalid JSON body")
)
