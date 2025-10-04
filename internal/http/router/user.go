package router

import (
	"github.com/go-chi/chi/v5"

	"auth_service/internal/http/handler"
	"auth_service/package/utils/permission"
)

func userRouter(pattern string, r chi.Router, h *handler.Handler, jwtSecret string) {
	r.Route(pattern, func(r chi.Router) {
		r.Use(permission.AuthRequired(jwtSecret))

		r.With(permission.Require(isAdmin)).Post("/", h.UserCreate)
		r.With(permission.Require(isAdmin, isModerator)).Get("/", h.UserList)
		r.With(permission.Require(isOwner)).Get("/{id}", h.UserGetByID)

		r.Group(func(r chi.Router) {
			r.Use(permission.Require(isOwner, isAdmin))
			r.Put("/{id}", h.UserUpdateByID)
			r.Delete("/{id}", h.UserDeleteByID)
		})
	})
}
