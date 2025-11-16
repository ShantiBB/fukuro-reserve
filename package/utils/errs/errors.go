package errs

import "errors"

var (
	InternalServer      = errors.New("internal server error")
	UserNotFound        = errors.New("user not found")
	UserRetrieving      = errors.New("error retrieving users")
	UniqueUserField     = errors.New("username or email already exists")
	Unauthorized        = errors.New("unauthorized")
	Forbidden           = errors.New("forbidden")
	FieldRequired       = errors.New("field is required")
	InvalidID           = errors.New("invalid user ID")
	InvalidEmail        = errors.New("invalid email format")
	InvalidPassword     = errors.New("minimum length 8 characters")
	InvalidCredentials  = errors.New("invalid credentials")
	InvalidRefreshToken = errors.New("invalid token")
	InvalidJSON         = errors.New("invalid JSON body")
	PasswordHashing     = errors.New("error hashing password")
)
