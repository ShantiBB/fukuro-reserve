package router

import (
	"github.com/go-chi/chi/v5"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/middleware"
)

type authRoutes struct {
	h       AuthHandler
	pattern string
}

func NewAuthRoutes(pattern string, h AuthHandler) RouteRegistrar {
	return authRoutes{pattern: pattern, h: h}
}

func (ar authRoutes) Register(r chi.Router) {
	r.Route(ar.pattern, func(r chi.Router) {
		r.Post("/register", ar.h.Register)
		r.Post("/login", ar.h.Login)
		r.Post("/refresh", ar.h.RefreshToken)

		r.Group(func(r chi.Router) {
			r.Use(middleware.AuthMiddleware)
			r.Get("/users", ar.h.GetUsers)
			r.Post("/users", ar.h.CreateUser)
			r.Get("/users/{id}", ar.h.GetUser)
			r.Put("/users/{id}", ar.h.UpdateUser)
			r.Patch("/users/{id}/activity", ar.h.UpdateUserActivity)
			r.Patch("/users/{id}/role", ar.h.UpdateUserRole)
			r.Delete("/users/{id}", ar.h.DeleteUser)
		})
	})
}
