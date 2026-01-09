package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpswagger "github.com/swaggo/http-swagger"

	_ "auth/docs"
	"auth/internal/http/handler"
	"auth/internal/http/utils/permission"
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

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/swagger/*", httpswagger.WrapHandler)
		authRouter("/auth", r, h)
		userRouter("/users", r, h, jwtSecret)
	})
}
