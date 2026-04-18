package router

import (
	"github.com/go-chi/chi/v5"

	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/handler"
	"github.com/ShantiBB/fukuro-reserve/services/gateway/internal/http/middleware"
)

func hotelRouter(pattern string, r chi.Router, h *handler.HotelHandler) {
	r.Route(pattern, func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)

		r.Post("/", h.CreateHotel)
		r.Get("/", h.GetHotels)
		r.Get("/{countryCode}/{citySlug}/{hotelSlug}", h.GetHotel)
		r.Put("/{countryCode}/{citySlug}/{hotelSlug}", h.UpdateHotel)
		r.Patch("/{countryCode}/{citySlug}/{hotelSlug}/title", h.UpdateHotelTitle)
		r.Delete("/{countryCode}/{citySlug}/{hotelSlug}", h.DeleteHotel)
	})
}
