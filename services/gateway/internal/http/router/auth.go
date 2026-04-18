package router

import (
	"github.com/go-chi/chi/v5"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/handler"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/middleware"
)

func authRouter(pattern string, r chi.Router, h *handler.AuthHandler) {
	r.Route(pattern, func(r chi.Router) {
		r.Post("/register", h.Register)
		r.Post("/login", h.Login)
		r.Post("/refresh", h.RefreshToken)

		r.Group(func(r chi.Router) {
			r.Use(middleware.AuthMiddleware)
			r.Get("/users", h.GetUsers)
			r.Post("/users", h.CreateUser)
			r.Get("/users/{id}", h.GetUser)
			r.Put("/users/{id}", h.UpdateUser)
			r.Patch("/users/{id}/activity", h.UpdateUserActivity)
			r.Patch("/users/{id}/role", h.UpdateUserRole)
			r.Delete("/users/{id}", h.DeleteUser)
		})
	})
}
