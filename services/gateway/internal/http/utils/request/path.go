package request

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func PositiveInt64Path(c *gin.Context, name string) (int64, error) {
	value := strings.TrimSpace(c.Param(name))
	parsed, err := strconv.ParseInt(value, 10, 64)
	if err != nil || parsed <= 0 {
		return 0, err
	}

	return parsed, nil
}
