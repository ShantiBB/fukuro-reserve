package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpswagger "github.com/swaggo/http-swagger"

	_ "hotel/docs"
	"hotel/internal/http/handler"
)

func New(r chi.Router, h *handler.Handler) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.StripSlashes)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/swagger/*", httpswagger.WrapHandler)

		hotelRouter("/{countryCode}/{citySlug}/hotels", r, h)
	})
}
