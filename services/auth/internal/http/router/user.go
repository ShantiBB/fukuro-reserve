package router

import (
	"github.com/go-chi/chi/v5"

	"auth/internal/http/handler"
	"auth/internal/http/lib/permission"
)

func userRouter(pattern string, r chi.Router, h *handler.Handler, jwtSecret string) {
	r.Route(pattern, func(r chi.Router) {
		r.Use(permission.AuthRequire(jwtSecret))

		r.With(permission.RequireRoles(isAdmin, isModerator)).Get("/", h.UserGetAll)
		r.With(permission.RequireRoles(isAdmin, isModerator, isOwner)).Get("/{id}", h.UserGetByID)

		r.Group(func(r chi.Router) {
			r.Use(permission.RequireRoles(isAdmin))
			r.Post("/", h.UserCreate)
			r.Put("/{id}/role", h.UserUpdateRoleStatus)
			r.Put("/{id}/status", h.UserUpdateActiveStatus)
		})

		r.Group(func(r chi.Router) {
			r.Use(permission.RequireRoles(isOwner, isAdmin))
			r.Put("/{id}", h.UserUpdateByID)
			r.Delete("/{id}", h.UserDeleteByID)
		})
	})
}
