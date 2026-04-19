package auth

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/consts"
)

var ErrAuthorizationHeaderRequired = errors.New(consts.ErrAuthorizationHeaderRequired)

func OutgoingContextWithAuthorization(c *gin.Context) (context.Context, error) {
	authHeader := c.GetHeader(consts.HeaderAuthorization)
	if authHeader == "" {
		return nil, ErrAuthorizationHeaderRequired
	}

	return metadata.NewOutgoingContext(c.Request.Context(), metadata.Pairs("authorization", authHeader)), nil
}
