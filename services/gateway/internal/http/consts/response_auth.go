package consts

const (
	ErrAuthorizationHeaderRequired = "authorization header is required"
	ErrInvalidAuthorizationHeader  = "invalid authorization header format"
	ErrInvalidToken                = "invalid token"
	ErrInvalidTokenClaims          = "invalid token claims"
	ErrInvalidJWTUserIDClaim       = "invalid user id in token"
	ErrInvalidUserID               = "invalid user id"
	ErrUserIDMustBePositiveInteger = "userId must be a positive integer"
)
