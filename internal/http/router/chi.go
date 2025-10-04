package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"auth_service/internal/http/handler"
	"auth_service/package/utils/permission"
)

var (
	isAdmin     = permission.IsAdmin
	isModerator = permission.IsModerator
	isOwner     = permission.IsOwner
)

func New(r chi.Router, h *handler.Handler, jwtSecret string) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	authRouter("/auth", r, h)
	userRouter("/users", r, h, jwtSecret)
}
