package handler

import (
	"context"
	"net/http"

	"google.golang.org/grpc/metadata"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/consts"
	authservice "github.com/ShantiBB/fukuro-reserve/services/gateway/internal/service/auth"
)

type AuthHandler struct {
	service *authservice.Service
}

func NewAuthHandler(service *authservice.Service) *AuthHandler {
	return &AuthHandler{service: service}
}

func authContext(r *http.Request) (context.Context, error) {
	authHeader := r.Header.Get(consts.HeaderAuthorization)
	if authHeader == "" {
		return nil, http.ErrNoCookie
	}

	return metadata.NewOutgoingContext(
		r.Context(),
		metadata.Pairs("authorization", authHeader),
	), nil
}
