package permission

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	Sub  int64
	Role string
	jwt.RegisteredClaims
}

const (
	RoleModerator = "moderator"
	RoleAdmin     = "admin"
)
