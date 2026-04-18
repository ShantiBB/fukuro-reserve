package request

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func OptionalUint64Query(c *gin.Context, key string) (uint64, error) {
	raw := strings.TrimSpace(c.Query(key))
	if raw == "" {
		return 0, nil
	}

	return strconv.ParseUint(raw, 10, 64)
}

func OptionalPositiveInt64Query(c *gin.Context, key string) (int64, error) {
	raw := strings.TrimSpace(c.Query(key))
	if raw == "" {
		return 0, nil
	}

	value, err := strconv.ParseInt(raw, 10, 64)
	if err != nil || value <= 0 {
		return 0, err
	}

	return value, nil
}

func FirstNonEmptyQuery(c *gin.Context, keys ...string) string {
	for _, key := range keys {
		value := strings.TrimSpace(c.Query(key))
		if value != "" {
			return value
		}
	}
	return ""
}
