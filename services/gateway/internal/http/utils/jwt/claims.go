package jwt

import (
	"strconv"

	"github.com/golang-jwt/jwt/v5"
)

func ExtractUserID(claims jwt.MapClaims) (int64, bool) {
	for _, key := range []string{"sub", "Sub"} {
		switch value := claims[key].(type) {
		case float64:
			return int64(value), true
		case string:
			id, err := strconv.ParseInt(value, 10, 64)
			if err == nil {
				return id, true
			}
		}
	}

	return 0, false
}

func ExtractStringClaim(claims jwt.MapClaims, keys ...string) string {
	for _, key := range keys {
		if value, ok := claims[key].(string); ok {
			return value
		}
	}

	return ""
}
