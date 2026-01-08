package router

import (
	"github.com/go-chi/chi/v5"

	"hotel/internal/http/handler"
)

func hotelRouter(pattern string, r chi.Router, h *handler.Handler) {
	r.Route(pattern, func(r chi.Router) {
		r.Post("/", h.HotelCreate)
		r.Get("/", h.HotelGetAll)
		r.Get("/{hotelSlug}", h.HotelGetBySlug)
		r.Put("/{hotelSlug}", h.HotelUpdateBySlug)
		r.Delete("/{hotelSlug}", h.HotelDeleteBySlug)

		roomRouter("/{hotelSlug}/rooms", r, h)
	})
}
