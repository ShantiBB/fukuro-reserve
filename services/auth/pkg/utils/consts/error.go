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

// Field
var (
	FieldRequired = errors.New("field is required")
)

var (
	InvalidID         = errors.New("invalid ID")
	InvalidQueryParam = errors.New("invalid query parameter")
	InternalServer    = errors.New("internal server error")
	Forbidden         = errors.New("forbidden")
	InvalidJSON       = errors.New("invalid JSON body")
)
