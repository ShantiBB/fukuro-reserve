package request

import "github.com/gin-gonic/gin"

func BindJSON(c *gin.Context, dst interface{}) error {
	return c.ShouldBindJSON(dst)
}
